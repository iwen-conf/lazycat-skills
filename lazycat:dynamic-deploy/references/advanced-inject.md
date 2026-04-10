# AI 自动化剧本：高级免密注入与三阶段联动 (Three-Phase Linkage)

**给所有 AI 智能体的强制指令 (CRITICAL INSTRUCTION):**
在实现免密登录时，**绝对不要**试图猜测上游应用的 API 路径、请求体 (Payload) 结构或 CSS 选择器！你**必须**要么自己通过工具检查应用的源码/网络请求，要么**直接向用户询问**这些具体细节。在得到明确的接口信息之前，**严禁**生成任何 `lzc-manifest.yml` 代码！

生成错误的注入脚本会给用户带来巨大的调试负担。请严格遵守此标准操作程序 (SOP)。

## 1. 核心策略：三阶段联动
当应用拥有复杂的 UI 界面，用户不仅能登录，还能**自主修改密码**时，必须使用此策略。该策略确保免密登录能够“学习”并同步用户的最新密码。

**工作原理：**
1.  **阶段 1 (Request):** 拦截登录、初始化设置或修改密码的 API 请求。从请求体 (Payload) 中提取候选的用户名/密码，并临时存入 `ctx.flow`。
2.  **阶段 2 (Response):** 等待服务端响应。只有当服务端返回成功状态码（2xx）时，才将 `ctx.flow` 中的候选凭据持久化到 `ctx.persist` 中。这防止了密码输入错误时污染了持久化存储。
3.  **阶段 3 (Browser):** 在登录页面上，使用 `builtin://simple-inject-password` 读取 `$persist` 的值来自动填充。在“修改密码”页面上，自动填充“当前密码”输入框（但不自动提交）。

## 2. 强制信息收集阶段 (MANDATORY GATHERING)
在编写 YAML 之前，你必须掌握以下信息：
1.  **登录 API 路径与 Method：** (例如：`POST /api/login`)
2.  **改密 API 路径与 Method：** (例如：`PUT /api/user/password`)
3.  **JSON Payload 结构：** 请求体中的账号、密码、新密码字段具体叫什么？ (例如：`{"usr": "...", "pwd": "..."}`)
4.  **前台登录页路由：** (例如：`/login` 或 `/#/signin`)
5.  **前台改密页路由：** (例如：`/settings` 或 `/#/profile`)
6.  *(可选但推荐)* **CSS 选择器：** 如果 `simple-inject-password` 无法自动检测，需要提供账号、密码、"当前密码"输入框的精确 `#id` 或 `.class`。

*如果你没有这些信息，停下来，立刻向用户提问！*

## 3. 黄金防呆模板 (Boilerplate Template)
请使用以下精确的模板结构。**禁止遗漏** `auth_required: false`，**禁止**在 request/response 的 `when` 规则中使用 `#` 路由哈希。

