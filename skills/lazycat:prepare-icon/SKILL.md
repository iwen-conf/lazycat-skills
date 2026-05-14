---
name: lazycat:prepare-icon
description: 面向 Lazycat 应用图标准备与外部 AI 图像生成交接的 skill。只要用户提到 app icon、应用图标、商店 icon、PNG 图标、给其他 AI 生成图标、Apple 风格图标、iOS/macOS 图标、图标提示词、图标 brief 等请求，就必须使用此 skill。负责从项目名称和项目功能出发，整理图标语义、检查是否已到素材阶段，并输出一份可直接交给外部图像模型生成 1024x1024 PNG 的英文 prompt。
---

# Lazycat Project Icon Handover

You are responsible for preparing App Icon generation handovers during the asset phase of a Lazycat app. You do not claim the image is generated; instead, you organize the icon semantics, design constraints, and final English prompt for the user to use with other image models to generate a 1024x1024 PNG.

## Overview

This skill is used when a project enters the "Store Assets / Submission Assets / Icon Completion" phase. The goal is not to discuss design in broad terms, but to output a copyable prompt that aligns with the project's actual functionality, store positioning, and review context.

## Quick Contract

- **Trigger**: User mentions app icon, store icon, Apple-style icon, icon prompt, letting other AIs generate PNGs, or when a project is in the asset prep phase but lacks a formal icon.
- **Inputs**: Project name, project function, industry or workflow, current asset stage, presence of an old icon.
- **Outputs**: Icon semantic summary, filled English prompt, handover instructions, and technical constraints for post-generation verification.
- **Quality Gate**: The prompt must include the actual project name and function, and retain requirements like 1024x1024, no rounded corners, no text, no transparency, and Apple App Store aesthetics.
- **Decision Tree**: Determine if the project is in the asset phase, then decide if it's a new icon, an upgrade, or just a prompt export.

## When to Use

**Primary Triggers**

- User explicitly asks for an app icon, store icon, PNG icon, or icon prompt.
- `lazycat:ship-app` reaches the asset prep phase and finds the icon missing, of poor quality, or requiring AI generation.
- User says "Give me a prompt, I'll have another AI draw it."

**Typical Scenarios**

- A new project has a confirmed name and function, needing a formal 1024x1024 App Icon.
- An old icon is too rough and needs an upgrade to Apple App Store style.
- Project is ready for submission but lacks icon assets; the prompt needs to be solidified first.

**Boundary Notes**

- This skill handles icon prompts and handover; it does not generate images.
- If the project name or function is not yet finalized, return to the initiation phase; do not generate a random prompt.
- If the user needs screenshots, taglines, or a full set of store assets, integrate with `lazycat:ship-app`.

## Announce

Upon execution, provide a brief summary of:

- Whether the icon prep phase has been reached.
- Whether it's a new icon, an upgrade, or just a prompt export.
- Where you are extracting the project name and function from.

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `project_name` | string | Recommended | Official project name (from repo, config, or user). |
| `project_function` | string | Recommended | Core function; must summarize management objects and industry workflow. |
| `asset_stage` | enum(`Not Ready`/`Asset Prep`/`Pre-submission`) | Optional | Determines if it's time to output the prompt formally. |
| `icon_status` | enum(`None`/`Upgrade Needed`/`Export Only`) | Optional | Current icon status. |
| `industry_hint` | string | Optional | Visual metaphors: e.g., warehousing, logistics, inventory, approval, automation, monitoring. |

## The Iron Law

1. Extract project name and function from the repo, README, or store description before deciding on the prompt; do not use placeholders.
2. Icons must serve real functionality; do not sacrifice product semantics for "high-end feel."
3. Do not omit technical requirements: `1024x1024`, square, NO rounded corners, NO text, NO transparency.
4. You deliver the prompt and handover instructions, not the PNG file.
5. If the project is not yet in the asset phase, you can provide a template for later use.

## Workflow

### 1. Confirm Asset Phase
Determine if the project is in asset prep, pre-submission, or store organization. If not, inform the user they can save the template for later.

### 2. Extract Project Semantics
Extract `project_name` and `project_function` from context. Focus on the management objects, industry scenarios, and core actions (e.g., organization, control, monitoring, flow, automation).

### 3. Select Icon Metaphors
Select 1 to 3 suitable visual metaphors (do not stack them all):
- Data nodes
- Database stacks
- Management dashboards
- Flow arrows
- Inventory boxes
- Logistics symbols
- Scanning devices
- Checklists
- Warehouse elements
- Documents
- Automation indicators

### 4. Output Handover Prompt
- Load [references/app-icon-prompt.md](./references/app-icon-prompt.md).
- Replace `<ProjectName>` and `<ProjectFunction>` with real content.
- Provide the complete English prompt for easy copying.

### 5. Post-generation Verification Reminder
Remind the user to check the generated PNG for:
- `1024x1024` dimensions.
- No rounded corners, text, letters, UI screenshots, or transparency.
- Actual expression of project functionality (not just a generic tech logo).
- Suitability as a professional SaaS App Store icon.

## Quality Gates

- `project_name` is not a placeholder.
- `project_function` clearly expresses what the project manages and solves.
- Prompt retains constraints: Apple App Store level, glass texture, soft gradients, centered composition.
- Prompt retains all technical limitations.
- Format is ready for copying.

## Red Flags

- Generating a prompt without knowing what the project does.
- Outputting a prompt with placeholders like `<ProjectName>`.
- Pursuing visual flair while losing industry or workflow metaphors.
- Forgetting `NO rounded corners`, `NO text`, `NO transparency`.
- Delivering without a post-generation verification reminder.

## Bundled References

- App Icon Prompt Template: [references/app-icon-prompt.md](./references/app-icon-prompt.md)

## Outputs

```text
Phase: Icon Preparation
Project: <ProjectName>
Function: <ProjectFunction>

Icon Semantics
- ...

Handover Instructions
- Copy the English prompt below to an image model to generate a PNG.
- Check dimensions, background, rounded corners, text, and semantics after generation.

App Icon Prompt
<Complete English prompt>
```
