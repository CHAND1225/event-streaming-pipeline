package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a new Kafka consumer
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "events",
		GroupID: "consumer-group-1",
	})

	defer reader.Close()

	log.Println("Consumer started... waiting for messages")

	for {
		// Read messages continuously
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		// Print the consumed message
		log.Printf("Consumed message | key=%s | value=%s\n", string(msg.Key), string(msg.Value))
	}
}
