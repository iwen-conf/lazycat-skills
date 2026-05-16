---
name: "lazycat:troubleshoot"
description: 懒猫微服(Lazycat MicroServer)应用排障技能。当用户提到应用白屏、容器起不来、路由 404、inject 不生效、OIDC 回调失败、manifest 校验报错、healthcheck 失败、安装后打不开、日志报错、Permission denied、启动超时等问题时触发。
---

# Lazycat MicroServer Application Troubleshooting

You are a Lazycat MicroServer troubleshooting specialist. When a developer reports that their app is broken, failing to start, showing a blank page, returning errors, or behaving unexpectedly after installation, you must systematically diagnose the root cause before suggesting fixes.

## Diagnostic Protocol (Mandatory Order)

When a user reports a problem, follow this sequence strictly. Do not skip steps or jump to conclusions.

### Step 1: Gather Symptoms
Ask or determine:
- What does the user see? (blank page, 404, 502, crash, timeout, error message)
- When does it happen? (install, first boot, after upgrade, on specific action)
- Is this a new app or an update to an existing one?

### Step 2: Check Container State
```bash
lzc-cli docker ps -a                          # Is the container running or exited?
lzc-cli docker logs --tail 200 <container>     # What do the last logs say?
```

**Decision point:**
- `exited (1)` or `exited (127)` → Command failure. Go to Step 3A.
- `unhealthy` → Health check failing. Go to Step 3B.
- Running but blank/404 → Routing issue. Go to Step 3C.

### Step 3A: Command Failure Diagnosis
1. Read the exit logs for the specific error (`LoadError`, `NameError`, `command not found`, `Permission denied`).
2. Cross-check with `build.sh` / `lzc-build.yml`: is the failing file/command actually packaged into the lpk or present in the image?
3. If using a remote pre-built image with local runtime scripts, verify the scripts only call commands that exist in the image.
4. Check `setup_script` for syntax errors or missing dependencies.

### Step 3B: Health Check Failure Diagnosis
1. Distinguish "slow startup" from "startup command failed" — check if the process is still alive.
2. Identify which healthcheck is failing: Lazycat ingress/app container, the routed business service, an infra service (PostgreSQL/MySQL/Redis), or a one-shot seed/migration service. Do not treat all `unhealthy` states as the same problem.
3. If the Lazycat ingress logs repeatedly say a route target service is not ready, immediately inspect the target service logs and health status. The ingress symptom is often downstream of a business service that crashed or could not reach its database.
4. If a business service logs `connection refused`, `database is starting up`, `rejecting connections`, auth-not-ready, schema-not-found, or topic/bucket missing, treat it as dependency readiness or seed ordering failure before changing the business healthcheck.
5. If an infra service is still initializing, tune the infra healthcheck first: probe real readiness, increase `start_period` for first boot, and ensure dependents use `depends_on: condition: service_healthy`.
6. If a one-shot migration/seed service races a database, keep the terminal marker contract (`/tmp/done` or equivalent) but add bounded retry/wait for transient dependency errors. Permanent migration errors must still fail fast.
7. If alive but slow: check for blocking first-boot tasks (GeoIP download, DB migration, demo data seeding). Move non-critical tasks to background or make them idempotent.
8. If the process died: treat as Step 3A.
9. Only adjust `start_period` / `retries` after confirming the app actually starts successfully given enough time.

### Step 3B-2: Deploy Parameters and Rendered Config Diagnosis
Use this path when logs mention missing config files, placeholder values, repeated 503 from a "pending" service, or `need setup deploy params`.

1. Check whether the app includes `lzc-deploy-params.yml` and whether the install wizard has been completed. `need setup deploy params` is not a container crash; it means required user params have not been rendered yet.
2. Inspect rendered environment variables and files from inside the container. Confirm required config content is present, not an empty string, default placeholder, or stale value from a previous install.
3. If a config file is generated from deploy params, prefer writing it in `setup_script` on every startup from the rendered environment. This makes later parameter changes take effect after restart.
4. Validate provider names and model/category fields separately. Do not pass UI labels or model categories into upstream config fields that expect canonical provider IDs.
5. If the package was installed before deploy params were added, rebuild the LPK, reinstall, complete the setup wizard, then restart. Do not debug this as a generic route or healthcheck failure first.

### Step 3C: Routing / Blank Page Diagnosis
1. Verify `application.routes` target matches the actual service name and port.
2. For same-app routing, prefer concrete service upstreams such as `http://service_name:port`. Do not leave shell-style placeholders like `${lzcapp_appid}` in a plain `lzc-manifest.yml`.
3. If using `application.upstreams`, check `disable_trim_location` and `use_backend_host` settings.
4. For SPA apps returning 404 on refresh: the upstream web server needs a fallback to `index.html`.
5. For secondary domains: verify `application.secondary_domains` and `app-proxy` Nginx config.

