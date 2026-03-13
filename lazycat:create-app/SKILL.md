---
name: lazycat:create-app
description: 面向 Lazycat 新项目创建和项目基线统一的 skill。只要用户提到从 0 创建懒猫应用、初始化项目、脚手架、项目标准、Go 后端、Vue、Element Plus、后台管理、Admin、管理台、登录、注册、JWT、access_token、refresh_token、无感刷新、认证改造、现金激励、对接微服账户系统、网盘右键菜单、需求分析文档、API 设计文档、原创应用如何融入 Lazycat 原生系统、AI 配置面板、模型配置、懒猫算力仓、AI 应用、AI 浏览器插件等请求，就必须使用此 skill。负责把项目创建阶段收敛到统一规范：第一步先建立 `docs/` 文档树，再落 Go + Vue + Element Plus，默认具备登录、注册、双 token 和无感刷新；普通业务型 Web 应用接 AI 时默认走 `BaseURL` 配置方案，只有明确做懒猫算力仓 / `AI应用` / AI 浏览器插件时才走官方 AI Pod 路线；若项目包含后台管理面，则进一步把它接到高质量管理 UI 质量链路，并在用户以现金激励为目标时优先满足官方激励门槛。
---

# Lazycat 项目创建基线

你负责把 Lazycat 项目从“一个想法”推进到“有统一技术基线、可继续开发、可进入资料准备与发布流程”的状态。重点不是只生成目录，而是把项目标准、认证基线、原创应用与 Lazycat 的融合方式、AI 接入基线、激励资格路径和后续发布兼容性一次性定清楚。

## Overview

这个 skill 用于新建项目或给现有项目补齐统一基线。默认标准是：

- 第一步先创建 `docs/` 文档目录，并拆分需求分析、API 设计等多个子目录与多个 `md` 文档
- 提供可执行的脚本入口，例如 `build.sh` 和 `Makefile`
- 后端使用 Go
- 前端使用 Vue
- UI 使用 Element Plus
- 如果项目包含后台管理面，默认要求高质量管理台 UI，可使用成熟模板打底但必须完成业务化改造
- 所有项目默认具备登录、注册
- 认证采用 `access_token + refresh_token`
- 前端默认支持无感刷新
- 原创应用默认需要评估怎样接入 Lazycat 原生能力，不能只是一个孤立网页
- 如果业务天然适合 AI，默认预留统一 AI 配置页，至少包含 `API BaseURL`、协议、模型拉取、模型选择和保存配置
- 如果这是普通业务型 Web 应用，AI 默认只是业务能力的一部分，不要自动切到 AI Pod 路线
- 只有当用户明确目标是懒猫算力仓、`AI应用` 或 AI 浏览器插件时，才额外评估官方 AI Pod 路线

如果现有仓库已经有更强的团队约束，优先沿用团队约束；如果没有，就按这个默认标准落地。

如果用户目标包含懒猫现金激励，还要额外按当前官方规则收敛：

- 优先做原创且有真实场景的应用
- 原创不只看“是否有人发过”，还要明确它为什么适合长在 Lazycat 原生环境里
- 避开官方明确不发红包的类型
- 需要账号密码的应用必须让普通用户可获得凭证
- 能对接微服账户系统或网盘文件关联时，优先做掉

## Quick Contract

