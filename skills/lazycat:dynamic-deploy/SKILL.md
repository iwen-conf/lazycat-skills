---
name: "lazycat:dynamic-deploy"
description: 处理懒猫微服(Lazycat MicroServer)应用的动态部署参数配置(lzc-deploy-params.yml)、清单文件 Go 模板渲染以及利用 application.injects 实现免密登录（三阶段联动）和前端页面脚本注入的专业指南。
---

# 懒猫微服动态部署与免密注入指南

你是一个专业的懒猫微服应用架构师。当开发者需要向用户索要自定义配置（如密码、远程 IP 等），或者需要在不修改原应用代码的情况下，强行向应用的前端页面注入 JavaScript 脚本（特别是实现 **免密登录**）时，请严格遵循本指南。

## 1. 动态部署参数与模板渲染 (v1.3.8+)
懒猫微服支持在安装应用前，弹出一个 UI 界面让用户填写参数，然后利用这些参数动态渲染 `lzc-manifest.yml`。

### 步骤 A: 编写 `lzc-deploy-params.yml`
在项目根目录创建此文件，定义需要用户填写的字段。
```yaml
params:
  - id: target_ip
    type: string
    name: "目标服务器 IP"
    description: "你要代理的内网服务器 IP"
  - id: enable_debug
    type: bool
    name: "开启 Debug"
    default_value: "false"
    optional: true
```
*类型支持:* `string`, `bool`, `secret`, `lzc_uid`

### 步骤 B: 在 `lzc-manifest.yml` 中使用模板渲染
使用 Go 模板语法 (`{{ ... }}`) 读取参数。
- 用户参数使用 `.U.参数ID` (例如: `{{ .U.target_ip }}`)。如果 ID 包含 `.`，需使用 `index` (如 `{{ index .U "my.param" }}`)。
- 系统参数使用 `.S` (例如: `.S.BoxDomain`, `.S.IsMultiInstance`)。
- 随机密码生成: `{{ stable_secret "admin_password" | substr 0 8 }}` (同一个微服，相同的 seed 永远生成相同的字符串)。

**示例:**
```yaml
services:
  myapp:
    image: xxx
    environment:
      - REMOTE_IP={{ .U.target_ip }}
      - DB_PASS={{ stable_secret "db_root_pass" }}
```

## 2. 网页脚本注入与免密登录 (`application.injects`) (v1.5.0+)
适用于在不修改第三方 Docker 镜像前端代码的情况下，实现自动登录和免密体验。

### AI 智能决策逻辑 (AI Decision Logic)
当用户要求实现“免密登录”或“自动填充密码”时，你**必须**首先分析上游应用的情况，选择以下三种策略之一：

1. **基础 Auth 注入 (Basic Auth)**：上游没有复杂的 HTML 登录页，仅仅使用标准的 HTTP Basic Auth。你只需要在 `request` 阶段注入 `Authorization` HTTP 头即可。
2. **简单自动填充 (Simple Autofill)**：应用有登录页面，但密码在部署后**几乎不改变**（例如固定的管理员密码）。直接在 `browser` 阶段使用 `builtin://simple-inject-password` 填充从部署参数传入的密码。
3. **三阶段联动 (Three-Phase Linkage)**：**最复杂的场景（如 Jellyfin）**。应用拥有复杂 UI，用户可以在应用内**自主修改密码**。你必须监听用户的改密动作，让免密系统“记住”新密码。

> **🛑 强制红线 (STOP)**: 如果你判断需要使用【三阶段联动】策略，**绝对不要**凭空捏造 API 路径和选择器！你**必须**主动读取 `references/advanced-inject.md` 中的黄金模板和防呆清单，并在写代码前向用户询问相关的 API 接口路径。

**简单自动填充示例 (静态密码):**
```yaml
application:
  injects:
    - id: auto-login
      mode: exact
      include:
        - "/login"
        - "/#signin"
      scripts:
        - src: builtin://simple-inject-password
          params:
            user: "admin"
            password: "{{ stable_secret "app_admin_pass" }}"
            autoSubmit: true
```

## 平台特定的规则与护栏 (Guardrails)
1. **不要在 request/response 的 `when` 里写 Hash 路由**：服务端看不见 `#` 后面的内容。`/#login` 只能用在 `browser` 阶段（通常是 `include` 或省略 `on` 默认浏览器阶段）。
2. **`auth_required: false` 是关键**：在 `request` 和 `response` 阶段抓取密码时，必须加上 `auth_required: false`，否则在用户未登录时，请求会被网关直接拦截，你的脚本根本不会执行。
3. **不要直接在 request 阶段污染 persist**：只在 `request` 阶段用 `ctx.flow.set` 暂存候选密码，必须等到 `response` 阶段且判断 `ctx.status >= 200 && ctx.status < 300` 时，才写入 `ctx.persist`。

## 平台兼容性说明
如果需要查看详细的内置模板函数列表、系统参数列表（`SysParams`）或了解脚本注入的 `builtin://simple-inject-password` 的详细参数配置，请读取 `references/injects.md`。
如果是处理免密登录的【三阶段联动】，请读取 `references/advanced-inject.md` (AI 剧本) 或 `references/advanced-inject-passwordless-login.md` (详细教程)。
