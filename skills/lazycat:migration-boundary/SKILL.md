---
name: lazycat:migration-boundary
description: "Non-invasive migration feasibility and workload gate for Lazycat. Use after migration-license to decide allowed write scope, runtime mapping, LPK/image boundary, workload level, and whether to proceed, POC, or stop. 迁移可行性、不改业务代码、运行适配、工作量评估、是否继续."
---

# Lazycat Migration Feasibility and Workload Gate

## 职责范围

只评估已通过 `lazycat:migration-license` 的候选能否在不修改上游业务代码的前提下迁移到懒猫，并给出工作量等级和下一步决策。它不做候选搜索、不做许可证结论、不实施完整打包、不提审。

通过后交给 `lazycat:ship-app`；风险较大时先做 POC。

## 输入

- `lazycat:migration-license` 输出的候选记录。
- 仓库路径或只读 clone 结果。
- README、Dockerfile、Compose、镜像、配置样例、启动方式和部署文档。
- 用户明确授权的业务代码修改范围；没有授权时视为零业务代码修改。

## 凭据环境变量约定

如流程需要登录懒猫微服、应用商店或开发者中心，优先读取当前进程环境变量；只有变量缺失或为空时才向用户报告缺失项。

- 微服设备：`LAZYCAT_USERNAME` / `LAZYCAT_PASSWORD`
- 懒猫应用商店：`LAZYCAT_APPSTORE_USERNAME` / `LAZYCAT_APPSTORE_PASSWORD`
- 懒猫开发者中心：`LAZYCAT_DEVELOPMENT_USERNAME` / `LAZYCAT_DEVELOPMENT_PASSWORD`

使用约束：

1. 只通过环境变量读取凭据，不把值写入仓库文件、报告、截图说明、提交信息或长期记忆。
2. 不为确认变量而执行 `echo $LAZYCAT_PASSWORD`、`printenv LAZYCAT_PASSWORD` 等会输出明文的命令；只允许用 `test -n "${LAZYCAT_PASSWORD:-}"` 这类方式检查是否存在。
3. 浏览器自动化登录时直接填入变量值，日志和回复只记录变量名、站点和操作结果，不记录值。
4. CLI 登录或 `copy-image` 需要凭据时优先复用已有登录态；确需传参时避免把密码放在命令行参数、shell history 或可见输出中。
5. 凭据只解锁认证，不扩大本技能权限；提交、发布、安装覆盖、远程写操作仍按本技能边界和当前任务授权执行。

## 输出

- 非侵入迁移结论。
- 允许修改的包装层范围。
- 运行模型映射：镜像、服务、路由、持久化、依赖、初始化、登录、文件能力。
- LPK 包体和远程镜像桥接方案。
- 工作量等级、主要风险和下一步建议。

## 前置条件

1. 候选已通过 `lazycat:migration-license`，许可证为 `Pass` 或 `Pass with Obligations`。
2. 应用商店和开发者中心查重为 `Clear`，或用户明确接受差异化理由。
3. 已读取仓库的运行文档、Docker/Compose 证据和关键配置。
4. 如需查看懒猫官方事实，优先读取 `lazycat:ship-app/references/lpk/` 和相关本地文档。

## 允许执行

- 读取和分析上游源码、配置、Dockerfile、Compose、README、release 和部署文档。
- 修改或规划包装层文件：`package.yml`、`lzc-build.yml`、`lzc-manifest.yml`、`lzc-deploy-params.yml`、`Makefile`、`build.sh`。
- 修改或规划 Docker wrapper、启动脚本、seed/setup 脚本、配置模板、图标、商店素材和迁移文档。
- 构建或试运行用于 POC 的包装层和镜像，前提是不修改上游业务代码。
- 调用 `arc:frontend` 仅处理新增包装层、配置向导、审核辅助页或独立管理台。

## 禁止执行

