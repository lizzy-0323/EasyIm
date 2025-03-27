# seq使用

## 使用的原因

- IM中需要保证消息的有序性，采用seq能保证消息的严格顺序
- 客户端可以采用seq实现消息同步

## 实现机制

```go
func (*seqRepo) Incr(objectType, objectId int64) (int64, error) {
    tx := db.DB.Begin()
    defer tx.Rollback()

    var seq int64
    // 使用SELECT FOR UPDATE锁住该行
    err := tx.Raw("select seq from seq where object_type = ? and object_id = ? for update", objectType, objectId).
        Row().Scan(&seq)
    
    if errors.Is(err, sql.ErrNoRows) {
        // 如果不存在则插入初始值
        err = tx.Exec("insert into seq (object_type,object_id,seq) values (?,?,?)", objectType, objectId, seq+1).Error
    } else {
        // 存在则递增
        err = tx.Exec("update seq set seq = seq + 1 where object_type = ? and object_id = ?", objectType, objectId).Error
    }

    tx.Commit()
}
```

## 一致性保证

1. 采用事务操作
2. 采用行级锁防止并发问题
3. 采用唯一索引，定位seq种类和用户id
4. 持久化存储： mysql

## 使用场景

- 单聊消息：每个用户都有自己的seq，确保消息顺序
- 群聊消息：每个群有独立的seq，保证群消息顺序
- 消息同步：客户端通过本地最大seq和服务器seq对比来同步消息
