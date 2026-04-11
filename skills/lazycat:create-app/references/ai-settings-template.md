# AI Settings Page Template for Standard Web Applications

This template is the default solution for integrating AI into "Standard Business Web Applications."

**Not applicable to:**
- Lazycat Computing Warehouse `AI Applications`
- AI Browser Extensions
- Projects requiring `ai-pod-service/`, `caddy-aipod`, or `extension.zip`

For these standard applications, the goal is to provide a stable, reusable AI configuration interface without introducing the additional complexity of the AI Pod package structure.

## 1. Required Fields

The first version must include these five essential elements:
1. `API BaseURL`
2. `API Protocol`
3. `Get Models` button
4. `Model` dropdown
5. `Save Configuration` button

## 2. Field Definitions

### `API BaseURL`
- **Type:** Input field
- **Required:** Yes
- **Purpose:** Specifies the base address for the model service.
- **Examples:**
  - `https://api.openai.com/v1`
  - `https://openrouter.ai/api/v1`
  - `<private-base-url>`

**Minimum Validation:**
- Cannot be empty.
- Must start with `http://` or `https://`.
- Trim leading and trailing whitespace before saving.

### `API Protocol`
- **Type:** Dropdown / Radio buttons
- **Required:** Yes
- **Fixed Options:**
  - `OpenAI Compatible`
  - `OpenAI Responses`
  - `Anthropic`

*Do not allow users to manually type the protocol name.*

### `Get Models` Button
- **Type:** Primary action button
- **Trigger Condition:** Both `API BaseURL` and `API Protocol` are filled.
- **Action:** Fetch the list of available models based on the current `BaseURL + Protocol`.

**Minimum Interaction:**
- Show a loading state during the request.
- Refresh the `Model` dropdown options upon success.
- Display a readable error message upon failure; do not fail silently.

### `Model` Dropdown
- **Type:** Select dropdown
- **Required:** Yes
- **Data Source:** The model list returned after clicking the `Get Models` button.

**Minimum Interaction:**
- Display placeholder text (e.g., "Please fetch models first") before data is loaded.
- Enable the dropdown only after a successful fetch.
- Ensure a valid option is selected before saving.

### `Save Configuration` Button
- **Type:** Primary button
- **Action:** Persist the current AI integration settings.

**Minimum Interaction:**
- Disable the button during the saving process to prevent duplicate submissions.
- Provide confirmation feedback upon success.
- Display a clear error message upon failure.

## 3. Recommended Page Layout

Arrange fields in this specific order:
1. `API BaseURL`
2. `API Protocol`
3. `Get Models`
4. `Model`
5. `Save Configuration`

*Do not combine "Get Models" and "Save Configuration" into a single button.*
*Do not use a free-text input for the "Model" field.*

## 4. Recommended State Flow

### Initial State
- `BaseURL` is empty or displays the previously saved value.
- `Protocol` uses the saved value or a default.
- `Model` dropdown is disabled.
- `Get Models` and `Save Configuration` buttons are visible.

### Successful Model Fetch
- Populate the `Model` dropdown.
- If a previously valid model still exists in the list, auto-select it.
- If the previous model is no longer available, prompt the user to re-select.

### Failed Model Fetch
- Retain current form values.
- Clear the `Model` dropdown or keep the old value marked as "to be confirmed."
- Display the reason for failure.

### Successful Save
- Persist the current `BaseURL / Protocol / Model`.
- Display a "Saved Successfully" notification or equivalent feedback.

## 5. Minimum Data Structure

```json
{
  "base_url": "https://api.openai.com/v1",
  "protocol": "openai_compatible",
  "model": "gpt-4.1-mini"
}
```

**Protocol Enum Suggestions:**
```text
openai_compatible
openai_responses
anthropic
```

## 6. Minimum API Recommendations

If implementing as a backend configuration API, the first version should provide equivalent capabilities:
- `GET /settings/ai`
- `PUT /settings/ai`
- `POST /settings/ai/models`

**Examples:**

### Get Saved Configuration
```http
GET /settings/ai
```

### Save Configuration
```http
PUT /settings/ai
Content-Type: application/json

{
  "base_url": "https://api.openai.com/v1",
  "protocol": "openai_compatible",
  "model": "gpt-4.1-mini"
}
```

### Fetch Models
```http
POST /settings/ai/models
Content-Type: application/json

{
  "base_url": "https://api.openai.com/v1",
  "protocol": "openai_compatible"
}
```

## 7. Optional Extensions

These are not mandatory for the first version and should only be added if required by the business:
- API Key
- Timeout settings
- Organization / Workspace ID
- Default System Prompt
- Advanced parameters (e.g., Temperature)

*Place these in an "Advanced Settings" section to avoid cluttering the primary configuration.*

## 8. Quality Thresholds

- All 5 required elements must be clearly present on the page.
- Protocol options must be fixed (no free-text input).
- Models must be sourced from a real fetch operation, not hardcoded.
- Saving logic must only persist the current selection without hidden defaults.
- Error messages must be user-friendly and actionable.
