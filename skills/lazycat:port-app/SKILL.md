---
name: lazycat:port-app
description: "Port an open-source/self-hosted project to Lazycat. Use for candidate selection (GitHub search), App Store de-duplication, incentive-space check, and driving the port to a packageable/installable/reviewable state via wrapper + runtime layers only (non-invasive: never edit upstream business code). Also decides whether an AI project fits better as an AI App or browser extension. Chains to lazycat:lpk-builder for packaging and lazycat:ship-app for delivery. 移植开源项目、选型、查重、激励判断、build.sh/Makefile、非侵入打包上架。"
---

# Lazycat App Porting

You are responsible for progressing a "search for a portable project" to a "portable and worthwhile" state. The focus is not just on porting itself, but on market de-duplication and incentive assessment, avoiding projects that are already duplicates, lack incentive potential, or are unsuitable for listing.

## Overview

This skill is for porting open-source or self-hosted software to Lazycat. Default requirements:

- Search GitHub first to determine upstream project, license, activity, and feature boundaries.
- GitHub use is read-only by default. Do not create issues, PRs, forks, comments, reviews, discussions, stars, watches, follows, or any other visible interaction using the user's GitHub account unless the user explicitly approves that exact action in the current conversation.
- Search the Lazycat App Store to confirm if a similar port already exists.
- If `lazycat_account` and `lazycat_password` exist on the local machine, prioritize using them to log in to the App Store before checking for duplicates. These variables are for MicroServer and App Store access, not the Developer Center or internal app accounts.
- If a similar port exists with no added value, do not proceed with a duplicate path.
- Implementation must provide `build.sh`, `Makefile`, `make build`, and `make install`.
- Once migration is confirmed, the agent must actually create or complete the repository `Makefile`; do not stop at giving advice, TODOs, or pseudo-targets.
- For image-based ports, settle final pullable image refs in `lzc-manifest.yml` during migration or `make update`; do not make `make install` responsible for `docker push`, `copy-image`, or manifest rewrite work.
- The porting boundary is non-invasive by default and must be enforced as a write-scope gate, not just a preference. Only change Lazycat packaging/runtime wrapper files unless the user explicitly approves product development scope in the current task.
- Before declaring a project suitable for porting, classify its Lazycat runtime model: delivery form, entry route, persistence, dependency layers, initialization, and login. Read `../lazycat:lpk-builder/references/runtime-model.md` for the canonical preflight model.
- For common web/server ports, declare unsupported platforms `android`, `ios`, and `tvos` in `package.yml` unless those clients have been verified.
- If the app requires an in-app account, integrate passwordless login and document the credential contract: account, password, and nickname. Prefer startup-created fixed initial credentials plus three-phase inject for later password changes.
- If the app has file open/save/upload/download flows, integrate Lazycat file picker selection by `application.injects` with the official file chooser inject; do not modify upstream business frontend for this integration.
- Must retain upstream address, license, and porting notes.
- In `package.yml`, set `author` from the GitHub owner segment in `homepage`; for example, `homepage: https://github.com/Makisuo/maple` maps to `author: Makisuo`.

For tool-based apps or those requiring unified login, evaluate `file_handler` and Microservice OIDC. If the project is AI-native, determine if it should be ported as a Computing Power Cabin `AI App` or AI Browser Extension.

To verify apps installed on Lazycat MicroServer, use `lazycat_account` / `lazycat_password` for entry; use app-specific credentials (e.g., `lazycat_gitea_account`) for internal app login testing.

## Quick Contract

