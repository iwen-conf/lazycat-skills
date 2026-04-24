# Changelog

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
