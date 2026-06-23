---
name: lazycat:ship-app
description: "Ship a ready Lazycat app only: verify metadata, copy-image, build LPK, install-test, prepare Developer Center materials, submit, and verify release. 上架发布、打包、copy-image、安装验证、提审、发布后检查."
---

# Lazycat App Shipping

## 职责范围

只处理已准备好的懒猫应用上架发布：打包、镜像同步、安装验证、开发者中心资料、提审和发布后检查。它不创建原创应用、不发现迁移候选、不做许可证门禁、不更新用户已安装应用。

## 输入

- 已准备好的原创应用项目，且已通过 `lazycat:original-app` 基线。
- 或已准备好的迁移项目，且已通过：
  - `lazycat:migration-license`
  - `lazycat:migration-boundary`
- 项目路径、目标版本、上架资料、截图、测试账号或免密登录说明。
- 上架平台选择；除非用户明确要求并提供移动端、智慧屏等验证证据，默认只选择桌面端。
- 如需提审：开发者中心登录态或凭据变量。

## 凭据环境变量约定

如流程需要登录懒猫微服、应用商店或开发者中心，优先读取当前进程环境变量；只有变量缺失或为空时才向用户报告缺失项。

- 微服设备：`LAZYCAT_USERNAME` / `LAZYCAT_PASSWORD`
- 懒猫应用商店：`LAZYCAT_APPSTORE_USERNAME` / `LAZYCAT_APPSTORE_PASSWORD`
- 懒猫开发者中心：`LAZYCAT_DEVELOPMENT_USERNAME` / `LAZYCAT_DEVELOPMENT_PASSWORD`

使用约束：

1. 只通过环境变量读取凭据，不把值写入仓库文件、报告、截图说明、提交信息或长期记忆。
2. 不为确认变量而执行 `echo $LAZYCAT_PASSWORD`、`printenv LAZYCAT_PASSWORD` 等会输出明文的命令；只允许用 `test -n "${LAZYCAT_PASSWORD:-}"` 这类方式检查是否存在。
3. 浏览器自动化登录时直接填入变量值，日志和回复只记录变量名、站点和操作结果，不记录值。
4. CLI 登录或 `copy-image` 需要凭据时优先复用已有登录态；确需传参时避免把密码放在命令行参数、shell history 或可见输出中。
5. 凭据只解锁认证，不扩大本技能权限；提交、发布、安装覆盖、远程写操作仍按本技能边界和当前任务授权执行。

## 输出

- 上架前检查结果。
- 镜像同步和 manifest 回写结果。
- `.lpk` 构建结果、包体大小和内嵌镜像检查结果。
- `lzc-cli lpk info <file.lpk>` 摘要。
- 安装验证、核心流程验证、平台选择、截图/资料状态。
- 提审状态或阻塞原因。

## 前置条件

1. 原创项目必须已完成 `lazycat:original-app` 后置条件。
2. 迁移项目必须有 `lazycat:migration-license` 和 `lazycat:migration-boundary` 输出。
3. 迁移项目许可证不得为 `Unclear` 或 `Blocked`。
4. 迁移项目不得依赖未授权的业务代码修改。
5. 迁移项目必须带有原作者名称和源项目或代码地址，且证据来自许可证门禁记录或上游仓库。
6. 提交开发者中心前，所有必填资料必须完整，不得有占位符或待补内容。

## 允许执行

- 修改上架相关 metadata、包装层、manifest、构建脚本、图标、截图和文档。
- 执行 `lzc-cli appstore copy-image` 并把返回的 `registry.lazycat.cloud/...` 写回源 manifest。
- 执行 `make build`、`lzc-cli project release` 或项目等价打包命令。
- 执行 `lzc-cli lpk info <file.lpk>`。
- 经用户确认后执行安装、提审或发布相关命令。
- 审核失败时修改包装层、运行时适配层、上架资料和审核说明。

## 禁止执行

- 不接收许可证不明、不可商用、不可再分发或查重不可接受的迁移项目。
- 未经当前任务明确授权，不修改迁移项目上游业务代码。
- 不把截图、网页表单或本地开发环境当作真实安装验证。
- 不使用 `lzc-build.yml.images`、`embed:<alias>`、包内 `images/` 或 `images.lock`。
- 不提交包含真实密码、Token、内部地址或真实微服域名的资料。
- 不上架色情、赌博、毒品、空投、破解软件或违法内容。
- 不处理已安装应用批量更新；改用 `lazycat:update-installed-app`。

