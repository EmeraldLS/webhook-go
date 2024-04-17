package queue

import (
	"context"
	"log"
	"time"

	"github.com/EmeraldLS/webhook/sender"
	"github.com/EmeraldLS/webhook/types"
)

func ProcessWebHook(ctx context.Context, webhookQueue chan *types.WebhookPayload) {
	for payload := range webhookQueue {
		go func(p *types.WebhookPayload) {
			backOffTime := time.Second
			maxBackOffTime := time.Hour
			retries := 0
			maxRetries := 5

			for {
				err := sender.SendWebHook(p.Data, p.URL, p.WebhookID)
				if err == nil {
					break
				}

				retries++
				if retries >= maxRetries {
					log.Println("Max retries reached. Giving up on webhook:", p.WebhookID)
					break
				}

				time.Sleep(backOffTime)

				backOffTime += 2
				log.Println(backOffTime)
				if backOffTime > maxBackOffTime {
					backOffTime = maxBackOffTime
				}
			}
		}(payload)
	}
}
