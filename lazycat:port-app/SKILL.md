---
name: lazycat:port-app
description: 面向 Lazycat 应用移植选型与落地的 skill。只要用户提到移植应用、移植开源项目、从 GitHub 选项目、查重、避免重复移植、看 App Store 里是否已有、移植激励、上游仓库、build.sh、Makefile、lpk 打包、file_handler、OIDC 等请求，就必须使用此 skill。负责从 GitHub 搜索候选项目、在 Lazycat App Store 查重、判断是否还有激励空间，并把移植项目收敛到可打包、可安装、可提审的状态。
---

# Lazycat 应用移植

你负责把“找一个可移植项目”推进到“可移植且值得移植”的状态。重点不只是移植本身，而是先做市场去重和激励判断，避免把时间花在已经有重复移植、没有红包空间、或根本不适合上架的项目上。

## Overview

这个 skill 用于移植开源或自托管软件到 Lazycat。默认要求：

- 先搜 GitHub，确定上游项目、许可证、活跃度和功能边界
- 再搜 Lazycat App Store，确认是否已有相关移植产品
- 如果已存在同类移植且没有新增价值，不走重复移植路径
- 项目落地后必须提供 `build.sh`、`Makefile`、`make build`、`make install`
- 必须保留上游地址、许可证和移植说明

如果项目是工具类或需要统一登录，还要额外评估 `file_handler` 和微服 OIDC。

## Quick Contract

- **Trigger**: 用户提到移植、port、从 GitHub 找项目、查重、避免重复移植、移植激励、上游仓库、Makefile、build.sh、file_handler、OIDC
- **Inputs**: 候选项目关键词、目标品类、激励目标、GitHub 上游信息、App Store 可访问状态、是否需要登录或文件关联
- **Outputs**: 候选项目清单、去重结论、激励判断、移植落地要求、脚本入口要求、后续交给 `lazycat:ship-app` 的发布入口
- **Quality Gate**: 必须完成 GitHub 搜索和 App Store 查重；如果查到重复移植且没有明显差异化，不继续按激励路径推进；落地方案必须包含 `build.sh` 和 `Makefile`
- **Decision Tree**: 先判断要移植什么，再做 GitHub 搜索、App Store 查重、激励判断、集成机会判断，最后决定继续移植还是换项目

## When to Use

**首选触发**

- 用户要把某个 GitHub 项目移植到 Lazycat
- 用户要你帮忙找值得移植的项目
- 用户明确说“先查有没有重复移植”
- 用户希望按移植激励规则来选项目

**典型场景**

- 从 GitHub 搜一批自托管应用，筛选适合移植到 Lazycat 的候选
- 移植工具类应用，并评估网盘右键菜单
- 移植需要登录的应用，并评估微服 OIDC
- 已有上游仓库，但不确定 App Store 里是否已经有同类移植

**边界提示**

- 如果用户要的是原创应用，不要误用本 skill，优先走 `lazycat:create-app`
- 如果用户只要写攻略文章，不要误用本 skill，优先走 `lazycat:write-guide`
- 如果 App Store 查重没有完成，不要对“值得移植”下结论

## Announce

开始执行后，先给用户一个短摘要：

- 你将搜哪些 GitHub 关键词
- 是否已经具备 App Store 搜索条件
- 当前最关键的 blocker 是上游质量、重复移植，还是激励空间

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `candidate_keywords` | string | 推荐 | GitHub 搜索关键词，可含品类、协议、技术方向 |
| `reward_target` | enum(`普通移植`/`现金激励优先`) | 推荐 | 如果目标是激励，重复移植和不奖励类型要尽早排除 |
| `appstore_access` | enum(`已登录可搜索`/`未登录`/`未知`) | 推荐 | App Store 查重在当前站点通常需要登录会话；没有会话时不要跳过 |
| `upstream_state` | enum(`已指定仓库`/`仅有方向`/`还未搜索`) | 推荐 | 用于判断是直接评估，还是先做 GitHub 搜索 |
| `integration_hint` | enum(`OIDC`/`file_handler`/`两者都评估`/`无`) | 可选 | 判断是否适合接微服账户系统或网盘文件关联 |

