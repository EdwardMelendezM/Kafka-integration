package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// StreamEvent repr. evento de streaming (mismo struct)
type StreamEvent struct {
	Event    string    `json:"event"`
	UserID   string    `json:"user_id"`
	StreamID string    `json:"stream_id"`
	Ts       time.Time `json:"timestamp"`
}

func main() {
	brokers := []string{"localhost:9092"}
	topic := "stream-events"
	group := "analytics-service"

	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("error creating consumer group: %v", err)
	}
	defer client.Close()

	handler := consumerGroupHandler{}
	ctx := context.Background()

	for {
		if err := client.Consume(ctx, []string{topic}, handler); err != nil {
			log.Printf("error from consumer: %v", err)
		}
	}
}

type consumerGroupHandler struct{}

func (c consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var evt StreamEvent
		if err := json.Unmarshal(msg.Value, &evt); err != nil {
			log.Printf("invalid message: %v", err)
			continue
		}
		// Procesar evento (p.ej. almacenar en BD, disparar notificación)
		log.Printf("[Consumed] %s by user %s at %s", evt.Event, evt.UserID, evt.Ts)
		sess.MarkMessage(msg, "")
	}
	return nil
}