- **Trigger**: 用户提到创建懒猫项目、初始化脚手架、统一项目标准、补登录注册、接入双 token、做无感刷新、Go + Vue + Element Plus 基线、后台管理、Admin、现金激励、微服账户系统、网盘右键菜单、原创应用怎么更像 Lazycat 原生应用、AI 配置页、模型配置、懒猫算力仓、AI 应用、AI 浏览器插件
- **Inputs**: 项目目标、仓库现状、是否新建项目、现有前后端结构、认证现状、是否已有用户系统、是否存在后台管理面、是否以现金激励为目标、是否原创应用、当前准备接哪些 Lazycat 原生能力、是否适合引入 AI、是否适合接入懒猫算力仓 / AI应用
- **Outputs**: 文档目录蓝图、命令入口蓝图、项目基线摘要、技术栈与目录建议、认证方案、原生融合策略、AI 配置基线、算力仓 / AI应用接入判断、后台管理 UI 基线、激励资格路径、必须完成的模块清单、后续交给 `lazycat:ship-app` 的发布准备接口
- **Quality Gate**: 最终方案必须先定义 `docs/` 目录与文档拆分、`build.sh` 与 `Makefile` 的最小命令入口，再明确 Go 后端、Vue + Element Plus 前端、登录/注册、`access_token + refresh_token`、无感刷新；如果项目是原创应用，还必须明确 Lazycat 原生能力结合面；如果项目适合 AI，还必须明确统一 AI 配置面板；只有当目标明确是懒猫算力仓 / `AI应用` / AI 浏览器插件时，才补官方 AI Pod 路线判断；如果项目包含后台管理面，还必须明确高质量管理台 UI 路线、模板改造策略和后续交给 `lazycat:admin-ui` 的收敛方式；在激励模式下还要补齐账号可获得性、应用类型合规性和附加对接机会
- **Decision Tree**: 先判断是新建项目还是补齐现有项目，再判断是全套基线落地、只补认证、只做结构收敛，还是还需要进入原生融合、AI 配置面、明确的算力仓 / AI应用链路或后台管理 UI 质量链路

## When to Use

**首选触发**

- 用户要从 0 创建 Lazycat 项目
- 用户明确要求项目统一使用 Go + Vue + Element Plus
- 用户要求管理后台、控制台或 Admin 部分也纳入统一基线
- 用户要求所有项目都必须具备登录、注册、双 token、无感刷新
- 用户希望应用尽量符合懒猫现金激励门槛
- 用户要求原创应用更像 Lazycat 原生应用，而不是单纯套壳网页
- 用户要求项目适当接入 AI，并把模型配置标准化
- 用户希望项目对接懒猫算力仓、AI 应用体系或 AI 浏览器插件
- 现有项目要补齐统一认证能力或前后端基线

**典型场景**

- 新建一个面向 Lazycat 上架的新应用，需要先把技术基线定下来
- 已有项目功能已经有了，但认证链路混乱，需要统一成登录/注册 + 双 token
- 前端已经有 Vue 页面，但没有 Element Plus、Pinia、路由守卫和无感刷新
- 后端已有 Go 服务，但没有 refresh token、token rotation 或统一认证接口
- 项目已经有后台管理功能，但界面还没有形成高质量管理台标准
- 工具类应用希望额外对接网盘右键菜单
- 应用希望对接微服账户系统，提高可用性并争取额外激励
- 团队打算做原创应用，但还没回答“为什么它更应该装在 Lazycat 里”
- 项目准备加入 AI 能力，但还没有统一的 `BaseURL / 协议 / 模型` 配置方式
- 项目本身是 AI 产品，但还没判断应该做成普通应用、懒猫 `AI应用`，还是 AI 浏览器插件

**边界提示**

- 这个 skill 负责“创建和基线”，不是最终上架；准备提审时要回到 `lazycat:ship-app`
- 如果用户只要图标或商店资料，不要误用本 skill
- 如果目标只是把现有后台界面做得更高级、更适合截图，优先切到 `lazycat:admin-ui`
- 如果现有项目已经采用别的语言栈，除非用户明确允许，不要擅自跨语言重写

## Announce

开始执行后，用一句短摘要说明：

