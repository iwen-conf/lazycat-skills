# Lazycat 命令入口约定

参考 `cloudmaze` 的做法，移植项目至少要提供这些入口：

- 根目录 `build.sh`
- 根目录 `Makefile`
- `make build`
- `make install`

## 1. `build.sh`

作用：

- 串起后端构建
- 串起前端构建
- 输出打包所需产物到统一目录

不要把构建步骤只写在文档里，必须让脚本可执行。

## 2. `Makefile`

最少目标：

- `build`
- `install`

推荐目标：

- `dev`
- `test`
- `clean`

## 3. 参考风格

以 `cloudmaze` 为例：

- `build.sh` 负责后端编译和前端构建
- `Makefile` 负责调用 `lzc-cli project build` 与 `lzc-cli app install`

核心原则不是完全复制，而是保证任何人拿到仓库后，都能通过统一命令完成构建和安装。
