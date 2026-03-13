---
name: lazycat:create-app
description: 面向 Lazycat 新项目创建和项目基线统一的 skill。只要用户提到从 0 创建懒猫应用、初始化项目、脚手架、项目标准、Go 后端、Vue、Element Plus、登录、注册、JWT、access_token、refresh_token、无感刷新、认证改造、现金激励、对接微服账户系统、网盘右键菜单、需求分析文档、API 设计文档、把现有项目补成统一基线等请求，就必须使用此 skill。负责把项目创建阶段收敛到统一规范：第一步先建立 `docs/` 文档树，再落 Go + Vue + Element Plus，默认具备登录、注册、双 token 和无感刷新，并在用户以现金激励为目标时优先满足官方激励门槛。
---

# Lazycat 项目创建基线

你负责把 Lazycat 项目从“一个想法”推进到“有统一技术基线、可继续开发、可进入资料准备与发布流程”的状态。重点不是只生成目录，而是把项目标准、认证基线、激励资格路径和后续发布兼容性一次性定清楚。

## Overview

这个 skill 用于新建项目或给现有项目补齐统一基线。默认标准是：

- 第一步先创建 `docs/` 文档目录，并拆分需求分析、API 设计等多个子目录与多个 `md` 文档
- 后端使用 Go
- 前端使用 Vue
- UI 使用 Element Plus
- 所有项目默认具备登录、注册
- 认证采用 `access_token + refresh_token`
- 前端默认支持无感刷新

如果现有仓库已经有更强的团队约束，优先沿用团队约束；如果没有，就按这个默认标准落地。

如果用户目标包含懒猫现金激励，还要额外按当前官方规则收敛：

- 优先做原创且有真实场景的应用
- 避开官方明确不发红包的类型
- 需要账号密码的应用必须让普通用户可获得凭证
- 能对接微服账户系统或网盘文件关联时，优先做掉

## Quick Contract

- **Trigger**: 用户提到创建懒猫项目、初始化脚手架、统一项目标准、补登录注册、接入双 token、做无感刷新、Go + Vue + Element Plus 基线、现金激励、微服账户系统、网盘右键菜单
- **Inputs**: 项目目标、仓库现状、是否新建项目、现有前后端结构、认证现状、是否已有用户系统、是否以现金激励为目标
- **Outputs**: 文档目录蓝图、项目基线摘要、技术栈与目录建议、认证方案、激励资格路径、必须完成的模块清单、后续交给 `lazycat:ship-app` 的发布准备接口
- **Quality Gate**: 最终方案必须先定义 `docs/` 目录与文档拆分，再明确 Go 后端、Vue + Element Plus 前端、登录/注册、`access_token + refresh_token`、无感刷新，以及在激励模式下的账号可获得性、应用类型合规性和附加对接机会
- **Decision Tree**: 先判断是新建项目还是补齐现有项目，再判断是全套基线落地、只补认证、还是只做结构收敛

## When to Use

**首选触发**

- 用户要从 0 创建 Lazycat 项目
- 用户明确要求项目统一使用 Go + Vue + Element Plus
- 用户要求所有项目都必须具备登录、注册、双 token、无感刷新
- 用户希望应用尽量符合懒猫现金激励门槛
- 现有项目要补齐统一认证能力或前后端基线

**典型场景**

- 新建一个面向 Lazycat 上架的新应用，需要先把技术基线定下来
- 已有项目功能已经有了，但认证链路混乱，需要统一成登录/注册 + 双 token
- 前端已经有 Vue 页面，但没有 Element Plus、Pinia、路由守卫和无感刷新
- 后端已有 Go 服务，但没有 refresh token、token rotation 或统一认证接口
- 工具类应用希望额外对接网盘右键菜单
- 应用希望对接微服账户系统，提高可用性并争取额外激励

**边界提示**

- 这个 skill 负责“创建和基线”，不是最终上架；准备提审时要回到 `lazycat:ship-app`
- 如果用户只要图标或商店资料，不要误用本 skill
- 如果现有项目已经采用别的语言栈，除非用户明确允许，不要擅自跨语言重写

## Announce

开始执行后，用一句短摘要说明：

