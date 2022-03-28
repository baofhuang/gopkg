package confluent_kafka_go

import (
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewProducer 实例化生产者
func NewProducer(host string) (*kafka.Producer, error) {
	// providing at least the `bootstrap.servers` configuration properties.
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	if err != nil {
		return nil, err
	}
	// defer producer.Close()
	// Delivery report handler for produced messages
	// go func() {
	// 	for e := range producer.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()
	return producer, nil
}

// Produce 开始生产
func Produce(producer *kafka.Producer, topic string) {
	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	var con = make(chan string)

ProducerLoop:
	for i := 0; i < 1000; i++ {
		con <- "hello you"
		select {
		case message := <-con:
			if err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(message),
			}, nil); err != nil {
				return
			}
		case <-signals:
			producer.Close() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}
}
