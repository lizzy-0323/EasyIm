# 项目中的Redis使用

## 案例

### 1. 消息队列

使用场景

1. 房间消息推送（PushRoom）
2. 全服推送（PushAll）
3. 优先级消息推送（PushPriority）

```go
// 订阅不同的消息主题
pushRoomPriorityChannel := db.RedisCli.Subscribe(mq.PushRoomPriorityTopic)
pushRoomChannel := db.RedisCli.Subscribe(mq.PushRoomTopic)
pushAllChannel := db.RedisCli.Subscribe(mq.PushAllTopic)
```

### 2. 用户缓存

1. 用户信息缓存
2. 设备ack缓存
3. 群组信息缓存
4. 群组成员管理（查找在线设备）

优化策略？

1. 采用json序列化： 高效的json序列化方式
2. 过期时间设置： 为了缓解压力
3. 缓存击穿防护： 使用空值缓存策略

## 数据结构

消息队列：Pub/Sub
用户缓存：String + JSON
设备ACK：Hash
