---
name: lazycat:admin-ui
description: 面向 Lazycat 应用后台管理、控制台和管理工作台 UI 收敛的 skill。只要用户提到后台管理、Admin、管理台、控制台、仪表盘、运营后台、Vue 后台、Element Plus 管理界面、后台模板、管理端改版、后台页面美化、管理列表、数据工作台、提审前页面升级、高质量后台等请求，就必须使用此 skill。负责把后台管理界面收敛成高质量、可截图、可提审的 Vue + Element Plus 管理体验，并允许以成熟模板为起点但必须完成品牌化和业务化改造。
---

# Lazycat Admin UI Quality Baseline

You are responsible for progressing admin interfaces, operational consoles, and B-side workspaces within Lazycat projects from "functional" to "high-quality, screenshot-ready, and submission-ready." The focus is not just a visual skinning, but ensuring information architecture, page hierarchy, tables, forms, feedback, and branding meet release standards.

## Overview

This skill is used for the design convergence and implementation constraints of admin-type pages. The default standards are:

- Maintain a Vue + Element Plus stack; do not switch to other UI frameworks without authorization.
- Mature admin templates are allowed as a starting point, but delivering raw templates is prohibited.
- Establish a real brand color, navigation structure, page hierarchy, and business module partitioning.
- Ensure workspaces, list pages, detail pages, form pages, settings pages, and login pages form a unified visual system.
- Cover loading, empty states, error states, permission restrictions, and long data scenarios.
- Final pages must be able to support store screenshots, submission materials, and real demonstrations.

If a project does not have an admin interface—only a few settings pages or standard user pages—do not force this skill. Enter this quality pipeline only when there is a clear management console, dashboard, or back-office surface.

## Quick Contract

- **Trigger**: User mentions admin interface, "Admin," management console, dashboard, back-office template, Vue Admin, Element Plus management UI, operational backend, workspace, dashboard, admin UI upgrade, high-quality UI.
- **Inputs**: Project directory, current admin state, page list, template presence, target screenshot scenarios, brand direction, core business flows.
- **Outputs**: Admin information architecture recommendations, template strategy, core page patterns, visual quality requirements, screenshot checklist, and UI quality conclusions for `lazycat:ship-app`.
- **Quality Gate**: Admin interfaces must undergo branding and business-specific customization. If a template is used, default branding, copy, data, and boilerplate charts must be removed. Core pages (workspace, list, detail, form, settings, login) must form a unified, high-quality, and screenshot-ready experience.
- **Decision Tree**: Determine if an admin interface actually exists, then decide whether to build from scratch, refactor an old one, or converge based on a template. Finally, decide whether to prioritize architecture, layout, or core page patterns.

## When to Use

**Primary Triggers**

- User requests a high-quality upgrade for an admin section.
- User wants to use a template to build an admin interface but doesn't want it to look like a template.
- User specifies Vue + Element Plus and needs a console, workspace, or dashboard.
- User is preparing for submission or screenshots, but the admin pages are immature.
- User needs to transform a "functional but ugly" admin UI into a release-level interface.

**Typical Scenarios**

- Building a new admin interface from scratch, requiring login, workspace, list, form, and settings pages.
- A template is already in use, but still retains default logos, layouts, and sample charts.
- Business logic exists, but the UI hierarchy is messy, tables are hard to use, forms are too long, and screenshots look poor.
- Preparing to list on the store and needing admin screenshots that actually showcase product value.
- Enhancing branding and professionalism for an Element Plus admin UI without changing the tech stack.

**Boundary Notes**

- If the task is primarily about unifying the stack, adding auth, or creating documentation, use `lazycat:create-app`.
- If the task is already at the stage of store assets, submission, or release, return to `lazycat:ship-app`.
- For standard user-facing front-end pages, do not categorize all UI tasks under this admin skill.

## Announce

Upon execution, provide a brief summary of:

- Whether the current admin is from scratch, an old UI, or a template upgrade.
- Whether you are checking information architecture, visual quality, or screenshot value first.
- The weakest area (workspace, list, form, login, or branding).
- If using a template, what you intend to keep vs. replace.

## Input Arguments

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `admin_surface_type` | enum(`None`/`Light`/`Heavy`) | Recommended | Determines if the admin UI quality pipeline is needed. |
| `ui_state` | enum(`From Scratch`/`Old UI Upgrade`/`Template Refactor`) | Recommended | Determines if it's a new build or a cleanup. |
| `template_strategy` | enum(`No Template`/`Template Scaffold`/`Continue Template Refactor`) | Recommended | Templates are for scaffolding, not final delivery. |
| `brand_state` | enum(`Has Guidelines`/`Name Only`/`None`) | Optional | Determines if color, icon, and copy baselines are needed first. |
| `release_pressure` | enum(`Dev Phase`/`Screenshot Prep`/`Final Submission`) | Optional | Prioritizes functional pages vs. screenshot/submission quality. |
| `critical_views` | string | Optional | Key page list (e.g., Dashboard, Order List, Detail Drawer, Form, Settings). |

## The Iron Law

