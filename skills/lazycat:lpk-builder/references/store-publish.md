# Lazycat Micro-service App Store Submission and Publishing Specification

## I. Submission Process

### 1. Developer Registration
1. Register an account on the [Lazycat Community](https://lazycat.cloud/login?redirect=https://developer.lazycat.cloud/).
2. Visit the [Developer Center](https://developer.lazycat.cloud/manage) and follow the prompts to submit a developer review application.
3. After submission, it is recommended to contact support or use the official [Contact Us](https://lazycat.cloud/about?navtype=AfterSalesService) channels to expedite review.

### 2. Push Images to the Official Registry
Before submitting to the store, you **must** push all images referenced in the lpk to the official Lazycat registry. Otherwise, reviewers will be unable to install the app, leading to **review failure**.

```bash
lzc-cli appstore copy-image <Publicly_accessible_image_name>
# After completion, the command prints registry.lazycat.cloud/<community-username>/<image-name>:<hash-version>
```

Example:
```bash
lzc-cli appstore copy-image alpine:3.18
# Output: registry.lazycat.cloud/snyh1010/library/alpine:d3b83042301e01a4
```

**Official Image Registry `registry.lazycat.cloud` Usage Restrictions:**
1. For stability, the generated image tag is replaced by the IMAGE_ID. Each `copy-image` execution triggers a mandatory `docker pull` on the server side.
2. Uploaded images **must be publicly accessible**. The `pull` operation occurs on the server; local-only images cannot be processed by `copy-image`.
3. Uploaded images must be referenced by at least one store application; the registry performs periodic garbage collection.
4. `registry.lazycat.cloud` is intended for use within Micro-services only. Using it outside of the Micro-service environment will result in **rate limiting**.

#### Recommended Path for Self-Built Images
If you have modified the upstream image locally, or the upstream only ships a `Dockerfile`, use this bridge workflow:

```bash
docker buildx build --platform linux/amd64 -t your-hub-user/app-name:v1.0 --load .
docker push your-hub-user/app-name:v1.0
lzc-cli appstore copy-image your-hub-user/app-name:v1.0
```

Rules for this path:
1. `copy-image` must receive a registry address that the server can pull. A local-only image tag is invalid.
2. Prefer Docker Hub for the public handoff unless the user explicitly provides another public registry.
3. After `copy-image` succeeds, backwrite `services.<name>.image` in the source `lzc-manifest.yml` with the returned `registry.lazycat.cloud/...` value. If the repo packages from manifest templates or staged manifests, backwrite those source files too.
4. The `lpk` should only retain manifests, runtime scripts, icons, and static assets. Do **not** try to embed the application image itself into the package.
5. If the user wants the full release closure, prefer a dedicated release target such as `release-build` / `release-install` that runs build image -> push public image -> `copy-image` -> backwrite source manifest -> build `.lpk` -> install `.lpk`.

After uploading, you **must manually update the image address in the source `lzc-manifest.yml`** to the official `registry.lazycat.cloud/...` address returned before packaging or publishing.

### 3. Submit for Review
Use `lzc-cli` (v1.2.54 or above) to submit:
```bash
lzc-cli project build
lzc-cli appstore publish ./your-app.lpk
```

## II. App Store Review Guidelines (7 Red Line Rules)

Before submitting, ensure all the following conditions are met:

### 1. Completeness of App Information
- `package.yml` must be complete with `package`, `version`, `name`, `description`, `author`, and `license`.
- App Icon and screenshots must be provided in the Developer Center.
- The name, description, and usage instructions **must support multiple languages** via the `locales` configuration in `package.yml`.
- Language key specifications follow the [BCP 47 standard](https://en.wikipedia.org/wiki/IETF_language_tag).

### 2. Installability and Loadability
- The app must install and load normally.
- Apps that fail to install, fail to load after installation, or become unresponsive will **not pass review**.
- Thoroughly test the installation process and initial loading before submission. **Specifically, verify that dependencies required for installation are accessible.**

### 3. Stability and Quality
- Avoid serious crashes or unexpected shutdowns.

### 4. Performance Metrics
- App startup speed and response time **must not exceed 5 minutes**.

### 5. Special Scenario Adaptation
- **Hardware-paired apps**: Must be tested in a real hardware environment; include test notes with hardware model information.
- **Special scenario apps** (e.g., browsers): Must be fully tested in their respective scenarios.
- **Update Prompts**: In-app update prompts should not severely disrupt normal use. If the app cannot complete in-app updates, consider removing the prompt.

### 6. Validity of Use Case
- The submitted app must provide a **genuine and valid use case** for the user.
- **Development libraries and middleware are generally not allowed**.
- Utility apps should be associated with relevant file types in the Lazycat Drive.

### 7. Data Persistence
- For apps requiring persistent data, you must verify that data is correctly preserved.
- **Ensure no data loss after app restarts or upgrades**.
- When upgrading an existing store app, avoid changing the instance definition unless necessary (as this changes the storage path). If changes are required, handle data migration and recovery properly.

### 8. Passwordless Auto-Login
- Applications MUST support passwordless auto-login (免密登录) to provide a seamless user experience, ensuring users do not need to manually enter credentials upon first launch or subsequent visits.
- **Implementation Methods:**
  1. **OIDC Standard Flow:** Integrate with Lazycat MicroServer's identity authentication system to achieve automatic user identification based on ingress injection.
  2. **Inject Autofill:** Use `builtin://simple-inject-password` or custom inject scripts to automatically fill and learn login credentials.
- For implementation details, refer to the "Store Submission Guide" and [Advanced Inject Passwordless Login](https://developer.lazycat.cloud/advanced-inject-passwordless-login.html) in the Developer Documentation.

## III. Prohibited App Types

- Apps involving illegal content (e.g., pornography, gambling, drugs), airdrops, cracked software, or anything violating Chinese laws **are prohibited**.
- For apps requiring a username and password, if ordinary users cannot obtain credentials through the Lazycat Store, the app **cannot be listed**.
