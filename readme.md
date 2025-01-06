# Go IM Server

## Service

- msg-server

## Todo

- [ ] 整合rpc请求到api-server中
- [ ] 支持TCP协议 
- [ ] kafka消息队列
- [ ] prometheus监控
- [ ] 前端支持
- [ ] id生成算法

## 编译proto命令

protoc --go_out=.. --go-grpc_out=.. ./pkg/protocol/proto/*.proto -I ./pkg/protocol/proto