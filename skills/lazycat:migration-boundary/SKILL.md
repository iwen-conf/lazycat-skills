---
name: lazycat:migration-boundary
description: "Second gate for migrating a GitHub project to Lazycat: decide whether it can be migrated without modifying upstream business code. 迁移第二关、不改业务代码、非侵入迁移."
---

# Lazycat Migration Boundary Gate

本技能是迁移第二关，只判断一个已通过许可证筛选的项目，能否在不修改上游业务代码的前提下迁移到懒猫。

## 使用场景

- 项目已通过 `lazycat:migration-license`。
- 用户问“能不能不改业务代码迁移”“能不能只包装一下”“能不能不动上游”。
- 需要在开始打包前判断是否会被登录、路径、数据库、初始化、端口、文件能力卡住。

## 强约束

1. 迁移默认零业务代码修改。除非用户在当前任务明确说明“允许/需要修改业务代码”并点名允许范围，否则一律禁止修改业务代码。
2. 只允许修改包装层和运行时适配层：
   - `package.yml`
   - `lzc-build.yml`
   - `lzc-manifest.yml`
   - `lzc-deploy-params.yml`
   - `Makefile`
   - `build.sh`
   - Docker wrapper 文件
   - 启动脚本、seed/setup 脚本、配置模板
   - 图标、商店素材、迁移文档
3. 禁止修改上游前端页面/组件/路由/状态、后端 handler/service/domain/auth、数据库 schema/migration/model、测试、fixture。
4. 如果唯一可行方案必须改业务代码，立即输出 `Blocked by business-code change requirement`，不要继续包装。
5. 用户只说“修好”“跑起来”“尽快上架”“可以处理问题”不构成业务代码修改授权；必须按迁移阻塞处理。
6. 登录、初始化、路径、健康检查、文件选择必须优先用环境变量、命令行参数、配置模板、启动脚本、seed 服务、OIDC、inject 或反向代理解决。
7. 文件打开/保存/上传/下载流程属于迁移能力门禁；迁移项目用 `application.injects` 接入文件选择，不改上游 UI。
8. 默认前端/跨端栈不适用于迁移项目的上游业务代码；不得为了统一到 React Web、React Native + Expo、Tauri 2 或 Taro 4 而重写上游前端、移动端、桌面端或小程序客户端。
9. 只有新增包装层、配置向导、审核辅助页、独立管理台或原创配套客户端等非上游业务前端时，才联动 `arc:frontend` 使用平台默认栈。

## 必查项

- 是否有官方 Docker image、Dockerfile 或 Compose。
- 进程是否长期运行，是否监听固定端口。
- 数据和配置是否能映射到 `/lzcapp/var`。
- 依赖服务是否能用健康检查和 `depends_on.condition: service_healthy` 排序。
- 首次初始化是否可用 CLI、环境变量、配置文件、seed 服务完成。
- 登录是否可以 OIDC、固定初始凭据、inject 或反向代理处理。
- 文件能力是否可用 inject 非侵入接入。
- 是否依赖特权容器、宿主机服务、外部 SaaS、不可分发二进制或手工安装步骤。
- 是否存在新增非业务包装前端面；如存在，交给 `arc:frontend`，不得混入上游业务前端/客户端修改。

## 决策

- `Can migrate non-invasively`: 可以只改包装层完成迁移。
- `Can migrate with wrapper risk`: 大体可行，但有启动、登录、文件或依赖风险，需要在工作量评估中放大。
- `Cannot determine yet`: 关键事实缺失，需要先运行或补充文档。
- `Blocked by business-code change requirement`: 必须改上游业务代码，不进入普通迁移。

## 输出格式

```text
Phase: Migration Boundary Gate
Repository: <owner/repo>

Decision: Can migrate non-invasively / Can migrate with wrapper risk / Cannot determine yet / Blocked by business-code change requirement

Allowed Write Scope
- ...

Runtime Model
- Delivery:
- Entry:
- Persistence:
- Dependencies:
- Initialization:
- Login:
- File flows:

Business-Code Risks
- ...

Frontend Stack Boundary
- Upstream frontend/client preserved:
- New wrapper/admin/client surface:
- arc:frontend handoff:

Next
- Proceed to lazycat:migration-workload / Stop
```
