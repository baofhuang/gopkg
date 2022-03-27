package lsaram

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// NewConsumer 实例化消费者
func NewConsumer(host string) (sarama.Consumer, error) {
	config := sarama.NewConfig()            // 实例化一个带默认配置的config
	config.Producer.Return.Successes = true // 开启生产消息响应
	config.Producer.Return.Errors = true
	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{host}, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Consume 开始消费
func Consume(consumer sarama.Consumer, topic string) error {
	// 获取sarama中所有topic
	topics, err := consumer.Topics()
	if err != nil {
		panic(err)
	}
	// 获取每个主题的分区id
	for _, topic := range topics {
		pIds, err := consumer.Partitions(topic)
		if err != nil {
			return err
		}
		fmt.Printf("topic %v pids:%v \n", topic, pIds)
	}
	// 获取主题内容
	count := 0
	con, err := consumer.ConsumePartition(topic, 0, -1)
	if err != nil {
		return err
	}
	for {
		message := <-con.Messages()
		count++
		fmt.Printf("消息数：%v,消息内容：%v\n", count, string(message.Value))
	}
}
