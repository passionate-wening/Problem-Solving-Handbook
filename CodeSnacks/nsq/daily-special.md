# NSQ
一款基于 go 语言开发实现的轻量级分布式消息队列组件。
轻量、低延迟、快速任务分发，适合内部通信。
不保证强一致，会有丢消息或消息重复的可能，适合非核心链路。

- [nsq 客户端源码：https://github.com/nsqio/go-nsq](https://github.com/nsqio/go-nsq)
- [nsq 服务端源码：https://github.com/nsqio/nsq](https://github.com/nsqio/nsq)
- [拓展：涉及到 nsq 服务端与磁盘的存储交互操作：https://github.com/nsqio/go-diskqueue](https://github.com/nsqio/go-diskqueue)

## 概述
重要组件：
- **nsqd**：最核心的消息队列后端模块，负责接收、存储和发送消息
- **nsqlookupd**：nsq 中的服务发现与注册中心，负责管理 nsqd 节点、topic、channel 之间拓扑映射关系
- nsqadmin：提供了用于实时监控 nsq 集群的 web 可视化界面

核心概念：
- producer：消息生产者
  - 生产时需要显式指定其从属的 topic
- consumer：消息消费者
  - 消费时（发起订阅时）需要显式指定主题 topic 以及频道 channel
- **topic**：自定义消息主题，用于消息分类。
- **channel**：消费者定义的消息频道。消费队列。
  - topic 与 channel 是一对多关系
  - 不同 channel 广播消费
    - 每个 channel 都会拥有一份 topic 下的全量完整数据：每当 topic 有新消息到达时，都会拷贝成多份，逐一发送到每个 channel 当中。
  - 同一 channel 内负载均衡
    - 同一 channel 下的消息会被随机推送给订阅了该 channel 的 1 名 consumer
    - 即相同 channel 的 consumer 之间自动形成了一种类似消费者组的机制（可类比其他消费队列组件中的 consumer group 消费者组的概念）
    - 拥有独立 channel 的 consumer 才能独吞全量数据

注： topic 和 channel 采用的都是懒创建的机制 ———— nsq 与其他消息队列组件的一大差异
- 无需显式执行 topic 或者 channel 的创建操作
- 首次推送消息的 producer 可自动创建 topic
- 首次发起订阅的 consumer 既可自动创建 topic ，也能自动创建 channel
- NSQ 是负载均衡模型，默认会将消息随机分发给 Channel 下的所有消费者。（Kafka 有 Partition 概念，可以通过 Key 保证同一消息的分片进入同一 Partition，从而被同一消费者组中的特定消费者处理）
- NSQ 是基于 TCP 长连接的，风险：
  - 大包阻塞连接： 发送/接收 5MB 的分片时，TCP 连接会被占用较长时间。 如果并发高，可能导致 nsqd 的文件描述符（FD）耗尽或连接池满。 建议：分片大小不要太大。虽然 max_msg_size 可以调大，但建议分片控制在 256KB ~ 512KB。这样网络传输快，减少连接占用时间。
  - RDY 背压（Backpressure）： NSQ 消费者通过 RDY 计数告诉服务端“我能处理多少消息”。 重组大消息非常消耗内存和 CPU。如果正在重组一个 50MB 消息，建议暂时将 RDY 设为 0 或降低，防止 nsqd 继续推送新消息导致内存爆炸。 在 Go 的 go-nsq 客户端中，可以动态调整 MaxInFlight。
  - 压缩传输： NSQ 支持 deflate 和 snappy 压缩。 在 Producer 和 Consumer 连接时协商开启压缩。 收益：50MB 的文本数据压缩后可能只有 5MB，网络传输时间减少 90%，极大降低阻塞风险。

## Reference
[万字解析 go 语言分布式消息队列 nsq ：https://zhuanlan.zhihu.com/p/665893174](https://zhuanlan.zhihu.com/p/665893174)


## NSQ 大消息应用层协议

- ~~协议标识 Magic ：固定 0x4C4D（即 "LM" = Large Message），用于标识这是大消息协议的数据包，帮助接收端区分该消息是否为分片消息。（2B=2*(8bit)=65535种组和）~~
- **协议标识Magic：以 0x4C 标识第一个消息，以 0x4D 标识最后一个消息（1B）**
- ~~协议版本 Version ：协议长期演进的保障，允许在不破坏现有系统的情况下添加新功能。（1B=8bit=256种组和）~~
- ~~MessageID ：消息唯一标识 (UUID)（16B）~~
- **MessageID：用整个消息的CRC64标识（8B）**
- ~~FragmentCount ：分片总数 （2B）~~
- ~~FragmentIndex ：当前分片索引 (0-based) （2B）~~
- **Index：当前分片索引只需要1B**
- Payload ：分片载荷 
- ~~CRC64 ：载荷 CRC64 校验和 （8B）~~

### 设计解释
1. 协议设计尽量简洁，只固定10B的头部数据即可；
2. CRC64即作为验证，也作为消息唯一索引 MessageID；
   - 1B能表示 256个消息，乘以50M，大概12G，够了，不用2B
3. 最后一个分片的index即可算出总分片数，不需要外围size字段：size=index+1（消息接收过程中验证，只有得到总数并且收齐，才重组并执行接收回调）；
4. 应用层做一层，不需要对分片验证CRC：
   - tcp有自己的CRC验证，NSQ应该也有验证，NSQ能保证消息完整性，否则作为一个MQ完全不合格，所以也不需要每个消息自己在应用层加crc验证；
   - 数据链路层(L2)的强力守护:绝大多数以太网(Ethernet)和Wi-Fi在帧尾部都带有一个CRC32校验。CRC32的碰撞概 率是1/232(大约40亿分之一)，且对连续的位错误极其敏感；
   - 网卡有CRC32验证，本地连接localhost不走网卡，就是程序间内存拷贝，有内存bit出错概率，但在这里我们不考虑，因为发生这种情况，程序都不正常了。所以不需要每条消息自己校验；
5. 设计兼容乱序消息；
   - tcp不保证不同消息，会顺序，也不保证不同消息每条都能收到。tcp只保证，先入缓存消息会先到，按时缓存先后顺序。
   - nsq能保证顺序发送的话一定能保证顺序接收，除非消息超时重发。设计兼容乱序是保险起见；
   - nsq不能保证消息不会被重复消费，所以要注意重复消息
     - map接收保证唯一，重组时转array进行sort格式化排序
     - （浪费点性能，其实还好，没有搞出N^2就行。大部分线上业务，不需要极致优化）
     - golang作者似乎很讨厌rb-tree。
6. 丢弃/重发
   - 消息组装没成功，直接删缓存，不用重试了，在重发也是错的（但是如果程序没问题，不可能出现 ）
     - 即 MessageID(CRC) 返回错误，最好就丢弃删除数据，因为重新消费还是错的。
   - 业务逻辑没成功，删除缓存，让NSQ重发
     - 即消息处理错误（即接收回调函数业务处理错误），可以重发。
       - 关于重发，nsq封装了单条消息的接收和发送，无法批量处理，外围存nsq的消息列表会使设计复杂化。
       - 设计：保留分片缓存，单个分片消息还是正常完成处理，只针对最后一个分片消息，如果业务逻辑失败，就重发最后一个分片（即仅重新处理业务逻辑）
       - 只有消息成功或nsq超过最大重试次数，才删分片缓存



## 相关Go标准库
- encoding/binary
  - 将整数编码为大端序字节序列
  - 网络协议序列化的标准写法
    - 大端序 (Big-Endian) ：高位字节在前；小端序 (Little-Endian)：低位字节在前
    - 大端序是网络协议标准中常用字节序，包括TCP/IP、HTTP/2等协议都使用大端序。大端序和十六进制表示一致（0x4C4D → 4C 4D）
  - binary.BigEndian.PutUint16、binary.BigEndian.PutUint32 
- 经典"向上取整除法"技巧: fragmentCount := (len(message) + fragmentSize - 1) / fragmentSize 
- sort
  - 自定义排序
  - sort.SliceIsSorted
- make
  - 数组make可以传三个参数，size为0，但是预分配内存，append不会触发更新内存，重新复制。append单个元素时间复杂度1