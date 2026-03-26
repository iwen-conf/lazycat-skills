# Docker Porting Pitfalls and Best Practices Guide

Developers often encounter issues at key points when porting existing Docker images or `docker-compose.yml` to Lazycat Micro-service (`lzc-manifest.yml`). Apply these best practices when assisting users.

## 1. User & Permissions
**Problem:** Many third-party Docker images run as non-root users (e.g., `node`, `abc`). However, in Lazycat Micro-service, persistent directories (`/lzcapp/var/`) and user document directories (`/lzcapp/run/mnt/home/`) require `root` privileges by default. This often results in `Permission denied` errors.

**Best Practices:**
- **Primary Solution:** Run the container as `root` whenever possible. This is the simplest path if the image documentation does not strictly forbid it.
- **Secondary Solution (App forbids root):** If the application (e.g., certain databases) prohibits `root`, use a `setup_script` with root privileges to handle directory permissions before the app starts, or specify `user: "1000"` (ensure the ID is a quoted string) in the `services` block and adjust permissions accordingly.

## 2. Config Files Initialization and R/W
**Problem:** An application requires an initial configuration file (e.g., `config.yml`) packaged in the lpk (located at `/lzcapp/pkg/content/`), which the app needs to modify at runtime. Mounting `/lzcapp/pkg/content/config.yml` directly via `binds` will cause the app to crash when it attempts to write, because `/lzcapp/pkg/content` is **Read-Only**.

**Best Practices:**
- **Never** mount files under `/lzcapp/pkg/content/` directly for read-write operations.
- **Correct Approach:** Use a `setup_script`. Before the container executes its main logic, check if the target writable path (e.g., `/lzcapp/var/config.yml`) exists. If not, copy the initial configuration from `/lzcapp/pkg/content/` to the writable path.
  
  ```yaml
  services:
    app:
      image: xxx
      binds:
        - /lzcapp/var/conf:/app/conf # Mount writable directory
      setup_script: |
        if [ ! -f /app/conf/config.yml ]; then
            cp /lzcapp/pkg/content/default-config.yml /app/conf/config.yml
        fi
  ```

## 3. Startup Order and Healthchecks
**Problem:** Heavy applications with databases may take a long time to initialize tables on the first run. Without proper healthchecks, the system might mark the container as `unhealthy` and kill it before initialization finishes.

**Best Practices:**
- **Do not rely solely on hard waits (`sleep`):** Extending `start_period` is not a perfect solution and results in a poor user experience.
- **Correct Approach:** Write semantic healthcheck probes. For Web services, use `curl` to check actual API endpoints. For databases like MySQL, use actual SQL queries (e.g., `select 1`) to determine if the service is truly ready.
- Use `services.[].healthcheck` (instead of the deprecated `health_check`) and properly configure `retries`, `interval`, and `start_period`.

## 4. Privileged & Capabilities
**Problem:** Certain applications (e.g., side routers, VPNs, apps requiring FUSE mounts) depend on Docker's privileged mode (`privileged: true`) or specific capabilities (`cap_add`).

**Best Practices:**
- If the original application requires privileged access, grant it in the micro-service.
- Use the `compose_override` field in `lzc-build.yml` to inject low-level Docker parameters like `privileged`, `cap_add`, and `devices`.
- **Store Review:** Apps requiring these privileges **are allowed** for submission and review in the official Lazycat App Store, provided the functionality is reasonable.

## 5. Host & CORS Headers
**Problem:** Container services often strictly validate the `Host` Header in HTTP requests. Incorrect headers can lead to errors or CORS issues.

**Best Practices:**
- **Default Case:** Lazycat's `lzc-ingress` is intelligent and handles Host headers and CORS automatically in most scenarios; no special configuration is usually required.
- **Special Case:** If the application reports domain or Host validation errors, configure forwarding rules in `application.upstreams` and set `use_backend_host: true` to ensure the upstream service receives the expected Host header.
