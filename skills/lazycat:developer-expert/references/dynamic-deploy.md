# Lazycat Micro-service Dynamic Deployment and Injection Guide

You are a professional Lazycat Micro-service application architect. Follow this guide when developers need to request custom configurations from users (e.g., passwords, remote IPs) or forcibly inject JavaScript scripts into frontend pages without modifying the original application code.

## 1. Dynamic Deployment Parameters and Template Rendering (v1.3.8+)
Lazycat Micro-service allows an interface to pop up before app installation for users to fill in parameters. These parameters are then used to dynamically render `lzc-manifest.yml`.

### Step A: Write `lzc-deploy-params.yml`
Create this file in the project root to define the fields for the user to fill.
```yaml
params:
  - id: target_ip
    type: string
    name: "Target Server IP"
    description: "The internal server IP you want to proxy"
  - id: enable_debug
    type: bool
    name: "Enable Debug"
    default_value: "false"
    optional: true
```
*Supported types:* `string`, `bool`, `secret`, `lzc_uid`

### Step B: Use Template Rendering in `lzc-manifest.yml`
Use Go template syntax (`{{ ... }}`) to read parameters.
- User parameters: Use `.U.param_id` (e.g., `{{ .U.target_ip }}`). If the ID contains `.`, use `index` (e.g., `{{ index .U "my.param" }}`).
- System parameters: Use `.S` (e.g., `.S.BoxDomain`, `.S.IsMultiInstance`).
- Random secret generation: `{{ stable_secret "admin_password" | substr 0 8 }}` (The same seed on the same micro-service will always generate the same string).

**Example:**
```yaml
services:
  myapp:
    image: xxx
    environment:
      - REMOTE_IP={{ .U.target_ip }}
      - DB_PASS={{ stable_secret "db_root_pass" }}
```

## 2. Web Script Injection (`application.injects`) (v1.5.0+)
Use this to forcibly inject JS scripts into specific web pages (e.g., to auto-fill default passwords) without modifying third-party Docker frontend code.

**Core Logic:** Scripts are injected only when `include` (whitelist) is met and `exclude` (blacklist) is not.

**Example: Auto-login for third-party systems**
```yaml
application:
  injects:
    - id: auto-login
      mode: exact # Supports exact or prefix
      include:
        - "/login"      # Inject when visiting /login
        - "/#signin"    # Also matches hash routes
      scripts:
        # Use Lazycat's built-in form-filling script
        - src: builtin://simple-inject-password
          params:
            user: "admin"
            password: "{{ stable_secret "app_admin_pass" }}"
            autoSubmit: true
```

**Custom Injection Scripts:**
To inject your own scripts, place the JS file in the packaging directory and reference it via `file:///lzcapp/pkg/content/myscript.js`. Inside the script, access the `params` via `__LZC_INJECT_PARAMS__`.

## Platform Compatibility
To view detailed lists of built-in template functions, system parameters (`SysParams`), or configurations for `builtin://simple-inject-password` (e.g., modifying selectors), read the relevant Markdown documents in the `references/` directory.
