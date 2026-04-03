# manifest.yml Rendering

lzcos v1.3.8+ supports dynamic rendering of the `manifest.yml` file, giving developers better control over deployment parameters.

The `manifest.yml` rendering process is as follows:

1. The developer creates an [lzc-deploy-params.yml](./spec/deploy-params) file in the project root and packages it into the lpk using `lzc-cli project build` (requires lzc-cli v1.3.7+).
2. Before execution, the system redirects to a UI for parameter supplementation, requiring the user to fill in all parameters defined in `lzc-deploy-params.yml`.
3. The system retrieves the user-provided parameters and uses them as template parameters (`U`) to render the final `lzc-manifest.yml` from the lpk.
4. The final manifest is stored at `/lzcapp/run/manifest.yml` (relative to the original file `/lzcapp/pkg/manifest.yml`) and is used as the definitive configuration.

--------------

1. Users can re-enter the deployment parameter modification page from the application list (Step 2). After each modification, the application instance is stopped and the process repeats.
2. In multi-instance applications, each user's deployment parameters are independent and filled by the user.
3. Manifest rendering occurs even if `lzc-deploy-params.yml` is not configured.

## Rendering Mechanism

`manifest.yml` is rendered using Golang's `text/template`. You should be familiar with the [official Go template syntax](https://pkg.go.dev/text/template). Below are built-in template functions and parameters.

## Built-in Template Functions

1. Functions supported by [sprig](https://masterminds.github.io/sprig/) (except for environment-related functions).
2. `stable_secret "seed"`: Generates a stable password. Requires an arbitrary string as a seed.
    1. The same seed results in different outputs for different applications.
    2. The same seed results in different outputs for different Micro-services of the same application.
    3. The same seed results in the same output for the same Micro-service instance across multiple calls (unless factory reset).

## Template Parameters

Mainly consists of two major parameter sets (aliases in parentheses):

- `.UserParams (.U)`: Parameters defined in `lzc-deploy-params.yml`.

- `.SysParams (.S)`: System-related parameters.
    - `.BoxName`: Name of the Micro-service.
    - `.BoxDomain`: Domain of the Micro-service.
    - `.OSVersion`: OS version. Beta versions are forced to `v99.99.99-xxx`.
    - `.AppDomain`: Application domain. Currently hardcoded by developers; future versions will support dynamic allocation and administrator adjustment.
    - `.IsMultiInstance`: Whether it is a multi-instance deployment. Currently hardcoded; future versions will allow dynamic adjustment by administrators.
    - `.DeployUID`: User ID during actual deployment (not available in single-instance mode).
    - `.DeployID`: Unique ID of the instance itself.

::: tip
During debugging, you can add the following anywhere in `lzc-manifest.yml` to render all available parameters:
```
xx-debug: {{ . }}
```

You can add a rule in `application.route` to view the final `manifest.yml`:
```
application:
    route:
        - /m=file:///lzcapp/run/manifest.yml
```
Alternatively, use `devshell` and run `cat /lzcapp/run/manifest.yml`.
:::

## Examples

For a complete demo, refer to [this repository](https://gitee.com/lazycatcloud/netmap).

### More Secure Internal Passwords

```yml
lzc-sdk-version: '2.0'
services:
  mysql:
    binds:
    - /lzcapp/var/mysql:/var/lib/mysql
    environment:
    - MYSQL_ROOT_PASSWORD={{ stable_secret "root_password" }}
    - MYSQL_USER=LAZYCAT
    - MYSQL_PASSWORD={{ stable_secret "admin_password" | substr 0 6 }}
    image: registry.lazycat.cloud/mysql
  redmine:
    environment:
    - DB_PASSWORD={{ stable_secret "root_password" }}
```

### Different Configurations for Multi-instance vs. Single-instance

In single-instance mode, store data in each user's document directory.
In multi-instance mode, store user data within the application.

```yml
# lzc-manifest.yml

services:
  some_service_name:
    binds:
    {{ if .S.IsMultiInstance }}
      - /lzcapp/run/mnt/home:/home/
    {{ else }}
      - /lzcapp/run/mnt/home/{{ .S.DeployUID }}/the_name:/home/
    {{ end }}
```

### Startup Parameters Configured by Users
```yml
# lzc-deploy-params.yml
params:
  - id: target
    type: string
    name: "target"
    description: "The target IP you want to forward to"

  - id: listen.port
    type: string
    name: "listen port"
    description: "The forwarder listen port (cannot be 80 or 81)"
    default_value: "33"
    optional: true
```

```yml
# lzc-manifest.yml
lzc-sdk-version: '2.0'
application:
  subdomain: netmap

  upstreams:
    - location: /
      backend_launch_command: /lzcapp/pkg/content/netmap -target={{ .U.target }} -port={{ index .U "listen.port" }}
```
```
