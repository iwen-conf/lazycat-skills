---
name: "lazycat:aipod-developer"
description: Build and package AI apps for the Lazycat AI Pod (Computing Power Cabin). Use for ai-pod-service docker-compose, Traefik routing, GPU container config, and AI browser-extension packaging. Only for the official AI Pod route; standard web apps that merely call an AI API stay in lazycat:create-app. AI算力舱应用、ai-pod-service、Traefik路由、GPU、AI浏览器插件打包。
---

# Lazycat AI Pod Application Construction Guide

You are an expert in application development for Lazycat MicroServer AI Pods. This skill focuses on **how to build an AI application deployable to the Computing Power Cabin (AI Pod)**.

## Core Concepts

- **AI Application** = Microservice App + Pod AI Service (and/or AI Browser Extension).
- Installing to a microservice automatically deploys the AI service to the Pod.
- AI Pods are based on NVIDIA Jetson; Docker uses `nvidia-runtime` by default. **GPUs are directly available inside containers without explicit configuration.**
- Traefik is used as the service gateway, forwarding via Host rules. **Domains must end with `-ai`.**

## Action Instructions

When a user needs to build an AI Pod application, read `references/aipod-app-spec.md` for complete packaging specifications, `docker-compose` requirements, Traefik routing configuration, environment variable descriptions, browser extension packaging, progress indicator integration, and app publishing details.

---
**Mandatory Constraints for AI Engine:**
1. AI Pod Docker uses `nvidia-runtime` by default; do not add `gpus` or `runtime: nvidia` to `docker-compose`.
2. Traefik Host domains **must** end with `-ai`, otherwise forwarding will fail.
3. Services must join the `traefik-shared-network` to be managed by Traefik.
4. Use `LZC_AGENT_DATA_DIR` for persistence, `LZC_AGENT_CACHE_DIR` for cache, and `LZC_SERVICE_ID` for routing names.
5. To get the microservice name, execute `lzc-cli box default`; do not ask the user.
