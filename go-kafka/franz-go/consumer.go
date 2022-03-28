package franzgo

import (
	"context"
	"fmt"
	
	"github.com/twmb/franz-go/pkg/kgo"
)

// NewConsumer 实例化消费者
func NewConsumer(host string) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(host),
		kgo.ConsumeTopics("no_topic"),
		// kgo.ConsumePartitions(par),
		// kgo.ConsumerGroup("123"), // 设置组消费时需要同时指定消费主体；消费主题可空，通过AddConsumeTopics添加
		// kgo.ConsumeTopics(""),
		// kgo.ConsumeTopics(topic), // consume partitions directly
	}
	kClient, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return kClient, nil
}

// Consume 开始消费
func Consume(consumer *kgo.Client, topic string) error {
	consumer.AddConsumeTopics(topic) // v1.4.0 中添加的新特性;
	// 不能开启正则匹配；且消费者组或者直接消费不能同时为空
	for {
		fetches := consumer.PollFetches(context.TODO())
		if fetches.IsClientClosed() {
			return fmt.Errorf("client closed")
		}
		var count int
		fetches.EachRecord(func(message *kgo.Record) {
			count++
			fmt.Printf("消息数：%v,topic:%v 消息内容：%v\n", count, message.Topic, string(message.Value))
		})
	}
}
