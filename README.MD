# wenci

基于 GoFrame v2 的单体仓库模板，已内置本地 MySQL（Docker Compose）与 `users/register` 示例接口。

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

## 启动服务（本地）

```bash
gf run main.go
```

默认监听：`http://127.0.0.1:8000`

- OpenAPI：`/api.json`
- Swagger：`/swagger`

## 接口示例

### 用户注册

- 路径：`POST /v1/users/register`
- Body：

```json
{
  "username": "test001",
  "password": "12345678",
  "email": "test001@example.com"
}
```

## 参考

- GoFrame Quick Start：`https://goframe.org/quick`