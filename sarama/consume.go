package main

import (
	"fmt"

	kafka "github.com/Shopify/sarama"
)

// Consume 开始消费
func Consume(consumer kafka.Consumer,topic string) error {
	// 获取kafka中所有topic
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
