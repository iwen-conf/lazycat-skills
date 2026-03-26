---
name: lazycat:create-app
description: 面向 Lazycat 新项目创建和项目基线统一的 skill。只要用户提到从 0 创建懒猫应用、初始化项目、脚手架、项目标准、Go 后端、Vue、Element Plus、后台管理、Admin、管理台、登录、注册、JWT、access_token、refresh_token、无感刷新、认证改造、现金激励、对接微服账户系统、网盘右键菜单、需求分析文档、API 设计文档、原创应用如何融入 Lazycat 原生系统、AI 配置面板、模型配置、懒猫算力仓、AI 应用、AI 浏览器插件等请求，就必须使用此 skill。负责把项目创建阶段收敛到统一规范：第一步先建立 `docs/` 文档树，再落 Go + Vue + Element Plus，默认具备登录、注册、双 token 和无感刷新；普通业务型 Web 应用接 AI 时默认走 `BaseURL` 配置方案，只有明确做懒猫算力仓 / `AI应用` / AI 浏览器插件时才走官方 AI Pod 路线；若项目包含后台管理面，则进一步把它接到高质量管理 UI 质量链路，并在用户以现金激励为目标时优先满足官方激励门槛。
---

# Lazycat Project Creation Baseline

You are responsible for advancing a Lazycat project from a "concept" to a state with a "unified technical baseline, ready for further development, documentation, and release." The focus is not just on generating a directory structure, but on defining project standards, authentication baselines, integration methods for original apps, AI integration baselines, incentive eligibility paths, and future release compatibility—all at once.

## Overview

This skill is used for creating new projects or aligning existing projects with a unified baseline. The default standard is:

- First, create the `docs/` directory and split it into multiple subdirectories and Markdown documents for requirements analysis, API design, etc.
- Provide executable script entries, such as `build.sh` and `Makefile`.
- Use Go for the backend.
- Use Vue for the frontend.
- Use Element Plus for the UI.
- If the project includes an admin interface, a high-quality admin UI is required by default. Mature templates can be used as a starting point but must undergo business-specific customization.
- All projects must include login and registration by default.
- Authentication uses `access_token + refresh_token`.
- The frontend supports "silent refresh" (seamless token renewal) by default.
- Original applications must evaluate how to integrate with Lazycat's native capabilities rather than being isolated web pages.
- If the business is naturally suited for AI, a unified AI configuration page is reserved by default, containing at least `API BaseURL`, protocol, model fetching, model selection, and configuration saving.
- For standard business web applications, AI is treated as a component of business capability; do not automatically switch to the AI Pod route.
- Only when the user explicitly targets Lazycat Computing Power Cabin (AI Pod), `AI App`, or AI Browser Extensions should the official AI Pod route be evaluated.

If an existing repository already has stronger team constraints, follow those; otherwise, implement this default standard.

If the user's goal includes Lazycat Cash Incentives, further align with official rules:

- Prioritize original applications with real-world scenarios.
- Originality is not just about "whether someone else has posted it," but also why it is suitable for the Lazycat native environment.
- Avoid app types explicitly excluded from official incentives.
- Applications requiring accounts must ensure regular users can obtain credentials.
- Prioritize integration with the Microservice Account System or Disk File Association if applicable.

## Quick Contract

- **Trigger**: User mentions creating a Lazycat project, initializing scaffolds, unifying project standards, adding login/registration, integrating dual tokens, implementing silent refresh, Go + Vue + Element Plus baseline, admin interfaces, cash incentives, Microservice Account System, Disk Context Menu, making original apps more "native," AI configuration pages, model configuration, Lazycat Computing Power Cabin, AI Apps, AI Browser Extensions.
- **Inputs**: Project goals, current repository state, whether it's a new project, existing stack, authentication status, existing user system, presence of an admin interface, incentive targets, originality, planned native integrations, AI suitability, AI Pod suitability.
- **Outputs**: Documentation blueprint, command entry blueprint, project baseline summary, stack and directory recommendations, authentication scheme, native integration strategy, AI configuration baseline, AI Pod integration judgment, admin UI baseline, incentive eligibility path, mandatory module list, and release preparation interfaces for `lazycat:ship-app`.
- **Quality Gate**: The final solution must first define the `docs/` directory structure and split documents, `build.sh` and `Makefile` entries, then specify Go backend, Vue + Element Plus frontend, login/registration, `access_token + refresh_token`, and silent refresh. Original apps must define native integration points. AI-suitable projects must define a unified AI config panel. AI Pod judgment is only required for specific targets. Admin interfaces must define high-quality UI routes. Incentive modes must address credential accessibility and compliance.
- **Decision Tree**: Determine if it's a new project or an update, then decide on a full baseline, authentication only, or structure alignment. Further evaluate native integration, AI config, AI Pod paths, or admin UI quality.

