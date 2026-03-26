# Lazycat Project Baseline

This baseline must be followed when a project enters the "Creation / Initialization / Tech Stack Convergence / Auth Implementation" phase.

## 0. Documentation First

The first step is to establish the `docs/` directory tree before starting scaffolding or business logic development. It should include at least:
- `docs/requirements/`
- `docs/api-design/`
- `docs/architecture/`
- `docs/release-prep/`

Refer to [docs-blueprint.md](./docs-blueprint.md) for detailed breakdown.

## 0.5 Unified Command Entry

Every project must provide executable command entry points, including:
- `build.sh` (root directory)
- `Makefile` (root directory)
- `make build`
- `make install`

Recommended additions:
- `make dev`
- `make test`
- `make clean`

Commands must be directly executable from the repository root, not just buried in chat logs or scattered documents.

## 1. Default Tech Stack

- **Backend:** Go
- **Frontend:** Vue
- **UI:** Element Plus

If there are existing team-specific standards, they can be refined within this baseline, but do not deviate from the core stack.

## 1.2 Admin UI Baseline

If the project includes an admin panel, operation console, or management dashboard:
- Keep the main stack as Vue + Element Plus.
- You may use mature templates, but remove default branding, menus, and example charts.
- Implement standard page patterns: Login/Register, Dashboard, Lists, Details/Forms, and Settings.
- Ensure admin pages are ready for screenshots and review (not just a collection of unstyled features).

Such projects should use `lazycat:admin-ui` to ensure quality.

## 1.5 Incentive-Priority Mode

If the goal is to qualify for Lazycat cash incentives, evaluate these during the creation phase:
- **Originality First:** Original applications typically have higher incentive potential than simple ports.
- **Lazycat Native Integration:** Explain the integration points with the Lazycat native system (not just "it hasn't been posted yet").
- **Avoid Excluded Types:** Steer clear of application types explicitly noted as ineligible for incentives.
- **User Credentials:** For apps requiring accounts, ensure ordinary users can obtain credentials easily.
- **File Association:** Prioritize file handlers for utility apps.
- **OIDC:** Prioritize Micro-service OIDC for applications suitable for unified accounts.

## 1.6 Native Integration for Original Apps

For original applications, answer these questions during creation:
- Why should a user install this in Lazycat instead of visiting a website?
- What is the primary native integration point: Micro-service OIDC, File Handler, Local File Workflow, Lazycat Computing Warehouse `AI Application`, or other system entries?
- Does the manifest reserve `application.oidc_redirect_path`, `application.file_handler`, or specific AI Pod routes?
- What is the first high-frequency scenario after installation? Can it provide a "ready-to-use" local experience?

If these questions cannot be answered, mark the project as "Weakly Integrated" rather than assuming it is natively compatible.

## 1.7 Computing Warehouse / AI Application Integration

For AI-native products, perform an extra evaluation step during creation:
- If it is a standard Web app, retain the `API BaseURL + Protocol + Model Discovery` configuration.
- Determine if it is better suited as a standard app using external APIs or as a Lazycat Computing Warehouse `AI Application`.
- Consider if it should be an AI Browser Extension rather than a standalone page.
- If following the `AI Application` route, plan for `ai-pod-service/`, `caddy-aipod`, and optional `extension.zip`.
- If not using the Computing Warehouse, justify why a standard app structure is more appropriate.

## 2. Default Authentication Requirements

All projects must support:
- Login
- Registration
- `access_token`
- `refresh_token`
- Silent Refresh
- Cleanup and Re-login after refresh failure

## 3. Recommended Auth Flow

### Login
1. User submits credentials.
2. Backend validates and issues an `access_token` and a `refresh_token`.
3. Frontend saves the authentication state and fetches user info.

*For incentive-priority apps, registration must be accessible to ensure reviewers and users can enter the application.*

### Startup Recovery
1. App attempts to restore user state on launch.
2. If `access_token` is expired but `refresh_token` is valid, perform a silent refresh.
3. If successful, proceed to the business page.
4. If failed, clear state and redirect to the login page.

### 401 Silent Refresh
1. A request returns 401 due to an expired `access_token`.
2. Frontend initiates a background refresh process.
3. Pending requests are queued until the refresh completes.
4. On success, replay the original request.
5. On failure, logout and clear session state.

## 4. Minimum Backend APIs

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `POST /auth/logout`
- `GET /auth/me`

## 5. Minimum Frontend Modules

- Login/Register pages
- User state store
- Route guards
- Request interceptors
- Refresh queue/mechanism
- Logout cleanup logic

## 5.5 AI Integration Baseline

For projects with AI scenarios, follow this minimum configuration interface:
- `API BaseURL`
- Protocol Type: `OpenAI Compatible`, `OpenAI Responses`, `Anthropic`
- "Get Models" button: Fetches model list based on current `BaseURL + Protocol`.
- Model Dropdown: Select from the fetched list (no hardcoded manual input).
- "Save Configuration" button: Persists the AI settings.

*Avoid scattering provider configs across multiple pages or hardcoding default model names.*

Standard Web apps only need this level of integration; do not introduce AI Pod structures unless necessary. Refer to [ai-settings-template.md](./ai-settings-template.md) for specifics.

For projects intended as `AI Applications`:
- Determine if core services belong in `ai-pod-service/`.
- Determine if exposure via `caddy-aipod` is required for the AI Browser or other frontends.
- Provide `extension.zip` if required as an AI Browser Extension.

## 6. Recommended Security Boundaries

- Short lifespan for `access_token`.
- Long lifespan for `refresh_token`.
- Implement token rotation on refresh.
- Immediate local session cleanup upon refresh failure.
- Explicitly define token secrets, expiration, and API URLs in environment variables.

## 6.5 Micro-service Auth and File Association

### OIDC
- Set `application.oidc_redirect_path` in the manifest.
- Pass system-generated OIDC env vars to the app.
- Map OIDC login to the application's internal session or user context.

### File Association
- Evaluate `application.file_handler` for utility apps.
- Declare `mime` types and `actions.open`.
- Implement `/open` or equivalent routes to parse incoming file parameters.

## 7. Transition to Release

Before entering `lazycat:ship-app`, ensure:
- Project is startable.
- Login and Registration work (or have clear business exemptions).
- Token refresh flow is functional.
- Admin UI is refined (if applicable).
- Incentive eligibility is clear.
- Native integration path is defined.
- AI configuration or AI Pod path is clear.
- Screenshots are available for the homepage, auth pages, and key features.
