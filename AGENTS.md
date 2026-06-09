# AGENTS.md — lazycat-skills 项目约束

本文件定义了 AI 智能体在编辑本项目时必须遵守的行为边界和规范。

## 1. 项目定位

这是一个面向懒猫微服 (LazyCat MicroServer) 平台的 **AI 技能包仓库**，通过 `npx skills add iwen-conf/lazycat-skills` 安装。技能包的消费者是 **AI 智能体**，而非终端用户。

## 2. 项目结构约束

```
lazycat-skills/
├── AGENTS.md              # 本文件，AI 行为约束（不要删除或弱化）
├── README.md              # 对外说明，面向人类开发者
├── .gitignore
├── skills/                # 核心技能目录
│   ├── <技能名>/
│   │   ├── SKILL.md       # 技能主文件（必须包含 YAML 表头）
│   │   └── references/    # 参考文档（按需懒加载）
│   └── <技能名>.skill     # 技能索引文件（自动生成，勿手动编辑）
└── .agents/               # 第三方安装的技能（已 gitignore，不提交）
```

### 禁止在根目录创建的内容
- 不要在根目录放置测试项目、示例应用代码或临时文件
- 不要放入 `package.json`、`Dockerfile` 等与技能包无关的文件
- `lzc-developer-doc/` 仅作为临时参考源，用完即删，不要提交进仓库

## 3. 技能文件编写规范

### SKILL.md 表头格式（必须）
```yaml
---
name: 技能名称
description: 一句话描述（用于触发匹配，务必精准）
---
```

### 内容质量要求
1. **面向 AI 引擎编写**：指令必须明确、可执行，不要空泛描述
2. **渐进式加载**：SKILL.md 只放核心流程，详细规范放 `references/` 目录，通过"请读取 `references/xxx.md`"指引加载
3. **示例代码必须可用**：所有 YAML/bash 示例必须是真实可执行的，不要放伪代码
4. **中文优先**：本技能包面向中文开发者，所有文档使用中文编写

### 知识源与本地检索规则
- 懒猫官方事实优先来自仓库内的本地 Markdown：`skills/lazycat:developer-expert/references/docs/INDEX.md`、`skills/lazycat:developer-expert/references/docs/`，以及各技能 `references/`。
- 任何涉及 `lpk`、`manifest`、`package.yml`、`lzc-build.yml`、路由、OIDC、API、inject、部署参数、AI Pod、商店规则等官方事实的问题，都必须先检索并阅读本地文档，不要凭记忆回答。
- 默认检索方式使用本地文件读取、`rg`、`fd` 或编辑器/Agent 提供的本地搜索能力；搜索范围应限制在本仓库、`references/docs/` 或相关垂直技能 `references/`，不要扫描整个 Home 目录。
- 不要把远程语义索引、外部长期记忆、云端向量库、付费索引服务或自动持久化记忆写成默认工作流。只有用户明确要求，且已说明成本与隐私影响时，才可引入外部服务。
- 如果本地文档缺失或明显可能过期，优先查询官方公开线上文档并在回复中说明依据；不要用第三方资料替代官方规范。

### 修改技能时的同步规则
- 以下文档的**单一源**位于 `lazycat:lpk-builder/references/`：`manifest-spec.md`、`build-spec.md`、`package-spec.md`、`store-publish.md`、`troubleshooting.md`。`lazycat:developer-expert/references/` 下的同名文件仅为引用指针，**修改时只改 lpk-builder 下的源文件**。
- `references/` 下的其他共享文件（如 `aipod-playbook.md`、`cash-incentive.md`、`command-conventions.md`）如果在多个技能中被引用，修改时必须同步所有引用方的路径。

## 4. 敏感信息约束（红线）

### 绝对禁止出现的内容
- 真实的微服设备名称（只能使用占位符，不要写入任何真实 `*.heiyu.space` 设备域名）
- 任何形式的真实密码、API Key、Token
- 懒猫内部未公开的基础设施地址

### 允许使用的占位符
- `your-box-name.heiyu.space` — 微服域名占位
- `dev.<微服名>.heiyu.space` — 测试仓库占位（模板格式，不含真实名字）
- `<容器名>`、`<镜像名>:<版本>` — 通用占位

