# lzc-copy-image-go

一个面向 `lzc-cli appstore copy-image` 的 Go 版批量包装器参考实现。

## 设计目标

- 不依赖 `npx`
- 不启动常驻后台进程
- 固定并发，避免无限放大内存和网络占用
- 每个镜像任务单独超时
- 输出结构化 JSON，方便后续自动回填 `lzc-manifest.yml`

## 用法

```bash
go run . \
  --images '["nginx:1.27","yuepaiji/myapp:v1"]' \
  --concurrency 2 \
  --format json
```

也支持从文件读取：

```bash
go run . --images-file ./images.json
```

`images.json` 可以是：

```json
["nginx:1.27", "yuepaiji/myapp:v1"]
```

## 输出

```json
{
  "results": [
    {
      "image": "nginx:1.27",
      "imageLzcUrl": "registry.lazycat.cloud/...",
      "ok": true,
      "durationMs": 1823
    }
  ],
  "successCount": 1,
  "failureCount": 0
}
```

失败项会保留 `error` 字段。

## Homebrew 发布

推荐走独立仓库 + GitHub Releases + Homebrew tap：

```bash
brew tap your-github-user/tap
brew install lzc-copy-image
```

参考文件：

- `.goreleaser.yaml`
- `release-homebrew.md`
- `.github/workflows/release.yml`

这套方案的特点：

- 安装时不依赖 `npx`
- 运行时不依赖 Node
- 可以直接分发原生二进制
- tap formula 更适合这个无签名 CLI 的分发方式

## 发布建议

- 这个目录是参考实现，便于从 skill 包中抽离。
- 真正发布到外部仓库时，建议把 `module` 路径替换成你自己的仓库地址。
- 默认并发建议保持在 `2` 或 `3`，不要盲目拉高。
