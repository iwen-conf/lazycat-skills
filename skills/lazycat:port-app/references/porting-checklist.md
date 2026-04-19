# Lazycat Porting Checklist

## 1. Before Selection
- [ ] GitHub search completed.
- [ ] License allows porting and commercial use (non-commercial projects prohibited).
- [ ] Upstream URL recorded.
- [ ] App Store duplication check completed.
- [ ] If highly duplicative, terminated incentive path or clarified differentiation.

## 2. Project Execution
- [ ] Created `docs/requirements`.
- [ ] Created `docs/api-design`.
- [ ] Created `docs/architecture`.
- [ ] Created `docs/release-prep`.
- [ ] Created `build.sh`.
- [ ] Created `Makefile`.
- [ ] `make build` implemented.
- [ ] `make install` implemented.

## 3. Lazycat Adaptation
- [ ] Prepared `lzc-build.yml`.
- [ ] Prepared `lzc-manifest.yml`.
- [ ] If the app uses bridged images, the final pullable image refs have already been written back to `lzc-manifest.yml` before `make install`.
- [ ] Evaluated OIDC requirement.
- [ ] Evaluated `file_handler` requirement.
- [ ] Clarified credential acquisition path for apps requiring login.
- [ ] Distinguished credential scopes: `lazycat_account` / `lazycat_password` for Lazycat OS and App Store; `lazycat_developer_center_account` / `lazycat_developer_center_password` for the Developer Center. In-app login uses app-level variables.
- [ ] Application MUST support passwordless auto-login (免密登录) via OIDC or Inject (`builtin://simple-inject-password`) to provide a seamless user experience.
- [ ] For AI projects, determined if AI Pod or AI Browser Plugin is more suitable.

## 4. Incentive Path
- [ ] Does not belong to non-rewarded types.
- [ ] Original/Porting path clarified.
- [ ] Prepared upstream attribution (for ports) and ensured the `author` field strictly matches the original project's author.
- [ ] Planned OIDC or `file_handler` if applicable.
- [ ] Planned to verify core capabilities by installing and opening the app within Lazycat OS.
- [ ] Planned Standard App / AI Pod / AI Browser Plugin route for AI-native projects.
