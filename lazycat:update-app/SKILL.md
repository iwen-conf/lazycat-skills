---
name: lazycat:update-app
description: 面向已上架 Lazycat 应用的版本更新、镜像升级、LPK 重构与开发者中心重新提交的 skill。当用户提到更新应用、升级版本、更新镜像、make update、修改 manifest 镜像、重新打包 LPK、提交新版本到开发者中心等请求时，必须使用此 skill。负责从镜像同步（copy-image）、manifest 修改、LPK 构建到开发者中心提审的完整更新闭环。
---

# Lazycat 应用更新与升级

你负责把“有一个新版本或新镜像”推进到“已提交至开发者中心并等待审核”的状态。重点是确保镜像已同步到懒猫私有仓库、manifest 已指向新镜像、LPK 已重新构建、且用户已在真实环境验证过更新后的版本。

## Overview

这个 skill 用于处理 Lazycat 应用的生命周期更新。核心流程包括：
1. **镜像更新**：使用 `lzc-cli appstore copy-image` 将 DockerHub 镜像同步到懒猫，获取内网镜像名。
2. **配置更新**：修改 `manifest.yml`，将镜像地址更新为同步后的地址。
3. **构建与验证**：执行 `make update` 和 `make install`，在真实懒猫环境验证新版本。
4. **准备提审**：执行 `make release-prep` 生成截图和报告。
5. **提交更新**：确保应用名称、简介、描述、关键词完整，并提交至开发者中心。

## Quick Contract

- **Trigger**: 更新应用、升级版本、镜像升级、copy-image、make update、提交 LPK、开发者中心更新、manifest 镜像修改。
- **Inputs**: 原始镜像名、版本号、仓库路径、开发者中心凭证、懒猫微服凭证。
- **Outputs**: 更新后的 manifest.yml、新的 LPK 文件、开发者中心提交记录、验证报告。
- **Quality Gate**: **必须先执行 `make install` 验证更新后的应用，严禁未验证直接提交。** 提交时所有元数据（名称、简介、描述、关键词）必须齐备。
- **Decision Tree**: 判断是仅代码更新、仅镜像更新还是两者都有，决定是否需要执行 `copy-image`。

## The Iron Law

1. **镜像同步优先**：如果有新镜像，必须先通过 `lzc-cli appstore copy-image` 同步，严禁在 `manifest.yml` 中直接使用外部 DockerHub 地址。
2. **真实环境验证**：更新后必须通过 `make install` 安装到懒猫微服进行体验验证，确保升级路径正常（如数据库迁移、配置兼容）。
3. **元数据完整性**：在开发者中心提交时，应用名称、简介、详细描述、关键词等必须填充满，且与本地 manifest 保持一致。
4. **Makefile 驱动**：必须使用 `make update` 自动化同步镜像和打包流程，使用 `make release-prep` 准备提审素材。

## Workflow

### 1. 镜像同步与 Manifest 更新
- 获取上游最新镜像名（如 `bitnami/gitea:latest`）。
- 执行 `lzc-cli appstore copy-image <image_name>`。
- 获取返回的懒猫私有镜像地址。
- 修改 `manifest.yml` 中的 `image` 字段。

### 2. 自动化构建
- 执行 `make update`：该 target 应包含镜像同步（若脚本支持）和 `lzc-cli project build`。
- 确认生成的 `.lpk` 文件版本号已递增。

### 3. 安装与体验验证
- 执行 `make install`：将新包安装到本地或真机环境。
- **强制要求：用户必须打开应用进行真实体验，确认功能正常。**
- **数据持久化校验**：必须验证从旧版本升级到新版本后，原有用户数据、配置、数据库内容是否完好。严禁为了更新镜像而强制清空挂载卷（Volumes）。

### 4. 准备提审素材
- 执行 `make release-prep`：运行测试、生成截图（尤其是新功能截图）。
- 整理 `changelog`，说明本次更新的具体改动。

### 5. 提交至开发者中心
- 访问 `developer.lazycat.cloud`。
- 上传新的 `.lpk`。
- **完整填写元数据：应用名称、简介、描述（含更新说明）、关键词、分类、截图。**
- 提交审核并记录提交证据。

## Outputs

```text
阶段: 应用更新 / 镜像升级
目标版本: <Version>

更新详情
- 镜像同步: <Old> -> <New (Lazycat)>
- Manifest 修改: 已完成 / 待处理
- LPK 构建: 已完成 / 待处理

验证状态
- 本地安装验证: <已验证 / 待验证>
- 核心功能体验: <正常 / 异常>

提交准备
- 应用名称/简介/描述/关键词: <已齐备 / 缺失>
- make release-prep: 已执行 / 待执行

下一步
1. ...
2. ...
```
