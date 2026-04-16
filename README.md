# Lazycat Skills

面向 LazyCat MicroServer 平台的 AI 技能包仓库。安装后，AI 可以围绕懒猫应用的创建、移植、打包、路由、认证、后台管理和上架交付提供稳定的工程化能力。

## 技能分组

### 生命周期
- `lazycat:ship-app`：覆盖从立项、打包、自测、提审到发布后的完整交付流程。
- `lazycat:create-app`：统一新项目基线，包括文档树、技术栈、认证方案和 AI 配置面板。
- `lazycat:update-app`：处理已上架应用的版本升级、镜像同步和重新提审。
- `lazycat:port-app`：负责开源项目移植、选型、查重、构建入口和上架落地。

### 开发与打包
- `lazycat:developer-expert`：Lazycat 微服开发总控技能。
- `lazycat:lpk-builder`：负责 `.lpk` 打包规范和构建细节。
- `lazycat:advanced-routing`：处理多域名、四层转发和复杂代理规则。
- `lazycat:auth-integration`：处理 OIDC、用户身份透传和 API 鉴权。
- `lazycat:aipod-developer`：处理 AI Pod 应用、算力舱能力和浏览器插件打包。
- `lazycat:dynamic-deploy`：处理动态部署参数、模板渲染和脚本注入。
- `lazycat:troubleshoot`：应用排障，覆盖容器启动失败、路由异常、inject 不生效、OIDC 回调失败等常见问题。

### 资产与文档
- `lazycat:admin-ui`：收敛后台管理界面质量。
- `lazycat:prepare-icon`：生成应用图标设计提示词。
- `lazycat:write-guide`：生成应用攻略和官方风格文档。
- `lazycat:ui-ux-pro-max`：提供可检索的 UI/UX 设计知识库。

## 仓库结构

GitHub 根目录只保留仓库说明和约束文件，所有技能统一收拢到 `skills/`，避免根目录堆满技能目录或出现重复副本。

```text
lazycat-skills/
├── README.md
├── AGENTS.md
├── .gitignore
└── skills/
    ├── lazycat:create-app/
    │   ├── SKILL.md
    │   └── references/
    ├── lazycat:ship-app/
    └── ... 其他技能目录
```

## 安装

```bash
npx skills add whoamihappyhacking/lazycat-skills
```

安装后，AI 会自动发现这些技能。常见入口包括：
- “帮我把这个 Docker 项目打包成懒猫应用”
- “帮我给懒猫项目补一个高质量后台管理界面”
- “帮我整理应用上架需要的截图、图标和提审材料”

## 贡献约定

- 根目录不放测试工程、临时目录或与技能包无关的构建文件。
- 技能主文档保留核心流程，详细说明放在各自的 `references/` 目录中。
- 修改共享参考文档时，需要同步受影响的其他技能副本。
