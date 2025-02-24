# grpc鉴权逻辑

## 思路

微服务下，一个整体的应用可能分为多个微服务，当有多个服务端时，都进行鉴权会非常冗余，因此考虑单点登录的方案，只要登录一次，就可以访问所有微服务

采用interceptor进行鉴权，这里采用单一拦截器的方案

同时，grpc允许在metadata中传递自定义数据，go中采用在context中传递metadata的策略

## 流程

### 初始化

1. 通过NewInterceptor函数创建拦截器
2. 每个服务启动时，配置自己的白名单

### 请求处理流程

1. 拦截请求，从ctx中获取metadata
2. 调用handleWithAuth函数进行鉴权

### HandleWithAuth函数

1. 检查是否是内部服务调用（以"Int"结尾的服务名）
2. 检查URL是否在白名单中
3. 如果需要鉴权：
   - 从context获取userId、deviceId和token
   - 调用business-server的Auth服务进行验证
4. 验证通过后继续处理请求