## When to Use

**Primary Triggers**

- User wants to create a Lazycat project from scratch.
- User specifies Go + Vue + Element Plus for the project.
- User requires admin consoles or dashboards to follow unified baselines.
- User requires login, registration, dual tokens, and silent refresh for all projects.
- User wants the app to meet Lazycat Cash Incentive thresholds.
- User wants original apps to feel like Lazycat native apps, not just wrapped websites.
- User wants to standardize AI integration and model configuration.
- User targets Lazycat Computing Power Cabin, AI App ecosystem, or AI Browser Extensions.
- Existing projects need unified authentication or stack alignment.

**Typical Scenarios**

- Starting a new app for the Lazycat Store, requiring technical baseline definition.
- Existing project has features but messy authentication, needing unified login/registration + dual tokens.
- Frontend has Vue but lacks Element Plus, Pinia, route guards, or silent refresh.
- Backend has Go but lacks refresh tokens, token rotation, or unified auth APIs.
- Project has admin features but lacks high-quality UI standards.
- Tool apps want to integrate with Disk Context Menus.
- Apps want to integrate with the Microservice Account System for better UX and incentives.
- Team is building an original app but hasn't answered "why it should live in Lazycat."
- Project is adding AI but lacks a standardized `BaseURL / Protocol / Model` config.
- Project is an AI product but needs to decide between a standard app, `AI App`, or AI Browser Extension.

**Boundary Notes**

- This skill handles "creation and baseline," not the final listing; use `lazycat:ship-app` for submission.
- Do not use this skill if the user only wants icons or store metadata.
- If the goal is just to make an existing admin UI look better for screenshots, switch to `lazycat:admin-ui`.
- If the existing project uses a different stack, do not rewrite it unless explicitly authorized.

## Announce

Upon execution, provide a brief summary of:

- Whether this is a new project or an update.
- Which `docs/` subdirectories you will create or complete.
- How you will verify the implementation of Go / Vue / Element Plus.
- Whether the current auth gap is in pages, APIs, token mechanisms, or silent refresh.
- Which Lazycat native capabilities the original app will integrate with.
- Whether a unified AI config panel is needed and which fields will be prioritized.
- Whether AI capabilities fit better as a standard app, `AI App`, or AI Browser Extension.
- For admin interfaces, the strategy for UI quality and template customization.
- For incentive goals, which eligibility path you are prioritizing.

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `project_mode` | enum(`New Project`/`Update Existing`/`Auth Only`) | Recommended | Determines if it's a full init or a structural reinforcement. |
| `backend_stack` | string | Recommended | Defaults to Go; do not deviate unless requested. |
| `frontend_stack` | string | Recommended | Defaults to Vue + Element Plus; prioritize upgrading existing Vue projects. |
| `auth_state` | enum(`None`/`Login Only`/`Single Token`/`Incomplete Dual Token`/`Complete`) | Recommended | Determines the depth of auth refactoring. |
| `user_system_state` | enum(`No User Table`/`Has User Table`/`Third-party Auth`) | Optional | Determines registration flow and minimal user fields. |
| `admin_surface_state` | enum(`No Admin`/`Needs Design`/`Needs Upgrade`) | Optional | Determines if admin UI quality path is needed. |
| `product_origin` | enum(`Original`/`Ported`/`Hybrid`) | Recommended | Determines if "Native Integration" evaluation is mandatory. |
| `native_fit_state` | enum(`Not Evaluated`/`OIDC Only`/`File Assoc Only`/`Multiple Points`) | Optional | Current state of Lazycat native integration. |
| `ai_surface_state` | enum(`No AI`/`Needs Design`/`Needs Unification`/`Complete`) | Optional | Determines if unified AI config and model discovery are needed. |
| `aipod_fit_state` | enum(`Not Evaluated`/`Standard App Best`/`AI App Candidate`/`AI Browser Ext Candidate`/`Pod Integrated`) | Optional | AI Pod ecosystem fit. |
| `release_target` | enum(`Dev First`/`Submit Ready`/`Release Baseline`) | Optional | Syncs release-oriented metadata and configuration constraints. |
| `reward_target` | enum(`Standard`/`Incentive Priority`) | Optional | If incentive-focused, avoid non-reward types and prioritize account/file integration. |
| `docs_state` | enum(`No Docs`/`Incomplete`/`Complete`) | Optional | Determines if directory creation is needed before coding. |

