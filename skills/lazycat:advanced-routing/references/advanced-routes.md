# Advanced Routing

## Introduction
An official `APP Proxy` image is maintained to help developers implement complex routing functions and view request logs. The `APP Proxy` is based on OpenResty.
Image address: `registry.lazycat.cloud/app-proxy:v0.1.0`.

## Usage
There are two primary configuration modes:
- **Environment Variables**: Suitable for cases with a single HTTP upstream service.
- **setup_script**: Directly overwrites the OpenResty configuration file, allowing any configuration supported by OpenResty.

::: danger
Do **NOT** mix these two modes.
:::

### 1. Environment Variables
APP Proxy abstracts specific features for developers unfamiliar with Nginx/OpenResty, allowing quick configuration via environment variables:

| Environment Variable | Function | Example |
| - | - | - |
| `UPSTREAM` (Required) | Sets the upstream HTTP service for the proxy | `UPSTREAM=http://whoami:80` |
| `BASIC_AUTH_HEADER` | Sets the `Authorization` header to bypass Basic Auth | `BASIC_AUTH_HEADER="Basic dXNlcjpwYXNzd29yZA=="` |
| `REMOVE_REQUEST_HEADERS` | Removes HTTP request headers (semicolon-separated) | `REMOVE_REQUEST_HEADERS="Origin;Host;"` |

### 2. setup_script
Before starting, ensure you understand the [principles of setup_script](advanced-setupscript.md) and Nginx configuration.

You can overwrite the OpenResty configuration file or write Lua scripts within the `setup_script` for complex configurations.

**Simple Example:**
```yaml
lzc-sdk-version: '2.0'
application:
  routes:
    # Forward requests to APP Proxy (app-proxy service)
    - /=http://app-proxy:80
  subdomain: app-proxy-test
services:
  app-proxy:
    image: registry.lazycat.cloud/app-proxy:v0.1.0
    setup_script: |
      # Overwrite OpenResty configuration
      cat <<'EOF' > /etc/nginx/conf.d/default.conf
      server {
         server_name  app-proxy-test.*;
         location / {
            root   /usr/local/openresty/nginx/html;
            index  index.html index.htm;
         }
      }
      EOF
```

## Examples

### Viewing Application Request Logs
When using APP Proxy, you can view request logs via `lzc-cli docker logs -f`.
For example: `lzc-cli docker logs -f cloudlazycatappapp-proxy-test-app-proxy-1`.

```yaml
lzc-sdk-version: '2.0'
application:
  routes:
    - /=http://app-proxy:80
  subdomain: app-proxy-test
services:
  app-proxy:
    image: registry.lazycat.cloud/app-proxy:v0.1.0
    environment:
      - UPSTREAM="http://whoami:80"
  whoami:
    image: registry.lazycat.cloud/snyh1010/traefik/whoami:c899811bc4a1f63a
```

### Bypassing Basic Auth
Set the `BASIC_AUTH_HEADER` to inject an `Authorization` header, enabling auto-login.
Value format: `Basic base64(username:password)`.

Example: For `user:password`, the base64 is `dXNlcjpwYXNzd29yZA==`.
```yaml
services:
  app-proxy:
    image: registry.lazycat.cloud/app-proxy:v0.1.0
    environment:
      - UPSTREAM="http://whoami:80"
      - BASIC_AUTH_HEADER="Basic dXNlcjpwYXNzd29yZA=="
```

### Removing Request Headers
Use `REMOVE_REQUEST_HEADERS` to strip specific headers like `Origin` or `Cache-Control`.
```yaml
services:
  app-proxy:
    image: registry.lazycat.cloud/app-proxy:v0.1.0
    environment:
      - UPSTREAM="http://whoami:80"
      - REMOVE_REQUEST_HEADERS="Origin;Cache-Control;"
```

### Multi-Domain Support
Lazycat supports [multiple domains per application](advanced-secondary-domains.md). Combined with `setup_script`, you can route different domains to different backends.

Example routing:
- `app-proxy-test.*` -> OpenResty default page
- `portainer.*` -> Portainer service
- `whoami.*` -> whoami service

```yaml
application:
  routes:
    - /=http://app-proxy.cloud.lazycat.app.app-proxy-test.lzcapp:80
  subdomain: app-proxy-test
  secondary_domains:
    - portainer
    - whoami
services:
  app-proxy:
    image: registry.lazycat.cloud/app-proxy:v0.1.0
    setup_script: |
      cat <<'EOF' > /etc/nginx/conf.d/default.conf
      server {
         server_name  app-proxy-test.*;
         location / {
            root   /usr/local/openresty/nginx/html;
            index  index.html index.htm;
         }
      }
      server {
         server_name  portainer.*;
         location / {
            proxy_pass http://portainer:9000;
         }
      }
      server {
         server_name  whoami.*;
         location / {
            proxy_pass http://whoami:80;
         }
      }
      EOF
```
