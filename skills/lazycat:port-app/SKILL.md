---
name: lazycat:port-app
description: 面向 Lazycat 应用移植选型与落地的 skill。只要用户提到移植应用、移植开源项目、从 GitHub 选项目、查重、避免重复移植、看 App Store 里是否已有、移植激励、上游仓库、build.sh、Makefile、lpk 打包、file_handler、OIDC、商店登录、lazycat_account、lazycat_password、懒猫算力仓、AI应用、AI 浏览器插件 等请求，就必须使用此 skill。负责从 GitHub 搜索候选项目、在 Lazycat App Store 查重、判断是否还有激励空间，并把移植项目收敛到可打包、可安装、可提审的状态，同时判断 AI 项目是否更适合移植为懒猫 `AI应用` 或 AI 浏览器插件。
---

# Lazycat App Porting

You are responsible for progressing a "search for a portable project" to a "portable and worthwhile" state. The focus is not just on porting itself, but on market de-duplication and incentive assessment, avoiding projects that are already duplicates, lack incentive potential, or are unsuitable for listing.

## Overview

This skill is for porting open-source or self-hosted software to Lazycat. Default requirements:

- Search GitHub first to determine upstream project, license, activity, and feature boundaries.
- Search the Lazycat App Store to confirm if a similar port already exists.
- If `lazycat_account` and `lazycat_password` exist on the local machine, prioritize using them to log in to the App Store before checking for duplicates. These variables are for MicroServer and App Store access, not the Developer Center or internal app accounts.
- If a similar port exists with no added value, do not proceed with a duplicate path.
- Implementation must provide `build.sh`, `Makefile`, `make build`, and `make install`.
- For image-based ports, settle final pullable image refs in `lzc-manifest.yml` during migration or `make update`; do not make `make install` responsible for `docker push`, `copy-image`, or manifest rewrite work.
- Must retain upstream address, license, and porting notes.

For tool-based apps or those requiring unified login, evaluate `file_handler` and Microservice OIDC. If the project is AI-native, determine if it should be ported as a Computing Power Cabin `AI App` or AI Browser Extension.

To verify apps installed on Lazycat MicroServer, use `lazycat_account` / `lazycat_password` for entry; use app-specific credentials (e.g., `lazycat_gitea_account`) for internal app login testing.

## Quick Contract

- **Trigger**: User mentions porting, "port," finding projects on GitHub, checking for duplicates, avoiding duplicate porting, porting incentives, upstream repository, Makefile, build.sh, file_handler, OIDC, store login, lazycat_account, lazycat_password, Computing Power Cabin, `AI App`, AI Browser Extension.
- **Inputs**: Candidate project keywords, target category, incentive goals, GitHub upstream info, App Store access status, local MicroServer credentials status, login/file association needs, AI-native status.
- **Outputs**: Candidate project list, de-duplication conclusion, incentive judgment, porting requirements, script entry requirements, AI Pod route judgment, and release entry for `lazycat:ship-app`.
- **Quality Gate**: Must complete GitHub search and App Store de-duplication. If local `lazycat_account` / `lazycat_password` are present, prioritize real login before checking. If a duplicate exists without differentiation, do not proceed with the incentive path. Deliverables must include `build.sh` and `Makefile`. AI projects must define a route (Standard, AI App, or Browser Ext).
## Decision Tree

- **Step 0 — Does a usable remote image exist?**
  - If the project publishes official Docker images on Docker Hub / GHCR / Quay (check README or Docker Hub page): **skip building entirely**. Use `lzc-cli appstore copy-image <image>` directly. This is the fastest and most reliable path.
  - If no remote image exists but a `Dockerfile` is provided: proceed to Step 1.
  - If neither exists: evaluate if the project is worth the build effort before proceeding.

- **Step 1 — Classify Complexity** (only when you must build from Dockerfile):

  | Signal | Points to Strategy A (Complex) | Points to Strategy B (Simple) |
  |--------|-------------------------------|-------------------------------|
  | Multi-stage Dockerfile | Yes | — |
  | Requires database (Postgres, MySQL, Redis, etc.) | Yes | — |
  | Multi-service `docker-compose.yml` (3+ services) | Yes | — |
  | Build output > 500MB image | Yes | — |
  | Single binary or static files only | — | Yes |
  | No external service dependencies | — | Yes |
  | Can run with `exec://` or a generic base image (alpine, nginx) | — | Yes |
  | Build completes in < 2 minutes on a standard machine | — | Yes |

  If 2+ signals point to Strategy A, use Strategy A. Otherwise, use Strategy B.

