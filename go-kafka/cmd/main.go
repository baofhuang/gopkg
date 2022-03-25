package main

import (
	"flag"
	"fmt"
	"time"

	// kafka "github.com/baofuhuang/gopkg/kafka/sarama"
	kafka "github.com/baofuhuang/gopkg/kafka/franz-go"
)

var host = flag.String("host", "localhost:9092", "kafka host")
var topic = flag.String("topic", "new_topic", "topic")

func main() {
	flag.Parse()
	// 创建生产者
	producer, err := kafka.NewProducer(*host)
	if err != nil {
		panic(err)
	}
	fmt.Println("生产者实例化完成")
	// 创建消费者
	consumer, err := kafka.NewConsumer(*host, *topic)
	if err != nil {
		panic(err)
	}
	fmt.Println("消费者实例化完成")
	// 生产数据
	go kafka.Produce(producer, *topic)
	// 消费数据
	go kafka.Consume(consumer, *topic)

	time.Sleep(time.Hour)
}
