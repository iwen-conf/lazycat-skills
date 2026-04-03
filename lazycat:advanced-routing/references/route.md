# Routing Rules
=========

The `application.routes` field is an array of `Rule` objects.

Rules are declared in the format `URL_PATH=UPSTREAM`, where `URL_PATH` is the actual URL path accessed by the browser (excluding the hostname) and `UPSTREAM` is the specific upstream service.

Supported protocols for `UPSTREAM`:
- `file:///$dir_path`
- `exec://$port,$exec_file_path`
- `http(s)://$hostname/$path`

**Note:** By default, `application.routes` strips the `URL_PATH` prefix during forwarding.
Example:
```yaml
routes:
  - /api/=http://backend:80
```
A request to `/api/v1` will be received by the backend as `/v1`. To preserve the prefix, use `application.upstreams` and set `disable_trim_location: true` (lzcos v1.3.9+).

## HTTP Upstream
Supports both internal and external services. For example, the built-in App Store lzcapp uses a single line:
```yaml
routes:
    - /=https://appstore.lazycat.cloud
```
Requests to `https://appstore.$MicroserviceName.heiyu.space` are forwarded to `https://appstore.lazycat.cloud`. In this case, client-side JS can still use `lzc-sdk/js` features. The installation logic runs in the microservice, while the code is hosted on the public cloud for centralized maintenance.

Typically, HTTP routes forward to a specific port of an internal service.
Example: `Bitwarden`
```yaml
lzc-sdk-version: '2.0'
application:
  routes:
  - /=http://bitwarden.cloud.lazycat.app.bitwarden.lzcapp:80
  - /short=http://bitwarden:80
  subdomain: bitwarden
services:
  bitwarden:
    image: bitwarden/nginx:1.44.1
```
1. `bitwarden:80` uses the service name, which resolves to the container's IP at runtime.
2. The full format is `$service_name.$appid.lzcapp`.

**Important Considerations:**
- In lzcos 1.3.x+, application isolation is enforced. Directly using the `service_name` is recommended for simplicity and ease of appid modification.
- Use the `$service_name.$appid.lzcapp` format if:
  1. There are potential name conflicts (e.g., common names like `app` or `db`) in environments without isolation.
  2. The upstream service requires the correct `Host` header. By default, the `Host` header will be the `service_name`. Use `upstreams.[].use_backend_host=true` to pass the backend host explicitly if needed.

## File Upstream
Used for serving static HTML files.
Example: `PPTist` (pure frontend app)
```yaml
application:
  subdomain: pptist
  routes:
    - /=file:///lzcapp/pkg/content/
  file_handler:
    mime:
      - x-lzc-extension/pptist
    actions:
      open: /?file=%u
```
Static assets are typically included in the LPK package and reside in `/lzcapp/pkg/content/` as read-only at runtime.

## Exec Upstream
Format: `exec://$port,$exec_file_path`. This is a special rule composed of:
1. `$port`: The port providing the service (implies `127.0.0.1`).
2. `$exec_file_path`: Path to the executable (script or ELF binary).

At startup, the system executes `$exec_file_path` and assumes the service is available at `http://127.0.0.1:$port`. This can also be used for initialization tasks.

Example: `Lazycat Cloud Drive`
```yaml
application:
  routes:
    - /api/=exec://3001,/lzcapp/pkg/content/backend
    - /files/=http://127.0.0.1:3001/files/
    - /=file:///lzcapp/pkg/content/dist
```

## UpstreamConfig (v1.3.8+)
The `application.upstreams` field allows for more granular control:

```yaml
subdomain: debug
routes:
  - /=http://app1.org.snyh.debug.whoami.lzcapp:80

upstreams:
  - location: /search
    backend: https://baidu.com/
    use_backend_host: true  # Required for most external servers
    disable_auto_health_checking: true
    remove_this_request_headers:
      - origin
      - Referer
    disable_url_raw_path: true # Normalizes the raw URL path

  - location: /other
    backend: https://app2.snyh.debug.lzcapp:4443
    disable_backend_ssl_verify: true # Skip SSL verification for self-signed certificates

  - location: /
    domain_prefix: config
    backend: http://config.snyh.debug.lzcapp:80

  - location: /api
    backend: http://127.0.0.1:3001/
    backend_launch_command: /lzcapp/pkg/content/my-super-backend -listen :3001
```
