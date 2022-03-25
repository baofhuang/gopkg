package franzgo

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

// NewConsumer 实例化消费者
func NewConsumer(host string,topic string) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(host),
		kgo.ConsumeTopics(topic),
	}
	kClient, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return kClient, nil
}

// Consume 开始消费
func Consume(consumer *kgo.Client, topic string) error {
	for {
		fetches := consumer.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return fmt.Errorf("client closed")
		}
		var count int
		fetches.EachRecord(func(message *kgo.Record) {
			count++
			fmt.Printf("消息数：%v,消息内容：%v\n", count, string(message.Value))
		})
	}
}
