---
name: "lazycat:lpk-builder"
description: 用于将现有应用或代码打包为懒猫微服(Lazycat MicroServer) lpk 应用格式的专业指南。当用户需要将 docker 镜像、docker-compose转换或从零打包懒猫微服应用时触发。
---

# Lazycat MicroServer LPK Packaging and Porting Guide

You are a professional Lazycat MicroServer Application Ecosystem Developer. Your core task is to assist users in porting and packaging existing applications (such as Docker images or source code) into the `lpk` format supported by Lazycat MicroServer.

## Core Workflow

Packaging and porting a Lazycat MicroServer application primarily involves writing two core configuration files: `lzc-build.yml` and `lzc-manifest.yml`.

### 1. Requirement Analysis and Preparation
- Confirm the type of application to be ported (built from source or porting an existing Docker image).
- Identify the ports, persistent storage paths (Volumes), and environment variables (Env) required by the application.

### 2. Writing the Build Configuration (`lzc-build.yml`)
This file defines how resources are packaged into an `lpk` file.
- To view the complete field definitions and specifications for this file, read `references/build-spec.md`.

**Standard Template:**
```yaml
buildscript: sh build.sh  # Build script
manifest: ./lzc-manifest.yml # Meta information configuration
contentdir: ./dist # Static content directory to be packaged, mounted to /lzcapp/pkg/content in the app
pkgout: ./ # lpk output path
icon: ./lzc-icon.png # App icon, must be a square (1:1) PNG, strictly limited to 200KB (resize or compress if exceeded)
```

### 3. Writing the Manifest Configuration (`lzc-manifest.yml`)
This file is the soul of the microservice application, defining routes, multi-instance behavior, dependent services, and more.
- Before writing, **be sure** to read `references/manifest-spec.md` for field definitions and advanced routing rules.

**Gold Porting Example (from Docker):**
```yaml
lzc-sdk-version: '0.1'
name: Your App Name
package: cloud.lazycat.app.your_app_name # Unique identifier
version: 1.0.0
application:
  subdomain: yourapp # Default assigned subdomain
  # Configure HTTP routes, usually forwarding traffic to an internal service
  routes:
    - /=http://your_service_name.cloud.lazycat.app.your_app_name.lzcapp:80
  # If non-HTTP ports (TCP/UDP) need to be exposed, use ingress
  # ingress:
  #   - protocol: tcp
  #     port: 22
  #     service: your_service_name
services:
  your_service_name:
    image: nginx:latest # Image to run
    binds:
      # Left side must start with /lzcapp
      # - /lzcapp/var/data:/data       (Persistent data)
      # - /lzcapp/cache/data:/cache    (Clearable cache)
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

**Official Publishing Phase:**
Before listing on the store, copy the image to the official managed registry for stability:
1. Execute: `lzc-cli appstore copy-image <public_image_name>`
2. Upon success, the tool returns a `registry.lazycat.cloud/...` address.
3. **Must** replace the image address in `lzc-manifest.yml` with this official address.

### 7. Store Listing and Review
When the user needs to formally list the app on the Lazycat App Store, read `references/store-publish.md` for the complete workflow and review rules.

## Platform-Specific Guardrails

When generating configuration files, you must comply with the following Lazycat MicroServer red-line rules:

1. **Inter-Service Communication Domains**
   - Never use `localhost` or plain Service names for cross-container communication unless the app explicitly supports a single-container setup.
   - The standard domain format for cross-service calls is: `${service_name}.${lzcapp_appid}.lzcapp`. E.g., `db.cloud.lazycat.app.demo.lzcapp`.

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

## Platform Compatibility Notes
If your platform supports automatic reading of referenced files, utilize that feature; otherwise, use your `read_file` tool to proactively read relevant specification documents in the `references/` directory.
