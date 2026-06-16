# Lazycat Skills

面向 LazyCat MicroServer 平台的 AI 技能包仓库。安装后，AI 只围绕原创应用、迁移评估、上架交付和已安装应用更新提供能力。

## 技能入口

### 原创

- `lazycat:original-app`：原创应用从想法或已有项目进入懒猫工程基线，补齐交付文件、登录、文件能力和上架前结构。
- 原创应用如包含 Web、移动端、桌面端或小程序前端，除非用户明确指定其他技术，默认联动 `arc:frontend`：Web 使用 React 19 + TypeScript + Vite + Tailwind CSS + shadcn/ui + Zustand + TanStack Query + TanStack Router + React Hook Form + Zod；移动端使用 React Native + Expo + TypeScript + NativeWind + Zustand + TanStack Query + Expo Router；桌面端使用 Tauri 2 + React Web 栈；小程序使用 Taro 4 + React + TypeScript + Zustand。

### 迁移

迁移候选按顺序通过两关：

1. `lazycat:migration-license`：候选发现与许可证门禁。支持用户给定 GitHub 仓库、主动搜索 Web/Agent 候选项目，或从开发者中心盘点未上架且非待审核应用；同时完成应用商店/开发者中心查重、本地目录对比、只读 clone 和商业许可证判断。
2. `lazycat:migration-boundary`：非侵入迁移可行性与工作量门禁。判断能否不修改上游业务代码，只通过包装层和运行时适配完成迁移，并给出工作量等级、POC/继续/停止建议。

迁移项目默认保留上游业务前端/客户端栈；不得为了统一默认栈而改写上游 React/Vue/Svelte/Angular/静态前端、移动端、桌面端或小程序客户端。只有新建包装层、管理台、审核辅助页或原创配套前端面时，才调用 `arc:frontend` 的平台默认栈。

### 上架

- `lazycat:ship-app`：接收已准备好的原创应用，或已通过迁移两关的迁移项目，执行打包、安装验证、商店资料、提审和发布后检查。
- 所有最终 `.lpk` 必须小于或等于 12 MB（按 `12,000,000` bytes 检查），且禁止内嵌镜像；不得使用 `lzc-build.yml.images`、`embed:<alias>`、包内 `images/` 或 `images.lock`。
- 镜像型项目必须使用可拉取远程镜像；提审前通过 `lzc-cli appstore copy-image` 同步到 `registry.lazycat.cloud/...` 并写回 manifest 后再打包。

### 更新

- `lazycat:update-installed-app`：通过用户提供的 LPK Inspector 页面获取已安装应用并下载当前 `.lpk`，对比 GitHub 最新 release/tag、更新时间和镜像，优先 `copy-image` 回写 manifest 后重新打包；没有上游镜像但可不改业务代码时，构建并推送公开镜像后再同步。

## 本地知识库

事实型问题优先读取仓库内本地 Markdown，不默认依赖远程语义索引、外部长期记忆、云端向量库或付费索引服务。

主要本地来源：

- `skills/lazycat:ship-app/references/docs/INDEX.md`：官方开发文档 URL 到本地文件的索引。
- `skills/lazycat:ship-app/references/docs/`：拆分后的官方开发文档 Markdown。
- `skills/lazycat:ship-app/references/lpk/`：`manifest`、`build`、`package`、上架和运行模型规范。
- 各保留技能目录下的 `references/`：迁移可行性、远程镜像桥接和命令入口约定。

推荐检索顺序：

1. 优先使用 `.ai-code-index/search.sh "query"` 做本地索引检索。
2. 查语法形态时使用 `.ai-code-index/struct-search.sh <language> '<pattern>'`。
3. 查符号时使用 `.ai-code-index/symbols.sh "SymbolName"`。
4. 索引缺失或结果不足时，再对 `skills/` 下相关目录做有范围的本地文本检索。
5. 只有本地文档缺失、规则明显可能过期，或用户明确要求联网核验时，才查询官方线上文档。

## 仓库结构

```text
lazycat-skills/
├── README.md
├── AGENTS.md
├── CHANGELOG.md
└── skills/
    ├── lazycat:original-app/
    ├── lazycat:migration-license/
    ├── lazycat:migration-boundary/
    ├── lazycat:update-installed-app/
    └── lazycat:ship-app/
```

## 安装

```bash
npx skills add iwen-conf/lazycat-skills
```

## 贡献约定

- 根目录不放测试工程、临时目录或与技能包无关的构建文件。
- 不新增技能入口，除非用户明确要求改变“原创 / 迁移两关 / 上架 / 更新”的强约束。
- 技能主文档只保留可执行流程；详细事实和规范放在各自 `references/` 目录中。
- 不在 Lazycat 技能中重新定义另一套前端或跨端技术选型；需要 Web、移动、桌面或小程序前端时引用 `arc:frontend`，迁移项目仍以“不修改上游业务代码”为最高优先级。
- 修改技能内容后同步更新 `CHANGELOG.md`。
