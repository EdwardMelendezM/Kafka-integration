package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// StreamEvent representa un evento de streaming
type StreamEvent struct {
	Event    string    `json:"event"`
	UserID   string    `json:"user_id"`
	StreamID string    `json:"stream_id"`
	Ts       time.Time `json:"timestamp"`
}

func main() {
	brokers := []string{"localhost:9092"}
	topic := "stream-events"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("error creating producer: %v", err)
	}
	defer producer.Close()

	// Ejemplo: emitir evento de inicio de stream
	evt := StreamEvent{
		Event:    "stream_started",
		UserID:   "user123",
		StreamID: "live_abc",
		Ts:       time.Now().UTC(),
	}
	payload, _ := json.Marshal(evt)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(evt.StreamID),
		Value: sarama.ByteEncoder(payload),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("error sending message: %v", err)
	}
	log.Printf("message sent to partition %d at offset %d", partition, offset)
}
