# Lazycat Skills

面向 LazyCat MicroServer 平台的 AI 技能包仓库。安装后，AI 只围绕原创应用、迁移评估和上架交付这三类工作提供能力。

## 技能入口

### 原创

- `lazycat:original-app`：原创应用从想法或已有项目进入懒猫工程基线，补齐交付文件、登录、文件能力和上架前结构。
- 原创应用如包含 Web、移动端、桌面端或小程序前端，除非用户明确指定其他技术，默认联动 `arc:frontend`：Web 使用 React 19 + TypeScript + Vite + Tailwind CSS + shadcn/ui + Zustand + TanStack Query + TanStack Router + React Hook Form + Zod；移动端使用 React Native + Expo + TypeScript + NativeWind + Zustand + TanStack Query + Expo Router；桌面端使用 Tauri 2 + React Web 栈；小程序使用 Taro 4 + React + TypeScript + Zustand。

### 迁移

迁移候选可以先从开发者中心盘点进入，再按顺序通过三关：

0. `lazycat:unlisted-candidate-audit`：对比开发者中心、应用商店和本地项目目录，找到未上架且非待审核的应用；对缺失的公开 GitHub 项目只读 clone，并进入许可证和非侵入上架可行性审查。
1. `lazycat:migration-license`：搜索 Web/Agent 候选项目，筛选“有页面 + 有后端”的 GitHub 项目，对比懒猫应用商店和开发者中心待审列表，再判断许可证是否允许商业使用和再分发。
2. `lazycat:migration-boundary`：判断能否不修改上游业务代码，只通过包装层和运行时适配完成迁移。
3. `lazycat:migration-workload`：判断迁移工作量大不大，给出继续、先 POC、换项目或停止的建议。

迁移项目默认保留上游业务前端/客户端栈；不得为了统一默认栈而改写上游 React/Vue/Svelte/Angular/静态前端、移动端、桌面端或小程序客户端。只有新建包装层、管理台、审核辅助页或原创配套前端面时，才调用 `arc:frontend` 的平台默认栈。

### 上架

- `lazycat:ship-app`：接收已准备好的原创应用，或已通过迁移三关的迁移项目，执行打包、安装验证、商店资料、提审和发布后检查。
- 所有最终 `.lpk` 必须小于或等于 12 MB（按 `12,000,000` bytes 检查），且禁止内嵌镜像；不得使用 `lzc-build.yml.images`、`embed:<alias>`、包内 `images/` 或 `images.lock`。
- 镜像型项目必须使用可拉取远程镜像；提审前通过 `lzc-cli appstore copy-image` 同步到 `registry.lazycat.cloud/...` 并写回 manifest 后再打包。

## 本地知识库

事实型问题优先读取仓库内本地 Markdown，不默认依赖远程语义索引、外部长期记忆、云端向量库或付费索引服务。

主要本地来源：

- `skills/lazycat:ship-app/references/docs/INDEX.md`：官方开发文档 URL 到本地文件的索引。
- `skills/lazycat:ship-app/references/docs/`：拆分后的官方开发文档 Markdown。
- `skills/lazycat:ship-app/references/lpk/`：`manifest`、`build`、`package`、上架和运行模型规范。
- 各保留技能目录下的 `references/`：迁移许可证、非侵入边界和工作量评估清单。

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
    ├── lazycat:unlisted-candidate-audit/
    ├── lazycat:migration-license/
    ├── lazycat:migration-boundary/
    ├── lazycat:migration-workload/
    └── lazycat:ship-app/
```

## 安装

```bash
npx skills add iwen-conf/lazycat-skills
```

## 贡献约定

- 根目录不放测试工程、临时目录或与技能包无关的构建文件。
- 不新增技能入口，除非用户明确要求改变“原创 / 迁移三关 / 上架”的强约束。
- 技能主文档只保留可执行流程；详细事实和规范放在各自 `references/` 目录中。
- 不在 Lazycat 技能中重新定义另一套前端或跨端技术选型；需要 Web、移动、桌面或小程序前端时引用 `arc:frontend`，迁移项目仍以“不修改上游业务代码”为最高优先级。
- 修改技能内容后同步更新 `CHANGELOG.md`。
