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

## 6. 预构建镜像、运行时脚本与首启排障
**Problem:** 端口迁移类应用经常复用远程预构建镜像，同时只把 `runtime/`、`lzc-manifest.yml`、`docs/` 等文件打进 lpk。此时如果你修改了运行时脚本，并让脚本去调用镜像里原本不存在的新 Ruby 文件、新 rake task 或新应用代码，容器会在首启时直接退出，而不是单纯“启动太慢”。

**Best Practices:**
- **先区分“慢启动”还是“启动命令失败”：** 如果日志里出现 `container ... exited (1)`、`Unrecognized command`、`LoadError`、`NameError`，先按命令失败处理，不要先去调大 `start_period`。
- **明确包里真正带了什么：** 检查 `build.sh`/`lzc-build.yml`。如果打包过程只复制 `runtime/` 和 manifest，而没有构建新的应用镜像，就不要让 `runtime/*.sh` 依赖镜像中不存在的新增源码。
- **复用远程镜像时，优先把自定义逻辑写在包内可交付层：** 例如 `runtime/*.sh`、`setup_script`、已存在的镜像命令参数。只有当你真的同步构建并发布了新镜像时，才去依赖新增的应用源码或 rake task。
- **排查顺序要固定：**
  1. 看 `startup-log-tail`，判断是卡在 `Waiting` 还是已经 `exited (1)`。
  2. 看失败容器的最后日志，抓具体异常字符串。
  3. 对照 `build.sh`，确认报错涉及的文件或命令是否真的会进入最终包或镜像。
  4. 只有在命令存在且进程仍存活时，才继续分析 healthcheck 与冷启动耗时。
- **缩短首启时，先移走非关键前置任务：** 像 GeoIP 下载、演示数据准备、可延后缓存构建，不要阻塞 Web 服务首次通过健康检查。优先放到后台 worker、应用就绪后的异步任务，或做成“已存在即跳过”的幂等步骤。
- **避免重复应用引导：** 如果启动链里先 `bundle exec rails <task>`，又额外跑一次 `bundle exec rails runner`，通常意味着支付了两次完整 Rails 冷启动成本。能合并成一次就合并成一次。
- **经验规则：** “healthcheck 失败”不等于“一定要调 healthcheck”。先确认应用有没有真的启动起来，再确认它是不是被不必要的前置步骤拖慢。