- **Trigger**: User mentions porting, "port," finding projects on GitHub, checking for duplicates, avoiding duplicate porting, porting incentives, upstream repository, Makefile, build.sh, file_handler, OIDC, store login, lazycat_account, lazycat_password, Computing Power Cabin, `AI App`, AI Browser Extension.
- **Inputs**: Candidate project keywords, target category, incentive goals, GitHub upstream info, App Store access status, local MicroServer credentials status, login/file association needs, AI-native status.
- **Outputs**: Candidate project list, de-duplication conclusion, incentive judgment, completed porting entry files (`build.sh`, `Makefile`), AI Pod route judgment, and release entry for `lazycat:ship-app`.
- **Quality Gate**: Must complete GitHub search, App Store de-duplication, and runtime runability gate. If local `lazycat_account` / `lazycat_password` are present, prioritize real login before checking. If a duplicate exists without differentiation, do not proceed with the incentive path. Deliverables must include an actually completed `build.sh` and `Makefile`, not just a suggested template. AI projects must define a route (Standard, AI App, or Browser Ext). Ports with login must include a passwordless-login design and documented initial credentials.
## Decision Tree

- **Rule 1 — Never Modify Original Business Code**: The entire porting process must rely on Docker images, wrapper images, startup scripts, seed scripts, and Lazycat manifest/package/build configuration. Do not edit or modify original upstream business code unless the user explicitly changes the task from "porting" to "feature development" and names the business files or feature scope that may be changed.
- **Rule 2 — Use Existing Image**: If the project publishes official Docker images (Docker Hub, GHCR, Quay, etc.), use the image directly. Run `lzc-cli appstore copy-image <image>`. This is the fastest and most reliable path.
- **Rule 3 — Build Image if Needed**: If no remote image exists but a `Dockerfile` is provided, build the image locally for `linux/amd64`, and then run `lzc-cli appstore copy-image <image>`.
- **Rule 4 — Write Back to YML**: Take the returned `registry.lazycat.cloud/...` address from the copy command and write it back into `lzc-manifest.yml`.
- **Rule 5 — Build LPK**: Once the `lzc-manifest.yml` is updated with the correct image addresses, run `make build` and `make install` to build the `.lpk` package.

- **Step 2 — Check Project Type**: Perform GitHub search, App Store check, incentive judgment, runtime-model preflight, integration assessment, and AI Pod route judgment. Finally, decide to proceed or switch projects.


## When to Use

**Primary Triggers**

- User wants to port a GitHub project to Lazycat.
- User needs help finding projects worth porting.
- User explicitly says "check for duplicate ports first."
- User wants to select projects based on porting incentive rules.
- User provides store login and wants to check App Store duplicates directly.
- User is preparing to port an AI product to the Computing Power Cabin or AI Browser.

**Typical Scenarios**

- Searching GitHub for self-hosted apps and filtering candidates for Lazycat.
- Porting tool apps and evaluating Disk Context Menus.
- Porting account-based apps and evaluating Microservice OIDC.
- Having an upstream repo but unsure if a similar port exists in the App Store.
- Porting an AI product but haven't decided between standard app and `AI App`.

**Boundary Notes**

- If the user wants an original app, use `lazycat:create-app`.
- If the user only wants to write a guide, use `lazycat:write-guide`.
- Do not conclude that a project is "worth porting" without completing the App Store check.

## Announce

Upon execution, provide a brief summary of:

- Which GitHub keywords you will search.
- Whether you have App Store search conditions and local environment variables.
- Whether the primary blocker is upstream quality, duplication, or incentive potential.

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `candidate_keywords` | string | Recommended | GitHub keywords (category, protocol, tech stack). |
| `reward_target` | enum(`Standard Port`/`Incentive Priority`) | Recommended | Exclude duplicates and non-reward types early for incentives. |
| `appstore_access` | enum(`LoggedIn`/`EnvVars Available`/`Not LoggedIn`/`Unknown`) | Recommended | App Store checks usually require MicroServer entry. Use `lazycat_account` / `lazycat_password`; do not use Dev Center or app-specific accounts. |
| `upstream_state` | enum(`Specified`/`Direction Only`/`Not Searched`) | Recommended | Direct assessment vs. GitHub search. |
| `integration_hint` | enum(`OIDC`/`file_handler`/`Both`/`None`) | Optional | Evaluate Microservice OIDC or Disk File Association. |
| `ai_native_hint` | enum(`None`/`Suspected AI`/`Definite AI`) | Optional | Evaluate AI Pod / AI App route. |

