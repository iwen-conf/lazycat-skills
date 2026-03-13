---
name: lazycat:prepare-icon
description: 面向 Lazycat 应用图标准备与外部 AI 图像生成交接的 skill。只要用户提到 app icon、应用图标、商店 icon、PNG 图标、给其他 AI 生成图标、Apple 风格图标、iOS/macOS 图标、图标提示词、图标 brief 等请求，就必须使用此 skill。负责从项目名称和项目功能出发，整理图标语义、检查是否已到素材阶段，并输出一份可直接交给外部图像模型生成 1024x1024 PNG 的英文 prompt。
---

# Lazycat 项目图标交接

你负责在 Lazycat 应用的资料阶段准备 App Icon 生成交接，不直接声称图片已经生成，而是把图标语义、设计约束和最终英文 prompt 组织好，交给用户拿去其他图像模型生成 PNG。

## Overview

这个 skill 用在项目进入“商店资料 / 提审资料 / 图标补齐”阶段时。它的目标不是泛泛聊设计，而是输出一份能直接复制的图标生成 prompt，并让这份 prompt 与项目真实功能、商店定位和审核语境保持一致。

## Quick Contract

- **Trigger**: 用户提到应用图标、App Icon、Apple 风格图标、商店 icon、图标 prompt、让其他 AI 生成 PNG，或当前项目已进入资料准备阶段但缺少正式图标
- **Inputs**: 项目名称、项目功能、应用所在行业或工作流、当前商店资料阶段、是否已有旧图标
- **Outputs**: 图标语义摘要、已填充的英文 prompt、交接说明，以及生成 PNG 后需要回收检查的技术约束
- **Quality Gate**: prompt 必须填入真实项目名和项目功能，并保留 1024x1024、无圆角、无文字、无透明、Apple App Store 质感等硬性要求
- **Decision Tree**: 先判断项目是否已进入资料阶段，再判断图标是全新创建、升级旧图标，还是仅导出 prompt 给外部模型

## When to Use

**首选触发**

- 用户直接说要做应用图标、商店 icon、PNG 图标或图标提示词
- `lazycat:ship-app` 走到资料准备阶段，发现 icon 缺失、质量不够或需要外部 AI 生成
- 用户明确说“给我一个 prompt，我去让别的 AI 画图”

**典型场景**

- 新项目已经确定名称和功能，需要一张正式的 1024x1024 App Icon
- 旧图标过于粗糙，需要升级为 Apple App Store 风格
- 项目准备提审，但还缺图标素材，需要先把 prompt 固化

**边界提示**

- 这个 skill 负责图标 prompt 和交接，不负责真正生成图片
- 如果项目名称或功能还没有收敛，先回到立项阶段补清楚，不要胡乱产出 prompt
- 如果用户需要的是截图、简介或整套商店资料，不要只停在图标，应该联动 `lazycat:ship-app`

## Announce

开始执行后，用一句短摘要说明：

- 现在是不是已经到图标准备阶段
- 当前图标是新建、升级还是仅导出 prompt
- 你将从哪里提取项目名称和项目功能

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `project_name` | string | 推荐 | 项目正式名称；优先从仓库、应用配置、商店资料或用户描述中获取 |
| `project_function` | string | 推荐 | 项目核心功能；必须能概括系统管理对象和行业工作流 |
| `asset_stage` | enum(`未到资料阶段`/`资料准备中`/`提审前补齐`) | 可选 | 判断现在是否应该正式输出图标 prompt |
| `icon_status` | enum(`无图标`/`旧图标待升级`/`仅导出 prompt`) | 可选 | 说明当前图标现状 |
| `industry_hint` | string | 可选 | 用于强化视觉隐喻，例如仓储、物流、库存、审批、自动化、监控 |

## The Iron Law

1. 先从仓库、README、应用简介、本地配置里提取项目名称和功能，再决定 prompt；不要空手向用户要两个占位词。
2. 图标必须服务真实功能，不要为了“高级感”脱离产品语义。
3. prompt 中的硬性技术要求不要丢：`1024x1024`、正方形、无圆角、无文字、无透明。
4. 你交付的是 prompt 和交接说明，不是假装已经生成好的 PNG。
5. 如果项目尚未进入资料阶段，可以提醒用户稍后再生成，但仍要保留一份可复用模板。

## Workflow

### 1. 确认是否已进入图标阶段

- 判断当前项目是否已经进入资料准备、提审前补齐或商店资料整理阶段
- 如果还没到这个阶段，明确告诉用户“可以先保存模板，等资料阶段再正式出图”

### 2. 提取项目语义

- 从项目配置、README、应用简介或用户上下文提取 `project_name`
- 用一句话概括 `project_function`，重点写清楚系统管理对象、行业场景和核心动作
- 如果产品是系统型应用，优先强调组织、控制、监控、流程、自动化这类语义

### 3. 选择图标隐喻

根据项目功能选择 1 到 3 个最合适的视觉隐喻，不要全部堆上去：

- data nodes
- database stacks
- management dashboards
- arrows showing flow
- inventory boxes
- logistics symbols
- scanning devices
- checklists
- warehouse elements
- documents
- automation indicators

### 4. 输出交接 prompt

- 读取 [references/app-icon-prompt.md](./references/app-icon-prompt.md)
- 把 `<ProjectName>` 和 `<ProjectFunction>` 替换为当前项目真实内容
- 直接把完整英文 prompt 发给用户，方便用户复制到其他图像模型

### 5. 回收检查提醒

在交付 prompt 的同时，提醒用户生成 PNG 后至少检查：

- 是否为 `1024x1024`
- 是否没有圆角、文字、字母、UI 截图、透明背景
- 是否真的表达了项目功能，而不只是一个泛科技 logo
- 是否适合作为专业 SaaS 产品的 App Store icon

## Quality Gates

- `project_name` 不是占位词，也不是模糊代号
- `project_function` 能清楚表达项目在管什么、服务谁、解决什么问题
- prompt 保留 Apple App Store 级别、玻璃质感、柔和渐变、居中构图等设计约束
- prompt 保留所有硬性技术限制
- 输出格式适合用户直接复制，不需要再次整理

## Red Flags

- 还不知道项目做什么，就急着出图标 prompt
- 产出的 prompt 没有填入真实项目名和功能，仍然保留占位符
- 一味追求炫光和质感，却没有任何行业或流程隐喻
- 忘记写 `NO rounded corners`、`NO text`、`NO transparency`
- 交付时没有提醒用户生成后还要做素材回收检查

## Bundled References

- 图标生成英文 prompt 模板： [references/app-icon-prompt.md](./references/app-icon-prompt.md)

## Outputs

```text
阶段: 图标准备
项目名: <ProjectName>
项目功能: <ProjectFunction>

图标语义
- ...

交接说明
- 把下面这段英文 prompt 直接发给外部图像模型生成 PNG
- 生成后回收检查尺寸、背景、圆角、文字和功能表达

App Icon Prompt
<完整英文 prompt>
```