1. Admin projects default to Vue + Element Plus; do not switch without explicit constraints.
2. Templates are allowed as scaffolds only; default branding, copy, icons, menus, charts, and data must be replaced.
3. An admin UI is not just a sidebar full of features. Navigation, grouping, and hierarchy must reflect real business flows.
4. Workspace, list, detail, form, settings, and login pages must share the same visual system.
5. Tables, filtering, pagination, batch actions, and status tags must serve real operational flows.
6. Forms must be controlled for length, grouping, validation feedback, and results; avoid long, raw input dumps.
7. Cover loading, empty, error, and restricted states; do not just polish the "happy path."
8. Submission screenshots must come from real business pages and data structures, not template boilerplate.
9. If an admin interface exists, provide a quality conclusion before release; functional completion is not UI completion.

## Workflow

### 1. Identify Admin Scope and Flow
- Inventory current modules: Dashboard, List, Detail, Approval, Config, Logs, Users, Permissions, Stats.
- Identify pages a reviewer is most likely to see and high-frequency user paths.
- Categorize as an operational console, management system, BI panel, or hybrid workspace.

### 2. Converge Information Architecture
- Organize top-level navigation, sub-grouping, and page naming.
- Merge low-frequency pages; split high-frequency hybrid pages.
- Determine content placement for Dashboard vs. List vs. Detail/Settings.
- If navigation is just template defaults, re-arrange based on real business needs.

### 3. Determine Template Strategy
If a template is needed, select one that:
- Is compatible with Vue + Element Plus.
- Is friendly to tables, forms, and detail panels.
- Has clear licensing.
- Is restrained enough to allow for business customization.

After scaffolding, perform these modifications:
- Replace brand colors, name, favicon, menu copy, and empty state copy.
- Delete template-specific sample charts, banners, and placeholder modules.
- Rebuild homepage cards, quick actions, and data blocks based on business modules.
- Adapt table columns, filters, status values, and detail structures to real business semantics.

### 4. Unify Visual System
- Establish admin color tokens, status colors, spacing, and card hierarchies.
- Standardize structures for page headers, action areas, filters, and detail areas.
- Unify button priorities, destructive actions, tag colors, and empty state illustrations.
- Ensure login/register pages and the admin shell share the same brand direction.

### 5. Converge Core Page Patterns
Check and refine these patterns:
- **Workspace**: Key metrics, primary entry points, pending tasks, recent activity.
- **List Page**: Filter, search, sort, batch actions, status tags, pagination.
- **Detail Page**: Summary info, primary data area, secondary info, timeline/logs.
- **Form Page**: Grouping, mandatory prompts, validation, feedback, multi-step flows.
- **Settings Page**: Security, integration, notifications, system params, members/permissions.
- **Auth Pages**: Login, Register, recovery entries, or business exemptions.

### 6. Add Interaction States and Error Paths
- Ensure skeleton screens, local loading, and disabled states are clear.
- Empty states should guide the next step, not just say "No Data."
- Failure messages must guide recovery, not just pop a generic error.
- Use secondary confirmation and feedback for risky actions (delete, sync, etc.).

### 7. Organize Screenshots and Submission Display
- Prepare screenshots for Workspace, List, Detail/Form, and Settings/Permissions.
- If the admin is a selling point, screenshots should showcase organization, control, and intelligence.
- Clean up sample charts, accounts, and boilerplate before taking screenshots.
- Hand over results to `lazycat:ship-app` for the release preparation phase.

## Quality Gates

- Admin interface has real business information architecture, not just template menus.
- Default branding, copy, charts, and sample data have been cleaned.
- Workspace, list, detail, form, settings, and login pages form a consistent visual system.
- Filtering, tags, empty states, and error handling are based on business semantics.
- Candidate screenshots showcase product value, not just an uncustomized template.
- UI quality conclusions have been handed over to `lazycat:ship-app`.

## Red Flags

- Delivering a raw admin template with just a name change.
- Retaining default template logos, menus, charts, or copyright info.
- All pages are just "Filter + Table" without workspace or detail logic.
- Forms are too long, lack grouping, or lack validation feedback.
- Empty, error, and restricted states are entirely unhandled.
- Screenshots include fake data, placeholder charts, or default accounts.
- Login pages look like a completely different project from the admin shell.

## Bundled References

- Admin UI Standards and Templates: [references/admin-ui-playbook.md](./references/admin-ui-playbook.md)
- Post-release Screenshots and Assets: [../lazycat:ship-app/references/store-assets.md](../lazycat:ship-app/references/store-assets.md)
- Pre-submission Quality Check: [../lazycat:ship-app/references/shipping-checklist.md](../lazycat:ship-app/references/shipping-checklist.md)

## Outputs

```text
Phase: Admin UI Planning / Refactor / Pre-submission Convergence
Admin Type: <Light / Heavy / N/A>

Confirmed
- ...

Information Architecture
- Primary Nav: ...
- Core Pages: ...

Template Strategy
- Source: <No Template / Template Name / Existing UI>
- Retain: ...
- Modify: ...

Visual and Interaction Conclusion
- Workspace: ...
- List Page: ...
- Detail/Form: ...
- Login/Register: ...

Screenshot Recommendations
- ...

Gaps / Risks
- ...

Next Steps
1. ...
2. ...

Deliverables
- Admin UI Quality Gates
- Candidate Screenshot List
- UI Quality Conclusion for lazycat:ship-app
```
