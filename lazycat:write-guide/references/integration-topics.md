# Lazycat 对接文章主题

如果要写“如何对接 Lazycat”的技术文章，优先围绕这些主题：

- 路由规则：`advanced-route`
- 微服账户系统：`advanced-oidc`
- 文件关联：`advanced-mime`
- 懒猫算力仓 `AI应用`
- AI 浏览器插件
- 应用启动与依赖
- 环境变量与配置注入
- 提审与发布流程

写法建议：

- 先讲为什么要接入
- 再讲 manifest / 配置怎么写
- 再讲应用代码怎么改
- 最后讲如何验证接入成功

如果主题是 AI Pod，再额外补：

- 为什么它更适合普通应用、`AI应用`，还是 AI 浏览器插件
- 如果是 `AI应用`，`ai-pod-service/`、`caddy-aipod`、`extension.zip` 分别承担什么角色
- 用户最终从哪里进入这项能力
