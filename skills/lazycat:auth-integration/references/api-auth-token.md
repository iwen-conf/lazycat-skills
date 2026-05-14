# API Auth Token

The API Auth Token is used for authentication when accessing system APIs via scripts or the command line, eliminating the dependency on browser session states. It is suitable for automation, O&M (Operations and Maintenance) scripts, CI/CD, and similar scenarios.

Requires lzcos v1.4.3+.

## Generation and Management

```bash
hc api_auth_token gen
hc api_auth_token gen --uid admin
hc api_auth_token list
hc api_auth_token show <token>
hc api_auth_token rm <token>
```

- `gen` creates a token in UUID format.
- `--uid` specifies the bound user; if omitted, it defaults to the administrator.

## Usage Example

```bash
curl -k -H "Lzc-Api-Auth-Token: <token>" "https://<box-domain>/sys/whoami"
```

## Behavior Description

- The Header name is fixed as `Lzc-Api-Auth-Token`.
- This header is used exclusively for system authentication and will be stripped before being forwarded to applications.
- Token permissions are equivalent to the bound user. Please store it securely and prevent leaks.
- Note: Since some LPKs use client information within the authentication data for reverse access, this feature is not supported under API TOKEN authentication.
- In this mode, the system does not automatically inject `X-HC-Device-PeerID` and `X-HC-Device-ID`.
- In this mode, `X-HC-Login-Time` is set to the creation time of the `Token`.
