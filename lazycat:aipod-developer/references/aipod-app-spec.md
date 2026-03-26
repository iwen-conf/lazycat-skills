# AI Application Packaging and Publishing Specification

## AI Application Structure

An `AI Application` is a Micro-service app with an added `AI Service` or `AI Browser Extension`. Upon installation, the service is automatically deployed to the AI Pod.

### lpk Package Directory Structure

```
ai-pod-service/              # AI Service directory
  docker-compose.yml          # Service orchestration file
  README.md                   # Optional documentation
  check_ollama.py             # Optional check scripts, etc.
content.tar                   # Micro-service application content
extension.zip                 # Optional - Browser extension
icon.png                      # App icon
package.yml                   # App metadata (LPK v2)
lzc-manifest.yml              # App runtime manifest
```

## Adding an AI Service

Add the `ai-pod-service` field in `lzc-build.yml`:

```yml
# ai-pod-service: Specifies the service directory for the AI Pod; content will be packaged into the lpk.
ai-pod-service: ./ai-pod-service
```

This directory must contain a `docker-compose.yml`. Resources within the directory can be referenced via relative paths in the docker-compose file.

### Environment Variables Provided by AI Pod

1. **Data Persistence Path** `LZC_AGENT_DATA_DIR`
   - Definition: `/ssd/lzc-ai-agent/data/<service_id>`
   - Example: `/ssd/lzc-ai-agent/data/cloud.lazycat.aipod.ai`

   ```yml
   ollama:
     volumes:
       - ${LZC_AGENT_DATA_DIR}/data:/root/.ollama
   ```

2. **Data Cache Path** `LZC_AGENT_CACHE_DIR`
   - Definition: `/ssd/lzc-ai-agent/cache/<service_id>`

   ```yml
   ollama:
     volumes:
       - ${LZC_AGENT_CACHE_DIR}/cache:/root/.cache
   ```

3. **Service ID** `LZC_SERVICE_ID`
   - Corresponds to the `appId` in the LPK, with dots (`.`) removed.
   - Example: appId `cloud.lazycat.aipod.fish-speech` → `cloudlazycataipodfishspeech`

   ```yml
   ollama:
     labels:
       - "traefik.http.routers.${LZC_SERVICE_ID}-ollama.rule=Host(`ollama-ai`)"
   ```

> **Important**: Docker in the AI Pod uses `nvidia-runtime` by default. GPU access is available directly within containers; explicit `gpus` configuration is not required.

### Configuring AI Service Traefik Access Rules

The AI Pod uses `traefik` to forward services via `Host` rules.

1. **Host Rules** - The domain **must** end with `-ai`.

   ```yml
   services:
     ollama:
       labels:
         - "traefik.enable=true"
         - "traefik.http.routers.${LZC_SERVICE_ID}-ollama.rule=Host(`ollama-ai`)"
   ```

2. **Traefik Network** - Must join `traefik-shared-network`.

   Recommended to set as the default network:

   ```yml
   networks:
     default:
       external: true
       name: traefik-shared-network
   ```

   Or specify at the service level:

   ```yml
   services:
     myservice:
       networks:
         - traefik-shared-network
   ```

### Complete docker-compose.yml Example

```yml
services:
  whoami:
    image: registry.lazycat.cloud/traefik/whoami:ab541801c8cc
    labels:
      - "traefik.http.routers.whoami.rule=Host(`whoami-ai`)"

networks:
  default:
    external: true
    name: traefik-shared-network
```

## Adding an AI Browser Extension

Add the `browser-extension` field in `lzc-build.yml`:

```yml
# browser-extension: Browser extension, supports zip files or directories.
browser-extension: ./my-awesome-chrome-extension.zip
```

## Configuring Shortcuts

Configure in `lzc-manifest.yml`:

```yml
aipod:
  shortcut:
    disable: false  # Set to true to hide the shortcut in the AI browser.
```

The `aipod` field only supports `.SysParams(.S)` template parameters: `.BoxName` and `.BoxDomain`.

## Adding Startup Progress Prompts (caddy-aipod)

To display the AI Pod service startup progress, use the `caddy-aipod` middleware:

```yml
# lzc-manifest.yml
name: ComfyUI
package: cloud.lazycat.aipod.comfyui
version: 1.0.5
description: The most powerful open-source node-based generative AI application.

aipod:
  shortcut:
    disable: false

application:
  subdomain: comfyui
  routes:
    - /=http://caddy:80

services:
  caddy:
    image: registry.lazycat.cloud/catdogai/caddy-aipod:65e058ce
    setup_script: |
      cat <<'EOF' > /etc/caddy/Caddyfile
      {
              auto_https off
              http_port 80
              https_port 0
      }
      :80 {
              handle {
                      route {
                              lzcaipod
                              root * /lzcapp/pkg/content/ui/
                              try_files {path} /index.html
                              header Cache-Control "max-age=60, private, must-revalidate"
                              file_server
                      }
              }
      }
      EOF
      cat /etc/caddy/Caddyfile
```

Note: The `lzcaipod` directive in the `Caddyfile` detects if the AI Pod service is running and provides progress prompts.

## AI App Dependency on AI Pod Service Health Checks (Micro-service 1.3.8+)

Use `upstreams` to access the AI Pod's health check interface:

```yml
application:
  subdomain: comfyui
  gpu_accel: true
  routes:
    - /=file:///lzcapp/pkg/content/dist
  health_check:
    test_url: http://127.0.0.1/version
    start_period: 5m

  upstreams:
    - location: /version
      backend: https://comfyui-ai.{{ .S.BoxDomain }}/api/manager/version
      trim_url_suffix: /
      use_backend_host: true
      dump_http_headers_when_5xx: true
```

## Publishing AI Applications

1. AI Pod Category: https://appstore.lazycat.cloud/#/shop/category/27
2. Publishing is identical to standard Micro-service apps, via `lzc-cli` or the developer platform.
3. Include keywords like `AI Pod` or `算力舱` during submission to facilitate review and categorization.

## Targeting a Specific AI Pod

If multiple AI Pods are available, access them via the domain format: `f-{Pod-Serial}-{Service-Name}-ai.{Box-Name}.heiyu.space`

Example: `https://f-1420225016421-dozzle-ai.your-box-name.heiyu.space`
