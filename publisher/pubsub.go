package publisher

import (
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

var (
	subscription *pubsub.Subscription
	countMu      sync.Mutex
	count        int
	PubSubClient *pubsub.Client
	PubSubCtx    context.Context
)

type PubsubMsg struct {
	Body        string                 `json:"body"`
	ID          string                 `json:"id"`
	OperationID string                 `json:"operation_id"`
	DeviceID    int                    `json:"device_id"`
	Region      string                 `json:"region"`
	Type        string                 `json:"type"`
	Output      string                 `json:"output"`
	Success     bool                   `json:"success"`
	Topic       string                 `json:"topic"`
	StatusTopic string                 `json:"status_topic"`
	Cmd         string                 `json:"cmd"`
	PublicToken string                 `json:"public_token"`
	Meta        map[string]interface{} `json:"meta"`
}

// Publish sends a message to PUBSUB
func Publish(msg []byte, topic string) {
	publish("my-topic", msg)
}

func publish(topic string, data []byte) {

	if shouldPublish() {
		log.Print("Publishing message to")
		if os.Getenv("PUBSUB_EMULATOR_HOST") != "" {
			fmt.Printf("USING PUBSUB EMULATOR: %s\n", os.Getenv("PUBSUB_EMULATOR_HOST"))
		}

		t := PubSubClient.Topic(topic)

		msgIDs, err := t.Publish(PubSubCtx, &pubsub.Message{
			Data: []byte(data),
		})

		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Published a message with a message id: %s\n", msgIDs[0])
		}
		return
	}
	// log.Print("Not publishing message, not allowed")
}

func shouldPublish() bool {

	env := os.Getenv("ENV")
	if env == "production" || env == "development" {
		return os.Getenv("PROJECT_ID") != ""
	}
	return false
}

func Publisher() {

	projectID := os.Getenv("PROJECT_ID")
	if shouldPublish() {
		PubSubCtx = context.Background()
		PubSubClient, _ = pubsub.NewClient(PubSubCtx, projectID)
	}
}