```yaml
application:
  injects:
    # -------------------------------------------------------------------
    # 阶段 1: REQUEST - 提取候选凭据暂存到 flow
    # -------------------------------------------------------------------
    - id: app-capture-password
      auth_required: false # 必须填写！否则未登录时的请求会直接被拦截，导致脚本不执行
      on: request
      when:
        - /api/login        # 请替换为真实的 API 路径 (绝对不能带 # 号！)
        - /api/user/pass*   # 支持通配符匹配
      do: |
        // 1. 安全路由：确保只处理正确的 Method 和 Path
        const path = String(ctx.request.path || "");
        const method = String(ctx.request.method || "").toUpperCase();
        
        // 请替换为目标应用的真实逻辑
        const isLogin = path === "/api/login" && method === "POST";
        const isChange = path.startsWith("/api/user/pass") && method === "PUT";
        
        if (!isLogin && !isChange) return;

        // 2. 安全解析：请求体可能为空或不是 JSON
        let payload = null;
        try { payload = ctx.body.getJSON(); } catch { return; }
        if (!payload || typeof payload !== "object") return;

        // 3. 提取字段：根据该应用的真实 JSON 结构修改 keys
        const pickString = (...values) => values.find((v) => typeof v === "string" && v.length > 0) ?? "";
        
        // 注意：这里的 payload.xxx 必须替换为真实的 key
        const username = pickString(payload.username, payload.email); 
        const password = pickString(payload.new_password, payload.password);

        // 4. 存入临时的 flow 上下文
        if (username) ctx.flow.set("app_pending_username", username);
        if (password) ctx.flow.set("app_pending_password", password);

    # -------------------------------------------------------------------
    # 阶段 2: RESPONSE - 当响应成功时持久化凭据
    # -------------------------------------------------------------------
    - id: app-commit-password
      auth_required: false # 必须填写！
      on: response
      when:
        - /api/login        # 必须与阶段 1 的路径完全一致
        - /api/user/pass*
      do: |
        // 1. 只在服务端返回成功(2xx)时才持久化，防止密码输错污染数据库
        if (ctx.status < 200 || ctx.status >= 300) return;

        // 2. 从 flow 中读取
        const username = ctx.flow.get("app_pending_username");
        const password = ctx.flow.get("app_pending_password");
        
        // 3. 存入长期的 persist 数据库
        // 请将 "app." 替换为该应用的唯一前缀标识
        if (typeof username === "string" && username.length > 0) {
          ctx.persist.set("app.username", username);
        }
        if (typeof password === "string" && password.length > 0) {
          ctx.persist.set("app.password", password);
        }

    # -------------------------------------------------------------------
    # 阶段 3: BROWSER - 在前台登录页自动填充
    # -------------------------------------------------------------------
    - id: app-login-autofill
      when:
        - /login       # 替换为真实的前台路由路径 (此处允许使用 # 号，如 /#login)
      do:
        - src: builtin://simple-inject-password
          params:
            user:
              $persist: app.username  # 必须与 ctx.persist.set 的 KEY 一致
            password:
              $persist: app.password  # 必须与 ctx.persist.set 的 KEY 一致
            # 如果自动寻找输入框失败，请显式指定选择器：
            # userSelector: "#usernameInput"
            # passwordSelector: "#passwordInput"

    # -------------------------------------------------------------------
    # 阶段 3: BROWSER - 在前台改密页自动填充“当前密码”
    # -------------------------------------------------------------------
    - id: app-change-password-autofill
      when:
        - /settings/profile  # 替换为真实的前台路由路径 (此处允许使用 # 号)
      do:
        - src: builtin://simple-inject-password
          params:
            password:
              $persist: app.password  # 必须与 ctx.persist.set 的 KEY 一致
            
            # 严重警告：我们只希望自动填充【旧密码】输入框。
            # 我们绝对不能填充【新密码】输入框，也绝对不能触发自动提交！
            autoSubmit: false
            
            # 通常必须显式指定“当前密码”的选择器，以防填充到“新密码”框。
            # passwordSelector: "#currentPasswordInput"
```

## 4. AI 强制自查清单 (Verification Checklist)
在向用户输出你的 YAML 代码前，**必须**在心中核对以下问题：
- [ ] 我是否在 request 和 response 两个阶段都加上了 `auth_required: false`？
- [ ] request/response 阶段的 `when` 匹配规则中，是否**完全没有** `#` 哈希符号？
- [ ] JS 中的 `ctx.body.getJSON()` 是否已经被 `try...catch` 包裹？
- [ ] response 阶段的注入中，是否包含了 `if (ctx.status < 200 || ctx.status >= 300) return;` 防御逻辑？
- [ ] JS 里 `ctx.persist.set` 使用的键名，和 YAML 里 `$persist:` 读取的键名，是否一字不差地完全一致？
- [ ] 在改密页面的自动填充阶段，我是否写上了 `autoSubmit: false` 并指定了旧密码的 `passwordSelector`？