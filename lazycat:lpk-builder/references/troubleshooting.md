# Docker Porting Pitfalls and Best Practices

When porting existing Docker images or `docker-compose.yml` files to Lazycat Micro-service (`lzc-manifest.yml`), developers often encounter issues at several key points. Refer to and apply these best practices when assisting users.

## 1. User & Permissions
**Problem:** Many third-party Docker images run as non-root users (e.g., `node`, `abc`) by default. However, in Lazycat Micro-service, persistent directories (`/lzcapp/var/`) and user document directories (`/lzcapp/run/mnt/home/`) require `root` privileges for read/write access by default, leading to `Permission denied` errors.

**Best Practices:**
- **Primary Solution:** Run the container as the `root` user whenever possible. This is the simplest path if the image does not strictly forbid it.
- **Secondary Solution (Root Forbidden):** If the application (e.g., certain databases) enforces a non-root policy, use `setup_script` to handle directory permissions with root privileges beforehand, or use `user: "1000"` (note: the UID must be a quoted string) in the `services` block and adjust permissions before startup.

## 2. Config File Initialization and R/W
**Problem:** Applications often require an initial configuration file (e.g., `config.yml`) packaged in the lpk (located at `/lzcapp/pkg/content/`), which the app needs to modify at runtime. Binding `/lzcapp/pkg/content/config.yml` directly via `binds` will fail because `/lzcapp/pkg/content` is **Read-Only**, causing the app to crash during modification attempts.

**Best Practices:**
- **Never** mount files under `/lzcapp/pkg/content/` directly as writable configurations.
- **Correct Approach:** Use `setup_script`. Before the container executes its original logic, check if the target writable path (e.g., `/lzcapp/var/config.yml`) exists. If not, copy the initial configuration from `/lzcapp/pkg/content/` to the writable location.
  
  ```yaml
  services:
    app:
      image: xxx
      binds:
        - /lzcapp/var/conf:/app/conf # Mount a writable directory
      setup_script: |
        if [ ! -f /app/conf/config.yml ]; then
            cp /lzcapp/pkg/content/default-config.yml /app/conf/config.yml
        fi
  ```

## 3. Startup Sequence & Health Checks
**Problem:** Heavyweight applications with databases may take a long time to initialize table structures during the first boot. Without proper health check configuration, the system may mark the container as `unhealthy` and kill it before initialization completes.

**Best Practices:**
- **Avoid hardcoded waits (`sleep`):** Arbitrarily extending the `start_period` is unreliable and provides a poor user experience.
- **Correct Approach:** Write semantically meaningful health check probes. For Web services, use `curl` to check actual API endpoints. For databases (e.g., MySQL), use actual SQL queries like `select 1` to determine if the service is truly ready.
- Use `services.[].healthcheck` (instead of the deprecated `health_check`) and properly configure `retries`, `interval`, and `start_period`.

## 4. Privileged Mode & Capabilities
**Problem:** Some applications (e.g., sidecars, VPNs, or apps requiring FUSE mounts) depend on Docker's privileged mode (`privileged: true`) or specific capabilities (`cap_add`).

**Best Practices:**
- If the original app requires privileges, you can grant them in the Micro-service.
- Use the `compose_override` field in `lzc-build.yml` to inject these low-level Docker parameters (e.g., `privileged`, `cap_add`, `devices`).
- **Store Review:** Apps requiring these privileges **are allowed** for submission and review in the official Lazycat App Store, provided the functionality is justified.

## 5. LAN Access, CORS, and Host Header Validation
**Problem:** Container services often strictly validate the `Host` header of HTTP requests. Incorrect headers can lead to errors or CORS issues.

**Best Practices:**
- **Default Behavior:** Lazycat's `lzc-ingress` is highly intelligent and automatically handles Host headers and CORS in most cases; manual configuration is usually unnecessary.
- **Special Cases:** If the app reports domain or Host validation errors, configure forwarding rules in `application.upstreams` and add `use_backend_host: true` to ensure the upstream service sees its expected Host header.
