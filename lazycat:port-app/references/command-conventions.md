# Lazycat 命令入口约定

所有 Lazycat 项目必须在根目录提供统一的命令入口，确保自动化构建、验证和提审流程一致。

## 1. 核心文件
- **build.sh**: 负责具体的前后端构建逻辑。
- **Makefile**: 负责串联流程，定义标准 target。

## 2. Makefile 标准模板

```makefile
# 定义文件名变量，方便修改
LPK_FILE = app.lpk

.PHONY: all build install update doctor backend-test ui-test ui-build ui-e2e capture-screenshots verify release-prep

# 默认执行 install
all: install

doctor:
	@echo "检查开发环境..."
	@lzc-cli version || (echo "错误: 未安装 lzc-cli" && exit 1)
	@lzc-cli user info || (echo "错误: lzc-cli 未登录，请先执行 lzc-cli login" && exit 1)
	@echo "环境就绪"

backend-test:
	@echo "运行后端测试..."
	# 示例：cd backend && go test ./...

ui-test:
	@echo "运行前端单元测试..."
	# 示例：cd ui && npm test

ui-build:
	@echo "构建前端..."
	# 示例：cd ui && npm run build

ui-e2e:
	@echo "运行前端 E2E..."
	# 示例：cd ui && ./node_modules/.bin/playwright install chromium && npm run test:e2e

capture-screenshots:
	@echo "生成提审截图..."
	# 示例：cd ui && ./node_modules/.bin/playwright install chromium && npm run capture:screenshots

verify: backend-test ui-test ui-build ui-e2e
	@echo "验证全部通过"

release-prep: verify capture-screenshots
	@echo "提审素材已生成"

# 1. 构建任务
build:
	@echo "构建中..."
	lzc-cli project build -o $(LPK_FILE)

# 2. 安装任务 (依赖于 build，所以 build 会先执行)
install: build
	@echo "安装中..."
	lzc-cli app install $(LPK_FILE)
	@echo "构建完成已经安装"

# 3. 更新任务 (用于镜像升级或代码同步)
update:
	@echo "执行更新流程..."
	# 1. 使用 lzc-cli appstore copy-image 获取最新镜像
	# 2. 修改 manifest.yml 指向新镜像
	# 3. 重新打包
	# 示例：IMAGE_NAME=$$(grep "image:" manifest.yml | awk '{print $$2}') && \
	#      NEW_IMAGE=$$(lzc-cli appstore copy-image $$IMAGE_NAME) && \
	#      sed -i "s|image:.*|image: $$NEW_IMAGE|" manifest.yml
	lzc-cli project build -o $(LPK_FILE)
```

## 3. 核心 Target 说明

- **make install**: 构建并安装到当前懒猫微服。**提审前必须执行并体验验证。**
- **make update**: 用于移植项目的版本升级。会自动处理镜像同步和 manifest 更新。
- **make release-prep**: 提审前的最后一步。包含单元测试、E2E 测试和自动截图。生成提审所需的元数据证据。
- **make verify**: 纯校验任务，不产生副作用，用于 CI/CD。

## 4. 镜像移植约定

对于移植项目，如果在 `manifest.yml` 中使用外部镜像，必须遵循：
1. 使用 `lzc-cli appstore copy-image` 将镜像同步到懒猫。
2. 将返回的 `private.ezer.heiyu.space/...` 地址填入 `manifest.yml`。
3. `make update` 应该能自动或辅助完成此过程。
