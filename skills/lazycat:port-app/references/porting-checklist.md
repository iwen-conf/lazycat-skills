# Lazycat Porting Checklist

## 1. Before Selection
- [ ] GitHub search completed.
- [ ] License allows porting and commercial use (non-commercial projects prohibited).
- [ ] Upstream URL recorded.
- [ ] App Store duplication check completed.
- [ ] If highly duplicative, terminated incentive path or clarified differentiation.

## 2. Project Execution
- [ ] Created `docs/requirements`.
- [ ] Created `docs/api-design`.
- [ ] Created `docs/architecture`.
- [ ] Created `docs/release-prep`.
- [ ] Created `build.sh`.
- [ ] Created or completed `Makefile` in the repo after migration, not just as a proposed template.
- [ ] `make build` implemented.
- [ ] `make install` implemented.
- [ ] No upstream business source files were modified for the port. Only Lazycat packaging/runtime wrapper files were changed, unless the user explicitly approved business-code changes.

## 3. Lazycat Adaptation
- [ ] Runtime model preflight completed before writing manifest: delivery form, entry path, persistence, dependency layers, initialization, and login path.
- [ ] Runability decision recorded as one of: `Can Run`, `Can Run After Packaging Fixes`, `Cannot Determine Yet`, `Not Suitable For Standard Lazycat Port`.
- [ ] Image/architecture evidence recorded: pullable image, build path, `linux/amd64` compatibility, or explicit advanced-runtime requirement.
- [ ] Long-running process model verified: entrypoint/command starts a foreground service rather than only a one-shot CLI or manual shell task.
- [ ] Actual listener protocol and port verified from upstream docs, Compose, Dockerfile, config, or runtime logs before writing `routes`/`upstreams`/`ingress`.
- [ ] Delivery path is justified: official image, Compose conversion, Dockerfile-built image, wrapper image, AI Pod, or source-only fallback.
- [ ] HTTP/TCP entry decision is explicit: `routes`, `upstreams`, `secondary_domains`, `ingress`, or browser extension.
- [ ] Persistent paths are mapped only to Lazycat-approved writable locations such as `/lzcapp/var` or `/lzcapp/cache`; runtime-writable configs are not left under `/lzcapp/pkg/content`.
- [ ] Service dependency graph is documented as infra -> middleware -> seed/migration -> business.
- [ ] External requirements classified: none, user-configurable, advanced capability, or blocking dependency.
- [ ] Prepared `lzc-build.yml`.
- [ ] Prepared `lzc-manifest.yml`.
- [ ] Prepared `package.yml` with `unsupported_platforms` declaring `android`, `ios`, and `tvos` unless those platforms were verified.
- [ ] Prepared `package.yml.locales` with BCP 47 keys such as `zh-CN` and `en-US`, including usage text when login credentials need explanation.
- [ ] If the app uses bridged images, the final pullable image refs have already been written back to `lzc-manifest.yml` before `make install`.
- [ ] Evaluated OIDC requirement.
- [ ] Evaluated `file_handler` requirement.
- [ ] Clarified credential acquisition path for apps requiring login.
- [ ] Distinguished credential scopes: `lazycat_account` / `lazycat_password` for Lazycat OS and App Store; `lazycat_developer_center_account` / `lazycat_developer_center_password` for the Developer Center. In-app login uses app-level variables.
- [ ] Application MUST support passwordless auto-login (免密登录) via OIDC or Inject (`builtin://simple-inject-password`) to provide a seamless user experience.
- [ ] For apps with an internal login page, a fixed initial account is created non-invasively at startup when supported by upstream CLI/CMD/env/admin API.
- [ ] Initial passwordless-login credentials are documented: `账号`, `密码`, `昵称`.
- [ ] Modifiable passwordless login follows the official three-phase inject pattern: request -> response -> browser, with `ctx.flow`, `ctx.persist`, and `builtin://simple-inject-password`.
- [ ] Login/init/change-password API paths, payload keys, and CSS selectors were verified from docs/runtime traffic or explicitly provided by the user; none were guessed.
- [ ] Startup order gates seed/login initialization behind app and infra healthchecks, and gates dependent business services on `condition: service_healthy` where applicable.
- [ ] Ingress/application healthchecks were tested as end-to-end symptoms, not treated as the root cause. If ingress reports a service is not ready, target service logs and dependency health were inspected.
- [ ] Database/cache first-boot readiness was validated with real probes and realistic `start_period`, especially PostgreSQL/MySQL cases that temporarily reject connections.
- [ ] One-shot migration/seed services use bounded retry for transient dependency failures, write a terminal success marker only after success, and are idempotent on reinstall/restart.
- [ ] Required deploy parameters were installed through the setup wizard and rendered config/env values were inspected inside the container. Missing config or `need setup deploy params` is resolved before route/healthcheck tuning.
- [ ] For AI projects, determined if AI Pod or AI Browser Plugin is more suitable.

## 4. Incentive Path
- [ ] Does not belong to non-rewarded types.
- [ ] Original/Porting path clarified.
- [ ] Prepared upstream attribution (for ports) and ensured the `author` field strictly matches the original project's author.
- [ ] Planned OIDC or `file_handler` if applicable.
- [ ] Planned to verify core capabilities by installing and opening the app within Lazycat OS.
- [ ] Planned Standard App / AI Pod / AI Browser Plugin route for AI-native projects.
