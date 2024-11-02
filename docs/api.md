# API 文档
以下 api 的请求 URL 均以 `/v1` 开头，下省略；且所有 api 的请求内容 / 返回内容均为 JSON 格式。
## 用户登录
### 请求 URL
- `/login`
### 请求方式
- POST
### 请求头
- 无需请求头
### 请求权限
- 允许任何人请求
### 请求内容
| 名称 | 是否必需 | 类型 | 说明 |
| :-----------: | :-----------: | :-----------: | :-----------: |
| `username` | 是 | `string` | 登录用户名 |
| `password` | 是 | `string` | 登录密码 |
### 返回内容
- HTTP 状态码 200：
    | 名称 | 类型 | 说明 |
    | :-----------: | :-----------: | :-----------: |
    | `token` | `string` | 鉴权 token |
- HTTP 状态码 4xx / 5xx：
    | 名称 | 类型 | 说明 |
    | :-----------: | :-----------: | :-----------: |
    | `message` | `string` | 错误信息 |

## 用户注册
### 请求 URL
- `/register`
### 请求方式
- POST
### 请求头
- 无需请求头
### 请求权限
- 允许任何人请求
### 请求内容
| 名称 | 是否必需 | 类型 | 说明 |
| :-----------: | :-----------: | :-----------: | :-----------: |
| `username` | 是 | `string` | 用户名 |
| `password` | 是 | `string` | 用户密码 |
| `email` | 是 | `string` | 用户邮箱 |
### 返回内容
| 名称 | 类型 | 说明 |
| :-----------: | :-----------: | :-----------: |
| `message` | `string` | 服务端返回信息 |

## 查询用户信息
查询指定 `id` 的用户信息。
### 请求 URL
- `/user/<id>`
### 请求方式
- GET
- POST
### 请求头
- `Authorization: <token>`，其中 `token` 为登录时返回的 `token`
### 请求权限
- 仅允许管理员或用户本人请求
### 请求内容
无
### 返回内容
- HTTP 状态码 200：
    | 名称 | 类型 | 说明 |
    | :-----------: | :-----------: | :-----------: |
    | `username` | `string` | 用户名 |
    | `email` | `string` | 用户邮箱 |
- HTTP 状态码 4xx / 5xx：
    | 名称 | 类型 | 说明 |
    | :-----------: | :-----------: | :-----------: |
    | `message` | `string` | 错误信息 |

## 更新用户信息
更新指定 `id` 的用户信息。
### 请求 URL
- `/user/<id>/update`
### 请求方式
- POST
### 请求头
- `Authorization: <token>`，其中 `token` 为登录时返回的 `token`。
### 请求权限
- 仅允许管理员或用户本人请求
### 请求内容
| 名称 | 是否必需 | 类型 | 说明 |
| :-----------: | :-----------: | :-----------: | :-----------: |
| `username` | 是 | `string` | 用户名 |
| `password` | 是 | `string` | 用户密码 |
| `email` | 是 | `string` | 用户邮箱 |
### 返回内容
| 名称 | 类型 | 说明 |
| :-----------: | :-----------: | :-----------: |
| `message` | `string` | 服务端返回信息 |

## 删除用户
删除指定 `id` 的用户。
### 请求 URL
- `/user/<id>/delete`
### 请求方式
- POST
### 请求头
- `Authorization: <token>`，其中 `token` 为登录时返回的 `token`。
### 请求权限
- 仅允许管理员请求
### 请求内容
无
### 返回内容
| 名称 | 类型 | 说明 |
| :-----------: | :-----------: | :-----------: |
| `message` | `string` | 服务端返回信息 |