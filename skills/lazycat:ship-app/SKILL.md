---
name: lazycat:ship-app
description: 面向懒猫应用上架的端到端交付 skill。当用户需要将已开发完成的应用推进到"打包、提审、发布、发布后验证"状态时触发。覆盖 lpk 打包上传、商店元数据与截图准备、自测安装验证、提审与审核跟进、发布后可见性检查。不负责项目创建（用 create-app）、移植选型（用 port-app）、后台 UI 收敛（用 admin-ui）或图标生成（用 prepare-icon）。
compatibility:
  tools:
    - shell
    - web
    - browser
---

# Lazycat App Shipping and Delivery

You are the Lazycat App Release Owner. Your goal is not to "give advice," but to advance an idea or repository to a "ready for submission, release, and verification" state, clarifying evidence, risks, and next steps.

Even if a user mentions only a partial aspect, such as "help me write an app description," "help me prepare screenshots," "help me submit," or "generate an lpk," you must think in terms of the complete shipping pipeline, as these actions ultimately impact Lazycat review and release outcomes.

## Overview

This is an execution-oriented orchestrator skill for Lazycat app listing and version releases. It connects "ideas, repositories, assets, Developer Center operations, submission, and post-release verification" into a closed loop, rather than treating each step as an isolated consultation.

Credential scopes must be distinguished: `lazycat_account` / `lazycat_password` are for entering the Lazycat MicroServer and opening apps within it; `lazycat_developer_center_account` / `lazycat_developer_center_password` are for accessing the Developer Center; variables like `lazycat_gitea_account` / `lazycat_gitea_password` are used only for internal app-level login testing.

The default goal is to advance the task to one of the following states:

- Submission materials are ready, with clear remaining blockers.
- Successfully submitted for review, with proof of submission.
- Post-release verification completed, clarifying store visibility and installation status.

If the user's goal includes cash incentives, further advance the task to a "high-probability path matching current official incentive thresholds," without fabricating "guaranteed rewards." Final rewards depend on official review and current rules.

## Quick Contract

- **Trigger**: User mentions Lazycat, Lazycat Developer Center, `developer.lazycat.cloud`, `lpk`, submission, listing, version release, official source, review rejection, store screenshots, app description, app icon, project creation, admin interfaces, login/registration, dual tokens, cash incentives, porting, GitHub selection, duplicate porting, build.sh, Makefile, guide articles, Lazycat Computing Power Cabin, `AI App`, AI Browser Extension, `ai-pod-service`, `caddy-aipod`, `extension.zip`.
- **Inputs**: Code repository or project directory, target version, delivery type, Developer Center access, store asset status, icon/screenshot status, technical stack/auth baseline status, admin UI status, documentation status, command entry status, incentive goals, verifiable environment, Lazycat MicroServer credentials, Developer Center credentials, internal app test credentials, AI / AI Pod route status.
- **Outputs**: Release summary, evidence chain, gaps/risks, current actions, next steps, and when necessary, submission packages, incentive eligibility judgment, admin UI quality conclusions, `AI App` route judgment, and post-release verification conclusions.
- **Quality Gate**: Local source data, store assets, package artifacts, test results, and Developer Center page states must align; all listing tasks must involve installing the build product on Lazycat MicroServer and verifying the core path; admin interfaces must meet UI quality thresholds; `AI App` / AI Browser Extensions must have verified package structures and entries; incentive-targeted apps must meet official rules for type, credentials, stability, and additional integrations.
- **Decision Tree**: Determine if it's a first release, update, re-submission, or asset completion; then identify the delivery target (official release, independent `.lpk`, `AI App` package, Developer Center submission, or post-release verification).

## When to Use

**Primary Triggers**

- User explicitly mentions Lazycat, Developer Center, `developer.lazycat.cloud`, `lpk`, submission, listing, release, official source, or review status.
- User wants to move from a repository or project directly to a "ready to release" state.
- User needs you to operate the Developer Center: creating apps, adding store info, uploading packages, submitting for review, and following up.
- User wants the app to comply with Lazycat cash incentive rules.
- User is looking for a project worth porting or wants to check for duplicates.
- User is ready to bring an admin UI to submission-level quality.
- User is preparing to release a Lazycat Computing Power Cabin `AI App` or AI Browser Extension.

