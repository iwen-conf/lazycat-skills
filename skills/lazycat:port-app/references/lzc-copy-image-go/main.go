package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"
)

var registryPattern = regexp.MustCompile(`registry\.lazycat\.cloud/[^\s]+`)

type config struct {
	images          string
	imagesFile      string
	concurrency     int
	timeout         time.Duration
	continueOnError bool
	format          string
	binary          string
}

type job struct {
	index int
	image string
}

type result struct {
	Image       string `json:"image"`
	ImageLzcURL string `json:"imageLzcUrl,omitempty"`
	OK          bool   `json:"ok"`
	Error       string `json:"error,omitempty"`
	DurationMS  int64  `json:"durationMs"`
}

type output struct {
	Results      []result `json:"results"`
	SuccessCount int      `json:"successCount"`
	FailureCount int      `json:"failureCount"`
}

func main() {
	os.Exit(run())
}

func run() int {
	cfg, err := parseFlags()
	if err != nil {
		writeError(err)
		return 2
	}

	if _, err := exec.LookPath(cfg.binary); err != nil {
		writeError(fmt.Errorf("未找到 %s，请先安装并确保其在 PATH 中", cfg.binary))
		return 2
	}

	images, err := loadImages(cfg)
	if err != nil {
		writeError(err)
		return 2
	}
	if len(images) == 0 {
		writeError(errors.New("镜像列表为空"))
		return 2
	}

	results, failed := executeBatch(cfg, images)
	out := output{
		Results:      results,
		SuccessCount: len(results) - failed,
		FailureCount: failed,
	}

	if err := writeJSON(out); err != nil {
		fmt.Fprintf(os.Stderr, "输出结果失败: %v\n", err)
		return 2
	}

	if failed > 0 {
		return 1
	}
	return 0
}

func parseFlags() (config, error) {
	var cfg config

	flag.StringVar(&cfg.images, "images", "", "JSON array of image strings")
	flag.StringVar(&cfg.imagesFile, "images-file", "", "Path to a JSON file or newline-delimited image list")
	flag.IntVar(&cfg.concurrency, "concurrency", 2, "Max number of concurrent copy-image tasks")
	flag.DurationVar(&cfg.timeout, "timeout", 20*time.Minute, "Per-image timeout")
	flag.BoolVar(&cfg.continueOnError, "continue-on-error", false, "Continue processing after a failure")
	flag.StringVar(&cfg.format, "format", "json", "Output format, only json is supported")
	flag.StringVar(&cfg.binary, "binary", "lzc-cli", "Path to lzc-cli binary")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "用法: lzc-copy-image --images '[\"nginx:1.27\"]' --concurrency 2\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if cfg.images != "" && cfg.imagesFile != "" {
		return cfg, errors.New("--images 与 --images-file 不能同时使用")
	}
	if cfg.concurrency < 1 {
		return cfg, errors.New("--concurrency 必须大于等于 1")
	}
	if cfg.timeout <= 0 {
		return cfg, errors.New("--timeout 必须大于 0")
	}
	if cfg.format != "json" {
		return cfg, errors.New("当前仅支持 --format json")
	}

	return cfg, nil
}

func loadImages(cfg config) ([]string, error) {
	switch {
	case strings.TrimSpace(cfg.images) != "":
		return decodeImageList(cfg.images)
	case strings.TrimSpace(cfg.imagesFile) != "":
		data, err := os.ReadFile(cfg.imagesFile)
		if err != nil {
			return nil, fmt.Errorf("读取 --images-file 失败: %w", err)
		}
		return parseImageInput(string(data))
	default:
		stat, err := os.Stdin.Stat()
		if err != nil {
			return nil, fmt.Errorf("读取 stdin 状态失败: %w", err)
		}
		if stat.Mode()&os.ModeCharDevice != 0 {
			return nil, errors.New("请通过 --images、--images-file 或 stdin 提供镜像列表")
		}
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("读取 stdin 失败: %w", err)
		}
		return parseImageInput(string(data))
	}
}