## The Iron Law

1. Perform GitHub search and App Store de-duplication before porting. Do not write code before checking duplicates.
2. Treat GitHub research as read-only. Public issue/PR status may be read for activity assessment, but creating issues, PRs, forks, comments, reviews, discussions, stars, watches, follows, or any other account-visible operation is forbidden without explicit user approval for that exact action.
3. If local `lazycat_account` and `lazycat_password` exist, prioritize using an interactive `zsh` to read them and log in to the App Store for checking. Do not assume lack of credentials in non-interactive shells.
4. If a duplicate exists in the App Store without differentiation, do not proceed with the incentive path.
5. Every port must include upstream address, license, and porting notes.
6. Every port must provide `build.sh`, `Makefile`, `make build`, and `make install`.
7. After migration is selected, you must finish the repo's `Makefile` yourself. Do not hand off with "please add a Makefile later" or leave placeholder targets unimplemented.
8. Prioritize OIDC, `file_handler`, and mandatory file picker selection for suitable projects, as they affect incentives and UX.
9. For AI-native projects, determine if they fit better as standard apps, `AI Apps`, or AI Browser Extensions.
10. If a ported project needs a static homepage, the priority must be: `Connection Entry`, `Status Check`, `Actions`, `Feedback`. Do not put "Why use it" or "Roadmap" in running pages; use `README` or store assets.
11. **Zero Modification to Original Business Code**: Absolutely DO NOT modify original source files for frontend pages, backend handlers, domain logic, auth logic, database schema/migrations, or tests during a port. Allowed scope is packaging/runtime wrapper files only: `package.yml`, `lzc-build.yml`, `lzc-manifest.yml`, `lzc-deploy-params.yml`, `Makefile`, `build.sh`, Docker wrapper files, startup/seed scripts, config templates, icons, store assets, and docs.
    - Before editing, state the intended write scope and keep it to the allowed wrapper files.
    - If a fix appears to require changing upstream business code, stop and report `Blocked by business-code change requirement`; offer a wrapper/image/config alternative if one exists.
    - Do not treat tests, UI copy, auth handlers, API handlers, migrations, seed data inside upstream app directories, or framework config as "small harmless changes"; they are business-source changes unless they are part of a wrapper directory created for the port.
    - User phrases such as "不要修改业务代码", "只做移植", "不要动上游", "包装一下", or an ordinary porting request mean business-source writes are forbidden.
12. **Image-Based Porting Flow**: If a project provides a Docker image, use the image directly. If no image exists but a Dockerfile is provided, build the image first. Once an image is available (remote or locally built), use `lzc-cli appstore copy-image <image>` to copy the image to the Lazycat registry.
13. **Write Back to YML and Build LPK**: After copying the image, the returned `registry.lazycat.cloud/...` address MUST be written back into `lzc-manifest.yml`. Only after this is done, run `make build` and `make install` to build the `.lpk` package.
14. **Strict Health Check and Startup Order**: You MUST configure `healthcheck` (with `test`, `start_period`, `interval`, `timeout`, `retries`) for all dependencies (like MySQL, Redis, etc.) and map `depends_on` with `condition: service_healthy`. Business containers must wait for infrastructure containers to be fully healthy to avoid startup crashes.
    - **Auto-Translation for `docker-compose.yml`**:
      - `ports: ["8080:80"]` -> Convert to `routes` in `lzc-manifest.yml` (e.g., `- /=http://service_name:80`). Do not leave `${lzcapp_appid}` or shell-style placeholders in a plain manifest.
      - `volumes: ["./data:/app/data"]` -> Convert to `binds` mapping to `/lzcapp/var/` (e.g., `- /lzcapp/var/data:/app/data`).
      - `depends_on` -> **KEEP IT**. Convert list form to map form with `condition: service_healthy` to guarantee correct startup sequences (e.g., app waits for DB to be healthy). DO NOT drop it, or apps will crash on boot.
      - `healthcheck` -> **MANDATORY FOR DEPENDENCIES**. If a service is depended upon, you MUST define a `healthcheck` with a robust `test` command (like `curl` or `mysqladmin ping`), and generous `start_period`, `interval`, `timeout`, and `retries`. Without health checks, `service_healthy` conditions will fail and dependents will hang forever.
    - Before writing the manifest, draw the service layers: infra -> middleware -> seed/migration -> business. If this graph is unclear, inspect upstream Compose, docs, env examples, and startup logs first.
    - Treat route health as end-to-end readiness. Lazycat can wait for route upstream ports, but it cannot make a business service survive a database race, missing schema, bad generated config, or absent initial account.
