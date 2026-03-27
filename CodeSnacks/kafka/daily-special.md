# KAFKA

## 实际问题
### Kafka Consumer 假存活
（consumer 看起来还活着，但实际上已经被 broker 踢出组了）

现象：tcp 连接正常，消息积压，消费者不消费
```shell
kafka-consumer-groups.sh --bootstrap-server <broker-list:9092> \
  --group <你的 consumer group.id> \
  --describe
```
- Consumer ID不存在
  - Consumer已被broker的Coordinator踢出消费组，Lag持续增长，是典型的TCP half-open导致的心跳超时间题  
    - Kafka broker 端根据超时（如 session.timeout.ms 或 max.poll.interval.ms）判定 consumer 失效，触发 rebalance，把当前 consumer 从 consumer group 中移除。但 consumer 自己可能不知道。
    - Half-open（半开连接）：一端（broker）已经断开或重启了，但另一端（consumer）的操作系统还不知道。
    - TCP Keepalive：是一种机制，让操作系统定期发送探测包，检测对方是否还活着。如果没开启，consumer 的 TCP 栈就不会主动发现连接已死。
```text
Consumer                    Broker
|                           |
|------ 已断开/重启 -------- |  ← Broker 端已清理连接
|                           |
|  (consumer OS 不知道)      |
|  连接状态 = ESTABLISHED    |  ← 假连接
```
  - 进程无错误输出（FetchMessage不报错）：Consumer被踢出组,broker不会再给其fetch请求返回正常数据，但FetchMessage内部是long-poll（默认最多等几秒到几十秒），底层TCP socket处于half-open状态（OS还没探测到断开）。所以当前 FetchMessage 还在阻塞等待,还没收到broker的错误响应（UNKNOWN_MEMBER_ID、REBALANCE_IN_PROGRESS、NOT_COORDINATOR等）。
    - Kafka consumer 的 poll() 底层会发送 fetch request 到 broker。
    - 现在情况是：要么请求还没发出去（consumer 正在等待上一批数据处理完），要么请求已经发出但响应还没回来。所以 consumer 还没收到 broker 的"已被踢出"的错误响应。
    - 应用调用 read() 或 poll() 时，实际上会阻塞在内核的 socket 读取操作上，由于 TCP 是 half-open 状态，操作系统没有收到任何 FIN 或 RST 包，所以 OS 认为连接还活着。因此，read() 会一直阻塞等待，不会返回错误，也不会返回数据。
    - 重启恢复：等下次心跳或fetch请求真正发出后，才会收到broker的错误响应(或超时)：当 consumer 真正发起新的请求（心跳或 fetch）时，会发生两件事之一：1）Broker 响应错误：比如 ILLEGAL_GENERATION、REBALANCE_IN_PROGRESS、UNKNOWN_MEMBER_ID ； 2）请求超时：如果 TCP 层面也断了，会收到 socket 超时或 write error
```text
时间线：
───────────────────────────────────────────────────────────►

T1: Broker 踢掉 consumer（rebalance 触发）
    │
    ├── Consumer 的 TCP 状态 = ESTABLISHED（假连接）
    │
    ├── Consumer 的 poll() 阻塞在 read() 上
    │   └── OS 不知道连接已断，不返回错误
    │
T2: Consumer 发起心跳/fetch 请求
    │
    └── 终于收到错误响应 / 超时
        └── 此时才暴露问题

```