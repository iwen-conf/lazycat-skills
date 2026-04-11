# lzc-build.yml Specification Document

## I. Overview

`lzc-build.yml` is used to define configurations related to application building. This document details its structure and field meanings.

## II. Top-Level Data Structure: `BuildConfig`

### 2.1 Basic Information {#basic-config}

| Field Name | Type | Description |
| ---- | ---- | ---- |
| `buildscript` | `string` | Path to a build script or a shell command. |
| `manifest` | `string` | Path to the `manifest.yml` file for the lpk package. |
| `contentdir` | `string` | Directory for static content to be packaged into the lpk. |
| `pkgout` | `string` | Output path for the lpk package. |
| `icon` | `string` | Path to the lpk package icon. Must be a PNG. Warns if missing. |
| `devshell` | `DevshellConfig` | Configuration for development dependencies. |
| `compose_override` | `ComposeOverrideConfig` | Advanced compose override configuration (**requires lzc-os >= v1.3.0**). |

## III. Development Dependencies: `DevshellConfig` {#devshell}

| Field Name | Type | Description |
| ---- | ---- | ---- |
| `routes` | `[]string` | Development routing rule configuration. |
| `dependencies` | `[]string` | Automatic installation of development dependencies. |
| `setupscript` | `string` | Manual installation script for development dependencies. |
| `image` | `string` | (Optional) Use a specified Docker image. |
| `pull_policy` | `string` | (Optional) Set to `build` to use a specified Dockerfile; use `${package}-devshell:${version}` for the `image` parameter. |
| `build` | `string` | (Optional) Path to the Dockerfile used for building the container. |

For details, see [Development Dependency Installation](../devshell-install-and-use.md).

::: warning ⚠️ Note

If both `dependencies` and `build` are present, `dependencies` takes precedence.

:::

## IV. Advanced Compose Override: `ComposeOverrideConfig` {#compose-override}

1. `compose_override` is supported in `lzc-cli@1.2.61` and above for specifying compose override configurations during build time.
2. It addresses runtime permission requirements in `lzcos v1.3.0+` that are not covered by current lpk specifications.

For details, see [Compose Override](../advanced-compose-override.md).

::: details Configuration Example
```yml
# Use ${var} to reference values defined in the manifest file.

# buildscript:
# - Path to a build script.
# - Can be a simple shell command.
# - ⚠️ Critical: Never execute `lzc project build` or related commands inside this script to avoid infinite loops.
buildscript: sh build.sh

# manifest: Path to the lzc-manifest.yml file.
manifest: ./lzc-manifest.yml

# contentdir: Static content to be packaged into the lpk.
contentdir: ./dist

# pkgout: Output path for the lpk package.
pkgout: ./

# icon: Path to the icon (must be PNG).
icon: ./lzc-icon.png

compose_override:
  services:
    # Service name
    some_container:
      # Capabilities to drop
      cap_drop:
        - SETCAP
        - MKNOD
      # Files/directories to mount
      volumes:
        - /data/playground:/lzcapp/run/playground:ro

# devshell: Development dependencies configuration.
# Uses alpine:latest as the base image. Add dependencies as needed.
# If both dependencies and build exist, dependencies takes precedence.
devshell:
  routes:
    - /=http://127.0.0.1:5173
  dependencies:
    - nodejs
    - npm
    - python3
    - py3-pip
  setupscript: |
    export npm_config_registry=https://registry.npmmirror.com
    export PIP_INDEX_URL=https://pypi.tuna.tsinghua.edu.cn/simple   
```
:::
