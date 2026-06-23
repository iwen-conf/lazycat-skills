# Changelog

## 2026-06-24

- docs: 新增迁移项目上架原创性表单约束，要求提审时取消勾选“应用程序为本人原创开发或本人是源作者”，并填写原作者名称和源项目或代码地址。

## 2026-06-16

- docs: 重构 Lazycat Skill 架构，合并 `lazycat:unlisted-candidate-audit` 到 `lazycat:migration-license`，合并 `lazycat:migration-workload` 到 `lazycat:migration-boundary`，并统一保留 5 个单一职责入口。
- docs: 新增 `lazycat:update-installed-app`，用于通过 LPK Inspector 下载已安装应用包，对比 GitHub 最新版本和镜像，执行 `copy-image`、manifest 回写、重打包和安装验证。
- docs: 新增 `lazycat:unlisted-candidate-audit`，用于盘点开发者中心未上架且非待审核应用，对比应用商店和本地项目目录，并对缺失的 GitHub 项目做许可证与非侵入上架可行性审查。
- docs: 新增 LPK 包体强约束，要求所有构建产物不超过 12 MB（`12,000,000` bytes），并禁止 `lzc-build.yml.images`、`embed:<alias>` 和包内 `images/` / `images.lock` 内嵌镜像。

## 2026-06-13

- docs: 强化上架信息完整性要求，提交前必须填写完开发者中心资料，并记录来自最终提审包的 LPK 信息。
- docs: 强化 `package.yml.author` 规则，要求 GitHub `homepage` 的 owner 段与 `author` 逐字符一致，大小写和符号错误都会导致审核打回。

## 2026-06-11

- docs: 将 Lazycat 原创/上架前端联动中的小程序默认栈统一为 Taro 4，不再限定为多厂家小程序。
- docs: 扩展 `arc:frontend` 联动到移动端、桌面端和多厂家小程序，要求原创新增前端面使用平台默认栈，迁移项目继续保留上游客户端。
- docs: 将原创、迁移边界、迁移工作量和上架技能联动到 `arc:frontend` 默认 Web 前端栈，同时明确迁移项目不得为了统一技术栈重写上游业务前端。
- docs: 精简 Lazycat 技能入口，仅保留原创、迁移许可证、迁移非侵入边界、迁移工作量和上架 5 个 SKILL，并删除旧的拆散技能入口。
- docs: 将官方文档和 LPK 规范迁移到 `lazycat:ship-app/references/`，同步更新 README 与 AGENTS 中的本地知识库路径。
- docs: 强化迁移业务代码红线，明确除非用户在当前任务说明允许或需要修改业务代码并点名范围，否则迁移、上架和排障都禁止修改上游业务源码。
- docs: 扩展 `lazycat:migration-license` 为迁移候选发现入口，要求搜索 Web/Agent 且有页面和后端的 GitHub 项目，并同时对比懒猫应用商店和开发者中心待审列表。

## 2026-06-09

- docs: 新增 GitHub 与第三方仓库操作红线，要求使用用户 GitHub 账号进行 issue、PR、fork、评论、review 等涉及他人的操作前必须得到当前对话中的显式允许。
- fix: 收紧 `lazycat:port-app` 的 GitHub 调研边界，明确默认只读，不得擅自代表用户对上游仓库发起可见互动。
- fix: 移除本地 `store-rule.md` 中鼓励 Fork/红包激励的操作性内容，并将 OIDC 示例中的真实微服域名替换为占位符。
- docs: 移除运行时中文 UI 作为上架硬性门槛的规则，将 `zh-CN` UI 改为按产品目标和用户群体决定的本地化建议。

## 2026-06-07

- docs: 补充 Lazycat `package.yml.author` 规则，要求 GitHub 项目从 `homepage` 的 owner 段推导作者，例如 `https://github.com/Makisuo/maple` 对应 `author: Makisuo`。

## 2026-06-01

- docs: 重写全部 15 个技能的 `description`，改为简洁英文摘要加少量中文触发词，去除冗长关键词堆砌；标准 React 前端技术栈串只在 `ui-ux-pro-max` 保留一次，其余技能改为引用，并补充技能间的衔接提示（create-app/port-app → admin-ui/lpk-builder/ship-app），同步更新 README 技能列表。
- docs: 新增文件选择器强制规则，要求迁移应用通过 `application.injects` 接入官方文件选择器自动拦截，原创应用在业务 UI/代码中内置懒猫文件选择能力，并同步到相关技能质量门禁。

