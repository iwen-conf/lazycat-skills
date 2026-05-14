# Lazycat AI App Submission Kit

Applicable to Lazycat AI Pod apps and AI Browser Plugins for organizing submission materials, preparing screenshots, and providing reviewer instructions.

Not applicable to standard web apps integrating AI. If your project is a standard web app with AI features, use the `BaseURL` configuration and follow `store-assets.md`. Do not use this AI Pod template.

## 1. Official References

- AI App Development Docs:
  `https://developer.lazycat.cloud/open/zh/guide/aipod/ai-application.html`
- AI Browser Plugin Testing Docs:
  `https://developer.lazycat.cloud/open/zh/guide/aipod/ai-browser-plugin-testing.html`
- App Submission/Review Guide:
  `https://developer.lazycat.cloud/open/zh/guide/publish/app-review.html`

Per official docs, AI Apps typically involve:
- `lzc-build.yml`
- `lzc-manifest.yml`
- `ai-pod-service/`
- `caddy-aipod/`
- Optional `extension.zip`

Special attention for AI Browser Plugin testing:
- `public_path`
- Isolation of login states between the plugin and normal pages.
- Independent authentication when cookies are ineffective.
- Troubleshooting network failures or `401` errors.

## 2. When to Use This Template

Prioritize this if:
- The project is a Lazycat AI Pod app.
- The project includes an AI Browser Plugin.
- The reviewer needs to understand why it isn't a standard web app.
- Metadata, screenshots, or instructions must explain AI entries, plugin entries, or the AI Pod directory structure.

## 3. Minimum Submission Package for AI Apps

Organize at least the following:
- App Name
- Version Number
- One-sentence Position
- Product Form: Standard App / AI App / AI Browser Plugin
- AI Entry Point location
- Key directory descriptions
- Reviewer reproduction path
- External dependencies and known limitations
- Actual screenshots
- Key verification conclusions

### Data Template (Copy-paste)

```text
App Name:
Version Number:
Product Form: <AI App / AI Browser Plugin / Standard App>

One-sentence Position:
<App Name> is a Lazycat <AI App / AI Browser Plugin> for <Target Users>, used for <Core Value>.

Why it's not a standard web app:
- ...

AI Entry Points:
- Main Entry:
- AI Browser Entry:
- Settings Entry:

Key Directory / Package Structure:
- lzc-build.yml: ...
- lzc-manifest.yml: ...
- ai-pod-service/: <Yes/No, Purpose>
- caddy-aipod/: <Yes/No, Purpose>
- extension.zip: <Yes/No, Purpose>

Reviewer Reproduction Path:
1. ...
2. ...
3. ...

Test Account / Initialization:
- ...

Dependencies & Limitations:
- External models or services:
- Network requirements:
- Known limitations:

Verification Conclusions:
- Installation:
- Startup:
- AI Entry:
- Core Workflow:
- Data Persistence:
```

## 4. Writing Reviewer Instructions

The goal is to let the reviewer understand and run the core path within 5 minutes.
Clarify:
- Where to access AI capabilities.
- Whether login, authorization, or configuration is needed before first use.
- Which step demonstrates the product value fastest.
- For AI Browser Plugins, which pages and content to test on.
- For AI Apps, the relationship between AI Browser, independent pages, and local services.

Do not just say "Open the app and see for yourself." Provide the shortest path to success.

## 5. AI Browser Plugin Screenshot Script

### Preparation
- Fix language, theme, and test data.
- Confirm the plugin is correctly installed and discoverable by the AI Browser.
- Confirm AI services are available.
- If the plugin requires separate login, verify the session is established.
- Remove debug markers, sensitive data, and invalid banners.

### Store Screenshot Priority
1. Entry point location within the AI Browser.
2. Main interface after opening the plugin.
3. Web context before triggering AI.
4. Plugin configuration or model settings page.
5. A representative AI operation in progress.
6. Results page after AI response.
7. A page showing the loop (history, saving, exporting, or re-invoking).

### Questions for Each Image
- Can the user tell at a glance this is an AI capability in the browser, not just a webpage?
- Does the image explain the core value?
- Is the image from the current submission version?

### Reviewer Evidence Screenshots
Prepare an extra set for internal verification (not necessarily for the store):
1. Plugin entry is visible.
2. Plugin successfully loads the target page.
3. Core AI request returns successfully.
4. Debugging screenshot for `401` or login isolation issues if applicable.

## 6. AI App Screenshot Script

For independent AI Apps, prioritize:
1. Homepage or main workflow entry after startup.
2. AI input or task configuration interface.
3. Key AI processing in progress.
4. Results page.
5. Settings or model/service configuration.
6. AI Browser entry if linked.

## 7. Self-Checklist

- AI App / AI Browser Plugin form clearly explained.
- AI entry point location clearly explained.
- Reviewer can reproduce core workflows within 5 minutes.
- Screenshots are from the actual running version.
- Directory descriptions (`ai-pod-service/`, etc.) are complete.
- Login or configuration methods for plugins are clear.
- External dependencies/network requirements are stated.