- 当前是新建项目还是补齐现有项目
- 你会先创建或补齐哪些 `docs/` 子目录
- 你将如何确认 Go / Vue / Element Plus 是否已经落地
- 当前认证链路缺的是页面、接口、token 机制，还是无感刷新
- 当前原创应用准备接哪几个 Lazycat 原生能力
- 当前是否需要统一 AI 配置面板，如果需要，先落哪几个字段
- 当前 AI 相关能力更适合普通应用、懒猫 `AI应用`，还是 AI 浏览器插件
- 如果存在后台管理面，你准备如何收敛后台质量和模板策略
- 如果目标是现金激励，你优先在检查哪条资格路径

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `project_mode` | enum(`新建项目`/`补齐现有项目`/`仅认证改造`) | 推荐 | 决定本次是从零初始化，还是基于现有仓库做结构补强 |
| `backend_stack` | string | 推荐 | 默认为 Go；除非用户明确要求，不要偏离 |
| `frontend_stack` | string | 推荐 | 默认为 Vue + Element Plus；如果已有 Vue 项目，优先补齐而不是推倒重来 |
| `auth_state` | enum(`无认证`/`仅登录`/`单 token`/`双 token 不完整`/`已完整`) | 推荐 | 用于判断认证改造深度 |
| `user_system_state` | enum(`无用户表`/`已有用户表`/`已有第三方登录`) | 可选 | 用于决定注册流程和用户资料最小字段 |
| `admin_surface_state` | enum(`无后台`/`有后台待设计`/`已有后台待升级`) | 可选 | 判断是否需要把后台管理面接入高质量 UI 质量链路 |
| `product_origin` | enum(`原创应用`/`移植应用`/`混合`) | 推荐 | 决定是否必须回答“如何融入 Lazycat 原生系统” |
| `native_fit_state` | enum(`未评估`/`仅 OIDC`/`仅文件关联`/`已有多项融合点`) | 可选 | 判断项目当前是否已经形成 Lazycat 原生能力结合面 |
| `ai_surface_state` | enum(`无 AI`/`有 AI 待设计`/`已有 AI 待统一`/`AI 已完整`) | 可选 | 判断是否需要补统一 AI 配置面板与模型发现流程 |
| `aipod_fit_state` | enum(`未评估`/`普通应用更合适`/`AI应用候选`/`AI 浏览器插件候选`/`已有算力仓接入`) | 可选 | 判断项目是否应该接入懒猫算力仓 / `AI应用` 生态 |
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
9. 项目必须提供可执行的命令入口；最少要有 `build.sh` 和 `Makefile`，不能把构建步骤散落在聊天记录里。
10. 如果项目包含后台管理、控制台或运营工作台，必须把它接入 `lazycat:admin-ui` 的高质量 UI 路径；可以用模板打底，但不能原样交付模板站。
11. 原创应用不能只满足“商店里没人发过”；必须说明它怎样和 Lazycat 原生能力结合，至少评估微服 OIDC、文件关联和安装后即用的本地工作流。
12. 只要项目存在明确 AI 场景，就优先提供统一 AI 配置面板；最小项固定为 `API BaseURL`、协议类型（`OpenAI Compatible` / `OpenAI Responses` / `Anthropic`）、获取模型按钮、模型下拉框、保存配置按钮。
13. 普通业务型 Web 应用接 AI 时，默认走 `BaseURL` 配置方案，不要自动要求 `ai-pod-service`、`caddy-aipod` 或 `extension.zip`。
14. 只有当用户明确目标是懒猫算力仓、`AI应用` 或 AI 浏览器插件时，才额外判断和落官方 AI Pod 路线。

## Workflow

### 1. 确认项目创建模式

- 判断是全新项目、现有项目补齐，还是只补认证能力
- 判断前后端是否已经存在，以及当前结构是否接近 Go + Vue + Element Plus 基线
- 判断这是原创应用还是移植应用，以及是否存在 AI 能力诉求
- 先判断它是不是普通业务型 Web 应用；如果是，AI 默认走 `BaseURL` 配置方案
- 只有用户明确要求懒猫算力仓 / `AI应用` / AI 浏览器插件时，才进入 AI Pod 判断
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

按需补充：

- 如果是原创应用，补 `docs/requirements/lazycat-native-fit.md`
- 如果项目要接 Lazycat 原生能力，补 `docs/architecture/lazycat-integration.md`
- 如果项目存在 AI 场景，补 `docs/api-design/ai-integration.md`
- 如果项目明确要走懒猫算力仓 / `AI应用` 路线，补 `docs/architecture/aipod-integration.md`

如果仓库里已经有文档，不要机械重建；先整理成这个结构或映射到等价结构。

### 3. 建立命令入口

参考 `cloudmaze` 的约定，至少补齐：

- 根目录 `build.sh`
- 根目录 `Makefile`

最小命令要求：

