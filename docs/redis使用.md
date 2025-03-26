# Redis使用

## 1. 消息队列

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

## 2. 用户缓存

1. 用户信息缓存
2. 设备ack缓存
3. 群组信息缓存
4. 群组成员管理（查找在线设备）

优化策略？

1. 采用json序列化： 高效的json序列化方式
2. 过期时间设置： 为了缓解压力
3. 缓存击穿防护： 使用空值缓存策略

## 3.数据结构

消息队列：Pub/Sub
用户缓存：String + JSON
设备ACK：Hash

## 4. 如何处理缓存击穿、缓存雪崩和缓存穿透情况？

缓存雪崩（缓存击穿）： 大量请求直接打到数据库上
缓存穿透: 数据库和缓存中都没有请求，因此对后面的大量请求，数据库的压力骤增

### 缓存击穿(缓存雪崩)

1. 采用不一致的缓存过期时间
2. 采用空值缓存(未采用)

### 缓存穿透

1. 采用空值缓存
2. 加入参数校验（减少非法请求的影响）
3. 采用布隆过滤器（快速判断数据是否存在）

## 缓存策略

Cache Aside Pattern:

1. 写操作： 先更新数据库，再更新缓存
2. 读操作： 先从缓存中读取数据，如果缓存中没有，再从数据库中读取
