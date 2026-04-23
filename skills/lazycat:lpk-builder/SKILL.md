---
name: "lazycat:lpk-builder"
description: 用于将现有应用或代码打包为懒猫微服(Lazycat MicroServer) lpk 应用格式的专业指南。当用户需要将 docker 镜像、docker-compose转换或从零打包懒猫微服应用时触发。
---

# Lazycat MicroServer LPK Packaging and Porting Guide

You are a professional Lazycat MicroServer Application Ecosystem Developer. Your core task is to assist users in porting and packaging existing applications (such as Docker images or source code) into the `lpk` format supported by Lazycat MicroServer.

## Core Workflow

Packaging and porting a Lazycat MicroServer application (LPK v2) involves writing three core configuration files: `package.yml`, `lzc-build.yml`, and `lzc-manifest.yml`.

### 1. Static Metadata (`package.yml`)
Define the application's identity, version (must strictly be `x.x.x` format), and localization.
- For field definitions and BCP 47 localization rules, read `references/package-spec.md`.

**Standard Template:**
```yaml
package: com.example.myapp
version: 1.0.0
name: My App
description: High-performance Lazycat application.
author: Developer Name
license: MIT
locales:
  zh:
    name: "我的应用"
    description: "高性能懒猫应用。"
```

### 2. Writing the Build Configuration (`lzc-build.yml`)
This file defines how resources are packaged into an `lpk` file.
- To view the complete field definitions and specifications for this file, read `references/build-spec.md`.

**Standard Template:**
```yaml
buildscript: sh build.sh  # Build script
manifest: ./lzc-manifest.yml # Runtime execution configuration
contentdir: ./dist # Static content directory to be packaged
pkgout: ./ # lpk output path
icon: ./lzc-icon.png # App icon (1:1 PNG, < 200KB)
```

### 3. Writing the Manifest Configuration (`lzc-manifest.yml`)
This file defines the runtime environment, services, and routing. **Static metadata should not be included here.**
- Before writing, **be sure** to read `references/manifest-spec.md` for field definitions and advanced routing rules.

**Gold Porting Example (from Docker):**
```yaml
lzc-sdk-version: '2.0'
application:
  subdomain: yourapp # Default assigned subdomain
  routes:
    - /=http://your_service_name:80
services:
  your_service_name:
    image: nginx:latest # Image to run
    binds:
      - /lzcapp/run/mnt/home:/home # Mount user documents directory
    environment:
      - ENV_KEY=ENV_VALUE
```

### 4. Building and Installing with lzc-cli
After writing the configuration, guide the user to use the `lzc-cli` command-line tool for packaging and installation.

**Packaging the App:**
```bash
# Execute in the project root containing lzc-build.yml
lzc-cli project build -o release.lpk
```

**Installing the App:**
```bash
# Install the packaged lpk into the microservice
lzc-cli app install release.lpk
```

These packaging and install entrypoints must consume an already-correct `lzc-manifest.yml`. Do **not** redesign `make install` or `lzc-cli app install` wrappers to perform `docker push`, `lzc-cli appstore copy-image`, or manifest rewrites on the fly. If the user explicitly requests an end-to-end release/install chain, implement or use a **separate** release target such as `release-build` / `release-install` that finishes image build, public push, `copy-image`, manifest backwrite, LPK build, and installation before returning.

**Entering Devshell (Development & Debugging Environment):**
If the user needs to debug locally or within a container, guide them into the devshell.
```bash
lzc-cli project devshell
```

### 5. Inspecting Deployed Apps
As an agent, if you need to **view deployed or running Lazycat apps** (e.g., checking status, logs, or troubleshooting errors), you **must proactively use `lzc-cli docker`** prefixed commands to operate within the microservice's Docker environment.
```bash
# View running containers in the microservice (find your app container name or ID)
lzc-cli docker ps -a

# View logs for a specific app to troubleshoot
lzc-cli docker logs -f --tail 100 <container_name>

# Enter a deployed app's container to troubleshoot
lzc-cli docker exec -it <container_name> sh
```

## 6. Image Handling Specifications
The source of images is critical during packaging and testing:

