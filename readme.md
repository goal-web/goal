# goal

一个继承了 laravel 思想的 golang web 框架

## 框架特点
goal 通过容器和服务提供者作为框架的核心，以 contracts 为桥梁，为开发者提供丰富的功能和服务，这点与 laravel 是相似的。
* [x] 强大的容器
* [x] 使用服务提供者对应用进行扩展
* [x] 支持依赖注入

## 功能特性

* [x] examples 各种程序的例子（还在完善中...）
* [x] contracts 定义模块接口
* [x] container 容器实现！！！
* [x] pipeline 简单但是很强大的洋葱模型的管道
* [x] supports 支持库
  * [x] logs 日志模块
  * [x] collection 集合库
* [x] config 配置模块
* [x] exceptions 异常处理模块
* [x] redis Redis模块
* [x] cache 缓存模块
  * [x] redis
  * [ ] memcached
  * [ ] file
  * [ ] database 数据库驱动
  * [ ] memory
  * [ ] multi 高可用多级缓存
* [x] encryption 加密模块
* [x] hashing 哈希模块
* [x] validation 数据校验模块
* [x] events 事件模块
* [x] filesystem 文件系统模块
  * [x] local 本地文件系统
  * [ ] qiniu 七牛文件系统
  * [ ] oss 阿里云文件系统
* [x] signal 信号监听，goal 实现了优雅关闭功能
* [x] database 数据库操作模块
  * [x] query builder 查询构造器
  * [ ] seeders 数据填充
  * [ ] migration 数据迁移
  * [x] drivers 数据库驱动
    * [x] mysql
    * [x] postgresql
    * [x] sqlite
    * [ ] sqlserver
* [ ] eloquent ORM模块，计划 golang 1.18 发布后完成，因为泛型
* [x] http http相关模块，请求、响应、中间件等
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
* [ ] rate limiter 限流器
* [x] auth 用户认证模块
* [x] serialize 序列化模块
  * [x] json
  * [x] xml
  * [x] gob
  * [ ] protobuf
* [x] queue 消息队列模块
  * [ ] redis
  * [x] kafka
  * [ ] rocketMQ
  * [ ] rabbitMQ
* [ ] gates 用户授权模块
* [ ] websocket socket通信模块
  * [ ] socket.io socket.io 实现
* [ ] sse server-sent-events模块(简称sse)
* [ ] rpc 远程调用模块
  * [ ] jsonrpc
  * [ ] grpc
* [ ] mail 邮件模块
* [ ] view 视图模块
* [ ] translation 多语言模块

## 参与项目

你可以通过以下方式参与到项目中来

* [x] 完善已有模块（优化或者改bug）
* [x] 完善或者修复测试用例
* [x] 开发新的模块（比如标记为未完成的模块）
* [x] 添加或者修改完善注释（用英语）
* [x] 修改错别字或者不当用词（文档和代码都可以，比如变量命名）
* [x] 帮助开发独立文档（readme是临时的，后面需要独立的文档项目）
* [x] 开发扩展包（goal 的扩展相当容易，后面我会写教程，现阶段进群聊）
* [x] 使用 goal 实现各种例子（放examples文件夹或者新建仓库在这里引用）
* [x] 更多方式进群聊吧

## 交流

扫码加入QQ群  
![qq_pic_merged_1635476228621.jpg](https://i.loli.net/2021/10/29/dpLvehizJCX7EUN.jpg)