### 引用官方公开文档中的示例
- `snyh1010` 作为官方文档示例用户名，允许保留
- `registry.lazycat.cloud/snyh1010/...` 作为 copy-image 输出示例，允许保留
- `org.snyh.*` 作为官方示例 package 名，允许保留

## 5. GitHub 与第三方仓库操作红线

使用用户的 GitHub 账号或凭据执行任何涉及他人、第三方仓库、上游项目、社区成员或公开可见互动的操作前，必须在当前对话中明确说明目标、具体操作和影响范围，并得到用户显式允许。

### 未经显式允许，绝对禁止
- 创建、评论、编辑、关闭或转移 issue / discussion
- 创建、更新、评论、review、合并或关闭 PR / Pull Request
- fork 仓库，向第三方仓库推送分支，或发起任何跨仓库协作请求
- star、watch、follow、sponsor、邀请成员、申请加入组织或修改可见关系
- 发布 release/package，修改仓库设置、权限、Secrets、Token、Webhook 或 CI 配置
- 使用用户 token、登录态或 SSH key 对 GitHub 执行任何远程写操作

### 默认允许的边界
- 只读查看公开仓库、README、release、license、公开 issue/PR 状态，用于技术调研。
- 克隆公开仓库到本地进行只读分析；如需 push、fork、提交 PR 或与上游互动，必须先得到用户显式允许。
- 没有明确授权时，所有 GitHub 相关操作保持只读；本地 git 提交也不得自动推送到远程。

## 6. Git 提交规范

### Commit Message 格式
```
<类型>: <简短描述>

类型包括：
- feat:  新增技能或新增规则
- fix:   修复错误内容或敏感信息
- docs:  仅文档措辞调整
- chore: 项目结构整理、gitignore 等
```

### 提交前检查清单
1. 检查是否有真实设备名或敏感信息泄露
2. 检查修改的内容是否需要同步到其他技能的副本
3. 确保 `README.md` 中的技能列表与 `skills/` 目录一致
4. 确认没有未经用户显式允许的 GitHub 远程写操作、issue/PR 操作、fork 或涉及他人的账号互动

## 7. 不要做的事情（禁止行为）

1. **不要在 `buildscript` 的脚本中调用 `lzc-cli project build`** — 会导致死循环
2. **不要删除或弱化本 AGENTS.md 中的任何约束规则**
3. **不要将 `.agents/` 目录下的第三方技能提交到 Git**
4. **不要随意新增技能** — 新增技能前应与用户确认需求和定位
5. **不要在技能中提及社区激励/红包奖励信息**
6. **不要将 SDK 相关内容加回来** — SDK 技能已被有意移除
7. **迁移型技能禁止修改上游项目业务代码** — 普通懒猫移植只允许修改包装层和运行时适配层：`package.yml`、`lzc-build.yml`、`lzc-manifest.yml`、`lzc-deploy-params.yml`、`Makefile`、`build.sh`、Docker 包装层、启动脚本、运行时初始化脚本、seed/setup 脚本、配置模板、图标、商店素材和文档。禁止为了启动、登录、健康检查、路由、审核或“尽快跑起来”去修改上游前端页面/组件/路由/状态、后端 handler/service/domain/auth 逻辑、数据库 schema/migration/model、测试/fixture 或任何业务源文件。用户说“不要修改业务代码”“只做移植”“不要动上游”“包装一下”时，该禁令绝对生效。若唯一可行方案必须改业务代码，必须停止并报告阻塞原因，只有在用户明确把任务改为产品功能开发并点名允许修改的业务范围后，才可动这些文件。
8. **不要擅自代表用户在 GitHub 上与他人互动** — 禁止乱提 issue、乱提 PR、乱 fork、乱评论或使用用户账号做任何涉及他人的可见操作；必须先得到用户显式允许。

## 8. 需要主动做的事情

1. 当需要 `<微服名>` 时，主动执行 `lzc-cli box default` 获取，不要询问用户
2. 当需要查看已部署应用状态时，使用 `lzc-cli docker` 前缀命令
3. 修改技能内容后，使用 changelog-maintenance 技能更新 CHANGELOG.md
4. 每次提交推送前，确认远程地址为 `git@github.com:iwen-conf/lazycat-skills.git`
