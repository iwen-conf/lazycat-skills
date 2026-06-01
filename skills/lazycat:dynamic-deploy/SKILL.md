---
name: "lazycat:dynamic-deploy"
description: Configure dynamic deploy params (lzc-deploy-params.yml) and Go-template manifest rendering, and use application.injects to inject frontend scripts non-invasively — passwordless login (three-stage flow) and file-picker auto-intercept for migrated apps. 动态部署参数、清单模板渲染、application.injects免密登录、文件选择器自动拦截、脚本注入。
---

# 懒猫微服动态部署与免密注入指南

你是一个专业的懒猫微服应用架构师。当开发者需要向用户索要自定义配置（如密码、远程 IP 等），或者需要在不修改原应用代码的情况下，强行向应用的前端页面注入 JavaScript 脚本（特别是实现 **免密登录** 或 **文件选择器自动拦截**）时，请严格遵循本指南。

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

1. **禁止修改源码**：绝对不要尝试修改应用自身的源代码来实现免密登录，无论是前端还是后端代码。
2. **统一使用 inject 三阶段联动**：始终采用 `application.injects` 机制，即参考 Jellyfin 示例的“三阶段联动”方案（Request拦截 -> Response确认持久化 -> Browser自动填充）。
3. **写 manifest 前先判断 inject 语法代际**：如果目标盒子或安装日志提示 `application.injects.0 when is required`，或者 `lzc-cli project build` 对 `mode/include/scripts` 报 unknown fields，就必须立刻回退到旧语法 `on / when / do`，不要继续赌新版字段会生效。
4. **禁止带着 lint warning 直接交付**：manifest 构建通过不代表安装能通过。只要 inject 字段出现 unknown field warning，就必须先解决兼容性再继续。

### AI 智能决策逻辑 (AI Decision Logic)
当用户要求实现“免密登录”或“自动填充密码”时，你**必须**统一且唯一地优先采用 **inject 三阶段联动**，并在不需要改动代码的前提下实现。具体策略如下：

1. **inject 三阶段联动 (Three-Phase Linkage - 👑 唯一首选方案)**：应用拥有复杂 UI，允许用户在应用内首次创建账号或自主修改密码（如 Jellyfin）。在不修改源码的前提下，你必须监听并联动三个阶段：
   - **第一阶段 (Request)**：在请求阶段拦截登录/注册/改密请求，提取用户名和密码存入临时变量 `ctx.flow`。
   - **第二阶段 (Response)**：在响应阶段检查 HTTP 状态码是否成功（2xx），只有成功时才将 `ctx.flow` 中的数据转存到持久化存储 `ctx.persist`。
   - **第三阶段 (Browser)**：在浏览器阶段从 `ctx.persist` 读取账号密码，并利用内置插件 `builtin://simple-inject-password` 自动填充到输入框。
   详细示例参考：https://developer.lazycat.cloud/advanced-inject-passwordless-login.html#%E7%A4%BA%E4%BE%8B%E4%B8%89-%E4%B8%89%E9%98%B6%E6%AE%B5%E8%81%94%E5%8A%A8-jellyfin
2. **简单自动填充 (Simple Autofill)**：如果密码在部署后**绝对固定不变**（例如通过部署参数传入固定管理员密码），可退化为直接在 `browser` 阶段使用 `builtin://simple-inject-password` 填充。
3. **基础 Auth 注入 (Basic Auth)**：上游没有 HTML 登录页，仅仅使用标准的 HTTP Basic Auth 时，直接在 `request` 阶段注入 `Authorization` HTTP 头即可。

### 固定初始账号 + 可修改式免密登录（迁移应用默认组合）
当第三方应用需要内部账号，并且上游支持通过启动命令、CLI、环境变量或管理 API 创建用户时，必须优先使用这个非侵入组合，而不是修改业务代码：

1. **启动期创建固定初始用户**：用 `setup_script`、wrapper `entrypoint`/`command`、或 one-shot seed service 在服务启动后幂等创建账号。seed 必须等待主应用 healthcheck 通过，成功后写入完成标记，并提供自己的 terminal healthcheck。
2. **固定凭据必须可见**：在 README、商店使用说明或 `package.yml.locales.<lang>.usage` 中写明：
   - 账号：`<固定账号>`
   - 密码：`<固定密码>`
   - 昵称：`<固定昵称>`
3. **后续允许用户改密**：使用三阶段联动监听登录、初始化和改密请求；只有 response 2xx 后才把新账号/密码写入 `ctx.persist`；登录页从 `ctx.persist` 自动填充。这样初始账号固定，但用户改密后免密登录会跟随新密码。
4. **健康检查和启动顺序绑定**：数据库/缓存先健康，主应用再健康，seed 最后创建用户并健康；浏览器注入只解决登录体验，不能掩盖后端未初始化或 seed 未完成的问题。
5. **禁止猜测**：创建用户命令、API 路径、payload 字段和输入框选择器必须来自上游文档、容器内 CLI help、运行时网络请求或用户提供的信息。

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
            user: "lazycat"
            password: '{{ stable_secret "app_admin_pass" }}'
            autoSubmit: true
