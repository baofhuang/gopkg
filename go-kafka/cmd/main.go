package main

import (
	"flag"
	"fmt"
	"time"

	// kafka "github.com/baofuhuang/gopkg/kafka/sarama"
	// kafka "github.com/baofuhuang/gopkg/kafka/franz-go"
	kafka "github.com/baofuhuang/gopkg/kafka/confluent-kafka-go"
)

var host = flag.String("host", "9.135.155.71:9092", "kafka host")
var topic = flag.String("topic", "what_topic", "topic")

func main() {
	flag.Parse()
	// 创建生产者
	producer, err := kafka.NewProducer(*host)
	if err != nil {
		panic(err)
	}
	fmt.Println("生产者实例化完成")
	// 创建消费者
	consumer, err := kafka.NewConsumer(*host)
	if err != nil {
		panic(err)
	}
	fmt.Println("消费者实例化完成")
	// 消费数据
	go kafka.Consume(consumer, *topic)
	// 生产数据
	go kafka.Produce(producer, *topic)

	time.Sleep(time.Hour)
}
