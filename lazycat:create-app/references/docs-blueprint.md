# Lazycat 文档目录蓝图

项目创建的第一步，不是直接写代码，而是先建立 `docs/` 文档树。至少拆成“需求分析”“API 设计”等多个目录，并在每个目录下放多个 `.md` 文档。

## 1. 推荐目录结构

```text
docs/
├── requirements/
│   ├── product-overview.md
│   ├── user-stories.md
│   └── scope-and-milestones.md
├── api-design/
│   ├── overview.md
│   ├── auth.md
│   └── domain-modules.md
├── architecture/
│   ├── system-overview.md
│   ├── frontend.md
│   └── backend.md
└── release-prep/
    ├── store-assets.md
    ├── test-plan.md
    └── submission-checklist.md
```

## 2. 每个目录最少写什么

### `docs/requirements/`

- `product-overview.md`：项目目标、用户、核心价值、边界
- `user-stories.md`：关键用户故事、主要流程、异常路径
- `scope-and-milestones.md`：MVP 范围、阶段目标、延期项

### `docs/api-design/`

- `overview.md`：接口风格、鉴权方式、返回结构、错误约定
- `auth.md`：登录、注册、刷新、退出、用户信息接口
- `domain-modules.md`：业务模块 API，按模块拆开写

### `docs/architecture/`

- `system-overview.md`：整体架构、模块关系、部署边界
- `frontend.md`：Vue、Element Plus、路由、状态管理、鉴权流
- `backend.md`：Go 服务划分、认证链路、数据库与外部依赖

### `docs/release-prep/`

- `store-assets.md`：应用简介、截图、图标、类目、卖点
- `test-plan.md`：测试范围、关键路径、认证与刷新验证
- `submission-checklist.md`：提审前检查项、证据项、 reviewer 复现路径

## 3. 写法要求

- 不要把所有内容堆在一个 `README.md`
- 每个目录至少 2 到 3 个 `.md`
- API 文档必须单独有 `auth.md`
- 需求文档里必须明确 MVP 和非目标
- 发布文档里必须覆盖登录、注册、token 刷新和提审路径

## 4. 与代码的关系

- 需求文档定范围
- API 文档定接口
- 架构文档定分层和职责
- 发布文档定上架和提审入口

代码实现要追着这四类文档走，不要反过来用代码去猜需求。
