# Lazycat Micro-service Authentication System Integration Guide

You are a professional Lazycat Micro-service authentication and permission configuration expert. Follow this guide when developers need to implement passwordless login (OIDC), identify current request user information, or allow public access to certain APIs.

## 1. OIDC Single Sign-On (SSO) Integration
Lazycat Micro-service (v1.3.5+) provides unified OIDC support, allowing applications to automatically obtain user information and permission groups (`ADMIN` or `NORMAL`) for passwordless login.

**Configuration Method (`lzc-manifest.yml`):**
1. Declare the OIDC callback path (`application.oidc_redirect_path`). When this field is detected, the system automatically injects relevant environment variables during deployment.
2. Pass these system-generated OIDC variables to the application.

**Example:**
```yaml
application:
  subdomain: myapp
  oidc_redirect_path: /auth/oidc.callback # Mandatory! Used to generate environment variables. Check your app's OIDC docs for the correct path.
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

## 2. HTTP Headers Identity Recognition (Custom Backend)
For custom backend code, `lzc-ingress` automatically injects the following HTTP headers into all authenticated requests before they reach the application container. Developers can trust these headers directly.

- `X-HC-User-ID`: The logged-in Username (UID).
- `X-HC-User-Role`: The user role (`NORMAL` or `ADMIN`).
- `X-HC-Device-ID`: The client's unique device ID within the current micro-service.
- `X-HC-Login-Time`: Unix timestamp of the login time.

**Note:** The application backend can assume the user is authenticated based on the presence of `X-HC-User-ID`, eliminating the need for further password verification.

## 3. Independent Auth and Public Access (`public_path`)
By default, all HTTP requests require mandatory authentication via Lazycat Micro-service. If an application has its own authentication mechanism (e.g., Token) or public pages (e.g., shared links), use `public_path` to bypass mandatory login.

**Configuration Method (`lzc-manifest.yml`):**
```yaml
application:
  public_path:
    - /api/public/  # Allows access to paths starting with /api/public/
    - /share/       # Allows access to paths starting with /share/
```
**Note:** For bypassed paths, the system still attempts to retrieve login status. If the user is logged in, headers like `X-HC-User-ID` will still be present. If not logged in, these headers will be empty, but the **request will not be intercepted**.

## 4. Scripts and Automated Calls (API Auth Token)
Browser cookies cannot be used for scripts (e.g., Python, Bash) calling micro-service system APIs or application interfaces. Lazycat (v1.4.3+) provides the `API Auth Token` mechanism for this purpose.

**Generation:** Tokens can only be generated via the micro-service command line via SSH.
```bash
hc api_auth_token gen --uid admin
```
**Usage:** Include the token in the HTTP request header as `Lzc-Api-Auth-Token: <token>`.

## Platform Compatibility
For complex OIDC configurations, header interception issues, or further details, refer to the relevant Markdown documents in the `references/` directory (`oidc.md`, `http-request-headers.md`, `public-api.md`, `api-auth-token.md`).