## 2026-05-19

- docs: 强化 `lazycat:troubleshoot` 排障写入边界，要求移植/包装语境下即使日志定位到后端 agent/API 等上游 bug，也必须先走部署参数、环境变量、wrapper/runtime 方案；无法非侵入修复时输出 `Blocked by business-code change requirement`，不得直接修改业务代码。
- docs: 将移植写入边界提升为 `developer-expert` 全局继承规则，并同步收紧 `ship-app`、`update-app`、`auth-integration`，防止提审修复、版本更新、OIDC 接入或后台截图质量流程绕过“不得修改上游业务代码”红线。

## 2026-05-18

- docs: 将 Lazycat 移植的“不要修改业务代码”升级为写入白名单、停止条件和交付自检，要求遇到必须改上游源码才能启动/登录/过健康检查时先阻塞并请求明确授权。

## 2026-05-16

- docs: 强化 Lazycat 迁移边界，要求移植默认只改包装层和运行时初始化，禁止 AI 默认修改上游业务代码。
- docs: 补充迁移应用默认声明 `android`、`ios`、`tvos` 不支持，并要求 `package.yml.locales` 使用 BCP 47 语言标签和登录使用说明。
- docs: 强化免密登录规范，新增固定初始账号/密码/昵称、启动期创建用户、三阶段 inject 学习改密、healthcheck 与启动顺序联动要求。
- docs: 精简技能仓库，删除聚合版官方文档快照和拆分脚本，将 `developer-expert` 下重复长参考文档改为垂直技能指针。
- docs: 将历史启动失败经验沉淀到排障和移植规则，强化 ingress healthcheck 二级症状、数据库首启 readiness、one-shot seed/migration、部署参数渲染缺失等诊断路径。
- docs: 新增懒猫运行机理与迁移前判定参考，要求移植前先明确交付形态、入口、存储、依赖层、初始化和登录风险。
- docs: 将“能否在懒猫上运行”固化为四态门禁，要求先验证镜像架构、长期进程、监听端口、持久化、依赖、初始化、登录和外部要求。

## 2026-05-14

- docs: 回滚远程语义索引知识源改造，恢复仓库内离线懒猫开发文档和历史技能目录。
- docs: 强化本地知识库检索协议，要求事实型问题优先读取 `references/docs/` 和垂直技能 `references/`。
- docs: 明确禁止把远程语义索引、外部长期记忆、云端向量库或付费索引服务作为默认工作流。

## 2026-04-24

- docs: 收紧懒猫项目的 `Makefile` 约定，明确 `make build` 只负责基于已准备好的镜像引用与交付物打包 `lpk`，不承担源码级构建。
- docs: 明确 `make install` 依赖一次 `make build` 后执行 `lzc-cli app install <lpk>`，并同步 `lazycat:lpk-builder` 与 `developer-expert` 的安装说明。
- docs: 强化 `lazycat:port-app`，要求 AI 在确认迁移后实际完成仓库里的 `Makefile`，不能只停留在建议或待办。

## 2026-04-21

- fix: 收紧 `lazycat:ui-ux-pro-max`，仅保留 `React + Vite + Tailwind CSS + shadcn/ui + Zustand + TanStack Query + React Router + React Hook Form + Zod + Framer Motion` 这一套前端基线入口。
- fix: 将 `create-app`、`admin-ui`、`ship-app`、`developer-expert` 等技能中的前端基线统一改为指定的 React 技术栈，并移除旧的抽象前端栈引用。
- docs: 更新 `README.md` 中对 `lazycat:ui-ux-pro-max` 的描述，明确其面向 React 技术栈。
- fix: 继续清理 `lazycat:ui-ux-pro-max` 知识库中的 3D / VisionOS / WebXR / WebGL / Spline 等残留映射，避免检索结果偏离 React 基线。
- docs: 再次对齐技能文档中的默认前端基线表述，统一使用 `React + Vite + Tailwind CSS + shadcn/ui + Zustand + TanStack Query + React Router + React Hook Form + Zod + Framer Motion`，移除额外语言层描述，确保技能入口口径一致。
