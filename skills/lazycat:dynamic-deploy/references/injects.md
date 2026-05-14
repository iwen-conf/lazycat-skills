# Script Injection (injects)

## Overview

`injects` is used to inject scripts into HTML pages matching specific URLs, suitable for minimal-intrusion adaptation of third-party applications. This feature requires lzcos 1.5.0+.
For field definitions, please refer to [manifest.md#injects](./spec/manifest.md#injects); this document focuses on behavioral details and practical recommendations.

## Compatibility First

There are at least two inject syntaxes in the field, and you must not assume the newer one is available:

- **Legacy syntax**: `on / when / do`
- **Newer syntax in newer docs**: `mode / include / exclude / scripts`

Operational rule:

1. If the target box, installer, or `lzc-cli project build` reports unknown fields for `mode/include/scripts`, immediately fall back to `on/when/do`.
2. If installation fails with `application.injects.0 when is required`, that is a strong signal the target environment expects the legacy syntax.
3. Do not treat a successful build as proof that the inject fields will install correctly.

Legacy syntax example:

```yml
application:
  injects:
    - id: login-autofill
      on: browser
      when:
        - "/login"
        - "/#login"
      do:
        - src: builtin://simple-inject-password
          params:
            user: "admin"
            password: "admin123"
```

Only use the newer `mode/include/scripts` form after confirming the target box really supports it.

## Quick Example

```yml
application:
  injects:
    - id: login-autofill
      on: browser
      when:
        - "/"
        - "/?version=1.2&channel=stable"
        - "/#login"
      do:
        - src: builtin://simple-inject-password
          params:
            user: "admin"
            password: "admin123"
```

## Rule Model

The rest of this document describes the matching semantics conceptually. If you are targeting a box that only supports legacy syntax, map those concepts back to `when/do` instead of emitting unsupported fields.

- `include`: Whitelist rules; any match enters candidacy.
- `exclude`: Blacklist rules; any match results in exclusion.
- `mode`: `exact` or `prefix` (default: `exact`), applied to `path/hash`.
- `prefix_domain`: If not empty, only matches requests with the `<prefix>-` domain prefix.
- `injects` entries are injected in the order of declaration; `scripts` within each entry are also injected sequentially.

## Rule Syntax

Single rule format:

`<path>[?<query>][#<hash>]`

Examples:

- `"/"`
- `"/?version=1.2"`
- `"/#login"`
- `"/app?version=1.2&channel=stable#signin"`

Parsing Details:

- `path` is required and must start with `/`.
- `query` tokens support two forms:
  - `key`: Requires the key to exist.
  - `key=value`: Requires the key to have at least one value equal to the specified value.
- Multiple query tokens within a single rule use AND logic (all must be satisfied).
- Query matching uses "contains" semantics: the request may include additional parameters.

## Matching Semantics

Overall Logic:

- `include` is OR: Any match enters candidacy.
- `exclude` is OR: Any match results in rejection.
- Final Result: `matched = includeMatched && !excludeMatched`.

Single Rule Logic (AND):

- `path` matches.
- `query` matches (if declared).
- `hash` matches (if declared).

`mode` Semantics (applied to `path/hash`):

- `exact`: Exact match.
- `prefix`: Prefix match.

## Hash Behavior (hard/soft)

- `path/query` are server-visible conditions (hard matching).
- `hash` is a server-invisible condition, automatically downgraded to client-side "soft" matching.

This means:

- The server first decides whether to inject the wrapper based on `path/query`.
- The wrapper then decides whether to execute the script on the browser side based on the full rule (including `hash`).
- It is possible that the wrapper is injected but the script is not executed because the hash does not match; this is expected behavior.

## Execution Timing and Runtime Parameters

Wrapper trigger timing:

1. Evaluation occurs once after page load (`trigger=load`).
2. Listens for `hashchange`; evaluation occurs after each hash change (`trigger=hashchange`).
3. Scripts are executed as long as the rule matches; no built-in deduplication is performed.

Scripts can read the following objects:

- `__LZC_INJECT_PARAMS__`: Parameters from `scripts[].params`.
- `__LZC_INJECT_RUNTIME__`:
  - `executedBefore`: Whether this script has been executed previously within the current page lifecycle.
  - `executionCount`: The current execution count (starts from `1`).
  - `trigger`: `load` or `hashchange`.

Example (Script Side):

```js
(() => {
  const runtime = __LZC_INJECT_RUNTIME__ || {};
  if (runtime.executedBefore) {
    return;
  }
  const params = __LZC_INJECT_PARAMS__ || {};
  console.log("inject params:", params);
})();
```

## Script Sources

`scripts[].src` supports:

- `builtin://name`: Use a script built into `lzcinit`.
- `file:///path`: Read a script from the application filesystem (common path: `/lzcapp/pkg/content/`).
- `http(s)://...`: Remote script (recommended for debugging only).

Remote script loading uses conditional request caching (`ETag`/`Last-Modified`).

## Built-in Scripts

### `builtin://hello`

Prints debugging information.

Parameters:

- `message`: Output content (default: `hello world`).

### `builtin://simple-inject-password`

Automatically fills in account/password and optionally submits the form. Recommended for injection only on explicit login paths.

Parameter Description (`params`):

| Parameter | Type | Description |
| ---- | ---- | ---- |
| `user` | `string` | Account value (default empty). |
| `password` | `string` | Password value (default empty). |
| `requireUser` | `bool` | Whether the account input must be found. Defaults to `false` if `allowPasswordOnly=true`, otherwise `true` if `user` is non-empty. |
| `allowPasswordOnly` | `bool` | Allows filling only the password (default `false`). |
| `autoSubmit` | `bool` | Whether to automatically submit (default `true`). |
| `submitMode` | `string` | Submission mode: `auto`, `requestSubmit`, `click`, or `enter` (default `auto`). |
| `submitDelayMs` | `int` | Delay before auto-submission (ms, default `50`, min `0`). |
| `retryCount` | `int` | Auto-submission retry count (default `10`). |
| `retryIntervalMs` | `int` | Auto-submission retry interval (ms, default `300`). |
| `observerTimeoutMs` | `int` | DOM/state observation timeout (ms, default `8000`). |
| `debug` | `bool` | Enables debug logging (default `false`). |
| `userSelector` | `string` | Explicitly specify the account input selector. |
| `passwordSelector` | `string` | Explicitly specify the password input selector. |
| `formSelector` | `string` | Limit input search to a specific container. |
| `submitSelector` | `string` | Explicitly specify the submit button selector. |
| `allowHidden` | `bool` | Allow filling invisible inputs (default `false`). |
| `allowReadOnly` | `bool` | Allow filling read-only inputs (default `false`). |
| `onlyFillEmpty` | `bool` | Only fill if the input is empty (default `false`). |
| `allowNewPassword` | `bool` | Allow filling password boxes with `autocomplete=new-password` (default `false`). |
| `includeShadowDom` | `bool` | Whether to search open Shadow DOMs (default `false`). |
| `shadowDomMaxDepth` | `int` | Maximum Shadow DOM recursion depth (default `2`). |
| `preferSameForm` | `bool` | Prioritize selecting an account field within the same form as the password field (default `true`). |
| `eventSequence` | `string` or `[]string` | Sequence of events to trigger (default `input,change,keydown,keyup,blur`). |
| `keyValue` | `string` | Key value for keyboard events (default `a`). |
| `userKeywords` | `string` or `[]string` | Additional account field keywords (comma-separated or array). |
| `userExcludeKeywords` | `string` or `[]string` | Additional account field exclusion keywords. |
| `passwordKeywords` | `string` or `[]string` | Additional password field keywords. |
| `passwordExcludeKeywords` | `string` or `[]string` | Additional password field exclusion keywords. |
| `submitKeywords` | `string` or `[]string` | Additional submit button keywords. |

## Practical Recommendations

- When multiple pages are needed, prioritize adding multiple `include` rules rather than relaxing the rules.
- For login redirection scenarios, use query conditions for constraints, e.g., `"/?version=1.2&channel=stable"`.
- For hash routing scenarios, use `__LZC_INJECT_RUNTIME__.executedBefore` within the script to control reruns.
- Strongly recommend injecting usernames and passwords via deployment parameters to avoid hardcoding weak passwords in code or the manifest.
- Non-HTML responses are naturally excluded from injection. Use `exclude` to further restrict the scope of HTML pages (e.g., paths like `/admin/debug`).
- Remote scripts are recommended for debugging only; use `builtin://` or `file://` for official releases.
