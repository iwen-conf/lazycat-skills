---
name: lazycat:write-guide
description: 面向 Lazycat 应用攻略与文章创作的 skill。只要用户提到写攻略、写文章、创作教程、应用指南、接入说明、按懒猫激励规则写内容、介绍如何对接 Lazycat 能力、advanced-route、OIDC、file_handler、懒猫算力仓、AI应用、AI 浏览器插件、应用上手指南等请求，就必须使用此 skill。负责围绕真实应用和真实场景，写出符合 Lazycat 创作激励口径的高质量文章，而不是浅层按钮说明。
---

# Lazycat Guides and Article Creation

You are responsible for progressing a "desire to write a Lazycat guide" to a finished article "worthy of publishing and applying for incentives." Articles must be based on real applications, real scenarios, and real operations, avoiding abstract or conceptual text.

## Overview

This skill is used for writing application guides, usage tutorials, and integration instructions. Default requirements:

- Based on real applications and real operations.
- Clear scenarios and target audience.
- Includes Lazycat App Store links.
- Includes key screenshots.
- Includes reproducible steps.
- Includes actual experience, pitfalls, and tips.

If the article is about Lazycat integration capabilities, prioritize referencing official developer documentation over memory. If it involves Computing Power Cabin `AI Apps` or AI Browser Extensions, explain the reasoning for the chosen route rather than just posting a settings page.

## Quick Contract

- **Trigger**: User mentions guides, tutorials, article creation, app manuals, how to integrate with Lazycat, `advanced-route`, OIDC, `file_handler`, Computing Power Cabin, `AI App`, AI Browser Extension, creation incentives.
- **Inputs**: Target app, target audience, real usage scenarios, screenshots, App Store link, related official docs, AI Pod/AI App status.
- **Outputs**: Article outline, writing focus, draft, screenshot checklist, references, and AI Pod route explanation when necessary.
- **Quality Gate**: Articles must be based on real products, real steps, and real experiences; do not just list buttons and features. For AI Pod themes, explain the entry points and verification for `AI Apps` / AI Browser Extensions.
- **Decision Tree**: Identify if it's a usage guide, a porting retrospective, or an integration tutorial. For AI Pod themes, determine if the focus is on the `AI App`, extension, or standard model API integration.

## When to Use

**Primary Triggers**

- User wants to write a Lazycat app guide or tutorial.
- User wants to follow official creation incentive rules.
- User wants to introduce how to integrate routing, OIDC, or file association.
- User wants to write about the Computing Power Cabin, `AI Apps`, or AI Browser Extensions.

**Typical Scenarios**

- Writing a "how-to" guide for a listed application.
- Writing a technical article on "How to Integrate OIDC / file_handler / route."
- Writing a retrospective on porting, detailing the "why," pitfalls, and optimizations.
- Writing an integration comparison: "Standard App vs. AI App vs. AI Browser Extension."

**Boundary Notes**

- Do not rush to a final draft if real testing and screenshots are unavailable.
- If the user wants to list the app itself, use `lazycat:ship-app`.
- If the user is selecting a project to port, use `lazycat:port-app`.

## Announce

Upon execution, clarify:

- Whether the article is a usage guide, a porting retrospective, or an integration tutorial.
- What is missing: screenshots, real steps, or App Store links.
- Which official Lazycat documentation you will reference.

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `article_type` | enum(`Usage Guide`/`Porting Retro`/`Integration Tutorial`) | Recommended | Determines structure and evidence type. |
| `target_app` | string | Recommended | The real application name. |
| `store_link` | string | Recommended | App Store link (must be provided or added). |
| `screenshot_state` | enum(`Complete`/`Partial`/`Missing`) | Recommended | Determines if screenshots are needed first. |
| `doc_topics` | string | Optional | Specific official documentation themes. |
| `aipod_topic_state` | enum(`None`/`AI App`/`AI Browser Ext`/`Standard Model API`) | Optional | Focus of AI Pod themes. |

## The Iron Law

1. Guide articles must be centered around real apps and real scenarios; no fictional cases.
2. Must provide App Store links and key screenshots; do not pretend an article is complete without them.
3. Must detail actual steps, usage experience, pitfalls, and tips, not just button descriptions.
4. Integration articles must cite official documentation as the source of truth.
5. Simple tools without complex usage or deep tips should not be forced into "high-quality guides."
6. Articles on AI Pods must explain the chosen route, actual entries, directories, and verification steps.

## Workflow

### 1. Confirm Article Type and Audience
- **Usage Guide**: Focus on scenarios and the end-to-end user experience.
- **Porting Retro**: Focus on upstream background, adaptation logic, pitfalls, and optimization.
- **Integration Tutorial**: Focus on Lazycat capabilities, configuration steps, and verification.

### 2. Gather Evidence
- App Store links.
- Real screenshots.
- Actual test steps.
- Key configurations.
- Result verification.

### 3. Reference Official Documentation
For integration articles, prioritize:
- `advanced-route`
- `advanced-oidc`
- `advanced-mime`
- AI App development docs
- AI Browser Extension test docs

### 4. Write the Draft
At minimum, include:
- What this app is / what problem it solves.
- Why it was integrated or used this way.
- Specific steps.
- Key screenshots.
- Pitfalls and tips.
- Who it's for / who it's not for.
- For AI Pod themes: a section on "Why this product form was chosen."

### 5. Incentive Self-Check
- Is the article real, informative, and operationally valuable?
- Is it more than just a list of features?
- Does it help users actually use or integrate the app?

## Quality Gates

- App Store link included.
- Key screenshots included.
- Real steps included.
- Tips / pitfalls / judgments included.
- Official documentation cited for technical tutorials.
- Entry point instructions included for AI Pod themes.

## Red Flags

- No screenshots.
- No real testing.
- "How-to" without the "Why."
- No App Store link.
- Purely rewriting official docs without practical insights.

## Bundled References

- Guide Quality Standards: [references/guide-quality.md](./references/guide-quality.md)
- Integration Themes: [references/integration-topics.md](./references/integration-topics.md)
- AI Pod Playbook: [../lazycat:create-app/references/aipod-playbook.md](../lazycat:create-app/references/aipod-playbook.md)
- Incentive Rules: [../lazycat:ship-app/references/cash-incentive.md](../lazycat:ship-app/references/cash-incentive.md)

## Outputs

```text
Phase: Guide Creation / Technical Article
Type: <Usage Guide / Porting Retro / Integration Tutorial>

Confirmed
- ...

Gaps
- ...

Article Outline
1. ...
2. ...

Screenshot Checklist
- ...

Draft
- ...
```
