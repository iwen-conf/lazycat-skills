# GitHub Search & App Store Duplication Check

Before starting a porting project, perform two types of searches:

## 1. GitHub Search
Record at least:
- Repository Name
- Upstream URL
- License
- Stars / Activity
- Last Update Time
- Deployment Complexity
- Inclusion of Web UI / API / Login system

## 2. App Store Duplication Check
Search on `https://appstore.ezer.heiyu.space/#/shop` for:
- English Project Name
- Chinese Project Translation
- Core Function Keywords
- Possible Aliases

The site usually requires a login session. Do not assume "no search results" means "no duplicate" if you are not logged in.

If local store environment variables are provided:
- `lazycat_account`
- `lazycat_password`

Prioritize using these for login before checking. These are for Lazycat OS and App Store access, not Developer Center or internal app accounts.

Do not mix variables with different scopes:
- `lazycat_developer_center_account` / `lazycat_developer_center_password`: Developer Center only.
- `lazycat_gitea_account` / `lazycat_gitea_password`: Internal login for the Gitea app only.

Note:
- If these variables are in `~/.zshrc`, non-interactive shells may not see them.
- Use an interactive shell `zsh -ic` to read them and drive the browser login.
- Do not misjudge "no credentials on this machine" just because they aren't in standard `env`.

Recommended Check Order:
1. Check for an existing login session.
2. If none, try reading `lazycat_account` / `lazycat_password`.
3. After successful login, search project names, translations, aliases, and keywords.
4. Record matching products, overlaps, and room for differentiation.
5. If aiming for incentives, verify the app by installing and testing it within Lazycat OS.

## 3. Handling Duplicates
- If a similar port exists with high overlap and no clear differentiation, do not pursue the incentive path.
- If a similar product exists but you can provide significant added value, clearly state the differentiation.
- Pure "re-shelling" is not worth pursuing.

## 4. Differentiation Examples
- Integrating OIDC when the existing port doesn't.
- Integrating `file_handler` to open Lazycat Drive files directly.
- UI/UX that better fits Lazycat scenarios.
- Significantly better stability, maintainability, or initialization experience.

Do not consider "OIDC planned" as established differentiation. If the app is weak or the experience is poor, incentives may still be denied.