15. **Default Platform Declaration**: In `package.yml`, add `unsupported_platforms: [android, ios, tvos]` for normal migrated web/server apps unless you have verified mobile/TV support. Keep `locales` in `package.yml`, not `lzc-manifest.yml`, and use BCP 47 keys such as `zh-CN` and `en-US`.
16. **Passwordless Login Contract**: If the app has an internal login page, provide a non-invasive passwordless-login path before considering the port complete.
    - Create a fixed initial user at startup using documented CLI/CMD/env/admin API, `setup_script`, wrapper `entrypoint`/`command`, or a one-shot seed service. Do not edit business auth code.
    - Document the initial credentials in README/store usage/locales: `账号`, `密码`, `昵称`.
    - Use official three-phase inject for modifiable credentials: request captures login/init/change-password credentials into `ctx.flow`; response writes `ctx.persist` only on 2xx; browser fills login and current-password fields with `builtin://simple-inject-password`.
    - Do not invent API paths, payload keys, or selectors. Inspect runtime traffic or ask the user before writing inject YAML.
17. **Mandatory File Picker Selection for Ports**: If a migrated app has any file open/save/upload/download entry, provide Lazycat file picker selection before considering the port complete.
    - Use `application.injects` with `lzc-file-chooser-inject.js` to intercept browser file entry points, including `showOpenFilePicker()`, `showSaveFilePicker()`, and `<input type="file">`.
    - Keep this non-invasive. Do not edit upstream pages/components/routes/state just to add Lazycat file selection.
    - Put inject parameters under `do[].params`, including `diskRoot`, `fallbackMime`, `locale`, optional `text`, and `hooks.fileSystemAccess` / `hooks.fileInput`.
    - Verify both branches where applicable: local filesystem still follows the original flow, and Lazycat MicroServer opens the Lazycat file picker.
    - If the app has no file entry, document `File picker: Not applicable, no file open/save/upload/download flow`.
    - Official reference: https://developer.lazycat.cloud/lazycat-file-picker-auto-intercept.html.

## Workflow

### 1. Search GitHub Candidates
- Use GitHub to find candidate projects.
- This step is read-only. Do not use the user's GitHub account for any visible upstream action without explicit approval.
- Check license for distribution and modification. Ensure the license explicitly permits commercial use; do not port projects with non-commercial licenses.
- Check activity, public issue/PR status, README completeness, and deployment complexity.
- Record repository name, upstream address, license, and core features.

### 2. Check App Store Duplicates
- Search for candidate project names, Chinese/English aliases, and keywords in `https://appstore.ezer.heiyu.space/#/shop`.
- App Store search usually requires MicroServer login. Use `lazycat_account` / `lazycat_password`.
- If variables are in `~/.zshrc`, use interactive `zsh -ic` to read them.
- `lazycat_developer_center_account` is only for the Developer Center, not App Store checks.
- App-specific accounts (e.g., `lazycat_gitea_account`) are only for internal app testing.
- Do not pretend to have finished the check if no session or credentials exist.
- Record product name and overlap for any duplicates found.

