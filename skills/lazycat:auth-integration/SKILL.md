---
name: "lazycat:auth-integration"
description: 用于处理懒猫微服(Lazycat MicroServer)应用接入官方认证体系（OIDC单点登录）、HTTP Header用户身份识别、API Auth Token 以及配置独立鉴权(public_path)的专业指南。
---

# Lazycat MicroServer Authentication Integration Guide

You are a professional expert in Lazycat MicroServer authentication and permission configuration. Follow this guide when developers need to implement password-less login (OIDC integration), identify user information for the current request, or allow public access to specific APIs.

## 1. OIDC Single Sign-On (SSO) Integration
Lazycat MicroServer (v1.3.5+) provides unified OIDC support, allowing applications to automatically obtain user information and permission groups (`ADMIN` or `NORMAL`) for password-less login.

**Configuration Method (`lzc-manifest.yml`):**
1. Declare the OIDC redirect path (`application.oidc_redirect_path`). When the system detects this field, it automatically injects related environment variables during deployment.
2. Pass the system-generated OIDC environment variables to the application.

**Example:**
```yaml
application:
  subdomain: myapp
  oidc_redirect_path: /auth/oidc.callback # Mandatory! The system uses this to generate env vars. Refer to your app's OIDC docs for the exact path.
services:
  myapp:
    image: xxx
    environment:
      - OIDC_CLIENT_ID=${LAZYCAT_AUTH_OIDC_CLIENT_ID}
      - OIDC_CLIENT_SECRET=${LAZYCAT_AUTH_OIDC_CLIENT_SECRET}
      - OIDC_ISSUER_URI=${LAZYCAT_AUTH_OIDC_ISSUER_URI}
      - OIDC_AUTH_URI=${LAZYCAT_AUTH_OIDC_AUTH_URI}
      - OIDC_TOKEN_URI=${LAZYCAT_AUTH_OIDC_TOKEN_URI}
      - OIDC_USERINFO_URI=${LAZYCAT_AUTH_OIDC_USERINFO_URI}
```

## 2. HTTP Header Identity Recognition (Custom Backend)
If a user is developing their own backend code, `lzc-ingress` automatically injects the following HTTP Headers into all authenticated requests before they reach the application container. Developers can trust these headers directly.

- `X-HC-User-ID`: The logged-in username (UID).
- `X-HC-User-Role`: The user role (`NORMAL` or `ADMIN`).
- `X-HC-Device-ID`: The unique device ID of the client within the current microservice.
- `X-HC-Login-Time`: The Unix timestamp of the login time.

**Note:** The application backend can consider the user logged in based on `X-HC-User-ID` without further password verification.

## 3. Public Path Access (`public_path`)
By default, all HTTP requests must pass Lazycat MicroServer's mandatory authentication. If an app has its own authentication mechanism (e.g., Tokens) or is a public page (e.g., a share link), use `public_path` to exempt it.

**Configuration Method (`lzc-manifest.yml`):**
```yaml
application:
  public_path:
    - /api/public/  # Exempts paths starting with /api/public/
    - /share/       # Exempts paths starting with /share/
```
**Note:** For exempted paths, the system will still attempt to retrieve login status. If logged in, headers like `X-HC-User-ID` will still be present; if not logged in, headers are cleared but the **request is not intercepted**.

## 4. Scripts and Automation (API Auth Token)
When writing scripts (e.g., Python, Bash) to call microservice system APIs or application interfaces, you cannot rely on browser cookies. Lazycat (v1.4.3+) provides the `API Auth Token` mechanism.

**Retrieval:** Can only be generated via the microservice command line over SSH.
```bash
hc api_auth_token gen --uid admin
```
**Usage:** Include `Lzc-Api-Auth-Token: <token>` in the HTTP request headers.

## Platform Compatibility Notes
For more complex OIDC configurations or header interception issues, proactively read the relevant Markdown documents in the `references/` directory of this skill (`oidc.md`, `http-request-headers.md`, `public-api.md`, `api-auth-token.md`).

## 5. OIDC Troubleshooting Quick Reference

When OIDC integration fails, diagnose in this order:

### Callback 404
1. Verify `application.oidc_redirect_path` in `lzc-manifest.yml` exactly matches the path the app expects (e.g., `/auth/oidc.callback` vs `/auth/oidc/callback` — trailing slashes and dots matter).
2. Confirm the app's web framework is actually listening on that callback path and not behind a route prefix that gets stripped by `application.routes`.
3. If using `application.upstreams` with `disable_trim_location: false` (default), the prefix is stripped before reaching the app — the callback path must account for this.

### redirect_uri Mismatch
1. The OIDC provider auto-generates the redirect URI from `oidc_redirect_path` + the app's subdomain. If the app hardcodes a different redirect URI, the provider will reject it.
2. Check the app's OIDC config for any hardcoded `redirect_uri` and ensure it matches `https://<subdomain>.<box-domain>/<oidc_redirect_path>`.
3. Some apps require the redirect URI to be registered in their own config file — ensure both sides agree.

### Token Refresh Failures
1. Verify the app stores and uses the `refresh_token` returned by the OIDC provider.
2. Check if the OIDC token endpoint (`LAZYCAT_AUTH_OIDC_TOKEN_URI`) is reachable from inside the container: `lzc-cli docker exec <container> curl -s $LAZYCAT_AUTH_OIDC_TOKEN_URI`.
3. Some apps silently discard refresh tokens — check the app's OIDC library configuration for `offline_access` scope.

### Role Mapping Not Working
1. `X-HC-User-Role` is injected by `lzc-ingress` and has values `ADMIN` or `NORMAL` (uppercase). If the app expects lowercase or different role names, add a mapping layer.
2. If the request goes through `public_path`, the header may be empty for unauthenticated users — this is expected behavior.
3. Verify the app reads `X-HC-User-Role` from the HTTP headers, not from the OIDC userinfo endpoint (which uses a different field).

### User ID Not Received
1. `X-HC-User-ID` is only injected for authenticated requests. If the path is in `public_path` and the user is not logged in, the header is cleared.
2. Check that the app reads the header name exactly as `X-HC-User-ID` (case-sensitive in some frameworks).
3. If using `app-proxy` (Nginx), ensure `proxy_pass` preserves upstream headers — custom Nginx configs can accidentally strip them.