**Testing Phase:**
If pulling from native registries (like Docker Hub) is slow or fails, push the image to the microservice's test registry:
1. **Get the Microservice Name**: As an agent, when you need the `<microservice_name>`, **proactively execute `lzc-cli box default`** instead of asking the user.
2. Re-tag the image: `docker tag <original_image> dev.<microservice_name>.heiyu.space/<image_name>:<version>`
3. Push the image: `docker push dev.<microservice_name>.heiyu.space/<image_name>:<version>`
4. Use this test image address in `lzc-manifest.yml`.

**Remote Image Bridge (Recommended):**

*When to choose this path — judge intelligently before packaging:*
- Upstream only provides a `Dockerfile`, or the user needs to publish a custom image containing local modifications.
- The project's build output is large (heavy artifacts, big `node_modules`/`vendor`/model weights, multi-GB image layers) — do **not** try to stuff that into the `lpk`.
- The project has many system-level or language-level dependencies (e.g. extensive `apt` / `yum` packages, multiple runtimes stacked, C extensions, CUDA/ML toolchains) where reproducing the build inside `lpk` would be fragile or slow.
- The build toolchain differs from the Lazycat runtime environment, or the user has already built a working image locally.

In any of these cases, do **not** try to pack the full build into the `lpk`. Take the image route:
1. Build the image locally for `linux/amd64`: `docker buildx build --platform linux/amd64 -t your-hub-user/app-name:v1.0 --load .`
2. Push the validated image to Docker Hub: `docker push your-hub-user/app-name:v1.0`
3. Sync it to the official Lazycat registry: `lzc-cli appstore copy-image your-hub-user/app-name:v1.0`
4. Backwrite `services.<name>.image` in the **source** `lzc-manifest.yml` with the returned `registry.lazycat.cloud/...` address. If the repo uses additional manifest templates or staged manifests during packaging, backwrite those sources too.
5. Keep the `lpk` focused on `package.yml`, `lzc-build.yml`, `lzc-manifest.yml`, icons, runtime scripts, and static assets. Do **not** attempt to pack the application image layers into the `lpk`.
6. Finish this image-sync and manifest-backfill work during migration, `make update`, or release preparation. By the time the user runs `make build` or `make install`, the source manifest used by packaging should already reference the final pullable image addresses.
7. When the user explicitly wants the whole release path to be executable in one command, prefer adding or using a dedicated `release-build` / `release-install` style command that runs: build image -> push public image -> `copy-image` -> backwrite source manifest -> build `.lpk` -> install `.lpk`.

**Official Publishing Phase:**
Before listing on the store, copy the image to the official managed registry for stability:
1. Execute: `lzc-cli appstore copy-image <public_image_name>`
2. Upon success, the tool returns a `registry.lazycat.cloud/...` address.
3. **Must** backwrite the image address in the source `lzc-manifest.yml` with this official address before packaging or publishing.

### 7. Store Listing and Review
When the user needs to formally list the app on the Lazycat App Store, read `references/store-publish.md` for the complete workflow and review rules.

## Platform-Specific Guardrails

When generating configuration files, you must comply with the following Lazycat MicroServer red-line rules:

1. **Inter-Service Communication Domains**
   - Never use `localhost` for cross-container communication.
   - For `application.routes` and most intra-app HTTP forwarding, prefer concrete service upstreams such as `http://your_service_name:80`.
   - Use the full internal domain `$service_name.$appid.lzcapp` only when you explicitly need the backend to receive that host or must disambiguate conflicting service names.
   - `lzc-manifest.yml` is not shell-templated. Do not leave `${lzcapp_appid}` or any other `${...}` placeholder in committed route targets.

2. **Persistent Storage Path Constraints**
   - Any application data that needs persistence **must** be mounted under `/lzcapp/var`.
   - To mount the microservice user's documents, use `/lzcapp/run/mnt/home`.
   - Never mount the system root or paths not starting with `/lzcapp` (unless using `compose_override`, which is not recommended for standard apps).

3. **HTTP Route Forwarding Prefix**
   - `application.routes` defaults to stripping the URL_PATH prefix. If the user needs to preserve it, suggest using `application.upstreams` with `disable_trim_location: true`.

