# goal
一个继承了 laravel 思想的 golang web 框架

## 框架特点
goal 通过容器和服务提供者作为框架的核心，以 contracts 为桥梁，为开发者提供丰富的功能和服务，这点与 laravel 是相似的。
* 强大的容器
* 服务提供者
* 契约精神

## 链接
* [goal 仓库](https://github.com/goal-web/goal)
* [快速开始](https://github.com/goal-web/goal/wiki/%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B---quick-start)

## 功能特性

* [x] examples 各种程序的例子（还在完善中...）
* [x] contracts 定义模块接口
* [x] container 容器实现！！！
* [x] pipeline 简单但是很强大的洋葱模型的管道
* [x] supports 支持库
  * [x] logs 日志模块
  * [x] collection 集合库
  * [x] utils 工具库，封装了包含字符串处理、默认参数处理、类型转换、反射等工具函数 
* [x] application 应用
  * [x] exceptions 异常处理模块
  * [x] signal 信号监听，goal 实现了优雅关闭功能
* [x] config 配置模块
* [x] redis Redis模块
* [x] cache 缓存模块
  * [x] redis
  * [x] memory 将数据存储在内存中，不支持持久化和分布式
  * [ ] memcached
  * [ ] file
  * [ ] database 数据库驱动
  * [ ] multi 高可用多级缓存
* [x] encryption 加密模块
* [x] hashing 哈希模块
* [x] validation 数据校验模块
* [x] mail 邮件模块
* [x] events 事件模块
* [x] filesystem 文件系统模块
  * [x] local 本地文件系统
  * [x] qiniu 七牛文件系统
  * [ ] oss 阿里云文件系统
* [x] database 数据库操作模块
  * [x] query builder 查询构造器
  * [ ] seeders 数据填充
  * [x] migration 数据迁移
  * [x] drivers 数据库驱动
    * [x] mysql
    * [x] postgresql
    * [x] sqlite
    * [x] clickhouse
    * [ ] sqlserver
    * [ ] oracle
* [ ] eloquent ORM模块，计划 golang 1.18 发布后完成，因为泛型
* [x] http http相关模块，请求、响应、中间件等
  * [x] sse server-sent-events模块(简称sse)
  * [x] routing http 路由服务
  * [x] session 会话服务
    * [x] cookie 将会话信息存储到加密的 cookie 中
    * [x] redis
    * [ ] file
    * [ ] database
    * [ ] memcached
* [x] console 命令行模块
  * [x] commands 自定义命令模块
  * [x] scheduling 任务调度模块
* [x] auth 用户认证模块
  * [x] gate 用户授权模块
* [x] serialize 序列化模块
  * [x] json
  * [x] xml
  * [x] gob
  * [ ] protobuf
* [x] queue 消息队列模块
  * [ ] redis
  * [x] kafka
  * [x] nsq
  * [ ] rocketMQ
  * [ ] rabbitMQ
* [x] rate limiter 限流器
* [x] bloom-filter 布隆过滤器
  * [x] file 持久化到文件
  * [x] redis 通过 redis bit 实现的过滤器，支持分布式
* [x] websocket socket通信模块
  * [ ] socket.io socket.io 实现
* [x] micro 远程调用模块（集成 go-micro）
  * [x] grpc
  * [x] 服务发现
  * [x] 负载均衡
  * [x] 自定义 go-micro
  * [x] [微服务demo](https://github.com/goal-web/microdemo)
* [ ] 第三方sdk
  * [x] [支付宝sdk](https://github.com/qbhy/goal-alipay)
  * [x] [微信sdk](https://github.com/qbhy/goal-wechat)
  * [x] [阿里云sdk](https://github.com/qbhy/goal-aliyun)
  * [ ] 极光推送 sdk
  * [ ] 字节跳动 sdk
  * [ ] QQ sdk

## 参与项目

你可以通过以下方式参与到项目中来

* 完善已有模块（优化或者改bug）
* 完善或者修复测试用例
* 开发新的模块（比如标记为未完成的模块）
* 添加或者修改完善注释（用英语）
* 修改错别字或者不当用词（文档和代码都可以，比如变量命名）
* 帮助开发独立文档（readme是临时的，后面需要独立的文档项目）
* 开发扩展包（goal 的扩展相当容易，后面我会写教程，现阶段进群聊）
* 使用 goal 实现各种例子（放examples文件夹或者新建仓库在这里引用）
* 更多方式进群聊吧

## 交流
微信群  
<img style="max-width: 200px" src="https://s2.loli.net/2022/03/03/n9jESh6r4e3WQsL.png"/>  
QQ群    
<img style="max-width: 200px" src="https://i.loli.net/2021/10/29/dpLvehizJCX7EUN.jpg"/>
