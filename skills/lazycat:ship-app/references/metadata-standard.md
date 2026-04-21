# Lazycat Store Metadata Quality Standards

For LPK v2, the following fields must be defined in `package.yml` according to these standards during submission. Leaving them blank or filling them perfunctorily is strictly prohibited.

## 1. App Name (Name)
*   **Standard**: [Chinese Name] - [English Name/Brand Name]
*   **Requirement**: Do not include suffixes like "Test", "Latest", or "V1".
*   **Example**: 云之幻 - CloudPsi

## 2. App Summary (Summary)
*   **Standard**: A single sentence describing the core value.
*   **Length**: 10 - 30 characters (Chinese) or equivalent.
*   **Example**: High-performance 3D rendering and collaboration platform for professional teams.

## 3. App Description (Description)
Must include the following three sections:
*   **Features**: List at least 3 core functionalities.
*   **Instructions**: Explain how to get started.
*   **Credentials**: For ported or authenticated apps, specify default accounts/passwords or registration methods.
*   **Example**:
    ```text
    [Features]
    - Real-time multi-device synchronization
    - Built-in AI rendering acceleration
    - Compatible with all major file formats

    [Instructions]
    Open the app after installation and configure your storage path in settings to begin.

    [Credentials]
    Default Username: admin, Password: lazycat123
    ```

## 4. Keywords (Keywords)
*   **Standard**: Business category, Tech stack, Core features (at least 5 keywords).
*   **Example**: 3D Rendering, GPU Acceleration, Go, Collaboration Tool, Design

## 5. Changelog (Changelog) - Updates Only
*   **Standard**: Clearly state what has been "Added", "Optimized", or "Fixed".
*   **Example**: Added OIDC login support; Optimized image loading speed; Fixed crash issue when exporting PDFs.

## 6. Version
*   **Standard**: The `version` string must strictly follow `x.x.x` semantic versioning. Suffixes or prefixes like `v1.0.0` or `1.0.0-beta` should be mapped carefully, but the `package.yml` field requires strict `x.x.x` format.
*   **Requirement**: Exactly three dot-separated numbers.
*   **Example**: `1.0.0`, `2.4.12`
