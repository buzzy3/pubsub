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
)

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
}

func (agent *Agent) shouldPublish() bool {
	if agent.Env == "production" || agent.Env == "development" {
		return agent.ProjectID != ""
	}
	return false
}
