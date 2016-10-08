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
	// PubSubClient *pubsub.Client
	// PubSubCtx    context.Context
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

type Agent struct {
	Verbose   bool
	ProjectID string
	Env       string
}

func NewAgent() *Agent {
	agent := &Agent{
		Verbose: true,
		Env:     "development",
	}
	return agent
}

// Publish sends a message to PUBSUB
func (agent *Agent) Publish(msg []byte, topic string) {
	if agent.shouldPublish() {

		PubSubCtx := context.Background()
		PubSubClient, _ := pubsub.NewClient(PubSubCtx, agent.ProjectID)

		if os.Getenv("PUBSUB_EMULATOR_HOST") != "" {
			fmt.Printf("USING PUBSUB EMULATOR: %s\n", os.Getenv("PUBSUB_EMULATOR_HOST"))
		}

		t := PubSubClient.Topic(topic)

		msgIDs, err := t.Publish(PubSubCtx, &pubsub.Message{
			Data: []byte(msg),
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

func (agent *Agent) shouldPublish() bool {
	if agent.Env == "production" || agent.Env == "development" {
		return agent.ProjectID != ""
	}
	return false
}