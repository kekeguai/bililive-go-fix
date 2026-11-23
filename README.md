# bililive-go 猫耳错误分类修复

基于 bililive-go v0.8.4 的猫耳（Missevan）错误分类与封禁重试行为修复。

## 变更

- API code `500030013` 分类为直播间封禁（`ErrRoomBanned`）。
- API code `500030004` 和 HTTP 404 分类为房间不存在。
- 风控、其他 API 错误、网络异常及 HTTP 5xx 保留可重试错误，不再误报为房间不存在。
- 明确封禁时停止当前录制器的取流重试；监听器持续轮询，房间恢复开播时可重新触发录制。

## 验证

在 Go 1.25 下执行 `go test ./...` 与 `go vet ./...` 均通过。使用项目开发构建目标对应命令 `go run build.go dev` 构建成功。
