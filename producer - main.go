package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "events",
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	for i := 1; i <= 10; i++ {
		event := Event{
			ID:        i,
			Message:   fmt.Sprintf("Event number %d", i),
			Timestamp: time.Now(),
		}

		data, err := json.Marshal(event)
		if err != nil {
			log.Println("marshal error:", err)
			continue
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: data,
		})
		if err != nil {
			log.Println("write error:", err)
		} else {
			log.Println("Produced:", string(data))
		}

		time.Sleep(1 * time.Second)
	}
}