## The Iron Law

1. All new projects default to Go backend, Vue frontend, and Element Plus UI; do not switch without evidence.
2. All projects must have login and registration; if registration is not needed, clarify it as a business exemption.
3. Authentication defaults to `access_token + refresh_token`; a single long-lived token is insufficient.
4. Frontend must implement silent refresh and failure fallback; manual re-login after token expiry is not the end state.
5. Consider future listing and submission during creation to avoid re-work on auth, menus, routes, or env vars.
6. For cash incentives, prioritize original, real-world, stable apps; avoid app types outside the reward scope.
7. Apps requiring accounts must ensure users can register, log in via Microservice OIDC, or get public test credentials.
8. Establish the `docs/` tree before coding; do not leave requirements and API design for later.
9. Projects must provide executable entry points; at minimum `build.sh` and `Makefile`.
10. Admin interfaces must use `lazycat:admin-ui` standards; templates are scaffolds, not final deliverables.
11. Original apps must justify their Lazycat native value; evaluate OIDC, file association, and local workflows.
12. If AI scenarios exist, provide a unified AI config panel; minimal fields: `API BaseURL`, protocol type, fetch models, model selector, save config.
13. Standard business web apps use the `BaseURL` scheme; do not automatically require `ai-pod-service`, `caddy-aipod`, or `extension.zip`.
14. AI Pod routes are only evaluated when explicitly targeting AI Pod, `AI App`, or AI Browser Extensions.

## Workflow

### 1. Confirm Creation Mode and Metadata (LPK v2)
Determine if it's a new project, an update, or auth-only. Assess existing stack and structure. Identify if it's original/ported and if AI is needed. **Establish `package.yml` early** with the unique package ID, version, and name.

### 2. Establish `docs/` Tree
Check and create `docs/` before coding. Required subdirectories with `md` files:
- `docs/requirements/`
- `docs/api-design/`
- `docs/architecture/`
- `docs/release-prep/`

### 3. Command Entries and Manifest Configuration
All projects must include:
- `build.sh`
- `Makefile` with standard targets: `build`, `install`, `verify`, `release-prep`.
- `lzc-manifest.yml` for runtime services (static metadata belongs in `package.yml`).
**Mandatory: No submission without `make install` and real functional verification.**

### 4. Incentive and Originality Path
For incentive goals, determine originality and native value. Avoid excluded app types. Prioritize account/file integration for tools and OIDC for existing account systems.

### 5. Native System Integration
For original apps, answer "why install this in Lazycat?" Prioritize OIDC, `file_handler`, and local workflows. Mark weak integration points as risks.

### 6. Solidify Technical Baseline
Backend: Go. Frontend: Vue + Element Plus. Unified state management and auth in a single store. Define anonymous vs. authenticated route boundaries.

### 7. Implement Auth Baseline
Required: Login/Register pages, state persistence, `access_token`, `refresh_token`, silent refresh, and failure fallback to login. Recommended: Short-lived access token, long-lived refresh token with rotation.