- 当前是新建项目还是补齐现有项目
- 你会先创建或补齐哪些 `docs/` 子目录
- 你将如何确认 Go / Vue / Element Plus 是否已经落地
- 当前认证链路缺的是页面、接口、token 机制，还是无感刷新
- 如果目标是现金激励，你优先在检查哪条资格路径

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `project_mode` | enum(`新建项目`/`补齐现有项目`/`仅认证改造`) | 推荐 | 决定本次是从零初始化，还是基于现有仓库做结构补强 |
| `backend_stack` | string | 推荐 | 默认为 Go；除非用户明确要求，不要偏离 |
| `frontend_stack` | string | 推荐 | 默认为 Vue + Element Plus；如果已有 Vue 项目，优先补齐而不是推倒重来 |
| `auth_state` | enum(`无认证`/`仅登录`/`单 token`/`双 token 不完整`/`已完整`) | 推荐 | 用于判断认证改造深度 |
| `user_system_state` | enum(`无用户表`/`已有用户表`/`已有第三方登录`) | 可选 | 用于决定注册流程和用户资料最小字段 |
| `release_target` | enum(`先开发`/`准备提审`/`同时为发布做基线`) | 可选 | 决定是否同步补齐面向发布的元数据和配置约束 |
| `reward_target` | enum(`普通上架`/`现金激励优先`) | 可选 | 如果目标是红包激励，要按官方规则避开不奖励类型，并优先做账户系统或文件关联对接 |
| `docs_state` | enum(`无 docs`/`docs 不完整`/`docs 已齐全`) | 可选 | 判断是否需要在编码前先创建 `docs/requirements`、`docs/api-design` 等目录与文档 |

## The Iron Law

1. 所有新项目默认使用 Go 后端、Vue 前端、Element Plus 组件库；不要在没有证据的情况下切换到别的栈。
2. 所有项目默认具备登录、注册；如果业务确实不需要开放注册，要在汇报里明确这是业务豁免，而不是遗漏。
3. 认证默认使用 `access_token + refresh_token`，不能只发一个长生命周期 token 就算完成。
4. 前端必须具备无感刷新与失败回退机制；不能把“token 过期后手动重登”当作完成态。
5. 创建阶段就要考虑后续上架与提审，避免后面再返工认证、菜单、路由、初始化或环境变量。
6. 如果以现金激励为目标，优先做原创、真实场景、稳定可用的应用；不要把明显不在奖励范围内的应用类型当主目标。
7. 如果应用需要账号密码，必须确保普通用户能自助注册、能通过微服账户系统登录，或能获得公开可用的测试凭证；否则会影响上架。
8. 进入编码前，必须先建立 `docs/` 文档树；不要一边想需求一边写业务，把需求分析和 API 设计留到后补。

## Workflow

### 1. 确认项目创建模式

- 判断是全新项目、现有项目补齐，还是只补认证能力
- 判断前后端是否已经存在，以及当前结构是否接近 Go + Vue + Element Plus 基线
- 如果用户只给了想法，没有仓库，先生成最小项目蓝图

### 2. 先创建 docs 文档树

编码前先检查并建立 `docs/`。至少要有这些子目录，并且每个目录下拆成多个 `md`：

- `docs/requirements/`
- `docs/api-design/`
- `docs/architecture/`
- `docs/release-prep/`

推荐最小文档集：

- `docs/requirements/product-overview.md`
- `docs/requirements/user-stories.md`
- `docs/requirements/scope-and-milestones.md`
- `docs/api-design/overview.md`
- `docs/api-design/auth.md`
- `docs/api-design/domain-modules.md`
- `docs/architecture/system-overview.md`
- `docs/architecture/frontend.md`
- `docs/architecture/backend.md`
- `docs/release-prep/store-assets.md`
- `docs/release-prep/test-plan.md`
- `docs/release-prep/submission-checklist.md`

如果仓库里已经有文档，不要机械重建；先整理成这个结构或映射到等价结构。

### 3. 确认激励路径

- 如果目标包含现金激励，优先判断该应用是不是原创应用，因为原创应用通常有更高的激励上限
- 排除官方明确不发红包的类型，例如纯网页游戏、纯书籍页面、纯教程网站、网页离线应用、游戏服务器不同模组、纯数据库软件
- 如果是移植应用，要预留上游地址，并在发布前检查商店是否已存在同款，避免落到非首发激励路径
- 如果是工具类应用，优先评估是否可对接网盘文件关联
- 如果应用已有账户体系，优先评估是否可接入微服 OIDC

### 4. 固化技术基线

- 后端固定为 Go
- 前端固定为 Vue
- UI 固定为 Element Plus
- 前端状态管理和鉴权状态建议统一到单一 store
- 路由层必须预留匿名页和鉴权页边界

如果仓库已有更细的框架约束，以仓库现状为准，但不能偏离这三项主基线。

### 5. 落认证基线

每个项目都要至少具备：

- 登录页
- 注册页
- 鉴权状态持久化与恢复
- `access_token`
- `refresh_token`
- 无感刷新
- 刷新失败后的清理与回登录页

