---
name: lazycat:unlisted-candidate-audit
description: "Inventory Lazycat Developer Center apps against App Store listings and local project folders, clone unlisted non-pending GitHub projects, then audit commercial license and non-invasive listing feasibility. 开发者中心盘点、未上架非待审核、项目目录对比、clone审查、可商用、无需修改业务代码可上架."
---

# Lazycat Unlisted Candidate Audit

本技能用于从懒猫微服开发者中心反向盘点可继续处理的迁移候选：找到“未上架、且不在待审核/审核中”的应用，和当前项目目录下已有文件夹对比；对缺失的公开 GitHub 项目执行只读 clone，并判断是否满足可商用、可再分发、可不改业务代码上架。

## 强约束

1. 开发者中心、应用商店和 GitHub 默认只读。不得创建、编辑、提交、撤回、删除、评论、fork、issue、PR、star、watch、follow 或改变任何审核状态，除非用户在当前对话明确授权具体动作。
2. 浏览器操作必须使用 `agent-browser`；如果该技能不可用，停止并报告浏览器盘点受阻，不得改用其他浏览器控制路径。
3. 开发者中心登录使用 `lazycat_developer_center_account` / `lazycat_developer_center_password` 或已有登录态；应用商店查询使用 `lazycat_account` / `lazycat_password` 或已有登录态。不得混用凭据。
4. 只处理开发者中心中状态为未上架、下架、草稿、驳回、待补充或等价非上架状态的应用；明确排除待审核、审核中、排队审核、已提交待处理等状态。
5. 应用商店必须按名称、别名、包名、上游仓库 URL、作者/组织和核心功能关键词查重。已上架或高度重合已上架的项目不得作为候选。
6. 对比本地项目目录时只读取目录名、`package.yml`、`lzc-manifest.yml`、README、迁移记录和 Git remote 等元数据；不要扫描无关大文件、构建产物或用户私密目录。
7. 只 clone 公开 GitHub 仓库或用户明确给出的仓库 URL。不得使用用户账号对第三方仓库做写操作或可见互动。
8. clone 目录必须放在用户指定的工作区；未指定时放入当前项目的同级或明确的候选工作目录，不能放在技能包仓库根目录。
9. 许可证结论必须沿用 `lazycat:migration-license`。没有许可证、许可证不明、非商用、禁止再分发或资产许可冲突，结论必须停止。
10. 业务代码边界必须沿用 `lazycat:migration-boundary`。如果上架必须修改上游业务代码，输出 `Blocked by business-code change requirement`，不得继续普通迁移。
11. 只有同时满足“未上架且非待审核”“GitHub 项目许可证可商用/可再分发”“不用修改业务代码即可迁移上架”的项目，才可进入后续打包或上架准备。

## 输入

- 开发者中心应用列表访问方式：已有浏览器登录态，或可用账号变量。
- 应用商店访问方式：公开页面、已有登录态，或可用账号变量。
- 本地项目父目录：用于和开发者中心应用做目录级对比。
- clone 目标目录：用户未指定时先确认或使用当前任务明确的候选工作区。

## 工作流

1. 读取项目上下文：确认本地项目父目录、clone 目标目录、是否存在 `.lark.json`、是否有用户明确授权任何写操作。
2. 使用 `agent-browser` 只读打开开发者中心，导出或记录应用清单字段：应用名、包名、状态、版本、上游仓库 URL、备注、更新时间。
3. 标准化状态：
   - `Published`: 已上架、已发布、商店可见。
   - `Pending`: 待审核、审核中、已提交、排队审核。
   - `Unlisted`: 未上架、下架、草稿、驳回、待补充、未提交或等价状态。
   - `Unknown`: 页面字段不足，不能判断。
4. 过滤候选：只保留 `Unlisted`；排除 `Published`、`Pending`、`Unknown`。
5. 对每个 `Unlisted` 候选查应用商店：按名称、包名、别名、仓库 URL、作者和功能关键词搜索。命中同名或高度重合已上架项目时移除。
6. 读取本地项目父目录，按目录名、`package.yml.package`、`package.yml.homepage`、Git remote、README 项目名匹配候选。已存在本地项目时记录路径，不重复 clone。
7. 对本地不存在且有公开 GitHub URL 的候选执行只读 `git clone` 到 clone 目标目录；没有仓库 URL 或仓库非公开时标为 `Blocked: missing source repository`。
8. 对每个已存在或新 clone 的候选运行 `lazycat:migration-license`：必须读取上游许可证、README 许可声明、资产/字体/模型/数据集许可和依赖许可风险。
9. 对许可证通过的候选运行 `lazycat:migration-boundary`：判断是否能只通过包装层、运行时配置、OIDC、inject、seed/setup、远程镜像桥接等方式迁移，且最终 `.lpk` 可保持 `<= 12,000,000` bytes 并禁止内嵌镜像。
10. 输出候选清单和阻塞原因；只有满足全部门槛的项目进入 `Ready for lazycat:migration-workload`。

## 本地目录匹配规则

优先级从高到低：

1. `package.yml.package` 与开发者中心包名完全一致。
2. `package.yml.homepage`、Git remote 或 README 上游链接与开发者中心仓库 URL 一致。
3. 目录名、应用名、英文名、中文名、别名规范化后相同。
4. 核心功能高度重合但名称不同：标为 `Possible Local Match`，需要人工确认，不得自动 clone 覆盖。

规范化只用于比较，不得改写文件名或项目 metadata：转小写、去空格、去常见分隔符、去 `lazycat` / `lzc` 包装前后缀。

## 决策

- `Ready`: 未上架、非待审核、商店无重复、本地已存在或已 clone、许可证可商用可再分发、可不改业务代码上架。
- `Existing Local`: 候选已在本地目录中存在，继续对本地路径跑许可证和边界审查。
- `Duplicate Published`: 应用商店已有同名或高重合应用。
- `Skip Pending`: 开发者中心状态为待审核或审核中。
- `Blocked License`: 许可证不明、非商用、禁止再分发或资产许可阻塞。
- `Blocked Source`: 没有公开 GitHub 仓库或无法只读 clone。
- `Blocked Business Code`: 必须修改上游业务代码才能上架。
- `Needs Manual Review`: 状态、商店重复、本地匹配或许可证证据不足。

## 输出格式

```text
Phase: Unlisted Candidate Audit
Developer Center Source: <account/login state/page>
Local Projects Root: <path>
Clone Target: <path>

Summary
- Developer Center apps scanned:
- Published excluded:
- Pending excluded:
- Unlisted candidates:
- App Store duplicates:
- Existing local matches:
- Newly cloned:
- Ready:
- Blocked:

Candidate: <name>
- Package:
- Developer Center Status: Published / Pending / Unlisted / Unknown
- App Store Check: Clear / Duplicate Published / Incomplete
- Local Match: Existing Local / New Clone / Missing / Possible Local Match
- Repository:
- License Gate: Pass / Pass with Obligations / Unclear / Blocked
- Boundary Gate: Can migrate non-invasively / Can migrate with wrapper risk / Cannot determine yet / Blocked by business-code change requirement
- Decision: Ready / Existing Local / Duplicate Published / Skip Pending / Blocked License / Blocked Source / Blocked Business Code / Needs Manual Review
- Evidence:
- Next:
```
