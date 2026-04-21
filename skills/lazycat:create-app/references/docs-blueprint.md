# Lazycat Documentation Directory Blueprint

The first step in project creation is not writing code, but establishing the `docs/` directory tree. It should be divided into sections like "Requirements Analysis" and "API Design," with multiple `.md` files in each.

## 1. Recommended Directory Structure

```text
docs/
├── requirements/
│   ├── product-overview.md
│   ├── user-stories.md
│   ├── scope-and-milestones.md
│   └── lazycat-native-fit.md
├── api-design/
│   ├── overview.md
│   ├── auth.md
│   ├── domain-modules.md
│   └── ai-integration.md
├── architecture/
│   ├── system-overview.md
│   ├── frontend.md
│   ├── backend.md
│   ├── lazycat-integration.md
│   └── aipod-integration.md
└── release-prep/
    ├── store-assets.md
    ├── test-plan.md
    └── submission-checklist.md
```

## 2. Minimum Content for Each Directory

### `docs/requirements/`
- `product-overview.md`: Project goals, target users, core value, and boundaries.
- `user-stories.md`: Key user stories, primary flows, and edge cases.
- `scope-and-milestones.md`: MVP scope, phase goals, and deferred items.
- `lazycat-native-fit.md`: Why the app fits Lazycat, planned native integrations, and the primary entry point.

### `docs/api-design/`
- `overview.md`: API style, authentication method, response structure, and error conventions.
- `auth.md`: Endpoints for login, registration, refresh, logout, and user info.
- `domain-modules.md`: Business logic APIs, separated by module.
- `ai-integration.md`: (Optional) AI config APIs, model discovery, protocol constraints, and error handling. Standard web apps can use `references/ai-settings-template.md`.

### `docs/architecture/`
- `system-overview.md`: High-level architecture, module relationships, and deployment boundaries.
- `frontend.md`: React + Vite + TypeScript + Tailwind CSS + shadcn/ui, plus Zustand for client state, TanStack Query for server state, React Router for navigation, React Hook Form + Zod for forms, and Framer Motion for animation; document the chosen low-saturation palette and the auth flow.
- `backend.md`: Go service structure, authentication chain, database, and external dependencies.
- `lazycat-integration.md`: Manifest constraints, OIDC / `file_handler` mapping, and internal route resolution.
- `aipod-integration.md`: (Optional) Planning for Computing Warehouse `AI Applications`, AI Browser Extensions, `ai-pod-service`, `caddy-aipod`, and extension packages.

### `docs/release-prep/`
- `store-assets.md`: App description, screenshots, icons, categories, and key features.
- `test-plan.md`: Testing scope, critical paths, and auth/refresh validation.
- `submission-checklist.md`: Pre-submission checks, evidence items, and reviewer reproduction steps.

## 3. Documentation Requirements
- Do not dump everything into a single `README.md`.
- Each directory must contain at least 2-3 `.md` files.
- API documentation **must** have a separate `auth.md`.
- Original applications **must** have a dedicated `lazycat-native-fit.md` (vague verbal plans are not sufficient).
- AI integrations must explicitly document `BaseURL`, protocol, model fetching, and persistence.
- AI-native products must specify the path (Computing Warehouse vs. Browser Extension).
- Requirements must clearly define the MVP and what is "out of scope."
- Release docs must cover login, registration, token refresh, and the submission path.

## 4. Relationship with Code
- **Requirements** define the scope.
- **API Design** defines the interfaces.
- **Architecture** defines layers and responsibilities.
- **Release Prep** defines the path to the store.

Implementation must follow these documents; do not use code to "guess" requirements.
