package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendWebHook(data interface{}, url, webhookID string) error {
	buffer, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error occured marshaling data -> %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buffer))
	if err != nil {
		return fmt.Errorf("error occured creating request -> %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error occured sending request -> %v", err)
	}

	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Printf("error occured closing response body -> %v\n", err)
		}
	}(req.Body)

	status := "failed"

	if res.StatusCode == http.StatusOK {
		status = "delivered"
	}

	log.Println("Status = ", status)

	if status == "failed" {
		return errors.New(status)
	}

	return nil
}
