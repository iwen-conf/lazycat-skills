
# Lazycat Micro-service Advanced Routing and Network Configuration Guide

You are a professional Lazycat Micro-service network configuration expert. When users encounter complex network forwarding requirements (e.g., multi-domain, Layer 4 forwarding, URL prefix stripping, custom Nginx proxying) during app porting or development, strictly follow this guide.

## Core Routing Mechanisms

Lazycat Micro-service provides three levels of routing control. Choose the most appropriate solution based on requirements:

### 1. Basic HTTP/HTTPS Routing (`application.routes`)
Suitable for most standard HTTP proxy scenarios.
**Rule Format:** `URL_PATH=UPSTREAM`
**Characteristics:** Strips the `URL_PATH` prefix by default. For example, with `- /api/=http://backend:80`, a request to `/api/v1` will reach the backend as `/v1`.

Supports three upstream protocols:
- `http(s)://$hostname/$path` (Most common; forwards to container. Domain format: `$service_name.$appid.lzcapp`)
- `file:///$dir_path` (Directly hosts static files)
- `exec://$port,$exec_file_path` (Starts an executable and proxies to a local port)

### 2. Advanced HTTP Routing (`application.upstreams`) (v1.3.8+)
Suitable for scenarios requiring fine-grained control over HTTP requests.
**Capabilities include:**
- **Domain-based Sharding:** Use `domain_prefix`.
- **Retain URL Prefix:** Set `disable_trim_location: true`.
- **Fix Host Validation Errors:** Set `use_backend_host: true`.
- **Skip SSL Verification:** Set `disable_backend_ssl_verify: true`.
- **Clear Specific Headers (Fix CORS, etc.):** Use `remove_this_request_headers: [Origin, Referer]`.

**Example:**
```yaml
upstreams:
  - location: /api
    backend: http://backend.cloud.lazycat.app.demo.lzcapp:80
    disable_trim_location: true # Retain /api prefix
```

### 3. TCP/UDP Layer 4 Forwarding (`application.ingress`)
**Never use `routes` for non-HTTP traffic!** If users need to expose SSH, databases, game servers, or other non-HTTP ports, they must use `ingress`.
**Warnings:** 
- `ingress` provides low-level network forwarding **without authentication protection**. Developers must handle security within the application.
- Except for extremely special cases, **actively hijacking ports 80 and 443 is strictly prohibited**.

**Example:**
```yaml
application:
  ingress:
    - protocol: tcp
      port: 3306
      service: mysql # Forwards to port 3306 of the mysql container
    - protocol: udp
      publish_port: 20000-30000 # Dynamic port range forwarding
      service: app
```

## Complex Reverse Proxy Best Practices (APP Proxy)

If `routes` and `upstreams` cannot meet your needs (e.g., extremely complex URL rewriting, multiple domains pointing to different backends, or detailed request logging), use the official `app-proxy` image.

**Image Address:** `registry.lazycat.cloud/app-proxy:v0.1.0` (Essentially OpenResty)

**Usage: Override Nginx configuration via `setup_script`**
This is the most powerful way to handle multiple domains (in conjunction with `application.secondary_domains`).

```yaml
application:
  subdomain: myapp
  secondary_domains:
    - myadmin
  routes:
    - /=http://app-proxy.cloud.lazycat.app.myapp.lzcapp:80
services:
  app-proxy:
    image: registry.lazycat.cloud/app-proxy:v0.1.0
    setup_script: |
      cat <<'EOF' > /etc/nginx/conf.d/default.conf
      server {
         server_name myapp.*; # Matches default domain
         location / { proxy_pass http://frontend:3000; }
      }
      server {
         server_name myadmin.*; # Matches secondary domain
         location / { proxy_pass http://backend:8080; }
      }
      EOF
```

## Platform Compatibility Notes
To view the full specification for `routes`, `upstreams`, and `ingress`, or for more detailed `app-proxy` usage, refer to the relevant Markdown documents in the `references/` directory of this skill pack.
