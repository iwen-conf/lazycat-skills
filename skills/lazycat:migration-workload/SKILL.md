---
name: lazycat:migration-workload
description: "Third gate for migrating a GitHub project to Lazycat: estimate whether the non-invasive migration workload is small, medium, large, or too large. 迁移第三关、工作量评估、值不值得做."
---

# Lazycat Migration Workload Gate

本技能是迁移第三关，只评估“在不修改业务代码的前提下，工作量大不大”。没有通过许可证和非侵入边界判断时，不使用本技能直接拍脑袋估工。

## 使用场景

- 项目已通过 `lazycat:migration-license`。
- 项目已通过或基本通过 `lazycat:migration-boundary`。
- 用户要决定是否开工、换项目、先做 POC，或直接进入上架打包。
- 候选项目来自 `lazycat:migration-license` 的搜索筛选，需要根据查重、许可证、形态和运行复杂度继续排序。

## 强约束

1. 工作量评估必须基于证据：README、Compose、Dockerfile、镜像、配置样例、启动日志、依赖服务、登录方式、文件流程。
2. 不把“能构建”当作“能上架”；必须覆盖安装、启动、登录、持久化、核心流程和商店材料。
3. 除非用户在当前任务明确说明“允许/需要修改业务代码”并点名范围，否则工作量评估只能按非侵入迁移估算。
4. 如果风险来自必须修改业务代码，返回边界门禁，不用工作量掩盖。
5. 工作量大时要明确建议：继续、先 POC、换项目或停止。
6. 输出必须能指导下一步执行，不写泛泛的“中等难度”。

## 分级

- `Small`: 官方镜像可用，单服务或简单 Compose，无内部登录或登录可配置，持久化清晰，1 天内可完成基本上架包。
- `Medium`: 有 2-4 个服务，需转换 Compose、补健康检查、初始化脚本、固定凭据或 inject，通常 1-3 天。
- `Large`: 多服务、多初始化阶段、登录复杂、文件流程复杂、外部依赖多、镜像需自建，通常 3-7 天，需要先做 POC。
- `Too Large`: 依赖特权能力、不可替代外部服务、缺失关键二进制/许可证、初始化不可自动化，或非侵入迁移不可控；不建议作为普通迁移项目。

## 评估维度

- 镜像路径：官方镜像、Dockerfile、自建镜像、源码构建。
- Compose 复杂度：服务数量、依赖顺序、健康检查、网络和端口。
- 持久化：数据目录、权限、缓存、配置生成。
- 初始化：数据库迁移、管理员创建、密钥生成、首启向导。
- 登录：无登录、固定账号、OIDC、inject、复杂二次验证。
- 文件能力：无文件流程、上传下载、文件选择器、文件关联。
- 外部要求：域名、邮件、对象存储、第三方 API、GPU、特权能力。
- 上架成本：图标、截图、描述、测试账号、复现步骤、版本来源。

## 输出格式

```text
Phase: Migration Workload Gate
Repository: <owner/repo>

Decision: Small / Medium / Large / Too Large
Recommendation: Proceed / POC first / Switch project / Stop

Evidence
- License gate:
- Duplicate checks:
- Boundary gate:
- Runtime complexity:
- Login and init:
- File flows:
- Store readiness:

Estimated Work
- Tasks:
- Main risks:
- Verification required:

Next
- Proceed to lazycat:ship-app / POC / Stop
```
