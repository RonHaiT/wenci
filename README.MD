# wenci

基于 GoFrame v2 的单体仓库模板，已内置本地 MySQL（Docker Compose）、用户注册/登录与账户信息示例接口。

**运行与工程化以 [GoFrame 官方文档](https://goframe.org/docs/cli) 为准**：开发启动使用 `gf run main.go`；代码生成使用 `gf gen ctrl` / `gf gen dao` 等（本仓库 `hack/hack.mk` 中的 `make ctrl`、`make dao` 与之对应）。`cmd/tools/watchctrl` 为可选本地辅助（监听 `api` 后调 `gf gen ctrl`），不属于框架自带流程。

## 本地依赖

- Go：以你本机 `go.mod` 为准
- Docker Desktop（用于本地 MySQL）

## 启动 MySQL（本地）

```bash
docker compose up -d
```

连接信息（例如 Navicat）：

- 主机：`127.0.0.1`
- 端口：`3306`
- 用户名：`root`
- 密码：`root123456`
- 数据库：`wenci`

> 数据库连接配置在 `manifest/config/config.yaml` 的 `database.default.link`。

## JWT 密钥（多机/多环境必读）

签名与校验使用 `internal/consts` 中的 `JwtKey`：

- 优先读取环境变量 **`WENCI_JWT_KEY`**；未设置时使用占位值 `dev-change-me`（仅适合本机随便跑）。
- **同一套环境里**，签发 token 的机器与校验 token 的机器必须使用**相同**的密钥，否则登录后访问鉴权接口会失败。
- **不要把真实密钥写进仓库**；生产环境请用环境变量或密钥管理服务注入。

Windows PowerShell（当前终端有效）：

```powershell
$env:WENCI_JWT_KEY = "你的随机长密钥"
gf run main.go
```

持久写入用户环境变量（新开终端生效）：

```powershell
setx WENCI_JWT_KEY "你的随机长密钥"
```

## 启动服务（本地）

```bash
gf run main.go
```

默认监听：`http://127.0.0.1:8000`

- OpenAPI：`/api.json`
- Swagger：`/swagger`

## 接口示例

路由前缀均为 **`/v1`**。需要鉴权的接口在 `internal/cmd/cmd.go` 中挂了 `middleware.Auth`。

### Swagger 鉴权（推荐）

访问 `http://127.0.0.1:8000/swagger`，点击右上角 **Authorize**，选择 `AuthToken`：

- 推荐填写：`Bearer <登录返回的 token>`
- 兼容旧调用：直接填写 `<token>`（未来可以按需下线）

### 用户注册（无需鉴权）

- 路径：`POST /v1/users/register`
- Body：

```json
{
  "username": "test001",
  "password": "12345678",
  "email": "test001@example.com"
}
```

### 用户登录（无需鉴权）

- 路径：`POST /v1/users/login`
- Body：

```json
{
  "username": "test001",
  "password": "12345678"
}
```

- 响应中的 `token` 用于后续鉴权。

### 账户信息（需要鉴权）

- 路径：`GET /v1/account/info`
- Header（推荐）：`Authorization: Bearer <登录返回的 token>`
- Header（兼容旧调用）：`Authorization: <登录返回的 token>`（未来可以按需下线）

### 单词（需要鉴权）

- 创建：`POST /v1/words`
- 更新：`PUT /v1/words/{id}`
- Header 同上（推荐 Bearer）

## 开发：自动生成 controller

官方方式：在项目根执行 `gf gen ctrl`，或使用 `make ctrl`（见 `hack/hack.mk`，会先检查 `gf` CLI）。

可选：修改 `api/**/v1/*.go` 后启动监听工具自动执行 `gf gen ctrl`（带防抖）。说明见 `cmd/tools/watchctrl/README.md`。

PowerShell（项目根目录）：

```powershell
go run .\cmd\tools\watchctrl
```

macOS / Linux：

```bash
go run ./cmd/tools/watchctrl
```

## 参考

- GoFrame Quick Start：`https://goframe.org/quick`
- GoFrame CLI（`gf run`、`gf gen`、`gf build` 等）：`https://goframe.org/docs/cli`