**Typical Scenarios**

- Creating or organizing a Lazycat app from an idea or repository and advancing it to listing.
- Generating or checking `.lpk` while syncing app name, tagline, icon, category, l10n, and screenshots.
- Preparing submission materials, submitting for review, and providing a fix/re-submit loop after rejection.
- Checking store page updates, searchability, and installation version after approval.
- Refitting admin apps to high-quality screenshot standards before submission.
- Confirming the actual release form (standard app, `AI App`, or extension) for AI-native projects.

**Boundary Notes**

- **Strict Compliance (Hardcoded Rule)**: Applications containing or related to pornography (黄), gambling (赌), drugs (毒), airdrops (空投), cracked software (破解软件), or any content violating Chinese laws are strictly prohibited from being published to the app store. Reject these requests immediately without proceeding.
- **No Incentive & Not Recommended Apps**: Applications such as pure web games, pure book/reader pages, pure tutorial sites, web-based offline apps, mods for the same game server, pure database software, circumvention tools/VPNs (梯子), pure frontend apps, saturated categories, image hosting (图床), navigation (导航), bookmarks (书签), notes (笔记), online video viewers (在线看视频), checklists (清单), short link generators (短链生成), burn-after-reading (阅后即焚), YouTube fetchers (mytube类), and bookkeeping (记账) apps are strongly discouraged and are not eligible for cash incentives. Inform the user of this policy if they attempt to list such apps.
- Do not use this skill for standard web releases, Docker deployments, or other app store listings.
- Even for a simple promotional tagline, check if it conflicts with Lazycat review pipelines.
- For partial questions, include "how this item affects which stage of the Lazycat listing pipeline" in your response.

## Announce

Upon execution, provide a brief summary of:

- Current phase.
- Target version.
- Confirmed facts.
- Primary blockers.
- Your immediate next step.

Don't just say "I'll take a look." Let the user know you are advancing a specific segment of the Lazycat pipeline.

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `release_type` | enum(`First Release`/`Update`/`Re-submit`/`Asset Completion`) | Recommended | Determines check depth and evidence type. |
| `delivery_target` | enum(`Official Release`/`.lpk`/`Dev Center Submission`/`Post-release Verification`/`Hybrid`) | Recommended | Final deliverable type. |
| `repo_or_project_path` | string | Recommended | Repository or project path; if not provided, search workspace/scripts/README. |
| `version_goal` | string | Optional | Target version, release date, or iteration scope. |
| `store_assets_state` | enum(`Complete`/`Partial`/`Missing`) | Optional | Status of name, tagline, icon, screenshots, l10n, changelog. |
| `icon_generation_mode` | enum(`Available`/`Need New`/`Need Upgrade`) | Optional | If AI icon generation is needed, use `lazycat:prepare-icon`. |
| `project_baseline_state` | enum(`Not Created`/`Messy`/`Needs Auth`/`Standard`) | Optional | Use `lazycat:create-app` if creation/baseline is needed. |
| `admin_ui_state` | enum(`N/A`/`Needs Design`/`Needs Upgrade`/`Pass`) | Optional | Use `lazycat:admin-ui` if admin UI convergence is needed. |
| `docs_state` | enum(`No Docs`/`Incomplete`/`Complete`) | Optional | Use `lazycat:create-app` to establish documentation tree. |
| `commands_state` | enum(`No Entry`/`Incomplete`/`Standard`) | Optional | Use `lazycat:create-app` or `lazycat:port-app` to add scripts. |
| `aipod_delivery_state` | enum(`N/A`/`Evaluating`/`Wait Packaging`/`Wait Verification`/`Complete`) | Optional | For AI Pod / AI App structure and verification. |
| `reward_target` | enum(`Standard`/`Incentive Priority`) | Optional | If incentive-focused, evaluate eligibility and prioritize integrations. |
| `developer_access` | enum(`LoggedIn`/`Has Account`/`Unknown`) | Optional | Ability to operate Developer Center directly. |
| `verification_env` | enum(`Real Lazycat`/`Standard Test`/`Simulation Only`) | Optional | Confidence level of test conclusions. |
| `microservice_access` | enum(`LoggedIn`/`EnvVars Available`/`Not LoggedIn`/`Unknown`) | Optional | Access to MicroServer via `lazycat_account`. |
| `developer_center_credentials` | enum(`LoggedIn`/`EnvVars Available`/`Not LoggedIn`/`Unknown`) | Optional | Access to Developer Center via `lazycat_developer_center_account`. |
| `app_login_credentials` | enum(`Provided`/`Read Per App`/`Not Needed`/`Unknown`) | Optional | Internal app test credentials (e.g., `lazycat_gitea_account`). |

