# Lazycat Store Assets Package

When tasks focus on "App Summary, Screenshots, Submission Materials, and Store Page Info," organize them per this template before filling them out in the Developer Center.

## Metadata Template

### One-sentence Position
- What is this app?
- Who is it for?
- What problem does it solve?

Template: `<App Name>` is a Lazycat app for `<Target User>` used to `<Core Value>`.

### App Description
Write in this order:
1. One-sentence Position.
2. 3 to 5 core features.
3. Prerequisites or dependencies.
4. Use cases.
5. Known limitations or notes.

Avoid:
- Exaggerated terms like "Best in the world," "Strongest," or "Unique."
- Promises that don't match actual features.
- Implying "out-of-the-box" without explaining prerequisites.
- Empty slogan-style sentences.

## Screenshot Script

Prepare at least these screens:
1. Homepage or main dashboard after startup.
2. Most critical operation step.
3. Feature page showing differentiated value.
4. Settings, management, or results page.

For AI Apps or AI Browser Plugins, add:
5. AI entry point or browser sidebar entry.
6. Core AI workflow page.
7. AI settings or model configuration page.

If the app focuses on admin/backend, add:
8. Dashboard or analytics view.
9. Most important list view.
10. Details page or key form page.

Each image must answer:
- Can a user understand what the app does from this image?
- Is this image identical to the actual submitted version?

## Screenshot Quality Requirements
- Use actual running screens, not design mockups.
- Maintain consistent language and visual theme.
- Remove debug markers, dummy data, sensitive info, and error prompts.
- If using a template, clean up default branding, sample charts, and placeholder copy.
- If a page has an empty state, prioritize showing a meaningful valid state.
- Follow fixed dimensions required by the Developer Center.

## Icon Outsourcing
- If the project lacks a formal icon or existing quality is low, use `lazycat:prepare-icon`.
- `lazycat:prepare-icon` should output an English prompt (Project Name + Function) for external image models.
- Check if the result is: `1024x1024`, no rounded corners, no text, no transparency, and clearly expresses the function.

## AI App Metadata Supplements
For AI Pod apps or plugins, also explain:
- Why it isn't a standard web app.
- Where the AI entry point is located.
- Whether users need extra configuration before first use.
- The role of the browser extension if included.

See `references/aipod-review-kit.md` for more detailed templates.

## Minimum Submission Package
- App Name
- Version Number
- One-sentence Position
- Full App Description
- Changelog or Update Notes
- Screenshot List
- Test Account, Reproduction Path, or Initialization Info
- External Dependencies and Limitations
- AI Pod / AI Browser entry points (if applicable)

## Check Before Re-submission
- Is the summary exaggerated, incorrect, or missing prerequisites?
- Are screenshots inconsistent with the current version?
- Are icons, links, screenshots, or copy broken or invalid?
- Does the actual experience support the promises made in the summary?