- 未经当前任务明确授权，不修改上游业务前端页面/组件/路由/状态、后端 handler/service/domain/auth、数据库 schema/migration/model、测试或 fixture。
- 不为了启动、登录、健康检查、路由、审核或“尽快跑起来”修改业务源文件。
- 不把上游前端、移动端、桌面端或小程序客户端重写成默认平台栈。
- 不使用 `lzc-build.yml.images`、`embed:<alias>`、包内 `images/` 或 `images.lock`。
- 不把“能构建镜像”当作“能上架”；必须覆盖安装、启动、登录、持久化、核心流程和商店资料。
- 不处理许可证不明或查重未通过的项目；退回 `lazycat:migration-license`。

## 必查项

- 镜像路径：官方镜像、公开镜像、Dockerfile、自建镜像、源码构建。
- LPK 边界：最终 `.lpk` 是否可保持 `<= 12,000,000` bytes，且不内嵌镜像。
- 服务模型：长期进程、固定端口、健康检查、依赖排序。
- 持久化：数据目录、配置目录、权限、缓存策略。
- 初始化：数据库迁移、管理员创建、密钥生成、首启向导。
- 登录：无登录、固定账号、OIDC、inject、反向代理、二次验证。
- 文件能力：无文件流程、上传下载、文件选择器、文件关联。
- 外部依赖：域名、邮件、对象存储、第三方 API、GPU、宿主机能力、特权容器。
- 上架成本：图标、截图、描述、测试账号、复现步骤、版本来源。

## 决策

非侵入迁移结论：

- `Can migrate non-invasively`: 可以只改包装层完成迁移。
- `Can migrate with wrapper risk`: 可以包装，但启动、登录、文件、依赖或数据风险较高。
- `Cannot determine yet`: 关键事实缺失，需要补证据或 POC。
- `Blocked by business-code change requirement`: 必须修改上游业务代码，停止普通迁移。

工作量等级：

- `Small`: 单服务或简单 Compose，官方/公开镜像可用，持久化清晰，通常 1 天内。
- `Medium`: 2-4 个服务，需要健康检查、初始化脚本、固定凭据或 inject，通常 1-3 天。
- `Large`: 多服务、复杂初始化、登录或文件流程、外部依赖多、镜像需自建，通常 3-7 天，先 POC。
- `Too Large`: 非侵入不可控、依赖特权/不可替代外部服务、关键二进制或许可证缺失，停止。

建议：

- `Proceed`: 进入 `lazycat:ship-app`。
- `POC first`: 先做最小包装验证。
- `Switch project`: 候选可做但性价比低，换项目。
- `Stop`: 被业务代码、许可证、运行模型或包体门禁阻塞。

## 参考资料

- `references/porting-checklist.md`
- `references/s2i-strategy.md`
- `references/command-conventions.md`

## 后置条件

- 输出允许写入范围和禁止写入范围。
- 输出运行模型、LPK/镜像边界、工作量等级和下一步建议。
- 如果进入上架，必须带上许可证结论、查重结论、非侵入结论和工作量结论。
- 如果阻塞，必须给出唯一阻塞原因或证据缺口。

## 输出格式

```text
Phase: Migration Feasibility and Workload Gate
Repository: <owner/repo>

Inputs
- License gate:
- Duplicate checks:
- Source evidence:

Allowed Write Scope
- Packaging:
- Runtime:
- Assets/docs:

Runtime Model
- Image:
- LPK size/embed boundary:
- Services/routes:
- Persistence:
- Dependencies:
- Initialization:
- Login:
- File flows:
- External requirements:

Decision
- Feasibility: Can migrate non-invasively / Can migrate with wrapper risk / Cannot determine yet / Blocked by business-code change requirement
- Workload: Small / Medium / Large / Too Large
- Recommendation: Proceed / POC first / Switch project / Stop

Risks
- Business-code risks:
- Data risks:
- Store readiness:

Next: lazycat:ship-app / POC / Stop
```
