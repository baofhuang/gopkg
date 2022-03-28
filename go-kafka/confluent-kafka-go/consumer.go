package confluent_kafka_go

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewConsumer 实例化消费者
func NewConsumer(host string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		// providing at least the `bootstrap.servers` and `group.id` configuration properties
		"bootstrap.servers": host,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	return consumer, nil
}

// Consume 开始消费
func Consume(consumer *kafka.Consumer, topic string) error {
	if err := consumer.Subscribe(topic, nil); err != nil {
		return err
	}
	var count int
	for {
		if message, err := consumer.ReadMessage(-1); err != nil {
			return err
		} else {
			count++
			fmt.Printf("消息数：%v,topic分区:%v 消息内容：%v\n", count, message.TopicPartition, string(message.Value))
		}

	}
	consumer.Close()
	return nil
}
