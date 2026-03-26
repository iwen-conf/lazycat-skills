---
name: "lazycat:dynamic-deploy"
description: 处理懒猫微服(Lazycat MicroServer)应用的动态部署参数配置(lzc-deploy-params.yml)、清单文件 Go 模板渲染以及利用 application.injects 实现前端页面脚本注入的专业指南。
---

# Lazycat MicroServer Dynamic Deployment and Injection Guide

You are a professional Lazycat MicroServer Application Architect. Follow this guide when developers need to request custom configurations from users (e.g., passwords, remote IPs) or forcibly inject JavaScript scripts into front-end pages without modifying the original application code.

## 1. Dynamic Deployment Parameters and Template Rendering (v1.3.8+)
Lazycat MicroServer supports displaying a UI for users to fill in parameters before installation, which are then used to dynamically render `lzc-manifest.yml`.

### Step A: Create `lzc-deploy-params.yml`
Create this file in the project root to define the fields for the user.
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

### Step B: Template Rendering in `lzc-manifest.yml`
Use Go template syntax (`{{ ... }}`) to access parameters.
- User parameters: `.U.ParamID` (e.g., `{{ .U.target_ip }}`). If the ID contains `.`, use `index` (e.g., `{{ index .U "my.param" }}`).
- System parameters: `.S` (e.g., `.S.BoxDomain`, `.S.IsMultiInstance`).
- Random password generation: `{{ stable_secret "admin_password" | substr 0 8 }}` (the same microservice and seed will always generate the same string).

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
Used to forcibly inject JS scripts into specific web pages (e.g., to auto-fill default passwords) without modifying third-party Docker image front-end code.

**Core Logic:** Scripts are injected only into HTML pages that match `include` (whitelist) and do not match `exclude` (blacklist).

**Example: Automated Login for Third-Party Systems**
```yaml
application:
  injects:
    - id: auto-login
      mode: exact # Supports exact or prefix
      include:
        - "/login"      # Inject when visiting /login
        - "/#signin"    # Also matches hash routes
      scripts:
        # Use built-in form-filling script
        - src: builtin://simple-inject-password
          params:
            user: "admin"
            password: "{{ stable_secret "app_admin_pass" }}"
            autoSubmit: true
```

**Custom Injection Scripts:**
To inject your own scripts, place the JS file in the package directory and reference it via `file:///lzcapp/pkg/content/myscript.js`. Within the script, access passed `params` via `__LZC_INJECT_PARAMS__`.

## Platform Compatibility Notes
To view the detailed list of built-in template functions, system parameters (`SysParams`), or parameters for `builtin://simple-inject-password` (such as modifying selectors), proactively read the relevant Markdown documents in the `references/` directory of this skill.
