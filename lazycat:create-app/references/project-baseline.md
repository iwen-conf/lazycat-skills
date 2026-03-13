# Lazycat 项目基线

当项目进入“创建 / 初始化 / 技术栈收敛 / 认证补齐”阶段时，默认按这个基线执行。

## 0. 先文档后代码

第一步先建立 `docs/` 文档树，再开始脚手架和业务开发。至少包含：

- `docs/requirements/`
- `docs/api-design/`
- `docs/architecture/`
- `docs/release-prep/`

具体拆分参考 [docs-blueprint.md](./docs-blueprint.md)。

## 0.5 统一命令入口

每个项目都必须提供可执行的命令入口，至少包含：

- 根目录 `build.sh`
- 根目录 `Makefile`
- `make build`
- `make install`

推荐继续补：

- `make dev`
- `make test`
- `make clean`

命令不要只写在聊天记录或零散文档里，必须能在仓库根目录直接执行。

## 1. 默认技术栈

- 后端：Go
- 前端：Vue
- UI：Element Plus

如果仓库已有更细的团队规范，可以在这三项主基线之内细化，但不要偏离主栈。

## 1.5 激励优先模式

如果目标是懒猫现金激励，创建阶段就先做这几个判断：

- 优先原创应用，其激励空间通常高于普通移植应用
- 避开官方明确不发红包的类型
- 对需要账号密码的应用，确保普通用户能获得凭证
- 工具类应用优先规划文件关联
- 适合统一账户的应用优先规划微服 OIDC

## 2. 所有项目默认具备的认证能力

- 登录
- 注册
- `access_token`
- `refresh_token`
- 无感刷新
- 刷新失败后的清理与重新登录

## 3. 推荐认证流

### 登录

1. 用户提交账号和密码
2. 后端校验成功后签发 `access_token`
3. 后端同时签发 `refresh_token`
4. 前端保存鉴权状态并拉取当前用户信息

如果目标是激励优先，普通用户必须可以通过注册或其他公开可获得方式进入应用，不要把凭证获取变成阻塞。

### 启动恢复

1. 应用启动时尝试恢复当前用户状态
2. 如果 `access_token` 已失效但 `refresh_token` 仍有效，走一次静默刷新
3. 刷新成功后继续进入业务页
4. 刷新失败后清理状态并跳到登录页

### 401 无感刷新

1. 业务请求因 `access_token` 过期返回 401
2. 前端进入单飞刷新流程
3. 其他待发请求先挂起，等待刷新结果
4. 刷新成功后重放原请求
5. 刷新失败后统一退出登录

## 4. 最小后端接口

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `POST /auth/logout`
- `GET /auth/me`

## 5. 最小前端模块

- 登录页
- 注册页
- 用户状态 store
- 路由守卫
- 请求拦截器
- 刷新队列或单飞机制
- 退出登录清理逻辑

## 6. 推荐安全边界

- `access_token` 使用短生命周期
- `refresh_token` 使用长生命周期
- refresh 时执行 rotation
- refresh 失败后立即清理本地会话
- 环境变量中明确 token 密钥、过期时间、前端 API 地址

## 6.5 微服账户系统与文件关联

### OIDC

- 在 manifest 中设置 `application.oidc_redirect_path`
- 把系统生成的 OIDC 环境变量透传给应用
- 应用侧将 OIDC 登录映射为自身会话或用户上下文

### 文件关联

- 工具类应用评估 `application.file_handler`
- 声明 `mime` 和 `actions.open`
- 应用侧实现 `/open` 或等价路由，解析传入文件参数

## 7. 与发布链路的衔接

进入 `lazycat:ship-app` 之前，至少要保证：

- 项目可启动
- 登录可用
- 注册可用或有明确业务豁免
- token 刷新链路可用
- 如果目标是激励优先，激励资格路径已明确
- 首页和认证页可截图