4. **Forbidden Ports**
   - Unless in extreme circumstances, do not take over ports `80` and `443` via `ingress` (this breaks microservice authentication and routing).

5. **Setup Script (`setup_script`)**
   - If a container requires special initialization (e.g., modifying permissions or copying presets), use the `setup_script` field within `services` rather than forcing the user to rewrite the Dockerfile.

6. **Avoid Build Script Recursion**
   - **Never** execute `lzc project build` or `lzc-cli project build` within the `buildscript` (e.g., `build.sh`) defined in `lzc-build.yml`. Since `buildscript` is called by the `build` command, calling it internally will cause an infinite loop.

7. **Prioritize Docker over Source Code**
   - If a project provides a Docker image or `docker-compose.yml`, base the porting ENTIRELY on these Docker artifacts. **Do NOT** read or analyze the application's source code, as this wastes context. Just configure the `image:` in `lzc-manifest.yml` to use the provided Docker image.
   - **Auto-Translation for `docker-compose.yml`**:
     - `ports: ["8080:80"]` -> Convert to `routes` in `lzc-manifest.yml` (e.g., `- /=http://service_name:80`).
     - `volumes: ["./data:/app/data"]` -> Convert to `binds` mapping to `/lzcapp/var/` (e.g., `- /lzcapp/var/data:/app/data`).
     - `depends_on` -> **Keep it**, and convert the list form to the map form with `condition: service_healthy` per rule 9. Do not drop it; without gating, dependents boot against half-ready infra and crash. Service-to-service HTTP still uses `http://service_name:port` — the `depends_on` block is only about startup ordering.

8. **Settle Final Image Refs Before Install**
   - If the app depends on copied or bridged images, the returned test/official registry addresses must be written into the source `lzc-manifest.yml` and any manifest templates used by packaging before `make build` / `make install`.
   - `make install` is for packaging + installing the current LPK. It must not silently own image upload, `copy-image`, or manifest backfill responsibilities.
   - If the user asks for a complete release closure, add or use a separate release target instead of overloading `make install`.

9. **Gate Dependents on Healthchecks, Not Just Start Order**
   - **Enforce a layered startup order**: classify every service into infra (MySQL/PostgreSQL/Redis/ZooKeeper/etcd), middleware (Nacos/Consul/Kafka/RocketMQ/MinIO), one-shot seeds (schema import, config push, bucket init), and business (gateway/auth/api/workers). Each layer must be gated `service_healthy` on the layers below it; never let a business container race against a still-booting infra container. Seed containers sit between infra and middleware/business — business waits on the seed's terminal healthcheck, not on its mere start.
   - A bare `depends_on: [svc_a, svc_b]` (list form) only waits for the dependency to be **started**, not **healthy**. For slow-starting stacks (MySQL, Nacos, RocketMQ, Redis with ACL init, Kafka, ZooKeeper, any JVM service) this lets dependents boot against a half-ready infra and crash with "connection refused" / "auth not ready" / "topic missing" / "schema not initialized".
   - Use the **map form** with an explicit `condition: service_healthy` (compose-style):
     ```yaml
     depends_on:
       mysql:
         condition: service_healthy
       nacos:
         condition: service_healthy
     ```
     For one-shot seed containers (schema import, config push), define a terminal healthcheck (`test -f /tmp/<name>-ready`) and gate with `condition: service_healthy` so dependents actually wait until the seed finishes.
   - **Every** long-running service must define a `healthcheck` whose `test` command actually probes the ready state, not just process liveness:
     - HTTP/Java: `curl -fsS http://127.0.0.1:<port>/actuator/health` (or equivalent ready endpoint).
     - MySQL: `mysqladmin ping -h 127.0.0.1 -p"$${MYSQL_ROOT_PASSWORD}"`.
     - Redis: `redis-cli -a "$$PASS" ping | grep -q PONG`.
     - RocketMQ: probe the actual listener port with `nc -z` or grep the running command, not just `pgrep`.
     - Nacos/Consul/etcd: hit the readiness endpoint, not the liveness one.
   - **Do not guess the actuator / readiness path.** Read the service's actual config before writing the healthcheck URL:
     - Spring Boot: check `server.servlet.context-path`, `management.server.base-path`, `management.endpoints.web.base-path`. A service with `context-path: /api` exposes actuator at `/api/actuator/health`; a service without a context path exposes it at `/actuator/health`.
     - **Gateway / BFF / API-aggregator services need special care**: the `/api/*` prefix in such services is usually a downstream routing rule, **not** a prefix on the gateway's own actuator. Probing `/api/actuator/health` on the gateway will 404 forever even though the gateway itself is healthy, which starves every dependent. Probe the gateway's own actuator directly (`/actuator/health`) and verify by the startup log line (`Exposing N endpoints beneath base path '/actuator'`).
     - Do not blindly copy a healthcheck line from a sibling service — two Spring Boot services in the same stack may have different `context-path` settings.
     - When in doubt, `curl` the candidate path from inside the container during verification and confirm `{"status":"UP"}` before committing the manifest.
   - Tune `start_period`, `interval`, `timeout`, `retries` to the service's real cold-start budget. JVM services routinely need `start_period: 30s–60s`; MySQL first-boot initialization can need minutes (`start_period: 5m+`). Do not copy a generic `start_period: 5s` from a web-app template onto a Spring Boot service.
   - Main business services (gateway, auth, api, business logic) must list **every** infra dependency they talk to at boot, each gated on `service_healthy`. "Only one container has the issue" is usually because a sibling forgot to gate on the same infra.
   - If a dependency is genuinely optional and the app should survive its absence, mark it with `condition: service_started` plus in-app retry/backoff — do **not** use `service_healthy` on an optional dep (the whole stack will hang when that dep naturally fails).

