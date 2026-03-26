# Source-to-Image (S2I) Migration Strategy for Lazycat

This document defines how to handle repositories that contain a `Dockerfile` but lack a pre-built remote image (e.g., no image name in `docker-compose.yaml` or local-only builds).

## Decision Matrix

| Criteria | Strategy A: External Build & Push | Strategy B: Direct Source Integration |
| --- | --- | --- |
| **Complexity** | High (Multi-stage builds, complex deps) | Low (Single binary, static files) |
| **Build Time** | Long (> 5 minutes) | Short (< 2 minutes) |
| **Stability** | High (Pre-validated image) | Medium (Dependent on LZC build environment) |
| **Recommended for** | Production-grade apps, SaaS ports | Tiny tools, scripts, simple web pages |

---

## Strategy A: External Build & Push (Recommended for Complexity)

If the project is large or has a complex `Dockerfile`, do **NOT** attempt to rebuild it inside the restricted LPK environment. Instead, follow this "Bridge" strategy:

1.  **Local Build**: Build the image locally ensuring the architecture is `linux/amd64`.
    ```bash
    docker buildx build --platform linux/amd64 -t your-hub-user/app-name:v1.0 .
    ```
2.  **Push to Hub**: Push the validated image to a public or private Docker Hub repository.
    ```bash
    docker push your-hub-user/app-name:v1.0
    ```
3.  **Sync to Lazycat**: Use the Lazycat CLI to pull the image into the platform's private registry.
    ```bash
    lzc-cli appstore copy-image your-hub-user/app-name:v1.0
    ```
4.  **Manifest Reference**: Update `lzc-manifest.yml` to use the image name returned by `copy-image`.

## Strategy B: Direct Source Integration

If the project is simple (e.g., a simple Go app or a static site), ignore the `Dockerfile` and use native Lazycat entry points:

1.  **Binary Mode**: Build the executable in `build.sh` and use `exec://` in `routes`.
2.  **Generic Base Image**: Use a standard base image (like `alpine` or `node:alpine`) and mount the source code into the container via the LPK structure.

---

## Red Flags

- **Wait for clarification**: If the `Dockerfile` relies on private base images or secret environment variables during build time, you **MUST** use Strategy A.
- **Architecture check**: Always ensure the image is `linux/amd64`. Arm64-only images will fail on most Lazycat nodes.
