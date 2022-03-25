package main

import (
	"flag"
	"time"

	kafka "github.com/Shopify/sarama"
)

func main() {
	host := flag.String("host", "localhost:9092", "kafka host")
	topic := flag.String("topic", "new_topic", "topic")
	flag.Parse()
	// 创建生产者
	config := kafka.NewConfig()             // 实例化一个带默认配置的config
	config.Producer.Return.Successes = true // 开启生产消息响应
	config.Producer.Return.Errors = true
	producer, err := kafka.NewAsyncProducer([]string{*host}, config)
	if err != nil {
		panic(err)
	}
	// 创建消费者
	consumer, err := kafka.NewConsumer([]string{*host}, config)
	if err != nil {
		panic(err)
	}
	// 消费数据
	go Consume(consumer, *topic)
	// 生产数据
	go Produce(producer, *topic)

	time.Sleep(time.Hour)
}
