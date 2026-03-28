package main

import (
	"fmt"
	"github.com/nsqio/go-diskqueue"
	"github.com/nsqio/go-nsq"
	"time"
)

func main() {
	fmt.Println(diskqueue.New("", "", 0, 0, 0, 0, time.Minute, nil))
	fmt.Println(nsq.Message{})
	fmt.Println(nsq.Consumer{})
	fmt.Println(nsq.Producer{})
}

/*
源码分析：
【核心设计】https://www.itart.cn/blogs/2025/note/nsq-core-design-nsqd-topic-channel-message-flow.html#%E5%89%8D%E8%A8%80
https://www.itart.cn/blogs/2025/practice/nsq-core-design--nsqd--topic--channel---message-flow.html
【优雅关闭】https://www.itart.cn/blogs/2025/practice/graceful-shutdown-of-nsq-consumers-source-code-deep-dive.html#%E5%89%8D%E8%A8%80
【DiskQueue 】https://www.itart.cn/blogs/2025/practice/nsq-diskqueue-persistent-queue-design-analysis.html#%E2%9C%A8-%E8%AE%BE%E8%AE%A1%E4%BA%AE%E7%82%B9
【message协议】https://www.itart.cn/blogs/2025/practice/tcp-stick-package-terminator-nsq-custom-protocol-s-core-breakthrough-design.html#%E5%89%8D%E8%A8%80
【消息延迟】https://www.itart.cn/blogs/2025/practice/nsq-delayed-messages.html#%E5%89%8D%E8%A8%80
【生产者封装】https://www.itart.cn/blogs/2025/practice/nsqd-producer-auto-failover.html
*/
