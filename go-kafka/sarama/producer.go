package lsaram

import (
	"log"
	"os"
	"os/signal"
	"sync"
	
	"github.com/Shopify/sarama"
)

// NewProducer 实例化生产者
func NewProducer(host string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()                     // 实例化一个带默认配置的config
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要leader和follow都确认
	config.Producer.Return.Successes = true          // 开启生产消息响应
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer([]string{host}, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// Produce 开始生产
func Produce(producer sarama.AsyncProducer, topic string) {
	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder("mcTest"),
		}
		select {
		case producer.Input() <- message:
			enqueued++
		
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}
	wg.Wait()
	log.Printf("Successfully produced: %d; errors: %d\n", enqueued, producerErrors)
}
