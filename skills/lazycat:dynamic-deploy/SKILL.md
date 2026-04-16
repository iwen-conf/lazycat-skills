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

### AI 强制决策顺序 (必须先判断，不能跳过)
当用户要求“免密登录 / 自动登录 / 自动填充密码”时，必须按下面顺序做决策：

1. **先判断源码是否可改**：如果当前仓库就是应用源码，或者登录页与改密页都在可修改的应用代码里，默认优先做**应用内三阶段联动**，不要先造 sidecar、代理换 token 页面或额外认证服务。
2. **只有源码不可改时，才优先走 inject**：inject 适合第三方镜像、闭源前端、或无法稳定改上游代码的场景。
3. **写 manifest 前先判断 inject 语法代际**：如果目标盒子或安装日志提示 `application.injects.0 when is required`，或者 `lzc-cli project build` 对 `mode/include/scripts` 报 unknown fields，就必须立刻回退到旧语法 `on / when / do`，不要继续赌新版字段会生效。
4. **禁止带着 lint warning 直接交付**：manifest 构建通过不代表安装能通过。只要 inject 字段出现 unknown field warning，就必须先解决兼容性再继续。

### AI 智能决策逻辑 (AI Decision Logic)
当用户要求实现“免密登录”或“自动填充密码”时，你**必须**首先分析上游应用的情况。**根据实际经验，大部分应用都应优先采用“三阶段联动”。** 具体策略如下：

1. **应用内三阶段联动 (App-native Three-Phase Linkage) - 👑源码可改时优先**：如果当前项目的登录页、登录成功动作、改密成功动作都可修改，就直接在应用代码内实现三阶段：登录成功持久化 -> 改密成功更新 -> 登录页自动填充/自动提交。这个方案通常比 runtime sidecar 更稳定，也更容易回退。
2. **inject 三阶段联动 (Three-Phase Linkage) - 👑第三方镜像的最常见方案**：应用拥有复杂 UI，允许用户在应用内首次创建账号或**自主修改密码**（如 Jellyfin），但源码不可改。你必须监听并联动三个阶段：在 `request` 阶段拦截登录/注册/改密请求并提取密码 -> 在 `response` 阶段确认请求成功后持久化 (`ctx.persist`) -> 在 `browser` 阶段从持久化存储读取并自动填充。这种方案适配绝大多数第三方现代应用。
3. **简单自动填充 (Simple Autofill)**：应用有登录页面，且密码在部署后**固定不变**（例如通过部署参数传入固定管理员密码）。直接在 `browser` 阶段使用 `builtin://simple-inject-password` 填充。
4. **基础 Auth 注入 (Basic Auth)**：上游没有 HTML 登录页，仅仅使用标准的 HTTP Basic Auth。只需要在 `request` 阶段注入 `Authorization` HTTP 头即可。

> **🛑 强制红线 (STOP)**: 既然“三阶段联动”是主流方案，**绝对不要**凭空捏造 API 路径和选择器！你**必须**主动读取 `references/advanced-inject-passwordless-login.md`（或 `advanced-inject.md`）中的教程模板，并在写代码前主动向用户询问或确认相关的 初始化/登录/改密 API 接口路径及表单选择器。

**简单自动填充示例 (静态密码):**
```yaml
application:
  injects:
    - id: auto-login
      # 省略 on 参数，默认为 browser 阶段
      when:
        - "/login"
        - "/#signin"
      do:
        - src: builtin://simple-inject-password
          params:
            user: "admin"
            password: "{{ stable_secret "app_admin_pass" }}"
            autoSubmit: true
```

## 平台特定的规则与护栏 (Guardrails)
1. **不要在 request/response 的 `when` 里写 Hash 路由**：服务端看不见 `#` 后面的内容。`/#login` 只能用在 `browser` 阶段（通常是 `when` 或省略 `on` 默认浏览器阶段）。
2. **`auth_required: false` 是关键**：在 `request` 和 `response` 阶段抓取密码时，必须加上 `auth_required: false`，否则在用户未登录时，请求会被网关直接拦截，你的脚本根本不会执行。
3. **不要直接在 request 阶段污染 persist**：只在 `request` 阶段用 `ctx.flow.set` 暂存候选密码，必须等到 `response` 阶段且判断 `ctx.status >= 200 && ctx.status < 300` 时，才写入 `ctx.persist`。

## 平台兼容性说明
如果需要查看详细的内置模板函数列表、系统参数列表（`SysParams`）或了解脚本注入的 `builtin://simple-inject-password` 的详细参数配置，请读取 `references/injects.md`。
如果是处理免密登录的【三阶段联动】，请读取 `references/advanced-inject.md` (AI 剧本) 或 `references/advanced-inject-passwordless-login.md` (详细教程)。
如果目标盒子存在 inject 新旧语法兼容问题，也必须读取 `references/injects.md` 中的兼容性说明，再决定用 `mode/include/scripts` 还是 `on/when/do`。