### 3. Judgment of Incentives and Feasibility
- Check against official non-reward types.
- Original, first-time port, or duplicate?
- Additional incentive opportunities (OIDC, `file_handler`).
- Can regular users obtain credentials?
- For AI projects: evaluate `AI App` route.
- Runtime feasibility preflight:
  - Delivery: official image, Compose, Dockerfile, or source-only.
  - Entry: HTTP route/upstream, TCP/UDP ingress, secondary domain, or browser extension.
  - Persistence: writable data paths under `/lzcapp/var`, cache paths, document access, and permissions.
  - Dependencies: infra, middleware, seed/migration, business service order and healthchecks.
  - Initialization: `setup_script`, deploy params, generated config, default account creation.
  - Login: none, OIDC, fixed initial credentials, or inject.
  - File selection: no file flow, existing file flow with file chooser inject, or `file_handler` plus file chooser inject.
  - Runability gate: image/arch, process model, network listener, storage, dependencies, init/login, and external requirements.
  - Risk: low / medium / high, with blockers recorded before implementation.
  - Decision: `Can Run`, `Can Run After Packaging Fixes`, `Cannot Determine Yet`, or `Not Suitable For Standard Lazycat Port`. Do not use vague conclusions like "should be possible".

### 4. Solidify Porting Repository Entries
At minimum, add:
- `docs/requirements/`, `docs/api-design/`, `docs/architecture/`, `docs/release-prep/`.
- `build.sh`, `Makefile` (with `build` and `install` targets).
- After migration lands, actually write or fix the `Makefile` in the repo so it can be executed immediately; do not leave the port in a "Makefile pending" state.

Before writing files, create a porting write-scope note in your working response or porting docs:

```text
Allowed write scope
- Lazycat packaging/runtime: package.yml, lzc-build.yml, lzc-manifest.yml, lzc-deploy-params.yml
- Build/install wrapper: Makefile, build.sh, Docker wrapper files
- Runtime wrapper: runtime/, setup scripts, seed scripts, config templates
- Port documentation/assets: docs/, README/store usage, icons

Forbidden write scope
- Upstream frontend pages/components/routes/state
- Upstream backend handlers/services/domain logic/auth logic
- Upstream database schema/migrations/models
- Upstream tests or fixtures used to justify changed behavior
- Any existing upstream application source file unless the user explicitly approves product-development scope
```

### 5. Plan Adaptation and Image Sync
- `lzc-build.yml`, `lzc-manifest.yml`.
- Write a short runtime-model note in the porting docs before implementation: delivery form, entry, persistence, dependency layers, init, login, and unresolved risks.
- **Image Porting (Mandatory)**: 
  1. **Zero Business Source Modification**: Only use existing Docker images or build directly from a provided Dockerfile.
  2. If the image needs to be built, build it locally first.
  3. Use `lzc-cli appstore copy-image <docker_image>` to get an official `registry.lazycat.cloud/...` image name.
  4. Write the returned image address back into `lzc-manifest.yml`.
  5. Put this sync + manifest-backfill step in migration flow or `make update`, not in `make install`.
  6. By the time `make build` / `make install` runs to build the LPK, `lzc-manifest.yml` must already contain the final pullable image refs.
- **Business-Code Stop Rule**:
  1. If the app cannot start because upstream code assumes paths, hosts, ports, secrets, or writable directories, first solve with env vars, command args, `setup_script`, binds, generated config, reverse-proxy/upstream settings, or wrapper entrypoint.
  2. If the app cannot provide login because upstream has only internal auth, first solve with documented CLI/CMD/env/admin API, one-shot seed service, OIDC, or inject.
  3. If the only viable solution is changing upstream auth, API, UI, schema, migrations, or application source, mark the port `Cannot Determine Yet` or `Not Suitable For Standard Lazycat Port` and ask for explicit product-development approval.
  4. Never silently make the business-code change just to pass build, healthcheck, login, or review gates.
