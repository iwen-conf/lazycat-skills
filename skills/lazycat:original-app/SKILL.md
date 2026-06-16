---
name: lazycat:original-app
description: "Build or align an original Lazycat app. Use for 原创应用、从零开发、原创项目基线、登录注册、文件选择、上架前准备."
---

# Lazycat Original App

本技能只处理原创应用。目标是把一个想法或已有原创项目推进到“可以进入上架准备”的工程状态，而不是做迁移、许可证判断或第三方项目包装。

## 使用场景

- 用户要从零创建一个懒猫应用。
- 用户强调“原创”“自己开发”“不是搬运 GitHub 项目”。
- 已有原创项目需要补齐懒猫应用基线、构建入口、登录、文件能力或上架前结构。

如果输入是 GitHub 第三方项目、开源项目或自托管项目，改用迁移三关：`lazycat:migration-license` -> `lazycat:migration-boundary` -> `lazycat:migration-workload`。

## 强约束

1. 原创应用必须先明确真实使用场景和核心用户，不写空泛营销页。
2. 必须建立可执行交付入口：`package.yml`、`lzc-manifest.yml`、`lzc-build.yml`、`build.sh`、`Makefile`。
3. 默认需要登录/注册或明确说明“不需要账号”的业务原因。
4. 有文件打开、保存、上传、下载流程时，原创应用必须在业务 UI/代码中接入懒猫文件选择能力；不要把原创应用的文件能力交给迁移用 inject 兜底。
5. 原创应用包含 Web、移动端、桌面端或小程序前端时，除非用户明确指定其他技术，必须联动 `arc:frontend` 并使用对应平台默认栈。
6. 不在 Lazycat 技能内另起一套前端或跨端默认选型；业务是文件工具、Agent、Dashboard、管理台、内容应用、移动客户端或桌面客户端都不改变平台默认栈。
7. 不引入无关技能主题：不单独展开高级路由、图标、后台 UI、攻略、更新、排障等旧入口；这些都只作为当前项目内的必要实现细节。
8. 交付设计必须保证最终 `.lpk` 小于或等于 12 MB（`12,000,000` bytes），且不得内嵌镜像。
9. 需要镜像时使用可拉取远程镜像地址；不得使用 `lzc-build.yml.images`、`embed:<alias>`、包内 `images/` 或 `images.lock`。
10. 不写真实微服域名、真实密码、API Key、Token 或内部未公开地址。

## 工作流

1. 确认这是原创应用，并记录目标用户、核心流程、是否需要账号、是否有文件能力。
2. 建立最小文档树：`docs/requirements/`、`docs/architecture/`、`docs/release-prep/`。
3. 建立懒猫交付文件：`package.yml`、`lzc-manifest.yml`、`lzc-build.yml`、`build.sh`、`Makefile`。
4. 如果包含 Web、移动端、桌面端或小程序前端，调用 `arc:frontend` 确认平台默认栈、页面、状态分层、路由、表单和 token 方案；只有用户明确指定时才记录例外栈。
5. 实现或校准核心业务入口，不做无关页面和装饰性内容。
6. 补齐登录/注册、初始化数据、持久化目录、健康检查和文件选择能力。
7. 本地构建并检查 `.lpk` 包体 `<= 12,000,000` bytes 且无内嵌镜像产物；满足后交给 `lazycat:ship-app`。

## 质量门禁

- 项目能解释“为什么应该安装在懒猫微服里”。
- 构建、安装、启动、核心流程都有可执行路径。
- Web、移动端、桌面端或小程序前端使用 `arc:frontend` 平台默认栈，或明确记录用户指定的例外技术和边界。
- `package.yml` 的 `package`、`version`、`name`、`description`、`author`、`license`、`locales` 完整。
- `version` 使用严格 `x.x.x` 格式。
- 没有真实设备名、真实凭据或敏感地址。
- 最终 `.lpk` 包体 `<= 12,000,000` bytes，且不包含内嵌镜像。
- 交付前至少运行一次项目内可用的构建或检查命令；无法运行时说明阻塞原因。

## 输出格式

```text
Phase: Original App
Project: <name/path>

Confirmed
- Scenario:
- Account model:
- File capability:

Implemented / Planned
- Delivery files:
- Core flow:
- Verification:

Ready for Shipping
- Yes / No
- Blockers:
```