## The Iron Law

1. 移植前必须先做 GitHub 搜索和 App Store 查重；不要先写适配代码再回头看是否重复。
2. 如果 App Store 已经存在同类移植，且你没有明确差异化理由，不要继续按激励路径推进。
3. 每个移植项目都必须带上上游地址、许可证和移植说明。
4. 每个移植项目都必须提供 `build.sh`、`Makefile`、`make build`、`make install`。
5. 如果项目适合 OIDC 或 `file_handler`，优先评估并记录，因为这会影响激励空间和使用体验。

## Workflow

### 1. 搜 GitHub 候选

- 使用 GitHub 搜索候选项目
- 看许可证是否允许分发和修改
- 看活跃度、issue 状态、README 完整度、部署复杂度
- 记录候选仓库名、上游地址、许可证、核心功能

### 2. 查 Lazycat App Store 是否重复

- 在 `https://appstore.ezer.heiyu.space/#/shop` 中搜索候选项目名、中文译名、英文别名、核心关键词
- 当前 App Store 搜索通常需要登录态；如果未登录，不要假装已完成查重
- 如果搜索到已有同类移植，记录商品名、功能重叠度和差异点
- 如果重叠度高且无明显差异化，不继续按激励移植

### 3. 判断激励与可行性

- 是否属于官方不奖励类型
- 是否属于原创、首发移植，还是重复移植
- 是否有额外激励机会，例如 OIDC 或 `file_handler`
- 是否能让普通用户获得凭证

### 4. 固化移植仓库入口

移植仓库至少补齐：

- `docs/requirements/`
- `docs/api-design/`
- `docs/architecture/`
- `docs/release-prep/`
- `build.sh`
- `Makefile`

命令基线至少包括：

- `make build`
- `make install`

### 5. 规划 Lazycat 适配点

- `lzc-build.yml`
- `lzc-manifest.yml`
- `build.sh`
- `Makefile`
- 如果适合，补 `application.oidc_redirect_path`
- 如果是工具类，补 `application.file_handler`

### 6. 交回发布链路

确定“值得移植”后，再交回 `lazycat:ship-app` 做打包、资料、提审和发布。

复杂任务先读 [references/market-research.md](./references/market-research.md)、[references/porting-checklist.md](./references/porting-checklist.md) 和 [references/command-conventions.md](./references/command-conventions.md)。

## Quality Gates

- GitHub 搜索已完成
- App Store 查重已完成，且不是凭空假设
- 如有重复移植，已明确差异化或已终止激励路径
- 上游地址、许可证、仓库状态已记录
- 已规划 `build.sh`、`Makefile`、`make build`、`make install`
- 已评估 OIDC 或 `file_handler`

## Red Flags

- 只看 GitHub，不查 App Store
- App Store 未登录，却假装完成了查重
- 已有重复移植，还继续按激励路径推进
- 不保留上游地址和许可证
- 没有 `build.sh` 或 `Makefile`
- 明明是工具类，却没评估文件关联

## Bundled References

- GitHub 搜索与 App Store 查重： [references/market-research.md](./references/market-research.md)
- 移植检查清单： [references/porting-checklist.md](./references/porting-checklist.md)
- 命令入口约定： [references/command-conventions.md](./references/command-conventions.md)
- 激励规则： [../lazycat:ship-app/references/cash-incentive.md](../lazycat:ship-app/references/cash-incentive.md)

## Outputs

```text
阶段: 移植评估 / 去重 / 落地准备
目标: <普通移植 / 现金激励优先>

GitHub 候选
- ...

App Store 查重
- ...

结论
- 继续移植 / 换项目 / 非激励路径继续

落地要求
- build.sh
- Makefile
- make build
- make install
- 上游地址

下一步
1. ...
2. ...
```