- `build.sh`, `Makefile` (must include `build`, `install`, `update`, `release-prep`).
- Add `unsupported_platforms` to `package.yml` for unverified platforms: `android`, `ios`, `tvos`.
- Add `locales.zh-CN` and `locales.en-US` usage text that explains how login works and lists the initial account/password/nickname when fixed credentials are used.
- For login apps, add passwordless login without changing business source: startup-created fixed user + health-gated seed + three-phase inject for login/change-password learning.
- For file-capable migrated apps, add `application.injects` for the official file chooser script and package `lazycat-injects/lzc-file-chooser-inject.js` into the LPK content. Treat missing file picker selection as an incomplete port unless the app has no file flow.
- Add `application.oidc_redirect_path` and `application.file_handler` if applicable.
- For AI-native: evaluate `ai-pod-service/`, `caddy-aipod`, and `extension.zip`.
- If a static homepage is needed, check if it's essential for runtime; if only for submission/promotion, use docs or store assets instead.

### 6. Hand over to Shipping Pipeline
Once "worth porting" is confirmed, hand over to `lazycat:ship-app` for packaging, assets, and submission. Use `lazycat:update-app` for future updates.

## Quality Gates

- GitHub search completed.
- App Store check completed (not an assumption).
- Real login used for checking if local credentials exist.
- Differentiation clarified or incentive path terminated if duplicates exist.
- Upstream address, license, and repo status recorded.
- `build.sh` and `Makefile` completed in the repo.
- OIDC or `file_handler` evaluated.
- For login apps, passwordless login designed, credentials documented (`账号`/`密码`/`昵称`), and API paths/selectors verified instead of guessed.
- For file-capable apps, file picker selection is either implemented with inject and verified, or explicitly marked not applicable because the app has no file flow.
- `package.yml` includes `unsupported_platforms` and BCP 47 `locales` when preparing a real LPK.
- For incentives: clarified that real installation and verification on Lazycat MicroServer is mandatory.
- Runtime runability gate completed and documented before implementation, with one of the four required decisions.
- AI Pod route evaluated for AI projects.

## Red Flags

- Searching GitHub without checking the App Store.
- Claiming a check is complete without being logged into the App Store.
- Not attempting login when `lazycat_account` / `lazycat_password` exist locally.
- Misusing Developer Center or app-specific accounts for App Store checks.
- Proceeding with incentives when duplicates exist.
- Relying on marketing talk for incentives when the app is weak.
- Missing upstream address or license.
- No runtime runability gate, especially for Compose stacks with databases or first-boot initialization.
- Claiming "can run" without evidence for image/arch, long-running process, listener port, writable data paths, dependencies, initialization, and login.
- No `build.sh` or `Makefile`.
- Missing file association for tool apps.

## Bundled References

- Market Research: [references/market-research.md](./references/market-research.md)
- Porting Checklist: [references/porting-checklist.md](./references/porting-checklist.md)
- Runtime Model and Porting Judgment: [../lazycat:lpk-builder/references/runtime-model.md](../lazycat:lpk-builder/references/runtime-model.md)
- Command Conventions: [references/command-conventions.md](./references/command-conventions.md)
- Native Batch Copy CLI (Go): [references/lzc-copy-image-go/README.md](./references/lzc-copy-image-go/README.md)
- S2I (Source-to-Image) Strategy: [references/s2i-strategy.md](./references/s2i-strategy.md)
- AI Pod Playbook: [../lazycat:create-app/references/aipod-playbook.md](../lazycat:create-app/references/aipod-playbook.md)
- Incentive Rules: [../lazycat:ship-app/references/cash-incentive.md](../lazycat:ship-app/references/cash-incentive.md)

## Outputs

```text
Phase: Porting Evaluation / De-duplication / Preparation
Target: <Standard Port / Incentive Priority>

GitHub Candidates
- ...

App Store Check
- ...

Conclusion
- Can Run / Can Run After Packaging Fixes / Cannot Determine Yet / Not Suitable For Standard Lazycat Port
- Proceed / Switch / Non-incentive path

Requirements
- build.sh
- Makefile
- make build
- make install
- Upstream address

Next Steps
1. ...
2. ...
```