```

## 3. 文件选择器自动拦截 (`lzc-file-chooser-inject.js`)
当迁移应用已经有浏览器原生文件能力，但没有接入懒猫网盘时，必须优先使用官方文件选择器自动拦截 inject，而不是修改上游前端业务代码。官方依据：https://developer.lazycat.cloud/lazycat-file-picker-auto-intercept.html。

### 强制适用条件

1. **迁移应用必须走 inject**：只要页面存在打开、保存、上传、下载、`showOpenFilePicker()`、`showSaveFilePicker()` 或 `<input type="file">` 等文件入口，就在 `application.injects` 中注入 `lzc-file-chooser-inject.js`，让用户选择“本地文件系统 / 懒猫微服”。
2. **原创应用必须内置**：如果是从零开发或可控业务源码的原创应用，应在应用 UI/代码中直接接入懒猫文件选择能力；不要把 inject 当作原创应用的默认实现。
3. **无文件入口则记录不适用**：如果应用没有文件打开、保存、上传或下载入口，在移植说明或提审自检中写明“不适用：无文件入口”，不要硬塞 inject。
4. **不得把参数写进脚本文件名**：`diskRoot`、`fallbackMime`、`locale`、`text`、`hooks` 必须写在 `do[].params` 下。

### Manifest 模板

```yaml
application:
  injects:
    - id: open-save-chooser
      on: browser
      when:
        - /*
      do:
        - src: file:///lzcapp/pkg/content/lazycat-injects/lzc-file-chooser-inject.js
          params:
            diskRoot: /_lzc/files/home
            fallbackMime: application/octet-stream
            locale: auto
            hooks:
              fileSystemAccess: true
              fileInput: true
```

如果应用还需要声明文件关联，同时配置 `application.file_handler`，例如：

```yaml
application:
  file_handler:
    mime:
      - x-lzc-extension/example
    actions:
      open: /?fileUrl=/%u
```

### 验证要求

1. 打开应用并触发一次打开/上传入口，确认出现“本地文件系统 / 懒猫微服”选择弹窗。
2. 选择“懒猫微服”，确认能打开懒猫文件选择器并把文件交还给原页面流程。
3. 触发一次保存/下载入口（如果应用支持），确认能选择保存到本地或懒猫微服。
4. 如果 inject 不生效，先核对 `on: browser`、`when` 命中路径、脚本是否被打进 `contentdir`、目标盒子的 inject 语法代际，再进入排障。

## 平台特定的规则与护栏 (Guardrails)
1. **不要在 request/response 的 `when` 里写 Hash 路由**：服务端看不见 `#` 后面的内容。`/#login` 只能用在 `browser` 阶段（通常是 `when` 或省略 `on` 默认浏览器阶段）。
2. **`auth_required: false` 是关键**：在 `request` 和 `response` 阶段抓取密码时，必须加上 `auth_required: false`，否则在用户未登录时，请求会被网关直接拦截，你的脚本根本不会执行。
3. **不要直接在 request 阶段污染 persist**：只在 `request` 阶段用 `ctx.flow.set` 暂存候选密码，必须等到 `response` 阶段且判断 `ctx.status >= 200 && ctx.status < 300` 时，才写入 `ctx.persist`。

## 平台兼容性说明
如果需要查看详细的内置模板函数列表、系统参数列表（`SysParams`）或了解脚本注入的 `builtin://simple-inject-password` 的详细参数配置，请读取 `references/injects.md`。
如果是处理免密登录的【三阶段联动】，请读取 `references/advanced-inject.md` (AI 剧本) 或 `references/advanced-inject-passwordless-login.md` (详细教程)。
如果目标盒子存在 inject 新旧语法兼容问题，也必须读取 `references/injects.md` 中的兼容性说明，再决定用 `mode/include/scripts` 还是 `on/when/do`。

## 官方规范参考文档 (Official Specifications)
在进行打包、构建、配置清单、设置部署参数及免密登录脚本注入时，必须严格参考并遵循以下官方规范文档：
- **Build Spec**: https://developer.lazycat.cloud/spec/build.html
- **Package Spec**: https://developer.lazycat.cloud/spec/package.html
- **Manifest Spec**: https://developer.lazycat.cloud/spec/manifest.html
- **Inject Context (免密登录抓取与持久化变量)**: https://developer.lazycat.cloud/spec/inject-ctx.html
- **Deploy Params**: https://developer.lazycat.cloud/spec/deploy-params.html
- **LPK Format**: https://developer.lazycat.cloud/spec/lpk-format.html
