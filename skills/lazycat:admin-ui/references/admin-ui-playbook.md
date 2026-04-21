# Lazycat Admin UI Playbook

Admin interfaces can start with established templates, but they must eventually become "your product's dashboard," not a "template demo site."

## 1. Objectives

Create an admin interface that:
- Projects a clear brand identity.
- Features a distinct information hierarchy.
- Reflects actual business workflows.
- Serves as high-quality App Store screenshots and reviewer demonstration material.

## 2. Template Usage Rules

### Prerequisites for Using Templates
- The template stack is compatible with React + Vite + Tailwind CSS + shadcn/ui, and can be cleanly wired into Zustand, TanStack Query, React Router, React Hook Form, Zod, and Framer Motion.
- Licenses are clear and permit current usage.
- The template is optimized for admin scenarios rather than marketing sites or BI dashboards.
- The structure is easy to prune and can be quickly replaced with real business pages.

### Elements That Must Be Modified
- Default brand name, logo, favicon, and copyright.
- Default menu groupings and example routes.
- Default charts, metrics, notifications, and announcements.
- Default empty state copy, example table columns, and detail fields.
- Default demo accounts, placeholder avatars, and mock data.

### Templates to Avoid
- Deeply coupled with other UI libraries, leading to high migration costs.
- Excessive animations or heavy decoration that reduces administrative efficiency.
- Templates where example content outweighs real business logic, resulting in high cleanup costs.
- Visual styles that clearly conflict with the current product direction.

## 3. Minimum Requirements for Admin Pages

Admin applications should include at least the following pages or equivalent capabilities:
- Login / Registration.
- Workbench or Home page.
- Core business list views.
- Detail pages or detail drawers.
- Create / Edit forms.
- Settings / Permissions / System parameters.

If a project lacks a workbench, screenshots and the landing experience often appear weak. Even for lightweight admin panels, provide a home page that illustrates core statuses and primary actions.

## 4. Workbench Design Essentials
- Retain only the most critical metrics and statuses at the top.
- Place primary task entries on the first screen; do not force users to navigate deep menus.
- Ensure card data relates to real business logic; avoid template-provided mock statistics.
- Prioritize "Recent Activity," "Pending Tasks," "Sync Status," and "Alerts" over flashy but non-functional charts.

## 5. List Page Design Essentials
- Keep only high-frequency filters visible; move low-frequency conditions to an expandable area.
- Maintain consistent placement for Search, Filter, Sort, Pagination, and Bulk Actions.
- Use semantic tags for status columns rather than plain text.
- Implement truncation, copy-to-clipboard, or tooltip strategies for long text, IDs, and timestamps.
- Order action columns by priority; separate dangerous actions from primary ones.

## 6. Detail Page and Form Design Essentials
- Provide a summary first, followed by details; avoid overwhelming the user with a full page of fields immediately.
- Group long forms into modules; use multi-step or collapsible sections when necessary.
- Ensure clear feedback for required fields, validation errors, successful saves, and failures.
- Provide explanations, default values, and rollback hints for high-risk configurations.

## 7. Visual System Recommendations
- Pick one of the five approved low-saturation palettes (① Slate, ② Warm Sand, ③ Dark Slate, ④ Sage, ⑤ Rose Mist) and drive everything through shadcn CSS variables (`--background`, `--primary`, `--border`, …). Do **not** hand-roll hex values or rely on shadcn's default palette.
- Maintain a consistent spacing rhythm (e.g., 8/16/24px) to avoid density fluctuations.
- Standardize borders and shadow intensities across cards, tables, drawers, and modals.
- Unify icon styles; avoid mixing too many different icon sets.
- Ensure the login page and the admin shell follow the same branding direction.

## 8. State and Exception Handling
- Use skeletons or local loading indicators; avoid full-screen flickering.
- Provide clear calls to action (CTAs) in empty states.
- Specify the cause and recovery steps for error prompts.
- Explain missing permissions on restricted pages instead of just showing an error.

## 9. Preparing Screenshots
A standard screenshot set for admin applications typically includes:
1. Workbench/Home page.
2. The primary business list page.
3. A detail page or key form.
4. A settings, permissions, or unique feature page.

Pre-screenshot checklist:
- No template default mock data.
- No debug markers or test buttons.
- Unified language and copy.
- Data fields consistent with real product semantics.
- Colors and branding successfully replaced.

## 10. Pre-Submission Self-Check
- Can a reviewer immediately identify the purpose of this admin panel?
- Might a reviewer mistake this for an unmodified template?
- Does the home page display truly important operational information?
- Do the lists, details, and forms demonstrate a complete business loop?
- Do the login page, home page, and settings page feel like part of the same product?