- `build.sh` 负责串起后端构建、前端构建、产物输出
- `make build` 负责调用标准打包链路
- `make install` 负责在本地 Lazycat 环境安装产物

如果项目复杂，可以继续加：

- `make dev`
- `make test`
- `make clean`

但最少不要低于 `build.sh + make build + make install`。

### 4. 确认激励路径与原创价值

- 如果目标包含现金激励，优先判断该应用是不是原创应用，因为原创应用通常有更高的激励上限
- 如果是原创应用，不要只回答“现在没人发过”，还要回答它的 Lazycat 原生价值是什么
- 排除官方明确不发红包的类型，例如纯网页游戏、纯书籍页面、纯教程网站、网页离线应用、游戏服务器不同模组、纯数据库软件
- 如果是移植应用，要预留上游地址，并在发布前检查商店是否已存在同款，避免落到非首发激励路径
- 如果是工具类应用，优先评估是否可对接网盘文件关联
- 如果应用已有账户体系，优先评估是否可接入微服 OIDC
- 如果用户明确要做懒猫算力仓相关产品，再判断它是否应该挂到懒猫算力仓生态

### 5. 先判断与 Lazycat 原生系统的结合面

- 对原创应用，至少回答“用户为什么要在 Lazycat 里安装它，而不是直接打开一个网站”
- 优先评估微服 OIDC、`file_handler`、安装后即用的本地文件 / 工作流入口
- 如果是普通业务型 Web 应用，默认保留普通应用接模型 API
- 如果用户明确要求懒猫算力仓 / `AI应用` / AI 浏览器插件，再补充评估三条路径
- 如果当前想不到任何原生结合点，要把它标记为弱融合项目，并在汇报里提示风险
- 相关结论要提前写进 `docs/requirements/` 和 `docs/architecture/`，不要等做完再补

### 6. 固化技术基线

- 后端固定为 Go
- 前端固定为 Vue
- UI 固定为 Element Plus
- 前端状态管理和鉴权状态建议统一到单一 store
- 路由层必须预留匿名页和鉴权页边界

如果仓库已有更细的框架约束，以仓库现状为准，但不能偏离这三项主基线。

### 7. 落认证基线

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

### 8. 设计最小接口与页面

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

### 9. 设计 AI 能力与配置面板

如果项目存在明确 AI 场景，例如智能问答、文本处理、内容生成、分析助手或模型驱动工作流，还要额外完成：

- 先判断它是不是普通业务型 Web 应用；如果是，默认停留在 `BaseURL` 配置方案
- 先判断 AI 是不是核心流程的一部分，而不是为了“看起来有 AI”而硬塞进去
- AI 设置页最小项固定为 `API BaseURL`
- 协议类型固定支持 `OpenAI Compatible`、`OpenAI Responses`、`Anthropic`
- 提供“获取模型”按钮，基于当前 `BaseURL + 协议` 拉取模型列表
- 提供模型下拉选择框，不要把模型名写死在代码里
- 提供“保存配置”按钮，统一持久化 AI 连接配置
- 普通业务 Web 应用的 AI 设置页直接按 [references/ai-settings-template.md](./references/ai-settings-template.md) 落，不要每个项目各写一套
- 只有当用户明确目标是懒猫算力仓 / `AI应用` / AI 浏览器插件时，才判断是否需要官方 AI Pod 结构
- 如果是 `AI应用` 候选，提前规划 `ai-pod-service/`、`caddy-aipod` 入口，以及是否需要 `extension.zip`
- 如果项目暂时不接 AI，要在结论里明确说明“不接 AI”的原因，避免后续反复讨论

### 10. 收敛后台管理 UI

如果项目存在后台管理面、Admin 控制台、运营后台或明显的 backoffice 区域，还要额外完成：

- 明确后台是否需要工作台、核心列表、详情视图、表单页和设置页
- 判断是从零建设后台，还是把现有后台交给 `lazycat:admin-ui` 做质量收敛
- 如果要用模板，只允许选择 Vue + Element Plus 兼容模板，并明确需要替换的 branding、菜单、示例图表和默认 copy
- 登录 / 注册页与后台主框架必须保持同一品牌方向，避免像两个独立项目
- 在进入资料准备前，至少准备一套可截图的后台候选页

