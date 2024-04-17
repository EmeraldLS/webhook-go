package redis_subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/EmeraldLS/webhook/types"
	"github.com/go-redis/redis/v8"
)

func Subscribe(ctx context.Context, client *redis.Client, webhookQueue chan *types.WebhookPayload) error {
	pubSub := client.Subscribe(ctx, "payments")

	defer func(PubSub *redis.PubSub) {
		if err := PubSub.Close(); err != nil {
			log.Printf("an error occured closing redis pub/sub -> %v", err)
		}
	}(pubSub)

	var payload types.WebhookPayload

	for {
		msg, err := pubSub.ReceiveMessage(ctx)
		if err != nil {
			return fmt.Errorf("error occured receiving messages from redis publisher -> %v", err)
		}

		err = json.Unmarshal([]byte(msg.Payload), &payload)
		if err != nil {
			log.Println("Error unmarshalling payload:", err)
			continue
		}

		webhookQueue <- &payload
	}
}
