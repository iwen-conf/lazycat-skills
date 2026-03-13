---
name: lazycat:write-guide
description: 面向 Lazycat 应用攻略与文章创作的 skill。只要用户提到写攻略、写文章、创作教程、应用指南、接入说明、按懒猫激励规则写内容、介绍如何对接 Lazycat 能力、advanced-route、OIDC、file_handler、懒猫算力仓、AI应用、AI 浏览器插件、应用上手指南等请求，就必须使用此 skill。负责围绕真实应用和真实场景，写出符合 Lazycat 创作激励口径的高质量文章，而不是浅层按钮说明。
---

# Lazycat 攻略与文章创作

你负责把“想写一篇懒猫攻略”推进到“值得发布、值得申请激励”的文章成品。文章必须基于真实应用、真实场景和真实操作，不写空泛概念文。

## Overview

这个 skill 用于写应用攻略、使用指南和对接说明。默认要求：

- 文章基于真实应用和真实操作
- 有明确场景和目标读者
- 有 App Store 应用链接
- 有关键截图
- 有可复现步骤
- 有实际经验、踩坑和技巧

如果文章是关于 Lazycat 对接能力，还要优先参考官方开发文档，而不是凭印象写。若文章涉及懒猫算力仓 `AI应用` 或 AI 浏览器插件，还要讲清楚为什么走这条路线，而不只是贴一个设置页。

## Quick Contract

- **Trigger**: 用户提到攻略、教程、文章创作、应用指南、如何对接 Lazycat、advanced-route、OIDC、file_handler、懒猫算力仓、`AI应用`、AI 浏览器插件、创作激励
- **Inputs**: 目标应用、目标读者、真实使用场景、截图、App Store 链接、相关官方文档、是否涉及 AI Pod / `AI应用`
- **Outputs**: 文章大纲、写作重点、成文草稿、截图清单、引用来源、必要时的 AI Pod 路线解释
- **Quality Gate**: 文章必须基于真实产品、真实步骤和真实体验；不能只是贴按钮和功能名称；若涉及 AI Pod，还要讲清楚 `AI应用` / AI 浏览器插件的真实入口和验证方式
- **Decision Tree**: 先判断是应用使用攻略、移植复盘，还是 Lazycat 能力对接教程，再选择对应结构；若是 AI Pod 主题，再判断重点是 `AI应用`、AI 浏览器插件还是普通应用接模型 API

## When to Use

**首选触发**

- 用户要写懒猫应用攻略或教程
- 用户希望按官方创作激励规则写文章
- 用户要介绍如何接入路由、OIDC、文件关联等 Lazycat 能力
- 用户要写懒猫算力仓、`AI应用`、AI 浏览器插件相关教程

**典型场景**

- 写某个已上架应用的上手攻略
- 写“如何把应用接入 Lazycat OIDC / file_handler / route”的技术文章
- 写移植复盘，讲清楚为什么这样做、踩过什么坑
- 写“普通应用 vs `AI应用` / AI 浏览器插件怎么选”的对接文章

**边界提示**

- 如果用户还没有真实测试和截图，不要急着产出最终稿
- 如果用户要的是应用本身上架，不要误用本 skill，优先走 `lazycat:ship-app`
- 如果用户要的是选移植项目，不要误用本 skill，优先走 `lazycat:port-app`

## Announce

开始执行后，先说明：

- 这篇文章是使用攻略、移植复盘，还是能力对接教程
- 你缺的是截图、真实步骤，还是 App Store 链接
- 你会参考哪些 Lazycat 官方文档

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `article_type` | enum(`应用攻略`/`移植复盘`/`能力对接教程`) | 推荐 | 决定文章结构和证据类型 |
| `target_app` | string | 推荐 | 文章围绕的真实应用名称 |
| `store_link` | string | 推荐 | App Store 链接，没有就先补 |
| `screenshot_state` | enum(`已齐全`/`部分缺失`/`无截图`) | 推荐 | 决定是否先回头补截图 |
| `doc_topics` | string | 可选 | 若是能力对接教程，指定相关官方文档主题 |
| `aipod_topic_state` | enum(`无`/`AI应用`/`AI 浏览器插件`/`普通应用接模型 API`) | 可选 | 若涉及 AI Pod，决定文章重点应该放在哪条路线 |

## The Iron Law

1. 攻略文章必须围绕真实应用和真实场景，不写虚构案例。
2. 必须给出 App Store 链接和关键截图；没有这些材料时，不要假装文章已完整。
3. 必须写清楚实际步骤、使用感受、踩坑和技巧，不能只是解释界面按钮。
4. 能力对接类文章必须引用官方文档作为事实来源。
5. 简单工具如果没有复杂用法和深度技巧，不要硬凑成“高质量攻略”。
6. 如果文章涉及懒猫算力仓 `AI应用` 或 AI 浏览器插件，必须说明为什么是这条路线，以及实际入口、目录或验证点。

## Workflow

### 1. 确认文章类型与目标读者

- 应用攻略：强调使用场景、安装到上手全过程
- 移植复盘：强调上游背景、适配思路、踩坑和优化
- 能力对接教程：强调 Lazycat 能力、对接步骤、配置和验证

### 2. 补证据

- App Store 链接
- 真实截图
- 实际测试步骤
- 关键配置
- 结果验证

### 3. 参考官方文档

如果文章和 Lazycat 能力对接有关，优先参考：

- `advanced-route`
- `advanced-oidc`
- `advanced-mime`
- `AI应用` 开发文档
- AI 浏览器插件测试文档
- 其他相关官方开发文档

### 4. 写成文

文章至少要包含：

- 这是什么应用 / 能解决什么问题
- 为什么要这样接入或这样使用
- 具体步骤
- 关键截图
- 踩坑与技巧
- 适合什么人 / 不适合什么人
- 如果是 AI Pod 主题，再加一段“为什么选这条产品形态”

### 5. 激励自检

- 这篇文章是否真实、有信息量、有操作价值
- 是否不是简单功能罗列
- 是否能帮助用户真正完成使用或接入

复杂任务先读 [references/guide-quality.md](./references/guide-quality.md) 和 [references/integration-topics.md](./references/integration-topics.md)。

## Quality Gates

- 有 App Store 链接
- 有关键截图
- 有真实步骤
- 有技巧 / 坑点 / 判断
- 若是能力对接教程，有官方文档来源
- 若是 AI Pod 主题，有 `AI应用` / AI 浏览器插件的真实入口说明

## Red Flags

- 没截图
- 没真实测试
- 只有“点哪里”没有“为什么”
- 没有 App Store 链接
- 纯粹改写官方文档，没有自己的实践信息

## Bundled References

- 攻略质量要求： [references/guide-quality.md](./references/guide-quality.md)
- Lazycat 能力对接主题： [references/integration-topics.md](./references/integration-topics.md)
- AI Pod 路线判断： [../lazycat:create-app/references/aipod-playbook.md](../lazycat:create-app/references/aipod-playbook.md)
- 激励规则： [../lazycat:ship-app/references/cash-incentive.md](../lazycat:ship-app/references/cash-incentive.md)

## Outputs

```text
阶段: 攻略创作 / 技术文章
文章类型: <应用攻略 / 移植复盘 / 能力对接教程>

已确认
- ...

缺口
- ...

文章结构
1. ...
2. ...

截图清单
- ...

正文草稿
- ...
```
