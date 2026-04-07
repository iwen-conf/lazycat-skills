# Lazycat Micro-service App Store Publishing and Review Standards

## I. Publishing Process

### 1. Developer Registration
1. Register a community account at [Lazycat Community](https://lazycat.cloud/login?redirect=https://developer.lazycat.cloud/).
2. Visit the [Developer Center](https://developer.lazycat.cloud/manage) and follow the prompts to submit your developer review application.
3. After submission, you can contact official support via customer groups to accelerate the review.

### 2. Pushing Images to the Official Registry
Before publishing, you **must** push all images referenced in the lpk to the official Lazycat registry. Failure to do so will prevent reviewers from installing the app, resulting in **review failure**.

```bash
lzc-cli appstore copy-image <publicly_accessible_image_name>
# On completion, the tool prints registry.lazycat.cloud/<community-username>/<image-name>:<hash-version>
```

**Example:**
```bash
lzc-cli appstore copy-image alpine:3.18
# Output: registry.lazycat.cloud/snyh1010/library/alpine:d3b83042301e01a4
```

**Official Registry (`registry.lazycat.cloud`) Constraints:**
1. To ensure stability, image tags are replaced with IMAGE_IDs. Every `copy-image` execution triggers a server-side `docker pull`.
2. Uploaded images **must be publicly accessible**. The `pull` operation happens on the server; local-only images cannot be processed by `copy-image`.
3. Uploaded images must be referenced by at least one store application; the registry performs periodic garbage collection.
4. `registry.lazycat.cloud` is for internal micro-service use only. Usage outside the micro-service is **rate-limited**.

After uploading, you **must manually update `lzc-manifest.yml`** to use the official `registry.lazycat.cloud/...` address.

### 3. Submission
Use `lzc-cli` (v1.2.54 or higher) to submit:
```bash
lzc-cli project build
lzc-cli appstore publish ./your-app.lpk
```

## II. App Store Review Guidelines (7 Red Lines)

Ensure all the following conditions are met before submission:

### 1. Metadata Completeness
- `package.yml` must be complete with `package`, `version`, `name`, `description`, `author`, and `license`. **For ported applications, the `author` field MUST exactly match the original project's author.**
- App Icon and screenshots must be provided in the Developer Center.
- Names, descriptions, and "Notes for Use" **must support multiple languages** via the `locales` configuration in `package.yml`.
- Language key standards should follow [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag).

### 2. Installability and Loadability
- The application must install and load normally.
- Apps that fail to install, fail to load after installation, or become unresponsive **will not pass review**.
- Thoroughly test the installation process and initial loading, **especially checking that all external dependencies are accessible during installation**.

### 3. Quality and Stability
- Avoid serious crashes or forced exits.

### 4. Performance Metrics
- Startup time and response time **must not exceed 5 minutes**.

### 5. Specialized Scenario Compatibility
- **Hardware-dependent Apps:** Must be tested in a real hardware environment; provide test notes including hardware models.
- **Special Scenario Apps (e.g., Browsers):** Must be fully tested in their respective scenarios.
- **Update Notifications:** In-app update prompts should not severely disrupt normal use. If in-app updates are not supported, remove the prompts.

### 6. Effective Use Case
- The submitted app must provide **genuinely useful scenarios** for users.
- **Development libraries and middleware are generally not permitted for publishing.**
- Utility apps should be associated with corresponding file types in Lazycat Drive.

### 7. Data Persistence
- Applications requiring persistent data must be tested to ensure data is correctly saved.
- **Ensure data is not lost after app restarts or upgrades.**
- Do not change instance definitions (which changes storage paths) during upgrades unless necessary. If changed, ensure data migration and recovery procedures are in place.

### 8. Passwordless Auto-Login
- Applications MUST support passwordless auto-login (免密登录) to provide a seamless user experience, ensuring users do not need to manually enter credentials upon first launch or subsequent visits.
- **Implementation Methods:**
  1. **OIDC Standard Flow:** Integrate with Lazycat MicroServer's identity authentication system to achieve automatic user identification based on ingress injection.
  2. **Inject Autofill:** Use `builtin://simple-inject-password` or custom inject scripts to automatically fill and learn login credentials.
- For implementation details, refer to the "Store Submission Guide" and [Advanced Inject Passwordless Login](https://developer.lazycat.cloud/advanced-inject-passwordless-login.html) in the Developer Documentation.

## III. Prohibited Application Types

- Software related to illegal activities (pornography, gambling, drugs), airdrops, cracked software, or anything violating Chinese laws **cannot be published**.
- For apps requiring usernames and passwords, if ordinary users cannot obtain credentials within the Lazycat Store, the app **cannot be published**.
