# Using Multiple Domains for a Single Application
===================

While `lzc-manifest.yml:application.subdomain` defines the developer's desired domain, the microservice system (v1.3.6+) may adjust it:

1. **Conflicts**: If multiple applications use the same `subdomain`, subsequently installed applications will have a suffix appended to their domain.
2. **Multi-instance Apps**: For multi-instance applications, each user is assigned a unique domain. Non-admin users will likely see a domain with a suffix.
3. **Domain Prefixes**: Domains like `prefix-subdomain` resolve identically to `subdomain`. Each application automatically owns an arbitrary number of prefixed domains.
4. **Environment Variables**: The final assigned `subdomain` can only be retrieved via the `LAZYCAT_APP_DOMAIN` environment variable.
5. **Ingress Exclusion**: Traffic entering via prefixed domains ignores `TCP/UDP Ingress` configurations (this does not affect traffic via the default application domain).

v1.3.8+ supports [domain-based traffic forwarding](./advanced-route#upstreamconfig).

Since `application.routes` does not support domain-based forwarding, use a special route rule like `- /=http://nginx.$appid.lzcapp` for fine-grained control. Note: You must use the `$service.$appid.lzcapp` format to ensure the upstream receives the complete hostname. [See details here](advanced-route.html#p2).

**Example Configuration:**
1. Default domain opened from the app list: `whoami.xx.heiyu.space` (assuming the assigned subdomain is `whoami`).
2. `nginx-whoami.xx.heiyu.space` returns the default Nginx "Hello World" page.
3. `anything-whoami.xx.heiyu.space` behaves the same as `whoami.xx.heiyu.space`.

```yaml
package: org.snyh.debug.whoami
name: whoami-lazycatmicroserver

application:
  subdomain: whoami
  routes:
    - /=http://nginx.org.snyh.debug.whoami.lzcapp:80

services:
  nginx:
    image: registry.lazycat.cloud/snyh1010/library/nginx:54809b2f36d0ff38
    setup_script: |
      cat <<'EOF' > /etc/nginx/conf.d/default.conf
      server {  # Route whoami.xxx.heiyu.space and all other prefixes to traefik/whoami
         server_name _;
         location / {
            proxy_pass http://app1:80;
         }
      }
      server {  # Route domains starting with "nginx" to Nginx default page
         server_name  ~^nginx.*-.*;
         location / {
            root   /usr/share/nginx/html;
            index  index.html index.htm;
         }
      }
      EOF
  app1:
    image: registry.lazycat.cloud/snyh1010/traefik/whoami:c899811bc4a1f63a
```
