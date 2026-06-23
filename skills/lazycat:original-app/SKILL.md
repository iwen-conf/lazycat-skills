---
name: lazycat:original-app
description: "Create or align an original Lazycat app baseline only. Use for original apps, first-party projects, delivery files, account model, file capability, and handoff to listing. 原创应用、从零开发、原创项目基线、登录模型、文件能力、上架前准备."
---

# Lazycat Original App

## 职责范围

只处理原创应用或用户自有项目的懒猫工程基线。目标是让项目具备进入 `lazycat:ship-app` 的条件，不处理第三方开源迁移、许可证筛选、已安装应用更新或商店提审。

第三方 GitHub、自托管开源项目或搬运项目必须改用 `lazycat:migration-license`。

## 输入

- 项目想法、需求说明或已有原创项目路径。
- 目标平台：Web、移动端、桌面端、小程序或纯服务。
- 账号模型：无账号、内置账号、注册登录、OIDC 或其他。
- 文件能力：打开、保存、上传、下载、文件关联或无文件流程。
- 期望交付范围和可运行验证方式。

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

- 原创应用基线结论。
- 必要交付文件清单和实现状态。
- 登录、初始化、持久化、健康检查、文件能力方案。
- 构建或检查结果。
- 是否可交给 `lazycat:ship-app`。

## 前置条件

1. 已确认项目不是第三方迁移项目。
2. 已确认核心使用场景、目标用户和最小可用流程。
3. 如果包含前端或客户端，已确定平台类型；默认技术栈由 `arc:frontend` 决定，除非用户明确指定其他栈。

## 允许执行

- 创建或修改原创项目的业务代码、配置、文档和交付脚本。
- 创建或修改 `package.yml`、`lzc-manifest.yml`、`lzc-build.yml`、`build.sh`、`Makefile`。
- 为原创应用实现登录、初始化、持久化、健康检查和文件选择能力。
- 调用 `arc:frontend` 处理原创 Web、移动端、桌面端或小程序前端。
- 构建本地 `.lpk` 并执行包体和内嵌镜像检查。

## 禁止执行

- 不处理第三方项目许可证、查重或迁移可行性；改用 `lazycat:migration-license`。
- 不提交开发者中心、不发布商店、不处理审核状态；改用 `lazycat:ship-app`。
- 不通过 LPK Inspector 更新已安装应用；改用 `lazycat:update-installed-app`。
- 不写入真实微服域名、真实密码、API Key、Token 或内部未公开地址。
- 不使用 `lzc-build.yml.images`、`embed:<alias>`、包内 `images/` 或 `images.lock`。

## 执行规则

1. 建立最小需求记录：场景、用户、核心流程、账号模型、文件能力。
2. 建立或校准交付文件：`package.yml`、`lzc-manifest.yml`、`lzc-build.yml`、`build.sh`、`Makefile`。
3. 原创前端或客户端必须通过 `arc:frontend` 继承平台默认栈；用户指定例外时记录原因和边界。
4. 文件流程必须在原创业务 UI/代码中接入懒猫文件选择能力；不要使用迁移用 inject 兜底。
5. 需要镜像时使用可拉取远程镜像；最终 `.lpk` 不得内嵌镜像。
6. 构建后检查 `.lpk`：大小 `<= 12,000,000` bytes，归档内无 `images/` 和 `images.lock`。

## 后置条件

- `package.yml` 包含 `package`、`version`、`name`、`description`、`author`、`license`、`locales`。
- `version` 为严格 `x.x.x`。
- 项目有可执行构建入口和可描述的安装验证路径。
- 构建或检查失败时明确阻塞原因。
- 满足门禁后交给 `lazycat:ship-app`，不在本技能内提审。

## 输出格式

```text
Phase: Original App Baseline
Project: <path/name>

Inputs
- Scenario:
- Target platform:
- Account model:
- File capability:

Actions
- Delivery files:
- Frontend handoff:
- Runtime setup:
- Verification:

Decision: Ready for lazycat:ship-app / Blocked
Blockers:
```