推荐默认做法：

- `access_token` 走短生命周期
- `refresh_token` 走长生命周期
- 刷新时执行 token rotation
- 前端只允许一个刷新流程在飞，其他请求排队等待结果
- 刷新失败后清理本地鉴权状态并跳转登录
- 如果目标是现金激励优先，尽量额外提供微服 OIDC 登录入口，或至少保证普通用户可以自助注册

### 6. 设计最小接口与页面

至少明确这些接口或等价能力：

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `POST /auth/logout`
- `GET /auth/me`

至少明确这些前端能力：

- 登录表单
- 注册表单
- 路由守卫
- 启动时恢复登录态
- 401 后单次刷新重试
- 刷新失败后的退出登录
- 如果接入微服账户系统，提供 OIDC 登录入口和回调处理逻辑
- 如果是工具类应用，提供文件打开入口以承接 `file_handler`

### 7. 对齐数据、安全边界与附加对接

- 用户表最小字段要能支撑注册、登录和状态判断
- refresh token 不能裸存明文长期复用，至少要具备可撤销或可轮换能力
- 认证相关环境变量、密钥、过期时间、Cookie/Header 约束要提前定下来
- 如果项目要上架 Lazycat，初始化说明里要明确首次登录、默认管理员或注册入口策略
- 如果接入微服 OIDC，要在 manifest 中规划 `application.oidc_redirect_path`，并把系统注入的 OIDC 环境变量正确透传给应用
- 如果是工具类应用，要在 manifest 中规划 `application.file_handler`、支持的 `mime`，以及应用侧 `/open` 或等价路由

### 8. 交给发布链路

当项目已经具备可开发、可登录、可注册、可刷新 token 的基线后，再交回 `lazycat:ship-app` 继续走资料、截图、提审和发布。

复杂基线收敛先读 [references/project-baseline.md](./references/project-baseline.md) 和 [references/docs-blueprint.md](./references/docs-blueprint.md)。

## Quality Gates

- 明确已经建立 `docs/` 文档树，且需求分析、API 设计等目录下不是空目录
- 明确采用 Go + Vue + Element Plus
- 明确具备登录、注册
- 明确具备 `access_token + refresh_token`
- 明确前端无感刷新流程和失败回退流程
- 明确最小接口、最小页面和最小用户数据要求
- 如果目标是现金激励优先，明确应用不属于不奖励类型
- 如果应用需要账号密码，明确普通用户如何获得凭证
- 若应用适合，明确是否对接微服 OIDC 或网盘文件关联
- 说明创建完成后如何进入 `lazycat:ship-app`

## Red Flags

- 项目已经开始做业务，但还没有统一认证基线
- 只做登录，不做注册，却没有业务豁免说明
- 只发一个 token，没有 refresh token
- 前端请求 401 后直接把用户踢掉，没有无感刷新
- 多个并发请求会重复触发 refresh，导致竞态或覆盖
- 项目已经开始写代码，但 `docs/requirements` 和 `docs/api-design` 还不存在
- 只有一个笼统的 `README.md`，没有拆分需求分析和 API 设计文档
- 目标是拿激励，却选了官方明确不发红包的应用类型
- 应用需要登录，但普通用户拿不到账号或无法自助注册
- 明明是工具类应用，却没有评估文件关联；明明适合账户统一，却没有评估 OIDC
- 创建阶段没有考虑后续上架和初始化路径

## Bundled References

- 项目技术栈与认证基线： [references/project-baseline.md](./references/project-baseline.md)
- 文档目录蓝图： [references/docs-blueprint.md](./references/docs-blueprint.md)
- 激励资格与对接机会： [../lazycat:ship-app/references/cash-incentive.md](../lazycat:ship-app/references/cash-incentive.md)

## Outputs

```text
阶段: 项目创建 / 基线补齐
项目模式: <新建项目 / 补齐现有项目 / 仅认证改造>

已确认
- ...

文档目录
- docs/requirements
- docs/api-design
- docs/architecture
- docs/release-prep

项目基线
- 后端: Go
- 前端: Vue + Element Plus
- 认证: access_token + refresh_token + 无感刷新

激励路径
- 目标: <普通上架 / 现金激励优先>
- 类型判断: <原创 / 移植 / 不建议走激励>
- 附加对接: <OIDC / file_handler / 无>

缺口 / 风险
- ...

当前执行
- ...

下一步
1. ...
2. ...

交付物
- 项目结构建议
- 认证链路清单
- 激励资格路径
- 交给 lazycat:ship-app 的后续入口
```
