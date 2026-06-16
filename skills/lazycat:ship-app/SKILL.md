---
name: lazycat:ship-app
description: "List a ready Lazycat app: build LPK, install-test, prepare metadata/screenshots, submit, and verify release. 上架、打包、提审、发布验证."
---

# Lazycat App Listing

本技能只处理上架。它接收已经准备好的原创应用，或已经通过迁移三关的迁移项目，推进到打包、安装验证、提审和发布后检查。

## 使用场景

- 用户要打包 `.lpk`、安装测试、准备商店资料、提交开发者中心或跟进审核。
- 原创项目已由 `lazycat:original-app` 补齐工程基线。
- 迁移项目已依次通过：
  1. `lazycat:migration-license`
  2. `lazycat:migration-boundary`
  3. `lazycat:migration-workload`

## 强约束

1. 不接收许可证不明、不可商用或不能再分发的迁移项目。
2. 迁移项目继承“默认禁止修改业务代码”：除非用户在当前任务明确说明“允许/需要修改业务代码”并点名范围，否则上架、审核修复和安装后排障都不得修改上游业务代码。
3. 不接收需要修改上游业务代码才能完成普通迁移的项目。
4. 不把截图、网页表单或本地开发环境当作真实验证；必须安装到懒猫微服并验证核心流程，除非明确说明环境阻塞。
5. 上架信息以本地项目文件为源头：`package.yml`、`lzc-manifest.yml`、`lzc-build.yml`、`README`、`docs/release-prep/`。
6. 镜像型项目必须先完成 `copy-image` 并把返回的 `registry.lazycat.cloud/...` 写回 manifest，再构建 `.lpk`。
7. 所有构建产出的 `.lpk` 必须小于或等于 12 MB；按 `12,000,000` bytes 检查，超过即停止交付。
8. 禁止内嵌镜像：不得使用 `lzc-build.yml.images`、`embed:<alias>`，最终 `.lpk` 不得包含 `images/` 或 `images.lock`。
9. 开发者中心上架资料必须填写完整，不得留下空字段、占位文案或待补信息；提交前必须包含最终 `.lpk` 的 LPK 信息，信息来源必须是实际提审包。
10. 提交资料不得包含真实密码、Token、内部地址或真实微服域名；测试账号只能使用明确可公开给审核的凭据。
11. 色情、赌博、毒品、空投、破解软件、违法内容，直接拒绝上架。
12. 原创或新增独立 Web、移动端、桌面端、小程序前端必须继承 `arc:frontend` 平台默认栈，除非用户明确指定其他技术；迁移项目仍不得为了默认栈改上游业务前端/客户端。
13. 不再分流到旧的 UI、图标、攻略、更新、排障等独立技能；这些都是当前上架任务内的必要检查项。

## 上架前检查

- `package.yml` 包含 `package`、`version`、`name`、`description`、`author`、`license`、`locales`。
- 如果 `package.yml.homepage` 是 GitHub 仓库 URL，`package.yml.author` 必须逐字符等于 URL 中的 owner 段，大小写和符号都不能变化；例如 `https://github.com/Sliverkiss/mimocode2api` 只能写 `author: Sliverkiss`，否则会被审核打回。
- `version` 是严格 `x.x.x`。
- `lzc-manifest.yml` 的服务、路由、持久化、健康检查、环境变量完整。
- `build.sh` 和 `Makefile` 可执行，至少包含 `build`、`install`、`verify` 或等价目标。
- `Makefile` 或等价构建流程包含 `.lpk` 大小检查和内嵌镜像检查；包体必须 `<= 12,000,000` bytes。
- `lzc-build.yml` 不包含顶层 `images`；`lzc-manifest.yml` 不包含 `embed:<alias>`；最终 `.lpk` 不包含 `images/` 或 `images.lock`。
- 迁移项目保留上游地址、应用商店查重结论、开发者中心待审列表查重结论、许可证结论、非侵入边界结论和工作量结论。
- 有账号的应用提供注册、OIDC、公开测试账号或明确初始化说明。
- 有文件能力的应用完成文件选择或文件关联验证。
- 原创前端或新增独立前端面记录 `arc:frontend` 平台默认栈验证；迁移项目记录上游前端/客户端保持不动。
- 图标、截图、描述来自真实运行版本，不使用模板占位或调试数据。
- 开发者中心所需资料全部填写完成；最终 `.lpk` 已执行 `lzc-cli lpk info <file.lpk>` 或等价检查，并把 LPK 信息随提审资料记录。

## 工作流

1. 确认项目来源：原创或迁移；迁移必须读取三关结论。
2. 读取本地项目文件，修正 metadata、版本、作者、许可证、locales 和 reviewer instructions。
3. 构建或同步镜像，确保 manifest 中使用最终可拉取镜像。
4. 运行 `make build` 或项目等价命令生成 `.lpk`。
5. 检查 `.lpk` 包体大小 `<= 12,000,000` bytes，并确认归档内没有 `images/` 或 `images.lock`。
6. 运行 `lzc-cli lpk info <lpk>` 或等价命令记录最终提审包的 LPK 信息。
7. 运行 `make install` 或 `lzc-cli app install <lpk>` 安装到懒猫微服。
8. 验证启动、登录、核心流程、持久化、文件能力、卸载/升级风险。
9. 准备截图、图标、描述、测试账号、复现步骤、限制说明，并确认开发者中心所有字段已填写完成。
10. 提交开发者中心；记录版本、LPK 信息、时间、状态、截图或页面证据。
11. 审核通过后验证商店可见性和安装版本；审核失败则按问题归类修复。

## 审核失败归类

- `Packaging/runtime`: manifest、镜像、路由、健康检查、持久化、启动脚本问题。
- `Store asset/docs`: 描述、截图、图标、版本、测试账号、复现说明问题。
- `Business source required`: 迁移项目必须改上游业务代码才能修复。此类问题停止普通迁移；只有用户在当前任务明确说明“允许/需要修改业务代码”并点名业务范围后，才可进入业务代码修改。

## 本地参考

- 官方文档索引：`references/docs/INDEX.md`
- LPK/manifest/build/package 规范：`references/lpk/`

## 输出格式

```text
Phase: Listing
Project: <name/path>
Source: Original / Migration

Preflight
- Metadata:
- Build entries:
- Migration gates:
- Frontend stack boundary:
- Blockers:

Verification
- Build:
- LPK size/embed check:
- LPK info:
- Install:
- Core flow:
- Screenshots/assets:

Submission
- Status:
- Evidence:
- Next action:
```
