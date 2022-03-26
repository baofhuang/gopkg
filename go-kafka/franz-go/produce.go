package franzgo

import (
	"context"
	"os"
	"os/signal"
	
	"github.com/twmb/franz-go/pkg/kgo"
)

// NewProducer 实例化生产者
func NewProducer(host string) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(host),
	}
	kClient, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return kClient, nil
}

// Produce 开始生产
func Produce(producer *kgo.Client, topic string) {
	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	var con = make(chan kgo.Record)

ProducerLoop:
	for i := 0; i < 1000; i++ {
		con <- kgo.Record{
			Topic: topic,
			Value: []byte("mcTest"),
		}
		select {
		case record := <-con:
			producer.Produce(context.Background(), &record, func(record *kgo.Record, err error) {})
		case <-signals:
			producer.Close() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}
}
