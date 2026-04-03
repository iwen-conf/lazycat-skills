# Lazycat App Shipping Checklist

Check items by stage. Do not skip steps.

## 1. Project Initiation & Prerequisites

- [ ] **Strict Compliance Check**: Ensure the application does not contain or relate to pornography (黄), gambling (赌), drugs (毒), airdrops (空投), cracked software (破解软件), or any content violating Chinese laws. Applications violating this rule are strictly prohibited from listing and must be rejected immediately.
- [ ] Clarify if it's a first-time launch, an update, or a re-submission after rejection.
- [ ] Define the target version number, release date, and minimum viable release scope.
- [ ] Verify developer certification, Developer Center permissions, and target app permissions.
- [ ] Clarify credential scopes: `lazycat_account` / `lazycat_password` for entering Lazycat OS; `lazycat_developer_center_account` / `lazycat_developer_center_password` for the Developer Center. In-app login should use specific app-level variables.
- [ ] Identify the repository, build commands, release entry point, and version source.
- [ ] Established `docs/requirements`, `docs/api-design`, `docs/architecture`, and `docs/release-prep`.
- [ ] Prepared `build.sh`, `Makefile`, `make build`, and `make install`.
- [ ] If the project includes an admin interface, completed `lazycat:admin-ui` quality convergence or noted pending items.
- [ ] Confirmed whether `.lpk`, `lzc-cli project publish`, or both are required.
- [ ] If the project is an AI-native product, clarified the path for Standard App / `AI App` / AI Browser Plugin.
- [ ] If the target is a cash incentive, clarified whether it is original or ported, and determined if it falls under non-rewarded/not-recommended types (e.g., pure web games, pure book pages, pure tutorials, web offline apps, game mods, pure database apps, VPNs, pure frontend apps, saturated categories, image hosting, navigation, bookmarks, notes, online video viewers, checklists, short link generators, burn-after-reading, YouTube fetchers, bookkeeping).

## 2. Local App Information

- [ ] App name matches the final store display name.
- [ ] App summary matches actual functionality.
- [ ] Icon is complete, clear, and contains no placeholders.
- [ ] Category settings are correct.
- [ ] Multilingual copy covers the default language with a clear fallback strategy.
- [ ] Version description or changelog is prepared.
- [ ] If an admin interface exists, visual direction is consistent across login/registration, dashboard, lists, forms, and settings pages.
- [ ] If the app requires a username/password, regular users can self-register, use unified login, or obtain publicly available credentials.

## 3. Packaging & Uploading

- [ ] Generated release artifacts using the actual build pipeline.
- [ ] Recorded package path, filename, size, and version number.
- [ ] Recorded release commands, CLI output, or upload page status.
- [ ] Confirmed current field and format requirements in the Developer Center before uploading.
- [ ] If using `.lpk`, confirmed the version in the package matches the store profile.
- [ ] If porting open-source software, prepared the upstream author's URL.
- [ ] For `AI Apps`, checked that structures like `ai-pod-service/`, `caddy-aipod/`, and `extension.zip` are complete.

## 4. Store Metadata & Screenshots

- [ ] No conflicts between manual fields and auto-populated fields.
- [ ] App summary contains no exaggerated claims or unverifiable statements.
- [ ] Screenshots are taken from the actual running version.
- [ ] Screenshots cover the homepage, core workflows, and key differentiators.
- [ ] If the admin interface is a key selling point, screenshots cover the dashboard, core lists, details/forms, and settings/permissions pages.
- [ ] For `AI Apps` or AI Browser Plugins, screenshots cover AI entry points and core AI workflows.
- [ ] Screenshots contain no debugging info, sensitive data, placeholders, or error messages.
- [ ] If a dashboard template was used, default logos, menus, sample charts, and placeholder text have been cleaned up.
- [ ] Recorded actual count, dimensions, or aspect ratio requirements for screenshots.
- [ ] If aiming for incentives, metadata explains the app's real-world scenarios, originality, or porting source.

## 5. Pre-submission Testing

- [ ] Installed the build artifact into a real Lazycat OS instance.
- [ ] Launched the target app from the installed app entry in Lazycat OS.
- [ ] New installation starts correctly.
- [ ] Upgrade process completes successfully.
- [ ] Core functionalities work as expected.
- [ ] Failure paths (unlogged, network issues, empty data) behave predictably.
- [ ] Dependencies like permissions, network, and directory mounts are normal.
- [ ] Store metadata matches the actual running application.
- [ ] If an admin interface exists, verified the dashboard, filtered lists, details/forms, empty states, and restricted permission scenarios.
- [ ] Verified login/registration/credential acquisition paths.
- [ ] If integrated with OIDC or File Handler, successfully tested related scenarios.
- [ ] For `AI Apps` or AI Browser Plugins, verified AI services, browser entries, or plugin entries are functional.

## 6. Submission

- [ ] Organized version number, changelog, app summary, screenshots, and artifact info.
- [ ] Organized reviewer reproduction steps, test accounts, or initialization steps.
- [ ] Recorded submission time, status, and reviewer feedback.
- [ ] Submission evidence is traceable to specific versions and build sources.
- [ ] If aiming for incentives, organized original/porting documentation, upstream URLs, and OIDC/File Handler integration notes.
- [ ] For `AI Apps`, organized AI Pod route descriptions and key directory explanations.
- [ ] For `AI Apps` or AI Browser Plugins, prepared reviewer reproduction materials and screenshot scripts per `aipod-review-kit.md`.

## 7. Post-release Verification

- [ ] App is searchable or visible in the store.
- [ ] Name, icon, summary, screenshots, and version number are correct.
- [ ] Version installed from the store runs correctly.
- [ ] If the store is not updated, compared release time, version number, and upload records.
- [ ] If aiming for incentives, preserved the evidence chain required for subsequent reward claims.

Do not consider "local development environment," "static screenshot previews," or "web form submission" as "actual testing completed." Pre-submission testing must use the version installed within Lazycat OS.
