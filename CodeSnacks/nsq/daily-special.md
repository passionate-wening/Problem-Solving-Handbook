# NSQ
一款基于 go 语言开发实现的分布式消息队列组件

- [nsq 客户端源码](https://github.com/nsqio/go-nsq)
- [nsq 服务端源码](https://github.com/nsqio/nsq) 
- [拓展：涉及到 nsq 服务端与磁盘的存储交互操作](https://github.com/nsqio/go-diskqueue)

## 概述
重要组件：
- nsqd：最核心的消息队列后端模块，负责接收、存储和发送消息 
- nsqlookupd：nsq 中的服务发现与注册中心，负责管理 nsqd 节点、topic、channel 之间拓扑映射关系 
- nsqadmin：提供了用于实时监控 nsq 集群的 web 可视化界面

核心概念：
- topic：自定义消息主题
- producer：消息生产者
    - 生产时需要显式指定其从属的 topic
- consumer：消息消费者
    - 消费时（发起订阅时）需要显式指定主题 topic 以及频道 channel
- channel：消费者定义的消息频道。
  - topic 与 channel 是一对多关系
  - 每个 channel 都会拥有一份 topic 下的全量完整数据：每当 topic 有新消息到达时，都会拷贝成多份，逐一发送到每个 channel 当中。
  - 数据分治与负载均衡：channel 下的消息会被随机推送给订阅了该 channel 的 1 名 consumer
    - 即相同 channel 的 consumer 之间自动形成了一种类似消费者组的机制（可类比其他消费队列组件中的 consumer group 消费者组的概念）
    - 拥有独立 channel 的 consumer 才能独吞全量数据

注：nsq 与其他消息队列组件的另一大差异是，其中的 topic 和 channel 采用的都是懒创建的机制，使用方无需显式执行 topic 或者 channel 的创建操作，channel 由首次针对该频道发起 subscribe 订阅操作的 consumer 创建；而 topic 则由首次针对该主题发起 publish 操作的 producer 或者 subscribe 操作的 consumer 创建.

## Reference
[万字解析 go 语言分布式消息队列 nsq ](https://zhuanlan.zhihu.com/p/665893174)
