---
name: "lazycat:advanced-routing"
description: 处理懒猫微服(Lazycat MicroServer)应用的高级路由、多域名配置、TCP/UDP四层转发(ingress)、跨域处理以及使用 app-proxy 进行复杂代理规则配置的专业指南。
---

# Lazycat MicroServer Advanced Routing and Network Configuration Guide

You are a professional expert in Lazycat MicroServer network configuration. Follow this guide when users encounter complex network forwarding requirements (such as multi-domain, Layer 4 forwarding, URL prefix stripping, custom Nginx proxying, etc.) during application porting or development.

## Core Routing Mechanisms

Lazycat MicroServer provides three levels of routing control. Select the most appropriate solution based on requirements:

### 1. Basic HTTP/HTTPS Routing (`application.routes`)
Suitable for most standard HTTP proxy scenarios.
**Rule Format:** `URL_PATH=UPSTREAM`
**Behavior:** By default, the `URL_PATH` prefix is **stripped**. For example, with `- /api/=http://backend:80`, a request to `/api/v1` reaches the backend as `/v1`.

Three upstream protocols are supported:
- `http(s)://$hostname/$path` (Most common; forwards to a container. Domain format must be `$service_name.$appid.lzcapp`).
- `file:///$dir_path` (Directly hosts static files).
- `exec://$port,$exec_file_path` (Starts an executable and proxies to a local port).

### 2. Advanced HTTP Routing (`application.upstreams`) (v1.3.8+)
Suitable for scenarios requiring fine-grained control over HTTP requests.
**Capabilities include:**
- **Domain-based Splitting**: Use `domain_prefix`.
- **Preserve URL Prefix**: Set `disable_trim_location: true`.
- **Resolve Host Validation Errors**: Set `use_backend_host: true`.
- **Skip SSL Verification**: Set `disable_backend_ssl_verify: true`.
- **Clear Specific Headers (Resolve CORS, etc.)**: Use `remove_this_request_headers: [Origin, Referer]`.

**Example:**
```yaml
upstreams:
  - location: /api
    backend: http://backend.cloud.lazycat.app.demo.lzcapp:80
    disable_trim_location: true # Preserves the /api prefix
```

### 3. TCP/UDP Layer 4 Forwarding (`application.ingress`)
**Never use `routes` to handle non-HTTP traffic!** If a user needs to expose SSH, databases, game servers, or other non-HTTP ports, you must use `ingress`.
**Warnings:**
- `ingress` provides only low-level network forwarding and has **no authentication protection**. Developers must handle security within the application.
- Unless in extreme circumstances, **it is strictly forbidden** to take over ports 80 and 443.

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

If `routes` and `upstreams` cannot meet requirements (e.g., extremely complex URL rewriting, different domains pointing to different backends, or the need for detailed request logs within the microservice), use the official `app-proxy` image.

**Image Address:** `registry.lazycat.cloud/app-proxy:v0.1.0` (essentially an OpenResty).

**Usage: Overriding Nginx Configuration via `setup_script`**
This is the most powerful way to handle multiple domains (coordinated with `application.secondary_domains`).

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
         server_name myadmin.*; # Matches additional domain
         location / { proxy_pass http://backend:8080; }
      }
      EOF
```

## Platform Compatibility Notes
To view the complete specifications for `routes`, `upstreams`, and `ingress`, or more detailed usage for APP Proxy, proactively read the relevant Markdown documents in the `references/` directory of this skill.