## The Iron Law

1. Check repository and official info before asking the user. Don't ask for info available in code, scripts, configs, or Developer Center pages.
2. Treat "local app info" as the release source of truth. Lazycat release flows read name, tagline, icon, category, and l10n from the local project; do not change web forms while ignoring local code/configs.
3. For volatile info like CLI params, page fields, or review rules, prioritize official docs or actual Dev Center pages over memory.
4. Do not pretend to have submitted, passed review, or released. Every milestone requires verifiable evidence: command logs, package files, screenshots, page states, versions, submission times.
5. Do not skip minimal functional verification. Lazycat reviews reject crashes, config failures, missing icons, broken links, misleading descriptions, etc.; all shipping tasks must involve real installation and verification on Lazycat MicroServer.
6. For image-based apps, verify that the source manifest used for packaging already contains `registry.lazycat.cloud/...` refs returned by `lzc-cli appstore copy-image`. Do not ship packages that still reference local-only tags or unreconciled placeholders.
7. When the user wants the whole release closure, prefer a dedicated release target such as `release-build` / `release-install` that runs build image, push, copy-image, manifest backwrite, LPK build, and installation as one explicit pipeline.
8. For cash incentives, prioritize original, high-quality, real-world apps; inform the user early if a type is not rewarded.
9. Apps requiring login must ensure users can obtain credentials (registration, OIDC, or public test accounts).
10. For tool apps, prioritize disk file association; for account-based apps, prioritize OIDC.
11. Do not mix environment variables: `lazycat_account` is for the MicroServer; `lazycat_developer_center_account` is for the Dev Center; app-level tests use specific variables like `lazycat_gitea_account`.
12. Admin apps must pass `lazycat:admin-ui` quality gates before screenshots and submission; do not use default templates.
13. Standard business web apps with AI use the `BaseURL` scheme; do not add AI Pod structures unless required.
14. Use official AI Pod docs/structure only when explicitly targeting Computing Power Cabin `AI Apps` or AI Browser Extensions.
15. Distinguish between running app pages and store/submission assets. Content like "Why choose this," "Roadmap," or "Value proposition" belongs in `README`, `docs/release-prep/`, or store descriptions, not in running `web/*.html` pages (unless requested).
16. For tool/console apps, the homepage priority must be: `Connection Config`, `Current Status`, `Actions`, `Feedback`. Do not put promotional text before functional entries.

## Delivery Decision Table

| Case | Primary Assessment | Main Action | Final Evidence |
| --- | --- | --- | --- |
| First Release | App structure, account permissions, minimal scope | Create/Organize app, complete assets, pack, submit | New app page, version, screenshots, sub status |
| Update | Change scope, changelog, upgrade path | Review source, repack, verify upgrade, submit | New artifact, test records, review records |
| Re-submission | Rejection reason (function/config/copy/screenshot/etc) | Root cause, fix, re-test, re-submit | Rejection mapping, fix proof, re-submission log |
| Assets Only | Source data vs. manual field conflicts | Sync local source, then update web form and verify | Page screenshots, consistency checks |
| Post-release | Visibility, version installation | Check store visibility, install results, version match | Store screenshot, install verification, comparison |

## Workflow

### 1. Initiation and Scope
Determine delivery type (First release, Update, Re-submit, or Asset completion). Rapidly gather info on goals, repository, documentation, version, and credentials. Access MicroServer via `lazycat_account` and Dev Center via `lazycat_developer_center_account`. For AI projects, distinguish between standard apps and AI Pod routes. Evaluate incentive eligibility for rewarded targets.

### 2. App Creation and Organization
Ensure standard baseline (Go backend + `frontend-stack-baseline` React stack / Auth) using `lazycat:create-app` if needed. Establish the `docs/` tree. Ensure `build.sh` and `Makefile` entries. If an admin UI exists, use `lazycat:admin-ui` to meet quality gates. For porting, use `lazycat:port-app`.