### 8. Design Minimal APIs and Pages
APIs: `POST /auth/register`, `/login`, `/refresh`, `/logout`, `GET /auth/me`.
Pages: Login/Register forms, route guards, state recovery, 401 retry, OIDC integration (if applicable), file open entry (for tools).

### 9. AI Capability and Config Panel
If AI is involved: Default to `BaseURL` for standard web apps. AI must be a core process. Config panel: `API BaseURL`, protocol (OpenAI/Anthropic), model fetching, model selector, save button. Use standard templates from [references/ai-settings-template.md](./references/ai-settings-template.md).

### 10. Admin UI Convergence
If an admin interface exists: Define dashboard, list, detail, form, and settings views. Use `lazycat:admin-ui` for quality convergence. Admin UI must match the main brand direction.

### 11. Data, Security, and Additional Integration
Define minimal user fields. Refresh tokens must be revocable or rotatable. Define auth env vars and secrets. Plan `application.oidc_redirect_path` and `application.file_handler` in the manifest.

### 12. Hand over to Release Pipeline
Once the baseline (dev-ready, login, dual token) is established, hand over to `lazycat:ship-app`. If admin UI is present, hand over to `lazycat:admin-ui` first.

## Quality Gates

- `docs/` tree established with content in requirements, API design, etc.
- `build.sh` and `Makefile` created.
- Go + Vue + Element Plus implemented.
- Login and registration present.
- `access_token + refresh_token` implemented.
- Silent refresh and fallback flow verified.
- Original apps have clear native integration points.
- Minimal APIs, pages, and user data defined.
- AI-suitable projects have a unified config panel.
- AI products have a defined route (Standard, AI App, or Browser Ext).
- Admin interfaces have a quality convergence path.
- Incentive goals are checked against official rules.
- Test credentials provided for account-based apps.
- OIDC or file association planned if applicable.

## Red Flags

- Developing features without a unified auth baseline.
- Login without registration (without exemption).
- Single token without refresh token.
- 401 kicks user out without silent refresh.
- Concurrent requests trigger multiple refreshes (race conditions).
- Coding without `docs/requirements` or `docs/api-design`.
- Only a vague `README.md` without split docs.
- No `build.sh`, `make build`, or `make install`.
- Admin interface present but UI quality path ignored.
- Targeting incentives with excluded app types.
- Account-based app without registration or test credentials.
- Tool app without file association; account app without OIDC evaluation.

## Bundled References

- Project Stack and Auth Baseline: [references/project-baseline.md](./references/project-baseline.md)
- Documentation Blueprint: [references/docs-blueprint.md](./references/docs-blueprint.md)
- AI Settings Template: [references/ai-settings-template.md](./references/ai-settings-template.md)
- AI Pod / AI App Decision: [references/aipod-playbook.md](./references/aipod-playbook.md)
- Admin UI Quality: [../lazycat:admin-ui/SKILL.md](../lazycat:admin-ui/SKILL.md)
- Command Conventions: [../lazycat:port-app/references/command-conventions.md](../lazycat:port-app/references/command-conventions.md)
- Incentive Eligibility: [../lazycat:ship-app/references/cash-incentive.md](../lazycat:ship-app/references/cash-incentive.md)

## Outputs

```text
Phase: Project Creation / Baseline Update
Mode: <New Project / Update / Auth Only>

Confirmed
- ...

Docs Tree
- docs/requirements
- docs/api-design
- docs/architecture
- docs/release-prep

Command Entries
- build.sh
- Makefile
- make build
- make install

Project Baseline
- Backend: Go
- Frontend: Vue + Element Plus
- Auth: access_token + refresh_token + Silent Refresh
- Admin UI: <N/A / lazycat:admin-ui / High Quality>

Incentive Path
- Target: <Standard / Incentive Priority>
- Type: <Original / Ported / Not Recommended>
- Integration: <OIDC / file_handler / None>

Gaps / Risks
- ...

Current Action
- ...

Next Steps
1. ...
2. ...

Deliverables
- Project structure recommendations
- Auth chain checklist
- Incentive eligibility path
- Entry point for lazycat:ship-app
```
