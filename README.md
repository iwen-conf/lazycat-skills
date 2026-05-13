# Lazycat Skills

面向 LazyCat MicroServer 平台的 AI 技能包仓库。安装后，AI 可以围绕懒猫应用的创建、移植、打包、认证、后台管理和上架交付提供稳定的工程化能力。

## 文档检索：依赖 OpenViking 知识库

本仓库不再夹带懒猫官方文档的离线副本。所有事实型问题（lpk / manifest / spec / API / 路由 / 认证 / AI Pod 等）都由 `lazycat:developer-expert` 通过用户机器上的 [OpenViking](https://github.com/volcengine/OpenViking) 知识库实时检索：

- `viking://resources/lazycat-developer-docs/`：58 篇开发者手册（developer.lazycat.cloud）
- `viking://resources/lazycat-aipod-docs/`：93 篇算力舱手册（developer.lazycat.cloud/aipod）

入库方式见 `lazycat:developer-expert/SKILL.md` 中的 "Knowledge Base Protocol"。如果机器上没装 OpenViking，技能仍可用，但 AI 会退回到一般知识，准确度会下降。

## 技能分组

### 生命周期
- `lazycat:ship-app`：覆盖从立项、打包、自测、提审到发布后的完整交付流程。
- `lazycat:create-app`：统一新项目基线，包括文档树、技术栈、认证方案和 AI 配置面板。
- `lazycat:update-app`：处理已上架应用的版本升级、镜像同步和重新提审。
- `lazycat:port-app`：负责开源项目移植、选型、查重、构建入口和上架落地。

### 开发与打包
- `lazycat:developer-expert`：Lazycat 微服开发总控技能；负责把事实问题路由到 OpenViking，把工作流问题分发给下面的垂直技能。
- `lazycat:dynamic-deploy`：动态部署参数、模板渲染和脚本注入（含免密登录强制决策顺序）。
- `lazycat:troubleshoot`：应用排障，覆盖容器启动失败、路由异常、inject 不生效、OIDC 回调失败等常见问题。

### 资产与文档
- `lazycat:admin-ui`：收敛后台管理界面质量。
- `lazycat:prepare-icon`：生成应用图标设计提示词。
- `lazycat:write-guide`：生成应用攻略和官方风格文档。
- `lazycat:ui-ux-pro-max`：提供面向 `React + Vite + Tailwind CSS + shadcn/ui + Zustand + TanStack Query + React Router + React Hook Form + Zod + Framer Motion` 的 UI/UX 设计知识库。

> 历史版本曾包含 `lpk-builder` / `advanced-routing` / `auth-integration` / `aipod-developer` 四个技能。它们的内容都是官方文档的二次转抄，已统一删除并改为通过 OpenViking 检索。

## 仓库结构

```text
lazycat-skills/
├── README.md
├── AGENTS.md
├── .gitignore
└── skills/
    ├── lazycat:developer-expert/
    │   ├── SKILL.md
    │   └── lzc-manifest.yml  # transmission 示例 manifest
    ├── lazycat:create-app/
    │   ├── SKILL.md
    │   └── references/
    └── ... 其他技能目录
```

## 安装

```bash
npx skills add iwen-conf/lazycat-skills
```

安装后，AI 会自动发现这些技能。常见入口包括：
- "帮我把这个 Docker 项目打包成懒猫应用"
- "帮我给懒猫项目补一个高质量后台管理界面"
- "帮我整理应用上架需要的截图、图标和提审材料"

## 贡献约定

- 根目录不放测试工程、临时目录或与技能包无关的构建文件。
- **不要把懒猫官方文档拷贝进 `references/`**——通过 OpenViking 实时检索。
- 技能 `references/` 只放：(a) 仅本技能使用的私有 cheat-sheet；(b) 官方文档没有的经验沉淀。