10. **No Cross-Runtime Artifacts in the LPK**
   - Do **not** compile language-specific artifacts on the host (`.class`, `.pyc`, Go/Rust native binaries, `.so`/`.dylib`, CGO outputs) and drop them into the `lpk` to be executed inside a service container whose runtime differs from the host. This is how you produce `UnsupportedClassVersionError` / `GLIBC_X.Y not found` / wrong-arch crashes the moment a container loads them.
   - The lpk is a neutral carrier for `package.yml`, `lzc-build.yml`, `lzc-manifest.yml`, icons, shell scripts, static assets, and config. Anything that needs a specific language runtime belongs **inside the service image**.
   - Preferred order for helpers and healthchecks:
     1. Use tools already present in the service image (`curl`, `wget`, `sh`, `nc`) — no cross-env handoff at all.
     2. If a custom helper is unavoidable, build it **inside the service image's Dockerfile** (`RUN javac ...`, `RUN go build ...`) so it's produced by the matching runtime.
     3. Host-side compilation is a last resort and must pin the output to the lowest runtime that will execute it (`javac --release <service_jre>`, `GOOS/GOARCH` matching the image, static-linked binaries), and the produced artifact's target level must be verified, not assumed.
   - When auditing an existing repo, flag any `javac` / `pip install` / `go build` / `cargo build` step in host-facing scripts (`build.sh`, `Makefile`, `scripts/*.sh`) that writes into the lpk payload rather than into an image build context. That is a regression waiting to happen.

## Platform Compatibility Notes
If your platform supports automatic reading of referenced files, utilize that feature; otherwise, use your `read_file` tool to proactively read relevant specification documents in the `references/` directory.

## 官方规范参考文档 (Official Specifications)
在进行打包、构建、配置清单、设置部署参数及免密登录脚本注入时，必须严格参考并遵循以下官方规范文档：
- **Build Spec**: https://developer.lazycat.cloud/spec/build.html
- **Package Spec**: https://developer.lazycat.cloud/spec/package.html
- **Manifest Spec**: https://developer.lazycat.cloud/spec/manifest.html
- **Inject Context (免密登录抓取与持久化变量)**: https://developer.lazycat.cloud/spec/inject-ctx.html
- **Deploy Params**: https://developer.lazycat.cloud/spec/deploy-params.html
- **LPK Format**: https://developer.lazycat.cloud/spec/lpk-format.html
