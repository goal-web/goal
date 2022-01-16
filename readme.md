# goal

一个继承了 laravel 思想的 golang web 框架

## 框架特点
goal 通过容器和服务提供者作为框架的核心，以 contracts 为桥梁，为开发者提供丰富的功能和服务，这点与 laravel 是相似的。
* 强大的容器
* 使用服务提供者对应用进行扩展
* 支持依赖注入

## 功能特性

* examples 各种程序的例子（还在完善中...）
* contracts 定义模块接口
* container 容器实现！！！
* config 配置模块
* logs 日志模块
* exceptions 异常处理模块
* redis Redis模块
* cache 缓存模块
  * redis
  * memcache [未完成]
  * file [未完成]
  * memory [未完成]
  * multi 高可用多级缓存 [未完成]
* encryption 加密模块
* hashing 哈希模块
* validation 数据校验模块
* events 事件模块
* filesystem 文件系统模块
  * local 本地文件系统
  * qiniu 七牛文件系统 [未完成]
  * oss 阿里云文件系统 [未完成]
* signal 信号监听，goal 实现了优雅关闭功能
* database 数据库操作模块
  * query builder 查询构造器 [开发中...] 
  * drivers 数据库驱动
    * mysql
    * postgresql
    * sqlite
    * sqlserver [未完成]
* http http相关模块，请求、响应、中间件等
  * routing http 路由服务
  * session 会话服务
* console 命令行模块
  * scheduling 任务调度模块
* auth 用户授权模块 [未完成]
* websocket socket通信模块 [未完成]
  * socket.io socket.io 实现 [未完成]
* sse server-sent-events模块(简称sse) [未完成]
* rpc 远程调用模块 [未完成]
  * jsonrpc [未完成]
  * grpc [未完成]
* serialize 序列化模块 [未完成]
  * json [未完成]
  * xml [未完成]
  * gob [未完成]
  * protobuf [未完成]
* mail 邮件模块 [未完成]
* queue 消息队列模块 [未完成]
  * redis [未完成]
  * kafka [未完成]
  * rocketMQ [未完成]
  * rabbitMQ [未完成]
* view 视图模块 [未完成]
* translation 多语言模块 [未完成]
* eloquent ORM模块，计划 golang 1.18 发布后完成，因为泛型 [未完成]

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

扫码加入QQ群  
![qq_pic_merged_1635476228621.jpg](https://i.loli.net/2021/10/29/dpLvehizJCX7EUN.jpg)
