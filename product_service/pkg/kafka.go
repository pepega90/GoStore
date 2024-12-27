package pkg

import "github.com/segmentio/kafka-go"

func ConnectKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Balancer: &kafka.LeastBytes{},
	}
}

func ConnectKafkaReader(topics ...string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "product-service",
		GroupTopics: topics,
		Partition:   0,
		MinBytes:    10e3,
		MaxBytes:    10e6,
	})
}
