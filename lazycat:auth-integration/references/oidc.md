# OIDC Integration for lzcapp
====================

v1.3.5+ provides unified OIDC support. Once an lzcapp is adapted for OIDC, it can automatically retrieve the `uid` and corresponding permission groups (e.g., `ADMIN` for administrators).

Developers only need to provide the following two pieces of information in `manifest.yml` to complete the adaptation.

1. **Set the correct OIDC callback address in `application.oidc_redirect_path`**
   This path is typically `/oauth2/callback` or `/auth/oidc.callback`. Please refer to the application's documentation. If the documentation does not provide this information, you can enter a placeholder; the actual path will be visible in the browser error message during login.

2. **Retrieve system-generated environment variables during deployment and pass them to the application.**
   Required fields are `client_id` and `client_secret`. Some applications may only require an `ISSUER` for automatic endpoint discovery. Others might require multiple specific `ENDPOINT` details. For a full list of supported variables, refer to [Deployment Environment Variables](./advanced-envs#deploy_envs).

::: warning oidc_redirect_path
The system will only dynamically generate OIDC client-related environment variables if `application.oidc_redirect_path` is set.

If you are unsure of the value, you can enter a placeholder. Typically, the application's error page will indicate the correct redirect URI.
:::

### Example: Outline OIDC Adaptation
According to the [Outline official documentation](https://docs.getoutline.com/s/hosting/doc/oidc-8CPBm6uC0I), the following environment variables need to be set:
* `OIDC_CLIENT_ID` – OAuth client ID
* `OIDC_CLIENT_SECRET` – OAuth client secret
* `OIDC_AUTH_URI`
* `OIDC_TOKEN_URI`
* `OIDC_USERINFO_URI`

In `manifest.yml`, it can be configured as follows:
```yml
lzc-sdk-version: '2.0'
application:
  subdomain: outline
  # The official Outline documentation does not specify this, but it can be determined from error logs.
  oidc_redirect_path: /auth/oidc.callback
  routes:
    - /=http://outline.cloud.lazycat.app.outline.lzcapp:3000
services:
  outline:
    image: registry.lazycat.cloud/tx1ee/outlinewiki/outline:fb0e2ef4f32f3601
    environment:
      - OIDC_CLIENT_ID=${LAZYCAT_AUTH_OIDC_CLIENT_ID}
      - OIDC_CLIENT_SECRET=${LAZYCAT_AUTH_OIDC_CLIENT_SECRET}
      - OIDC_AUTH_URI=${LAZYCAT_AUTH_OIDC_AUTH_URI}
      - OIDC_TOKEN_URI=${LAZYCAT_AUTH_OIDC_TOKEN_URI}
      - OIDC_USERINFO_URI=${LAZYCAT_AUTH_OIDC_USERINFO_URI}
```

# OIDC Issuer Info
===============

Access `https://$MicroserviceName.heiyu.space/sys/oauth/.well-known/openid-configuration#/` to retrieve the complete issuer information.

Then use `https://$LAZYCAT_BOXDOMAIN/$endpoint_path` to resolve any endpoint address.

```json
{
"issuer": "https://your-box-name.heiyu.space/sys/oauth",
"authorization_endpoint": "https://your-box-name.heiyu.space/sys/oauth/auth",
"token_endpoint": "https://your-box-name.heiyu.space/sys/oauth/token",
"jwks_uri": "https://your-box-name.heiyu.space/sys/oauth/keys",
"userinfo_endpoint": "https://your-box-name.heiyu.space/sys/oauth/userinfo",
"device_authorization_endpoint": "https://your-box-name.heiyu.space/sys/oauth/device/code",
"introspection_endpoint": "https://your-box-name.heiyu.space/sys/oauth/token/introspect",
"grant_types_supported": [
"authorization_code",
"refresh_token",
"urn:ietf:params:oauth:grant-type:device_code",
"urn:ietf:params:oauth:grant-type:token-exchange"
],
"response_types_supported": [
"code"
],
"subject_types_supported": [
"public"
],
"id_token_signing_alg_values_supported": [
"RS256"
],
"code_challenge_methods_supported": [
"S256",
"plain"
],
"scopes_supported": [
"openid",
"email",
"groups",
"profile",
"offline_access"
],
"token_endpoint_auth_methods_supported": [
"client_secret_basic",
"client_secret_post"
],
"claims_supported": [
"iss",
"sub",
"aud",
"iat",
"exp",
"email",
"email_verified",
"locale",
"name",
"preferred_username",
"at_hash"
]
}
```
