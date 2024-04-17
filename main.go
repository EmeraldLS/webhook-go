package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/EmeraldLS/webhook/queue"
	"github.com/EmeraldLS/webhook/redis_subscriber"
	"github.com/EmeraldLS/webhook/types"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})

	webhookQueue := make(chan *types.WebhookPayload, 100)
	fmt.Println("Application is running and listening for incoming req")

	go queue.ProcessWebHook(ctx, webhookQueue)

	if err := redis_subscriber.Subscribe(ctx, client, webhookQueue); err != nil {
		log.Println("Error:", err)
	}

	<-make(chan struct{})
}