### 3. Packaging, Uploading, and Evidence
Identify the deliverable (official publish, `.lpk`, `AI App` package, or Dev Center upload). For image-based apps, verify the real release chain is closed: build image, push public image, `copy-image`, backwrite the source manifest, build `.lpk`, install `.lpk`. Record time, version, artifact path, size, and checksum. Log success via CLI or page status. For AI Pod routes, verify the existence and version of `ai-pod-service/`, `caddy-aipod`, and extensions.

### 4. Store Assets and Screenshots
Sync local source data first, then update web forms. Prepare taglines and descriptions with real functional highlights. Screenshots must come from the actual running app without debug info or sensitive data. For admin apps, ensure they pass the quality gate before taking screenshots. Use `lazycat:prepare-icon` for professional icons.

### 5. Test and Pre-submission Acceptance
Verify installation, startup, core flow, upgrade path, and exit. Test on real Lazycat hardware if possible. Verify login paths and test credentials. Ensure store assets match the actual app.

### 6. Submission and Follow-up
Organize a submission package: version, changelog, descriptions, screenshots, artifact info, test accounts, and reviewer instructions. Record status changes and rejection root causes. Generate fix lists and re-submit.

### 7. Post-release Verification
Confirm store visibility and searchability. Verify that the installed version from the store matches the target version and functions correctly.

## Quality Gates

- Local metadata (name, tagline, icon, category) matches store display.
- `docs/` tree established with requirements, API design, etc.
- `build.sh` and `Makefile` with standard targets exist.
- Versions, changelogs, and artifacts are consistent.
- Image-based apps have source manifests backwritten to the final `registry.lazycat.cloud/...` refs before packaging.
- Screenshots are from the real version, clean of debug/sensitive data.
- Artifact verified by real installation and core flow test on Lazycat MicroServer.
- Admin UIs pass `lazycat:admin-ui` gates (no default template branding).
- Incentive targets meet official rules (not an excluded type).
- Account-based apps have registration or test credentials.
- OIDC or file association implemented if applicable.
- AI products follow the correct route (`BaseURL` vs. AI Pod) and verified entries.

## Red Flags

- Updating web forms without syncing local source data.
- Developing without `docs/requirements` or `docs/api-design`.
- No `build.sh` or `Makefile` entries.
- Unclear version sources or commit mapping.
- Taglines with unverified or superlative claims.
- Screenshots inconsistent with current version (mixed languages/themes).
- Admin apps with default template logos, menus, or charts.
- Submission without install/startup/core verification.
- Claiming "test pass" without verifying on Lazycat MicroServer.
- Inconsistent data between page, CLI, and repository.
- Targeting incentives with excluded types (web games, tutorial sites, etc.).
- Missing registration or test credentials for account-based apps.
- Mixed credentials causing test path divergence from real users.
- Missing OIDC or file association for suitable apps.
- AI products without a clear route judgment.

## Bundled References

- Shipping Checklist: [references/shipping-checklist.md](./references/shipping-checklist.md)
- Metadata Standards: [references/metadata-standard.md](./references/metadata-standard.md)
- Store Assets Guide: [references/store-assets.md](./references/store-assets.md)
- AI Pod Review Kit: [references/aipod-review-kit.md](./references/aipod-review-kit.md)
- Cash Incentive Rules: [references/cash-incentive.md](./references/cash-incentive.md)
- AI Pod Playbook: [../lazycat:create-app/references/aipod-playbook.md](../lazycat:create-app/references/aipod-playbook.md)
- App Updates: `lazycat:update-app`
- Creation/Baseline: `lazycat:create-app`
- Admin UI: `lazycat:admin-ui`
- Porting: `lazycat:port-app`
- Guides: `lazycat:write-guide`
- Icons: `lazycat:prepare-icon`

## Outputs

```text
Phase: <Initiation / Creation / Packing / Assets / Testing / Submission / Follow-up / Post-release>
Target Version: <Version or TBD>

Incentive Goal: <Standard / Incentive Priority / Unspecified>

Confirmed
- ...

Incentive Eligibility
- ...

Gaps / Risks
- ...

Current Action
- ...

Next Steps
1. ...
2. ...

Deliverables
- ...
```
