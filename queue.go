package godiscord

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type QueueEntry struct {
	WebhookPayload *Embed
	WebhookUrl     string
	RetryCount     int
}

var (
	WebhookQueue   = map[string][]*QueueEntry{}
	QueueMutex     = &sync.RWMutex{}
	WebhookMutexes = map[string]*sync.RWMutex{}
)

//SendToWebhook sents the Embed to a webhook.
//Returns error if embed was invalid or there was an error posting to the webhook.
func (e *Embed) SendToWebhook(webhookUrl string) error {
	QueueMutex.Lock()

	defer QueueMutex.Unlock()

	if _, ok := WebhookQueue[webhookUrl]; !ok {
		WebhookQueue[webhookUrl] = make([]*QueueEntry, 0)
		WebhookMutexes[webhookUrl] = &sync.RWMutex{}
	}

	WebhookQueue[webhookUrl] = append(WebhookQueue[webhookUrl], &QueueEntry{
		WebhookPayload: e,
		WebhookUrl:     webhookUrl,
		RetryCount:     3,
	})

	return nil
}

func processWebhookQueue(webhookUrl string) {
	WebhookMutexes[webhookUrl].Lock()

	defer WebhookMutexes[webhookUrl].Unlock()

	if len(WebhookQueue[webhookUrl]) == 0 {
		return
	}

	QueueMutex.Lock()
	entry := WebhookQueue[webhookUrl][0]
	WebhookQueue[webhookUrl] = WebhookQueue[webhookUrl][1:]
	QueueMutex.Unlock()

	postWebhook(entry.WebhookUrl, entry.WebhookPayload, entry.RetryCount)

	if len(WebhookQueue[webhookUrl]) > 0 {
		time.Sleep(2 * time.Second)
	}
}

func validateWebhookPayload(webhookPayload *Embed) *Embed {
	for _, embed := range webhookPayload.Embeds {
		for _, field := range embed.Fields {
			if field.Value == "" {
				field.Value = "Empty"
			}
			if field.Name == "" {
				field.Name = "Empty"
			}
		}
	}

	return webhookPayload
}

func postWebhook(webhookUrl string, webhookPayload *Embed, retryCount int) {
	webhookPayload = validateWebhookPayload(webhookPayload)

	for i := 0; i < retryCount; i++ {

		payloadJson, err := json.Marshal(*webhookPayload)
		if err != nil {
			continue
		}

		req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(payloadJson))
		if err != nil {
			continue
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}

		defer resp.Body.Close()

		break
	}
}
