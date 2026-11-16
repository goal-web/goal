# 缓存配置说明

## 概述

本文档说明了Goal Web框架中缓存组件的配置选项和使用方法。

## 配置结构

缓存配置支持三种驱动：
- **memory**: 内存缓存（RAM驱动）
- **file**: 文件缓存
- **redis**: Redis缓存

## 配置选项

### 基本配置

```go
cache.Config{
    Default: "memory",  // 默认缓存驱动
    Stores: map[string]contracts.Fields{
        // 各种缓存驱动的配置
    },
}
```

### 环境变量

| 环境变量 | 默认值 | 说明 |
|----------|--------|------|
| `CACHE_DEFAULT` | `memory` | 默认缓存驱动 |
| `CACHE_PREFIX` | `""` | 缓存键前缀 |
| `CACHE_TTL` | `86400` | 默认缓存TTL（秒） |
| `CACHE_FILE_PATH` | `./storage/cache` | 文件缓存路径 |
| `CACHE_CONNECTION` | `default` | Redis连接名称 |

## 驱动配置

### 1. 内存缓存 (memory)

```go
"memory": {
    "driver": "ram",
    "prefix": env.GetString("cache.prefix"),
    "ttl":    utils.IntOr(env.GetInt("cache.ttl"), 24*int(time.Hour)),
}
```

**配置参数：**
- `driver`: 固定为 "ram"
- `prefix`: 缓存键前缀
- `ttl`: 默认过期时间（秒）

**特点：**
- 高性能，数据存储在内存中
- 应用重启后数据丢失
- 支持统计和监控功能

### 2. 文件缓存 (file)

```go
"file": {
    "driver": "file",
    "path":   utils.StringOr(env.GetString("cache.file.path"), "./storage/cache"),
    "prefix": env.GetString("cache.prefix"),
}
```

**配置参数：**
- `driver`: 固定为 "file"
- `path`: 缓存文件存储路径
- `prefix`: 缓存键前缀

**特点：**
- 持久化存储，应用重启后数据保留
- 适合开发环境和小型应用
- 自动创建缓存目录

### 3. Redis缓存 (redis)

```go
"redis": {
    "driver":     "redis",
    "connection": utils.StringOr(env.GetString("cache.connection"), "default"),
    "prefix":     env.GetString("cache.prefix"),
}
```

**配置参数：**
- `driver`: 固定为 "redis"
- `connection`: Redis连接名称
- `prefix`: 缓存键前缀

**特点：**
- 高性能，支持分布式
- 持久化存储
- 支持复杂数据结构

## 使用示例

### 1. 基本使用

```go
// 获取默认缓存实例
cache := cache.Store()

// 获取指定驱动缓存实例
memoryCache := cache.Store("memory")
fileCache := cache.Store("file")
redisCache := cache.Store("redis")
```

### 2. 环境配置

```bash
# 设置环境变量
export CACHE_DEFAULT=memory
export CACHE_PREFIX=myapp_
export CACHE_TTL=3600
export CACHE_FILE_PATH=/tmp/cache
export CACHE_CONNECTION=redis1
```

### 3. 配置文件

```toml
# config.toml
[cache]
default = "memory"
prefix = "myapp_"
ttl = 3600

[cache.file]
path = "/tmp/cache"

[cache.redis]
connection = "redis1"
```

## 配置最佳实践

### 1. 开发环境
```go
// 使用内存缓存，快速开发
cache.default = "memory"
```

### 2. 测试环境
```go
// 使用文件缓存，便于调试
cache.default = "file"
cache.file.path = "./test_cache"
```

### 3. 生产环境
```go
// 使用Redis缓存，高性能和可靠性
cache.default = "redis"
cache.connection = "production_redis"
```

### 4. 混合使用
```go
// 不同类型的数据使用不同的缓存
userCache := cache.Store("memory")    // 用户会话
configCache := cache.Store("file")    // 配置数据
dataCache := cache.Store("redis")     // 业务数据
```

## 性能考虑

### 1. 内存缓存
- **优点**: 最高性能，零延迟
- **缺点**: 内存占用，重启丢失
- **适用**: 临时数据，会话缓存

### 2. 文件缓存
- **优点**: 持久化，简单部署
- **缺点**: 磁盘I/O，性能中等
- **适用**: 配置缓存，开发环境

### 3. Redis缓存
- **优点**: 高性能，分布式，持久化
- **缺点**: 需要额外服务，网络延迟
- **适用**: 生产环境，大数据量

## 监控和调试

### 1. 统计信息
```go
// 获取缓存统计（仅内存缓存支持）
if memoryCache, ok := cache.Store("memory").(*drivers.Memory); ok {
    stats := memoryCache.GetStats()
    fmt.Printf("命中率: %.2f%%\n", stats["hit_rate"])
    fmt.Printf("内存使用: %.2f MB\n", stats["memory_usage"].(map[string]any)["total_mb"])
}
```

### 2. 键管理
```go
// 获取所有键（仅内存缓存支持）
if memoryCache, ok := cache.Store("memory").(*drivers.Memory); ok {
    keys := memoryCache.GetKeys()
    fmt.Printf("当前有 %d 个缓存键\n", len(keys))
}
```

### 3. 清理过期项
```go
// 手动清理过期项（仅内存缓存支持）
if memoryCache, ok := cache.Store("memory").(*drivers.Memory); ok {
    cleaned := memoryCache.CleanupAllExpired()
    fmt.Printf("清理了 %d 个过期项\n", cleaned)
}
```

## 故障排除

### 1. 常见问题

**问题**: 缓存不生效
**解决**: 检查驱动名称是否正确，确保缓存服务正常运行

**问题**: 内存使用过高
**解决**: 调整TTL设置，定期清理过期项

**问题**: 文件权限错误
**解决**: 确保缓存目录有写入权限

**问题**: Redis连接失败
**解决**: 检查Redis服务状态和连接配置

### 2. 调试技巧

```go
// 启用调试日志
logs.SetLevel(logs.DebugLevel)

// 检查缓存状态
stats := cache.GetStats()
logs.Debug("Cache stats", stats)
```

## 总结

缓存配置提供了灵活的选项来满足不同场景的需求：

1. **开发环境**: 使用内存缓存，快速迭代
2. **测试环境**: 使用文件缓存，便于调试
3. **生产环境**: 使用Redis缓存，高性能和可靠性

通过合理配置缓存驱动和参数，可以显著提升应用程序的性能和用户体验。


