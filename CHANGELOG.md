# Changelog

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
