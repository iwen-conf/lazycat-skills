# 普通业务 Web 应用 AI 设置页模板

适用于“普通业务型 Web 应用接 AI”的默认方案。

不适用于：

- 懒猫算力仓 `AI应用`
- AI 浏览器插件
- 需要 `ai-pod-service/`、`caddy-aipod`、`extension.zip` 的项目

这类普通业务应用默认只做一套稳定、可复用的 AI 连接配置面，不额外引入 AI Pod 包结构。

## 1. 固定必备字段

第一版固定保留这 5 项，不要随意删减：

1. `API BaseURL`
2. `API 协议`
3. `获取模型` 按钮
4. `模型` 下拉框
5. `保存配置` 按钮

## 2. 字段定义

### `API BaseURL`

- 类型：输入框
- 必填：是
- 用途：指定模型服务基础地址
- 示例：
  - `https://api.openai.com/v1`
  - `https://openrouter.ai/api/v1`
  - `<private-base-url>`

最小校验：

- 不能为空
- 必须是 `http://` 或 `https://`
- 保存前去掉首尾空格

### `API 协议`

- 类型：下拉框 / 单选
- 必填：是
- 固定选项：
  - `OpenAI Compatible`
  - `OpenAI Responses`
  - `Anthropic`

不要让用户手填协议名。

### `获取模型` 按钮

- 类型：主操作按钮
- 触发条件：`API BaseURL` 和 `API 协议` 都已填写
- 动作：基于当前 `BaseURL + 协议` 拉取模型列表

最小交互：

- 请求中展示加载状态
- 成功后刷新模型下拉框选项
- 失败后展示可读错误，不要静默失败

### `模型` 下拉框

- 类型：下拉选择
- 必填：是
- 数据来源：点击 `获取模型` 按钮后返回的模型列表

最小交互：

- 未拉取前显示占位文案，例如“请先获取模型”
- 拉取成功后可选
- 保存前必须有有效选项

### `保存配置` 按钮

- 类型：主按钮
- 动作：保存当前 AI 接入配置

最小交互：

- 保存中禁用重复点击
- 保存成功给出确认反馈
- 保存失败给出明确错误信息

## 3. 推荐页面布局

按这个顺序排：

1. `API BaseURL`
2. `API 协议`
3. `获取模型`
4. `模型`
5. `保存配置`

不要把“获取模型”和“保存配置”混成一个按钮。
不要把“模型”做成自由输入框。

## 4. 推荐状态流

### 初始状态

- `BaseURL` 为空或回显已保存值
- 协议使用已保存值或默认值
- 模型下拉框禁用
- `获取模型` 按钮可见
- `保存配置` 按钮可见

### 拉取模型成功

- 模型下拉框填充选项
- 保留上次有效模型时，若仍存在可自动选中
- 若上次模型已不存在，提示重新选择

### 拉取模型失败

- 保留当前表单值
- 模型下拉框清空或保持旧值但标记待确认
- 给出失败原因

### 保存成功

- 持久化当前 `BaseURL / 协议 / 模型`
- 页面提示“保存成功”或等价反馈

## 5. 最小数据结构

```json
{
  "base_url": "https://api.openai.com/v1",
  "protocol": "openai_compatible",
  "model": "gpt-4.1-mini"
}
```

协议枚举建议：

```text
openai_compatible
openai_responses
anthropic
```

## 6. 最小接口建议

如果要在项目里写成后端配置接口，第一版至少准备等价能力：

- `GET /settings/ai`
- `PUT /settings/ai`
- `POST /settings/ai/models`

示意：

### 获取已保存配置

```http
GET /settings/ai
```

### 保存配置

```http
PUT /settings/ai
Content-Type: application/json

{
  "base_url": "https://api.openai.com/v1",
  "protocol": "openai_compatible",
  "model": "gpt-4.1-mini"
}
```

### 拉取模型

```http
POST /settings/ai/models
Content-Type: application/json

{
  "base_url": "https://api.openai.com/v1",
  "protocol": "openai_compatible"
}
```

## 7. 可选扩展项

这些不是第一版硬门槛，只有业务明确需要时再加：

- API Key
- 超时时间
- 组织 / workspace 标识
- 默认 system prompt
- 温度等高级参数

如果要加，也放在“高级设置”区域，不要污染第一屏主配置。

## 8. 质量门槛

- 页面上明确存在 5 个固定要素
- 协议选项固定，不可随意输入
- 模型来自真实拉取结果，而不是写死
- 保存逻辑只保存当前选择，不夹带隐式默认模型
- 错误提示对用户可读
