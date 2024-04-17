package types

type Payload struct {
	Id      string
	Payment string
	Event   string
	Date    string
}

type WebhookPayload struct {
	URL       string  `json:"url,omitempty"`
	WebhookID string  `json:"webhook_id,omitempty"`
	Data      Payload `json:"data,omitempty"`
}
