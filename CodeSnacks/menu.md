# MENU

## To Do List
- 消息队列
- 幂等性
- Kafka
- LRU
- nuclei
- HTTP/HTTPS
- WEB鉴权
- TCP/IP
- LangChain/LangGraph
- ETCD
- ...

## In Progress
- NSQ
  - 概念
    - [x] nsq是什么——轻量级分布式消息队列组件
    - [x] nsq的生产者和消费者的作用，消息完整生命周期
    - [x] nsq会重复发送消息吗——会
    - [x] nsq的ack机制
    - [x] nsq的存储机制
    - [x] nsq消息有时序性吗——有，顺序读写append-only
    - [x] nsq退避算法
    - [ ] nsq的客户端和服务端都控制什么，有哪些重要参数
  - 大消息应用层分片协议
    - [ ] 如何通过制定应用层协议实现大块消息的拆分和重组
    - [ ] 了解nsq自身的收发协议
    - [ ] 区分它自身的协议和应用层协议，最简化需要设计几个要点
    - [ ] 应用层失败处理
    - [ ] 缓存清理应该如何解决
    - [ ] ...
- Kafka
  - [ ] kafka是什么
  - [ ] 为什么会偶发kafka不消费
  - [ ] kafka消息对齐的作用，如何做到的
  - [ ] kafka消息有时序性吗
  - [ ] ...

## Done