## 上架门禁

- `package.yml` 包含 `package`、`version`、`name`、`description`、`author`、`license`、`locales`。
- `version` 为严格 `x.x.x`。
- 除非用户明确要求并已验证移动端或智慧屏，开发者中心上架平台默认只选择桌面端。
- 桌面端默认上架时，`package.yml` 必须声明不支持移动端和智慧屏：

```yaml
unsupported_platforms:
  - ios
  - tvos
  - android
```

用户或旧文档称为 manifest 平台字段时，LPK v2 按 `package.yml` 静态元数据处理；只有旧版项目仍使用 manifest 字段时才同步保持一致。

- 如果 `package.yml.homepage` 是 GitHub URL，`package.yml.author` 必须逐字符等于 URL owner 段。
- `lzc-manifest.yml` 服务、路由、持久化、健康检查、环境变量完整。
- 镜像型项目已 `copy-image`，manifest 使用最终 `registry.lazycat.cloud/...` 地址。
- 最终 `.lpk` 大小 `<= 12,000,000` bytes。
- 最终 `.lpk` 不包含 `images/` 或 `images.lock`。
- 有账号的应用提供 OIDC、公开测试账号、固定初始凭据或明确初始化说明。
- 有文件能力的应用完成文件选择、文件关联、上传下载或对应说明。
- 开发者中心截图数量为 2-5 张；默认使用真实运行的桌面端截图。移动端或智慧屏截图只有在用户明确要求对应平台并完成验证后才准备。
- 截图、图标、描述来自真实运行版本。
- 开发者中心资料完整，最终 `.lpk` 信息来自实际提审包。
- 迁移项目在开发者中心必须取消勾选“应用程序为本人原创开发或本人是源作者”，并填写原作者名称和源项目或代码地址。

## 执行规则

1. 确认项目来源：原创或迁移；读取对应前置结论。
2. 读取项目文件：`package.yml`、`lzc-manifest.yml`、`lzc-build.yml`、README、`docs/release-prep/`。
3. 同步或确认镜像，确保 manifest 引用最终可拉取镜像。
4. 构建 `.lpk`。
5. 检查包体大小和内嵌镜像产物。
6. 运行 `lzc-cli lpk info <file.lpk>` 并记录摘要。
7. 安装到懒猫微服验证启动、登录、核心流程、持久化、文件能力、卸载/升级风险。
8. 准备上架平台选择、2-5 张截图、图标、描述、测试账号、复现步骤、限制说明和 LPK 信息；默认只选择桌面端。
9. 桌面端默认上架时，确认 `package.yml.unsupported_platforms` 包含 `ios`、`tvos`、`android`。
10. 若项目来源为迁移，提交前确认开发者中心未勾选“应用程序为本人原创开发或本人是源作者”，并已填写原作者名称和源项目或代码地址。
11. 提交开发者中心；记录版本、时间、状态和证据。
12. 审核通过后验证商店可见性和安装版本；审核失败按问题归类修复。

## 审核失败处理

- `Packaging/runtime`: 只改 manifest、镜像、路由、健康检查、持久化、启动脚本。
- `Store asset/docs`: 只改描述、截图、图标、版本、测试账号、复现说明。
- `Business source required`: 迁移项目必须修改上游业务代码；停止普通上架，等待用户明确授权业务范围。

## 本地参考

- `references/docs/INDEX.md`
- `references/lpk/`

## 后置条件

- 已产出通过包体和内嵌镜像检查的最终 `.lpk`，或明确阻塞原因。
- 已记录实际 `.lpk` 信息。
- 如已安装或提审，已记录目标、版本、时间和验证证据。

## 输出格式

```text
Phase: Shipping
Project: <path/name>
Source: Original / Migration

Preflight
- Metadata:
- Source gates:
- Image:
- Store platforms:
- Screenshots:
- Store materials:
- Originality checkbox:
- Original author:
- Source project/code:
- Blockers:

Verification
- Build:
- LPK size:
- Embedded image check:
- LPK info:
- Install:
- Runtime:
- Core flow:

Submission
- Status:
- Evidence:
- Next:
```
