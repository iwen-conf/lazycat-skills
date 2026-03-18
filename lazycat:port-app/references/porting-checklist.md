# Lazycat 移植检查清单

## 1. 选型前

- [ ] GitHub 搜索已完成
- [ ] 许可证允许移植
- [ ] 已记录上游地址
- [ ] App Store 查重已完成
- [ ] 如果重复度高，已终止激励路径或说明差异化

## 2. 项目落地

- [ ] 已建立 `docs/requirements`
- [ ] 已建立 `docs/api-design`
- [ ] 已建立 `docs/architecture`
- [ ] 已建立 `docs/release-prep`
- [ ] 已创建 `build.sh`
- [ ] 已创建 `Makefile`
- [ ] 已有 `make build`
- [ ] 已有 `make install`

## 3. Lazycat 适配

- [ ] 已准备 `lzc-build.yml`
- [ ] 已准备 `lzc-manifest.yml`
- [ ] 已评估是否需要 OIDC
- [ ] 已评估是否需要 `file_handler`
- [ ] 需要登录的应用已明确凭证获取路径
- [ ] 已区分凭证作用域：`lazycat_account` / `lazycat_password` 用于进入懒猫微服和 App Store，`lazycat_developer_center_account` / `lazycat_developer_center_password` 用于开发者中心，应用内登录按具体 app 级变量读取
- [ ] 如果上游是 AI 项目，已判断是否更适合懒猫算力仓 `AI应用` / AI 浏览器插件

## 4. 激励路径

- [ ] 不属于官方不奖励类型
- [ ] 原创 / 移植路径已明确
- [ ] 如为移植，已准备上游归因
- [ ] 如适合，已规划 OIDC 或 `file_handler`
- [ ] 如目标是激励，已规划后续把应用真实安装到懒猫微服并打开已安装版本验证核心能力
- [ ] 如为 AI 原生项目，已规划普通应用 / `AI应用` / AI 浏览器插件路线
