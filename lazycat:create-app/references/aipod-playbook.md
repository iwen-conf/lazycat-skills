# Lazycat AI Pod Decision Manual

This manual provides guidance on determining whether a project should follow the Lazycat Computing Warehouse `AI Application` path, the AI Browser Extension path, or simply integrate a standard Model API.

**Core Rule:**
- For standard business Web applications, use the `BaseURL` configuration approach by default.
- Only follow the official AI Pod route when explicitly developing for the Lazycat Computing Warehouse, an `AI Application`, or an AI Browser Extension.

## 1. Official Resources

- **AI Application Development Docs:**  
  `https://developer.lazycat.cloud/open/en/guide/aipod/ai-application.html`
- **AI Browser Extension Testing Docs:**  
  `https://developer.lazycat.cloud/open/en/guide/aipod/ai-browser-plugin-testing.html`
- **Lazycat AI Pod Product Page:**  
  `https://lazycat.cloud/ai-pod`

## 2. Three Common Integration Routes

### Standard Application + External Model API
**Best for:**
- Products that are essentially standard Web applications.
- AI is a feature, not the primary hosting environment.
- Users need to manually configure third-party model services.

This is the **default route**. It does not require `ai-pod-service/`, `caddy-aipod`, or `extension.zip`.

**Minimum requirements during the creation phase:**
- AI Settings Page
- `API BaseURL`
- Protocol Selection: `OpenAI Compatible`, `OpenAI Responses`, or `Anthropic`
- "Get Models" button
- Model dropdown
- "Save Configuration" button

### Lazycat Computing Warehouse `AI Application`
**Best for:**
- AI-native products.
- Deep integration with Lazycat AI Pod or AI Browser features.
- Needs for local AI services or independent AI workflow entry points.

Only enter this route if the user explicitly intends to build for the Lazycat Computing Warehouse or an `AI Application`.

**Key evaluations during creation and release:**
- Requirement for `ai-pod-service/`
- Requirement for `caddy-aipod`
- Inclusion of a browser extension in the package
- Need for specialized AI Pod routing and middleware

### AI Browser Extension
**Best for:**
- Capabilities that naturally occur within browser pages.
- Features like text selection, summarization, rewriting, page understanding, sidebars, or assistants.
- Scenarios where an extension form factor is more effective than a standalone page.

Only enter this route if the user explicitly intends to build an AI Browser Extension or browser-side capabilities.

**Key evaluations during creation and release:**
- Requirement for `extension.zip`
- Combination with `ai-pod-service/`
- Exposing services to the AI Browser or frontend via `caddy-aipod`

## 3. Official Package Structure Clues

According to official `AI Application` documentation, typical structures include:
- `lzc-build.yml`
- `lzc-manifest.yml`
- `ai-pod-service/`
- `caddy-aipod/`
- Optional `extension.zip`

*Do not hardcode additional detail fields in the skill; always refer to the current official documentation for field names and directory constraints.*

## 4. Decision Questions for the Creation Phase

At a minimum, answer the following:
- Is this a standard business Web application? If so, why is `BaseURL` configuration sufficient?
- Why should a user install this in Lazycat rather than visiting a web-based AI service?
- Does it function more like a standard app, an `AI Application`, or an AI Browser Extension?
- What is the primary, high-frequency entry point: a standalone page, an AI Browser entry, or a local service?
- If not using the AI Pod route, why is a standard application structure adequate?

## 5. Minimum Validation for Release

If the project follows the `AI Application` or AI Browser Extension route, validate the following:
- All required directories and package structures are complete.
- Services can be started successfully.
- AI Browser or extension entry points are visible.
- Critical AI workflows can be executed from end to end.
- Store metadata aligns with the actual AI entry points.

## 6. Documentation Requirements

Ensure the decision and integration details are recorded in:
- `docs/requirements/lazycat-native-fit.md`
- `docs/api-design/ai-integration.md`
- `docs/architecture/lazycat-integration.md`
- `docs/architecture/aipod-integration.md`
