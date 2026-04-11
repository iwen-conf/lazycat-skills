# deploy-params

`lzc-deploy-params.yml` is the configuration file used by developers to define installation-time parameters. This document details its structure and the meaning of each field.

# DeployParams

| Field Name | Type | Description |
| ---- | ---- | ---- |
| `params` | `[]DeployParam` | List of deployment parameters defined by the developer. |
| `locales` | `map` | Internationalization settings. |

-------------------------------

# DeployParam
| Field Name | Type | Description |
| ---- | ---- | ---- |
| `id` | `string` | Unique ID within the app, used for internationalization and references in `manifest.yml`. |
| `type` | `string` | Field type. Currently supports `bool`, `lzc_uid`, `string`, and `secret`. |
| `name` | `string`| Display name for the field, supports internationalization. |
| `description` | `string`| Detailed description for the field, supports internationalization. |
| `optional` | `bool` | Whether the field is optional. If true, the user is not forced to fill it; if all fields are optional, the deployment UI may be skipped. |
| `default_value`| `string`| Default value provided by the developer. Supports `$random(len=5)` to generate a random string (lzcos 1.5.0+). |
| `hidden` | `bool` | If true, the field is active but not rendered in the UI. |