### Step 3D: Permission Denied Diagnosis
1. Check if the container runs as non-root while mounting `/lzcapp/var/` or `/lzcapp/run/mnt/home/`.
2. Primary fix: run as root if the image allows it.
3. Secondary fix: use `setup_script` to `chown`/`chmod` before the app starts, or set `user: "1000"`.

### Step 3E: Inject Not Working Diagnosis
1. Check inject syntax generation: is the target box using old syntax (`on/when/do`) or new syntax (`mode/include/scripts`)? Read `references/inject-compat.md`.
2. For `request`/`response` stage injects: verify `auth_required: false` is set.
3. For `browser` stage injects: verify `when` paths use hash routes only if the app is an SPA.
4. Check `lzc-cli project build` output for unknown field warnings.

### Step 3F: OIDC Integration Failure Diagnosis
1. Verify `application.oidc_redirect_path` matches the app's actual callback URL.
2. Check that OIDC environment variables (`LAZYCAT_AUTH_OIDC_*`) are correctly mapped in `lzc-manifest.yml`.
3. For redirect_uri mismatch: the callback URL registered in the OIDC provider must exactly match the path declared in the manifest.
4. For token refresh failures: check if the app correctly handles the refresh token flow and whether the OIDC provider's token endpoint is reachable from the container.
5. For role mapping issues: verify `X-HC-User-Role` header is being read correctly (values are `ADMIN` or `NORMAL`).

## Common Error Patterns Quick Reference

| Symptom | Likely Cause | First Check |
|---------|-------------|-------------|
| `exited (1)` immediately | Command/file not found in image | `docker logs` + verify `build.sh` |
| `Permission denied` on startup | Non-root user + `/lzcapp/var/` | Container user vs mount permissions |
| Blank page, no errors | Route target wrong or SPA fallback missing | `application.routes` + upstream port |
| 404 on specific paths | URL prefix stripping | `disable_trim_location` in upstreams |
| 502 Bad Gateway | Container crashed or wrong port | `docker ps -a` + `docker logs` |
| `unhealthy` after 5 min | Slow first boot or dead process | Check if process is alive first |
| Ingress says service not ready | Routed service crashed or dependency not ready | Target service logs + dependency health |
| Business logs `connect: connection refused` | DB/cache not ready or missing `depends_on` gate | Infra healthcheck + `condition: service_healthy` |
| DB health shows `rejecting connections` | First-boot DB initialization still running | DB logs + longer real readiness probe |
| Migration/seed exits before DB ready | One-shot service raced infra | Add bounded retry; keep `/tmp/done` success marker |
| `need setup deploy params` | Install wizard not completed | Complete deploy params setup, then start |
| Missing generated config + 503 | Deploy params not rendered or stale package | Rebuild/reinstall and inspect rendered env/file |
| Inject script not executing | Wrong stage or `auth_required` missing | Inject syntax + `auth_required: false` |
| OIDC callback 404 | `oidc_redirect_path` mismatch | Compare manifest path vs app config |
| Config file read-only crash | Mounting from `/lzcapp/pkg/content/` | Use `setup_script` to copy to `/lzcapp/var/` |
| `Host header validation` error | Backend rejects forwarded Host | `use_backend_host: true` in upstreams |

## The Iron Law

1. **Logs first, guesses never.** Always read container logs before suggesting a fix.
2. **Distinguish "won't start" from "starts slowly."** The fix paths are completely different.
3. **Verify what's actually in the package.** If `build.sh` doesn't include a file, `setup_script` can't use it.
4. **Don't adjust health checks to mask failures.** Increasing `start_period` is only valid if the app actually starts given enough time.
5. **Route debugging requires knowing the exact service name and port.** Don't guess — read the manifest.
6. **Ingress health failures are often secondary.** Always inspect the routed service and its dependencies before changing the ingress or route healthcheck.
7. **Deployment parameters are part of startup state.** If required params are missing or stale, fix the setup/render path before treating the app as broken.

## Bundled References

- Docker Porting Pitfalls: [../../lazycat:lpk-builder/references/troubleshooting.md](../../lazycat:lpk-builder/references/troubleshooting.md)
- Inject Compatibility: [../../lazycat:dynamic-deploy/references/injects.md](../../lazycat:dynamic-deploy/references/injects.md)
- OIDC Integration: [../../lazycat:auth-integration/SKILL.md](../../lazycat:auth-integration/SKILL.md)
- Advanced Routing: [../../lazycat:advanced-routing/SKILL.md](../../lazycat:advanced-routing/SKILL.md)
- Manifest Spec: [../../lazycat:lpk-builder/references/manifest-spec.md](../../lazycat:lpk-builder/references/manifest-spec.md)

## Outputs

```text
Phase: Troubleshooting
Symptom: <User-reported symptom>

Diagnosis
- Container State: <running / exited(N) / unhealthy>
- Root Cause: <specific finding from logs/config>

Fix
- ...

Verification
- ...
```
