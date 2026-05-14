# Independent Authentication

## HTTP Services

When accessing an app via a browser in Lazycat OS, you must enter a username and password for security.

However, for low-security scenarios like public file downloads where login is cumbersome, you can add a `public_path` sub-field under the `application` field in `lzc-manifest.yml`.

Additionally, apps with their own authentication mechanism (e.g., tokens in URLs) can disable mandatory OS-level authentication.

```yml
application:
  public_path:
    - /api/public
```

This configuration allows the browser to access the `/api/public` route directly without OS-level credentials.

Notes:
1. `public_path` only disables OS-level HTTP credential authentication. Users still need to be logged into the Lazycat client to establish the virtual network.
2. `public_path` carries risk; do not expose sensitive APIs (e.g., file reading services).

You can also use the `!` exclusion syntax to bypass mandatory authentication for an entire path except for specific subpaths like `/admin`. (Not recommended for general bypass).

```yml
application:
  public_path:
    - /
    - !/admin
```

::: warning Exclusion syntax has the highest priority and does not support nested logic.
For example, adding `/admin/unsafe` to the above rules would not take effect.
:::
