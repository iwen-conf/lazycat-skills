---
name: "lazycat:aipod-developer"
description: 懒猫AI算力舱(AI Pod)应用开发与打包规范。当用户需要构建一个部署到算力舱的AI应用、编写ai-pod-service的docker-compose.yml、配置Traefik路由规则、打包AI浏览器插件、或发布AI应用到商店时触发。
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
