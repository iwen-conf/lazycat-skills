# Homebrew 发布说明

这份说明面向准备把 `lzc-copy-image` 单独抽成外部仓库并通过 Homebrew tap 发布的人。

## 选择说明

这里采用的是 **GitHub Releases + Homebrew tap formula**，而不是 npm。

原因：

- 运行时不需要 Node
- 安装后就是原生二进制
- 对无签名 CLI 来说，tap formula 比 cask 更省事
- 同一套发布物还能覆盖 Linuxbrew

说明：

- GoReleaser 官方从 v2.10 起把 Homebrew Formula 标记为 deprecated，更推荐 cask。
- 这里仍保留 formula 方案，是一个工程取舍：当前这个工具是无签名命令行程序，formula tap 更容易落地。

## 需要的仓库

至少准备两个 GitHub 仓库：

1. `your-github-user/lzc-copy-image`
   用来存放 CLI 源码和 GitHub Releases。
2. `your-github-user/homebrew-tap`
   用来存放 Homebrew formula。

## 需要的 Secrets

在源码仓库中配置：

- `HOMEBREW_TAP_GITHUB_TOKEN`

要求：

- 这个 token 需要对 tap 仓库有内容写权限。
- 不要只用 GitHub Actions 默认的 `GITHUB_TOKEN` 去改另一个仓库。

## 首次发布前要改的内容

在 `.goreleaser.yaml` 中替换这些占位值：

- `release.github.owner`
- `release.github.name`
- `brews[].homepage`
- `brews[].license`
- `brews[].repository.owner`
- `brews[].repository.name`
- `brews[].commit_author`

## 发布流程

1. 把这个目录抽到独立仓库。
2. 补 `LICENSE`。
3. 提交并打标签，例如：

```bash
git tag v0.1.0
git push origin v0.1.0
```

4. GitHub Actions 触发 `goreleaser release --clean`。
5. GoReleaser：
   - 构建 `darwin/linux` 的 `amd64/arm64` 二进制
   - 上传 GitHub Releases
   - 生成校验和
   - 更新 `homebrew-tap` 仓库中的 formula

## 用户安装

```bash
brew tap your-github-user/tap
brew install lzc-copy-image
```

升级：

```bash
brew update
brew upgrade lzc-copy-image
```

## 建议

- 默认先支持 `darwin/arm64`，这是你当前机器最直接的目标平台。
- 并发默认保持 `2` 或 `3`，不要把 Homebrew 发布和高并发运行混为一谈。
- 如果以后要做已签名 macOS 分发，再考虑迁移到 cask。
