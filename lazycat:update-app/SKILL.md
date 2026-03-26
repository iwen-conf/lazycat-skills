---
name: lazycat:update-app
description: 面向已上架 Lazycat 应用的版本更新、镜像升级、LPK 重构与开发者中心重新提交的 skill。当用户提到更新应用、升级版本、更新镜像、make update、修改 manifest 镜像、重新打包 LPK、提交新版本到开发者中心等请求时，必须使用此 skill。负责从镜像同步（copy-image）、manifest 修改、LPK 构建到开发者中心提审的完整更新闭环。
---

# Lazycat App Update and Upgrade

You are responsible for progressing a "new version or image" to the state of "submitted to the Developer Center and awaiting review." Key tasks include ensuring images are synced to the Lazycat private registry, updating the manifest to point to the new image, rebuilding the LPK, and verifying the update in a real environment.

## Overview

This skill handles the lifecycle updates of Lazycat apps. The core workflow includes:
1. **Image Update**: Use `lzc-cli appstore copy-image` to sync DockerHub images to Lazycat and obtain the internal image name.
1. **Configuration Update**: Modify `package.yml` to update the version and `manifest.yml` to update the image address to the synced version.
3. **Build and Verify**: Run `make update` and `make install` to verify the new version in a real Lazycat environment.
4. **Prepare Submission**: Run `make release-prep` to generate screenshots and reports.
5. **Submit Update**: Ensure metadata in `package.yml` (name, tagline, description, keywords) is complete and submit to the Developer Center.

## Quick Contract

- **Trigger**: Updating an app, upgrading a version, image upgrade, copy-image, make update, submitting LPK, Developer Center updates, manifest image modification.
- **Inputs**: Original image name, version number, repository path, Developer Center credentials, Lazycat MicroServer credentials.
- **Outputs**: Updated manifest.yml, new LPK file, Developer Center submission record, verification report.
- **Quality Gate**: **You must run `make install` to verify the updated app before submission; unverified submissions are strictly prohibited.** All metadata must be complete upon submission.
- **Decision Tree**: Determine if it's a code update, image update, or both, and decide if `copy-image` is required.

## The Iron Law

1. **Image Sync First**: New images must be synced via `lzc-cli appstore copy-image`; using external DockerHub addresses directly in `manifest.yml` is prohibited.
2. **Real Environment Verification**: After an update, you must install it via `make install` on Lazycat MicroServer for functional verification and upgrade path checks (e.g., database migrations).
3. **Metadata Integrity**: When submitting to the Developer Center, the app name, tagline, detailed description, and keywords must be complete and consistent with the local manifest.
4. **Makefile Driven**: Use `make update` to automate image syncing and packaging, and `make release-prep` to prepare submission assets.

## Workflow

### 1. Image Sync and Manifest Update
- Obtain the latest upstream image name (e.g., `bitnami/gitea:latest`).
- Execute `lzc-cli appstore copy-image <image_name>`.
- Retrieve the returned Lazycat private image address.
- Update the `image` field in `manifest.yml`.

### 2. Automated Build
- Run `make update`: This target should include image syncing (if supported by script) and `lzc-cli project build`.
- Confirm the version number in the generated `.lpk` file has incremented.

### 3. Installation and Verification
- Run `make install`: Install the new package to a local or physical environment.
- **Mandatory: Users must open the app for a real experience to confirm functionality.**
- **Persistence Check**: Verify that user data, configurations, and databases remain intact after upgrading from an older version. Forcing volume wipes for image updates is prohibited.

### 4. Prepare Submission Assets
- Run `make release-prep`: Run tests and generate screenshots (especially for new features).
- Organize the `changelog`, detailing specific changes in this update.

### 5. Submit to Developer Center
- Visit `developer.lazycat.cloud`.
- Upload the new `.lpk`.
- **Complete metadata: name, tagline, description (including update notes), keywords, category, and screenshots.**
- Submit for review and record proof of submission.

## Outputs

```text
Phase: App Update / Image Upgrade
Target Version: <Version>

Update Details
- Image Sync: <Old> -> <New (Lazycat)>
- Manifest Update: Done / Pending
- LPK Build: Done / Pending

Verification Status
- Local Installation: <Verified / Pending>
- Core Experience: <Normal / Abnormal>

Submission Prep
- Metadata (Name/Tagline/Desc/Keywords): <Complete / Incomplete>
- make release-prep: Executed / Pending

Next Steps
1. ...
2. ...
```
