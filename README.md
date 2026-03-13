# Lazycat Skills

Lazycat 专项技能目录，只收纳 `lazycat:*` 命名空间下的技能与配套参考资料。

- 技能命名空间：`lazycat:*`
- 仓库边界：仅保留 Lazycat 领域技能；通用调试或通用协作 skill 不放在这个仓库里
- 当前主技能：`lazycat:ship-app`
- 当前子技能：`lazycat:create-app`、`lazycat:prepare-icon`
- 当前目标：围绕懒猫开发者中心、`lpk`、商店资料、提审、发布和发布后核验形成闭环

## 仓库结构

```text
Lazycat-Skills/
├── README.md
├── lazycat:create-app/
│   ├── SKILL.md
│   └── references/
│       └── project-baseline.md
├── lazycat:prepare-icon/
│   ├── SKILL.md
│   └── references/
│       └── app-icon-prompt.md
└── lazycat:ship-app/
    ├── SKILL.md
    └── references/
        ├── shipping-checklist.md
        └── store-assets.md
```

## 当前形态

- `lazycat:ship-app` 是总控型 skill，负责从立项、资料整理、打包、提审到发布后核验的完整链路
- `lazycat:create-app` 负责新项目创建、第一步文档树、技术栈基线和统一认证能力落地
- `lazycat:prepare-icon` 负责在资料阶段输出可直接交给外部图像模型的 App Icon prompt
- `references/shipping-checklist.md` 用于提审前检查与发布后复核
- `references/store-assets.md` 用于应用简介、截图和提审资料包整理

## 演进方向

当前先维持一个强主技能，确保闭环稳定。后续如果 Lazycat skill 继续扩张，再把能力下沉为更细的子技能，例如：

- `lazycat:create-app`
- `lazycat:prepare-assets`
- `lazycat:prepare-icon`
- `lazycat:submit-review`
- `lazycat:post-release-check`