- **Strategy A** (Complex): Build locally for `linux/amd64` → Push to Docker Hub → `lzc-cli appstore copy-image` → Use returned `registry.lazycat.cloud/...` address in `lzc-manifest.yml`.
- **Strategy B** (Simple): Extract source → Direct build/run in LPK using `exec://` or generic base images. No Docker Hub round-trip needed.

- **Step 2 — Check Project Type**: Perform GitHub search, App Store check, incentive judgment, integration assessment, and AI Pod route judgment. Finally, decide to proceed or switch projects.


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
2. If local `lazycat_account` and `lazycat_password` exist, prioritize using an interactive `zsh` to read them and log in to the App Store for checking. Do not assume lack of credentials in non-interactive shells.
3. If a duplicate exists in the App Store without differentiation, do not proceed with the incentive path.
4. Every port must include upstream address, license, and porting notes.
5. Every port must provide `build.sh`, `Makefile`, `make build`, and `make install`.
6. Prioritize OIDC or `file_handler` for suitable projects, as they affect incentives and UX.
7. For AI-native projects, determine if they fit better as standard apps, `AI Apps`, or AI Browser Extensions.
8. If a ported project needs a static homepage, the priority must be: `Connection Entry`, `Status Check`, `Actions`, `Feedback`. Do not put "Why use it" or "Roadmap" in running pages; use `README` or store assets.
9. If the port relies on bridged images, the returned registry address must be written back into `lzc-manifest.yml` before `make build` / `make install`. `make install` is only for building and installing the ready LPK.
10. **Prioritize Docker over Source Code**: If a project provides a Docker image or `docker-compose.yml`, base the porting ENTIRELY on these Docker artifacts. **Do NOT** read or analyze the project's source code. Just use the Docker image directly.
    - **Auto-Translation for `docker-compose.yml`**:
      - `ports: ["8080:80"]` -> Convert to `routes` in `lzc-manifest.yml` (e.g., `- /=http://${service_name}.${lzcapp_appid}.lzcapp:80`).
      - `volumes: ["./data:/app/data"]` -> Convert to `binds` mapping to `/lzcapp/var/` (e.g., `- /lzcapp/var/data:/app/data`).
      - `depends_on` -> Not directly needed in Lazycat. Services communicate automatically via `${service_name}.${lzcapp_appid}.lzcapp`.

## Workflow

### 1. Search GitHub Candidates
- Use GitHub to find candidate projects.
- Check license for distribution and modification. Ensure the license explicitly permits commercial use; do not port projects with non-commercial licenses.
- Check activity, issue status, README completeness, and deployment complexity.
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

### 4. Solidify Porting Repository Entries
At minimum, add:
- `docs/requirements/`, `docs/api-design/`, `docs/architecture/`, `docs/release-prep/`.
- `build.sh`, `Makefile` (with `build` and `install` targets).

### 5. Plan Adaptation and Image Sync
- `lzc-build.yml`, `lzc-manifest.yml`.
- **Image Porting (Mandatory)**: 
  1. If the image is custom-built, first build it locally for `linux/amd64` and push it to Docker Hub.
  2. Use `lzc-cli appstore copy-image <docker_image>` to get an official `registry.lazycat.cloud/...` image name.
  3. Use the returned image in `lzc-manifest.yml`.
  4. Put this sync + manifest-backfill step in migration flow or `make update`, not in `make install`.
  5. By the time `make build` / `make install` runs, `lzc-manifest.yml` should already contain the final pullable image refs.
- `build.sh`, `Makefile` (must include `build`, `install`, `update`, `release-prep`).
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
- `build.sh` and `Makefile` planned.
- OIDC or `file_handler` evaluated.
- For incentives: clarified that real installation and verification on Lazycat MicroServer is mandatory.
- AI Pod route evaluated for AI projects.

## Red Flags

- Searching GitHub without checking the App Store.
- Claiming a check is complete without being logged into the App Store.
- Not attempting login when `lazycat_account` / `lazycat_password` exist locally.
- Misusing Developer Center or app-specific accounts for App Store checks.
- Proceeding with incentives when duplicates exist.
- Relying on marketing talk for incentives when the app is weak.
- Missing upstream address or license.
- No `build.sh` or `Makefile`.
- Missing file association for tool apps.

## Bundled References

- Market Research: [references/market-research.md](./references/market-research.md)
- Porting Checklist: [references/porting-checklist.md](./references/porting-checklist.md)
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