### 11. 对齐数据、安全边界与附加对接

- 用户表最小字段要能支撑注册、登录和状态判断
- refresh token 不能裸存明文长期复用，至少要具备可撤销或可轮换能力
- 认证相关环境变量、密钥、过期时间、Cookie/Header 约束要提前定下来
- 如果项目要上架 Lazycat，初始化说明里要明确首次登录、默认管理员或注册入口策略
- 如果接入微服 OIDC，要在 manifest 中规划 `application.oidc_redirect_path`，并把系统注入的 OIDC 环境变量正确透传给应用
- 如果是工具类应用，要在 manifest 中规划 `application.file_handler`、支持的 `mime`，以及应用侧 `/open` 或等价路由
- 如果项目接入 AI，要明确 AI 配置存储位置、是否按用户或按管理员生效，以及模型拉取失败时的降级策略
- 如果项目走懒猫算力仓 `AI应用` 路线，要明确 `ai-pod-service`、`caddy-aipod`、AI 浏览器插件或扩展入口的包结构约束

### 12. 交给发布链路

当项目已经具备可开发、可登录、可注册、可刷新 token 的基线后，再交回 `lazycat:ship-app` 继续走资料、截图、提审和发布；如果项目包含后台管理面，先交给 `lazycat:admin-ui` 收敛截图级页面，再进入资料阶段。

复杂基线收敛先读 [references/project-baseline.md](./references/project-baseline.md)、[references/docs-blueprint.md](./references/docs-blueprint.md) 和 [references/aipod-playbook.md](./references/aipod-playbook.md)。

## Quality Gates

- 明确已经建立 `docs/` 文档树，且需求分析、API 设计等目录下不是空目录
- 明确已经建立 `build.sh` 和 `Makefile`
- 明确采用 Go + Vue + Element Plus
- 明确具备登录、注册
- 明确具备 `access_token + refresh_token`
- 明确前端无感刷新流程和失败回退流程
- 如果项目是原创应用，明确为什么它适合融入 Lazycat 原生系统
- 明确最小接口、最小页面和最小用户数据要求
- 如果项目适合 AI，明确 AI 设置页至少包含 `API BaseURL`、协议、获取模型、模型选择、保存配置
- 如果项目是 AI 原生产品，明确是否走普通应用、懒猫算力仓 `AI应用` 或 AI 浏览器插件路径
- 如果项目存在后台管理面，明确后台是否已进入 `lazycat:admin-ui` 的高质量 UI 收敛链路
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
- 项目能运行，但没有 `build.sh`、没有 `make build` 或没有 `make install`
- 项目明明有后台管理面，却完全没有评估后台 UI 质量和模板改造策略
- 目标是拿激励，却选了官方明确不发红包的应用类型
- 应用需要登录，但普通用户拿不到账号或无法自助注册
- 明明是工具类应用，却没有评估文件关联；明明适合账户统一，却没有评估 OIDC
- 创建阶段没有考虑后续上架和初始化路径

## Bundled References

- 项目技术栈与认证基线： [references/project-baseline.md](./references/project-baseline.md)
- 文档目录蓝图： [references/docs-blueprint.md](./references/docs-blueprint.md)
- 普通业务 Web 应用 AI 设置页模板： [references/ai-settings-template.md](./references/ai-settings-template.md)
- AI Pod / `AI应用` / AI 浏览器插件决策： [references/aipod-playbook.md](./references/aipod-playbook.md)
- 后台管理 UI 质量链路： [../lazycat:admin-ui/SKILL.md](../lazycat:admin-ui/SKILL.md)
- 命令入口约定： [../lazycat:port-app/references/command-conventions.md](../lazycat:port-app/references/command-conventions.md)
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

命令入口
- build.sh
- Makefile
- make build
- make install

项目基线
- 后端: Go
- 前端: Vue + Element Plus
- 认证: access_token + refresh_token + 无感刷新
- 后台 UI: <不适用 / 进入 lazycat:admin-ui / 已具备高质量管理台>

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
