# HTTP Headers

All HTTPS/HTTP traffic initiated from the client is first routed through the `lzc-ingress` component for distribution.

`lzc-ingress` handles the following:
- Authenticates HTTP requests; redirects to login if unauthenticated.
- Routes traffic to different app backends based on the request domain.

Before forwarding traffic to a specific app, `lzc-ingress` sets the following additional HTTP headers:

- `X-HC-User-ID`: Logged-in UID (Username).
- `X-HC-Device-ID`: Unique ID of the client within this OS instance; can be used as a device identifier.
- `X-HC-Device-PeerID`: Client peer ID (Internal use only).
- `X-HC-Device-Version`: Client kernel version number.
- `X-HC-Login-Time`: Last login time of the client (Unix timestamp, int32).
- `X-HC-User-Role`: "NORMAL" for regular users, "ADMIN" for administrators.
- `X-Forwarded-Proto`: Fixed to "https" to ensure apps requiring HTTPS work correctly.
- `X-Forwarded-By`: Fixed to "lzc-ingress".

`lzc-ingress` uses the `HC-Auth-Token` cookie for authentication (client-side uses other internal methods).

If this cookie is invalid or empty, and the target address is not in `public_path`, it redirects to the login page.

If the target address is in `public_path`, `lzc-ingress` still performs authentication but will not redirect:
- If authentication fails, `X-HC-XX` headers are cleared to avoid security risks.
- If authentication succeeds, `X-HC-XX` headers are included.

Developers can simply trust `X-HC-User-ID` without checking if a path is in `public_path`.
