# Lazycat AI Pod 决策手册

适用于“项目要不要走懒猫算力仓 `AI应用` / AI 浏览器插件 / 普通应用接模型 API”这类判断。

先记住默认规则：

- 普通业务型 Web 应用接 AI，默认走 `BaseURL` 配置方案
- 只有明确做懒猫算力仓 / `AI应用` / AI 浏览器插件时，才走官方 AI Pod 路线

## 1. 官方入口

- `AI应用` 开发文档：
  `https://developer.lazycat.cloud/open/zh/guide/aipod/ai-application.html`
- AI 浏览器插件测试文档：
  `https://developer.lazycat.cloud/open/zh/guide/aipod/ai-browser-plugin-testing.html`
- Lazycat AI Pod 产品页：
  `https://lazycat.cloud/ai-pod`

## 2. 三种常见路线

### 普通应用 + 外部模型 API

适合：

- 你的产品本质上还是标准 Web 应用
- AI 只是其中一个功能，不是宿主入口
- 用户需要手动配置第三方模型服务

这是默认路线，不需要额外引入 `ai-pod-service/`、`caddy-aipod` 或 `extension.zip`。

创建阶段至少补：

- AI 设置页
- `API BaseURL`
- 协议选择：`OpenAI Compatible` / `OpenAI Responses` / `Anthropic`
- 获取模型按钮
- 模型下拉框
- 保存配置按钮

### 懒猫算力仓 `AI应用`

适合：

- 产品本身就是 AI 原生能力
- 需要和 Lazycat AI Pod / AI 浏览器深度联动
- 需要本地 AI 服务或独立 AI 工作流入口

只有当用户明确要做懒猫算力仓 / `AI应用` 时，才进入这条路线。

创建与发布阶段重点评估：

- 是否需要 `ai-pod-service/`
- 是否需要 `caddy-aipod`
- 是否需要在包中附带浏览器扩展
- 是否需要专门的 AI Pod 路由和中间件

### AI 浏览器插件

适合：

- 你的能力天然发生在浏览器页面上
- 需要划词、总结、改写、页面理解、网页侧栏或网页助手
- 插件形态比独立业务页面更贴近使用场景

只有当用户明确要做 AI 浏览器插件或浏览器侧能力时，才进入这条路线。

创建与发布阶段重点评估：

- 是否需要 `extension.zip`
- 是否需要和 `ai-pod-service/` 组合
- 是否需要通过 `caddy-aipod` 暴露服务给 AI 浏览器或前端

## 3. 官方包结构线索

根据官方 `AI应用` 文档，常见结构会包含这些元素：

- `lzc-build.yml`
- `lzc-manifest.yml`
- `ai-pod-service/`
- `caddy-aipod/`
- 可选的 `extension.zip`

不要在 skill 里硬编码更多细节字段；实际字段名和目录约束以当前官方文档为准。

## 4. 创建阶段的判断问题

至少回答这些问题：

- 这是不是普通业务型 Web 应用；如果是，为什么 `BaseURL` 配置已经足够
- 用户为什么要在 Lazycat 里安装它，而不是直接访问一个网页 AI 服务
- 它更像普通应用、`AI应用`，还是 AI 浏览器插件
- 它的首个高频入口是什么：独立页面、AI 浏览器入口，还是本地服务
- 如果不用 AI Pod 路线，为什么普通应用已经足够

## 5. 发布阶段的最小验证

如果项目走 `AI应用` / AI 浏览器插件路线，至少补这些验证：

- 相关目录和包结构齐全
- 服务能启动
- AI 浏览器或插件入口可见
- 关键 AI 流程可真实走通
- 商店资料与实际 AI 入口一致

## 6. 文档落点

推荐至少写进这些文档：

- `docs/requirements/lazycat-native-fit.md`
- `docs/api-design/ai-integration.md`
- `docs/architecture/lazycat-integration.md`
- `docs/architecture/aipod-integration.md`
