# watchctrl：监听 `api` 并自动执行 `gf gen ctrl`

**与官方分工**：业务服务请按 GoFrame 文档使用 **`gf run main.go`**（或 `go run .`）启动；本目录工具**不参与**服务运行，仅在开发时减少重复执行 **`gf gen ctrl`**（与官方 CLI 行为一致，只是自动触发）。

在项目根目录运行本工具后，会监视 `api/**/v1/**/*.go` 的变更（带防抖），稳定后自动执行 `gf gen ctrl`，避免每次改接口都手敲命令。

## 复制到其它项目用（推荐：整目录拷贝）

把本目录 **`watchctrl`** 原样放到目标仓库的相同路径下即可：

```text
<目标项目根>/
  cmd/tools/watchctrl/    ← 拷贝本 README 所在整个文件夹
    main.go
    README.md
```

然后：

1. **在目标项目根目录**（与 `go.mod`、`api/` 同级）执行一次依赖整理（会拉取 `fsnotify`）：

   ```bash
   go mod tidy
   ```

   若 `go mod tidy` 未自动加入 `fsnotify`，可手动：

   ```bash
   go get github.com/fsnotify/fsnotify@v1.10.1
   ```

2. **安装 GoFrame 命令行**（与项目无关，本机一次即可）：

   ```bash
   go install github.com/gogf/gf/cmd/gf/v2@latest
   ```

   确保 `gf` 在 `PATH` 中（`go env GOPATH` 下的 `bin` 已加入环境变量）。

3. **启动监听**（必须在项目根目录执行，因为工具用当前工作目录作为 `gf gen ctrl` 的工作目录）：

   - Windows（PowerShell）：

     ```powershell
     go run .\cmd\tools\watchctrl
     ```

   - macOS / Linux：

     ```bash
     go run ./cmd/tools/watchctrl
     ```

按 `Ctrl+C` 结束。

## 使用前提（不满足则无法「拷过去即用」）

| 条件 | 说明 |
|------|------|
| GoFrame 工程 | 目标项目需按 GoFrame 约定包含 `api/`，且接口定义在 `api/**/v1/**/*.go`（与当前硬编码规则一致）。 |
| `gf gen ctrl` 可用 | 根目录执行 `gf gen ctrl` 能正常生成 controller；数据库等其它 `gf gen` 与本工具无关。 |
| 从**模块根**运行 | 不要在子目录里 `go run`，否则监听路径与 `gf` 工作目录会错。 |

## 行为说明

- **监视目录**：仅递归监视项目根下的 `api/`。
- **触发文件**：路径中包含 `v1` 目录段的 `.go` 文件（避免误触 `api` 下非 v1 文件导致与生成器互相触发）。
- **防抖**：默认约 600ms 内多次保存只触发一次 `gf gen ctrl`。
- **Windows**：生成结束后短暂 sleep，减轻 fsnotify 连发导致的重复执行。

## 可移植性

- **跨平台**：依赖 Go 标准库与 `fsnotify`，Windows / macOS / Linux 均可编译运行。
- **非「零依赖单文件」**：除本目录外，目标项目需有合法 `go.mod` 并能解析 `fsnotify`；**不**替代安装 `gf` CLI。

## 与本仓库 README 的关系

根目录 `README.MD` 已写明：运行按官方 `gf run`；本工具为可选。若你希望「只带一个文件夹给别人」，把 **`cmd/tools/watchctrl` 整个目录** 连同本 `README.md` 一起复制即可。