func parseImageInput(raw string) ([]string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}
	if strings.HasPrefix(raw, "[") {
		return decodeImageList(raw)
	}

	lines := strings.Split(raw, "\n")
	images := make([]string, 0, len(lines))
	seen := make(map[string]struct{}, len(lines))
	for _, line := range lines {
		image := strings.TrimSpace(line)
		if image == "" {
			continue
		}
		if _, ok := seen[image]; ok {
			continue
		}
		seen[image] = struct{}{}
		images = append(images, image)
	}
	return images, nil
}

func decodeImageList(raw string) ([]string, error) {
	var input []string
	if err := json.Unmarshal([]byte(raw), &input); err != nil {
		return nil, fmt.Errorf("解析镜像 JSON 数组失败: %w", err)
	}

	images := make([]string, 0, len(input))
	seen := make(map[string]struct{}, len(input))
	for _, item := range input {
		image := strings.TrimSpace(item)
		if image == "" {
			continue
		}
		if _, ok := seen[image]; ok {
			continue
		}
		seen[image] = struct{}{}
		images = append(images, image)
	}
	return images, nil
}

func executeBatch(cfg config, images []string) ([]result, int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make([]result, len(images))
	jobs := make(chan job)
	var wg sync.WaitGroup
	var cancelOnce sync.Once

	workerCount := cfg.concurrency
	if workerCount > len(images) {
		workerCount = len(images)
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				if ctx.Err() != nil && !cfg.continueOnError {
					results[item.index] = result{
						Image: item.image,
						OK:    false,
						Error: "前序任务失败，当前任务已跳过",
					}
					continue
				}

				res := copyOne(ctx, cfg, item.image)
				results[item.index] = res
				if !res.OK && !cfg.continueOnError {
					cancelOnce.Do(cancel)
				}
			}
		}()
	}

	for idx, image := range images {
		jobs <- job{index: idx, image: image}
	}
	close(jobs)
	wg.Wait()

	failed := 0
	for _, res := range results {
		if !res.OK {
			failed++
		}
	}

	return results, failed
}

func copyOne(parent context.Context, cfg config, image string) result {
	started := time.Now()
	res := result{Image: image}

	cmdCtx, cancel := context.WithTimeout(parent, cfg.timeout)
	defer cancel()

	cmd := exec.CommandContext(cmdCtx, cfg.binary, "appstore", "copy-image", image)
	output, err := cmd.CombinedOutput()
	res.DurationMS = time.Since(started).Milliseconds()

	if err != nil {
		if errors.Is(cmdCtx.Err(), context.DeadlineExceeded) {
			res.Error = fmt.Sprintf("copy-image 超时: %s", cfg.timeout)
			return res
		}
		if errors.Is(parent.Err(), context.Canceled) && !cfg.continueOnError {
			res.Error = "任务被取消"
			return res
		}
		res.Error = buildCommandError(err, output)
		return res
	}

	match := registryPattern.FindString(string(output))
	if match == "" {
		res.Error = fmt.Sprintf("未从输出中解析到 Lazycat 镜像地址: %s", trimOutput(string(output)))
		return res
	}

	res.OK = true
	res.ImageLzcURL = match
	return res
}

func buildCommandError(err error, output []byte) string {
	message := strings.TrimSpace(err.Error())
	text := trimOutput(string(output))
	if text == "" {
		return message
	}
	return message + ": " + text
}

func trimOutput(raw string) string {
	cleaned := strings.Join(strings.Fields(strings.TrimSpace(raw)), " ")
	const limit = 400
	if len(cleaned) <= limit {
		return cleaned
	}
	return cleaned[:limit] + "..."
}

func writeJSON(v any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(v)
}

func writeError(err error) {
	_ = writeJSON(map[string]any{
		"results":      []result{},
		"successCount": 0,
		"failureCount": 0,
		"error":        err.Error(),
	})
}
