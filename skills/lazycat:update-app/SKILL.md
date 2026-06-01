---
name: lazycat:update-app
description: Update an already-published Lazycat app to a new version. Use for image sync (copy-image), manifest image/version bump, LPK rebuild, and Developer Center re-submission. Covers the full update loop from image to re-review; for first-time listing use lazycat:ship-app. 更新已上架应用、升级版本、更新镜像、make update、重打LPK、开发者中心重新提交。
---

# Lazycat App Update and Upgrade

You are responsible for progressing a "new version or image" to the state of "submitted to the Developer Center and awaiting review." Key tasks include ensuring images are synced to the Lazycat private registry, updating the manifest to point to the new image, rebuilding the LPK, and verifying the update in a real environment.

## Overview

This skill handles the lifecycle updates of Lazycat apps. The core workflow includes:
1. **Image Update**: Build the updated image, push it to a public registry, then use `lzc-cli appstore copy-image` to sync it to Lazycat and obtain the internal image name.
2. **Configuration Update**: Modify `package.yml` to update the version (must strictly be `x.x.x` format) and backwrite `lzc-manifest.yml` to the copied Lazycat registry address. If packaging uses manifest templates, update those sources too.
3. **Build and Verify**: Build the `.lpk` from the backwritten manifest, then run `make install` or the repo's dedicated release-install target to verify the new version in a real Lazycat environment.
4. **Prepare Submission**: Run `make release-prep` to generate screenshots and reports.
5. **Submit Update**: Ensure metadata in `package.yml` (name, tagline, description, keywords) is complete and submit to the Developer Center.

## Quick Contract

- **Trigger**: Updating an app, upgrading a version, image upgrade, copy-image, make update, submitting LPK, Developer Center updates, manifest image modification.
- **Inputs**: Original image name, version number, repository path, Developer Center credentials, Lazycat MicroServer credentials.
- **Outputs**: Updated `lzc-manifest.yml`, new LPK file, Developer Center submission record, verification report.
- **Quality Gate**: **You must run `make install` to verify the updated app before submission; unverified submissions are strictly prohibited.** All metadata must be complete upon submission.
- **Decision Tree**: Determine if it's a code update, image update, or both, and decide if `copy-image` is required.

## The Iron Law

1. **Image Sync First**: New images must be synced via `lzc-cli appstore copy-image`; using external DockerHub addresses directly in `lzc-manifest.yml` is prohibited.
2. **Real Environment Verification**: After an update, you must install it via `make install` on Lazycat MicroServer for functional verification and upgrade path checks (e.g., database migrations).
3. **Metadata Integrity**: When submitting to the Developer Center, the app name, tagline, detailed description, and keywords must be complete and consistent with the local manifest.
4. **Makefile Driven**: Prefer the repo's dedicated release target when it exists. If the repo is image-only and the user asks for the full closure, the standard sequence is `build image -> push public image -> copy-image -> backwrite source manifest -> build lpk -> install lpk`.
5. **Ported App Boundary**: For ported open-source/self-hosted apps, a "code update" means adopting an upstream release, rebuilding an image from already-approved upstream code, or applying a user-approved product-development patch with named business scope. Do not silently modify upstream frontend/backend/auth/API/schema/test code to make an update, upgrade path, healthcheck, login, or review pass.

## Workflow

### 1. Image Build, Sync, and Manifest Update
- Obtain the latest upstream image name or build output.
- If the image is locally modified, push it to a public registry first.
- Execute `lzc-cli appstore copy-image <public_image_name>`.
- Retrieve the returned Lazycat private image address.
- Backwrite the `image` field in the source `lzc-manifest.yml`. If the repo packages from manifest templates, backwrite those sources too.
- If the update is for a ported app, classify changed files before editing: packaging/runtime files are allowed; upstream business source files require explicit product-development authorization. If the only way to verify the update is to patch business code, report `Blocked by business-code change requirement`.

### 2. Automated Build
- Prefer an explicit release target such as `make release-build` / `make release-install` when the repo already supports the full chain.
- If the repo lacks that target but the user requests end-to-end closure, add or fix one instead of overloading `make install`.
- Confirm the version number in the generated `.lpk` file has incremented.

### 3. Installation and Verification
- Run `make install`: Install the new package to a local or physical environment.
- **Mandatory: Users must open the app for a real experience to confirm functionality.**
- **Persistence Check**: Verify that user data, configurations, and databases remain intact after upgrading from an older version. Forcing volume wipes for image updates is prohibited.
- Verification failures in ported apps must first be handled through manifest, deploy params, env/config rendering, wrapper entrypoints, setup/seed scripts, routes/upstreams, image version selection, or docs. Do not patch upstream business logic as part of the update loop without explicit authorization.

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
