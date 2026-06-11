---
name: lazycat:migration-license
description: "First gate for migrating to Lazycat: discover full-stack Web/Agent GitHub candidates, compare App Store and Developer Center pending list, then check commercial license. 迁移第一关、候选搜索、商店和待审查重、GitHub许可证、可商用."
---

# Lazycat Migration Discovery and License Gate

本技能是迁移第一关。它有两个模式：

1. 用户给出 GitHub 仓库时，判断该项目是否具备进入懒猫迁移评估的许可证基础。
2. 用户找不到可迁移应用时，主动搜索并筛选候选项目，再做商店与待审列表查重、许可证判断。

没有通过本关，不进入非侵入迁移和工作量评估。

## 使用场景

- 用户给出 GitHub 仓库，询问是否适合迁移到懒猫。
- 用户要找可迁移项目，需要先筛许可证。
- 用户没有候选项目，需要 AI 搜索和筛选可迁移项目。
- 用户明确要找 Web 相关、前后端分离、Agent 相关、有页面有后端的项目。
- 用户问某个开源项目能不能商用、能不能上架、能不能打包分发。

## 强约束

1. GitHub 调研默认只读。不得创建 issue、PR、fork、评论、review、discussion、star、watch、follow，除非用户在当前对话明确授权具体动作。
2. 候选搜索只选“有页面 + 有后端”的项目：必须能看到 Web UI、API/server/backend/worker 等后端证据。纯前端、纯后端、CLI-only、库、SDK、模板、纯 prompt、纯文档站默认过滤。
3. 候选优先类型：前后端分离 Web 应用、带 Web 控制台的服务端应用、Agent/多 Agent/RAG/自动化工作台等有页面和后端的项目。
4. 每个候选都必须和两个位置比较：懒猫应用商店已上架列表、懒猫开发者中心待审列表。没有完成两处比较时，不得宣称“没有重复”。
5. 应用商店检查使用 `lazycat_account` / `lazycat_password` 或已有登录态；开发者中心待审列表检查使用 `lazycat_developer_center_account` / `lazycat_developer_center_password` 或已有登录态。不得混用凭据。
6. 对应用商店和开发者中心只做只读查看；不得创建、编辑、提交、撤回、删除、评论或改变任何审核状态。
7. 必须找到并读取上游许可证来源：`LICENSE`、`COPYING`、README、package metadata、官网许可说明或 release 资产许可。
8. 没有许可证、许可证不明、只写“all rights reserved”、含非商用限制，结论必须是 `Blocked` 或 `Unclear`，不能默认可迁移。
9. 许可证必须允许商业使用和再分发；如果允许商业使用但有 copyleft、署名、源码公开、网络服务条款等义务，必须写清义务。
10. 不能只看代码许可证，还要检查 logo、图片、字体、模型、数据集、插件市场资源是否有单独非商用限制。
11. 不能用第三方文章替代上游仓库或官方许可证文本。

## 候选搜索标准

优先搜索这些方向：

- Web full-stack / 前后端分离：`frontend`、`web`、`api`、`server`、`backend`、`dashboard`、`admin`、`console` 同时存在。
- Agent 相关：有 Web UI、任务编排、工具调用、工作流、知识库或自动化后端，而不是纯 SDK 或示例脚本。
- 自托管成熟度：有 Docker image、Dockerfile、Compose、部署文档或明确服务启动方式。
- 用户可进入：有登录、无登录、默认账号、OIDC 或可初始化账号路径；完全依赖私有 SaaS 的项目降级。

默认排除：

- 纯前端静态站、纯教程站、纯文档站。
- 纯 CLI、库、SDK、starter template、代码生成模板。
- 没有可见 UI 或没有后端服务的项目。
- 只依赖外部闭源 SaaS 才能运行的项目。
- 许可证不允许商业使用或再分发的项目。
- 已在应用商店上架，或已在开发者中心待审列表中存在高度重合项目，且没有用户认可的差异化理由。

## 两处查重标准

必须分别记录：

- 懒猫应用商店：按项目名、中文译名、英文别名、核心功能关键词、上游仓库 URL、作者/组织搜索。
- 懒猫开发者中心待审列表：按相同维度检查待审、审核中、驳回待重提、草稿中已经占用的同类项目。

查重结论：

- `Clear`: 两处都查过，未发现同名或高重合项目。
- `Duplicate Published`: 应用商店已有同名或高度重合项目。
- `Duplicate Pending`: 开发者中心待审列表已有同名或高度重合项目。
- `Incomplete`: 缺少商店或待审列表访问权限，不能下最终无重复结论。

只有 `Clear` 或用户明确接受差异化理由的候选，才继续做许可证结论。

## 判断标准

- `Pass`: 明确允许商业使用和再分发，且没有明显资产许可阻塞。
- `Pass with Obligations`: 可商用，但必须遵守署名、许可证保留、源码提供、修改说明、网络服务条款等义务。
- `Unclear`: 没有明确许可证、存在多许可证冲突、资产许可缺失、上游说明自相矛盾。
- `Blocked`: 非商用、禁止再分发、禁止修改、闭源专有、许可证不允许上架所需行为。

常见可商用许可证包括 MIT、Apache-2.0、BSD、ISC、MPL-2.0、LGPL、GPL、AGPL。Copyleft 许可证不是“不可商用”，但会带来分发和源码义务，必须在结论中说明。

## 工作流

1. 如果用户未给出仓库，先搜索 GitHub 候选；每轮输出不超过 10 个短名单，避免泛滥。
2. 对候选做形态过滤：必须有页面和后端证据，并记录 Docker/Compose/部署证据。
3. 对每个短名单候选做两处查重：懒猫应用商店、懒猫开发者中心待审列表。
4. 排除 `Duplicate Published`、`Duplicate Pending` 和 `Incomplete` 中无法继续判断的候选；权限不足时明确把候选标为临时结论。
5. 读取仓库许可证文件和 README 中的许可声明。
6. 检查依赖、前端资产、字体、图标、模型、数据集、插件资源是否另有许可证。
7. 判断商业使用、修改、再分发、署名、源码公开、商标/品牌使用限制。
8. 输出结论；只有查重可接受且许可证为 `Pass` 或 `Pass with Obligations` 才能进入 `lazycat:migration-boundary`。

## 输出格式

```text
Phase: Migration License Gate
Repository: <owner/repo>
Mode: Provided Repo / Candidate Discovery

Decision: Pass / Pass with Obligations / Unclear / Blocked

Candidate Evidence
- UI:
- Backend:
- Deploy:
- Agent/Web fit:

Duplicate Checks
- Lazycat App Store: Clear / Duplicate Published / Incomplete
- Developer Center Pending List: Clear / Duplicate Pending / Incomplete

License Evidence
- Code:
- Assets:
- Dependencies:

Commercial Use
- Allowed:
- Obligations:
- Blockers:

Next
- Proceed to lazycat:migration-boundary / Stop
```
