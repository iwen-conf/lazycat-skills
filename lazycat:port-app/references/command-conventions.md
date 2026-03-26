# Lazycat Command Entry Point Conventions

All Lazycat projects must provide unified entry points in the root directory to ensure consistent automated build, verification, and submission workflows.

## 1. Core Files
- **build.sh**: Responsible for specific backend/frontend build logic.
- **Makefile**: Orchestrates workflows and defines standard targets.

## 2. Standard Makefile Template

```makefile
# Define filename variables for easy modification
LPK_FILE = app.lpk

.PHONY: all build install update doctor backend-test ui-test ui-build ui-e2e capture-screenshots verify release-prep

# Default to install
all: install

doctor:
        @echo "Checking development environment..."
        @lzc-cli version || (echo "Error: lzc-cli not installed" && exit 1)
        @lzc-cli user info || (echo "Error: lzc-cli not logged in, please run lzc-cli login" && exit 1)
        @echo "Environment ready"

backend-test:
        @echo "Running backend tests..."
        # Example: cd backend && go test ./...

ui-test:
        @echo "Running frontend unit tests..."
        # Example: cd ui && npm test

ui-build:
        @echo "Building frontend..."
        # Example: cd ui && npm run build

ui-e2e:
        @echo "Running frontend E2E tests..."
        # Example: cd ui && ./node_modules/.bin/playwright install chromium && npm run test:e2e

capture-screenshots:
        @echo "Generating submission screenshots..."
        # Example: cd ui && ./node_modules/.bin/playwright install chromium && npm run capture:screenshots

verify: backend-test ui-test ui-build ui-e2e
        @echo "All verifications passed"

release-prep: verify capture-screenshots
        @echo "Submission assets generated"

# 1. Build Task
build:
        @echo "Building..."
        lzc-cli project build -o $(LPK_FILE)

# 2. Install Task (depends on build)
install: build
        @echo "Installing..."
        lzc-cli app install $(LPK_FILE)
        @echo "Build complete and app installed"

# 3. Update Task (for image upgrades or code sync)
update:
        @echo "Executing update workflow..."
        # 1. Use lzc-cli appstore copy-image to fetch the latest image
        # 2. Modify manifest.yml to point to the new image
        # 3. Re-package
        # Example: IMAGE_NAME=$$(grep "image:" manifest.yml | awk '{print $$2}') && \
        #          NEW_IMAGE=$$(lzc-cli appstore copy-image $$IMAGE_NAME) && \
        #          sed -i "s|image:.*|image: $$NEW_IMAGE|" manifest.yml
        lzc-cli project build -o $(LPK_FILE)
```

## 3. Core Target Descriptions

- **make install**: Builds and installs to the current Lazycat OS instance. **Must be executed and verified before submission.**
- **make update**: Used for version upgrades in ported projects. Automatically handles image synchronization and manifest updates.
- **make release-prep**: The final step before submission. Includes unit tests, E2E tests, and automated screenshots. Generates metadata evidence for review.
- **make verify**: Pure validation task with no side effects, used for CI/CD.

## 4. Image Porting Conventions

For ported projects using external images in `manifest.yml`:
1. Use `lzc-cli appstore copy-image` to sync the image to Lazycat.
2. Fill the returned `private.ezer.heiyu.space/...` address into `manifest.yml`.
3. `make update` should automate or assist this process.
