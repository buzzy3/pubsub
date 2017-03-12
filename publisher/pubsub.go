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
	Client    *pubsub.Client
}

func NewAgent(projectID string) (*Agent, error) {
	var agent Agent
	agent.Verbose = true
	agent.Env = "development"
	agent.ProjectID = projectID

	ctx := context.Background()
	PubSubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Print("ERROR: %v", err)
		return &agent, err
	}

	agent.Client = PubSubClient
	return &agent, nil
}

// Publish sends a message to PUBSUB
func (agent *Agent) Publish(msg []byte, topic string) {
	if agent.shouldPublish() {

		ctx := context.Background()

		if os.Getenv("PUBSUB_EMULATOR_HOST") != "" {
			fmt.Printf("USING PUBSUB EMULATOR: %s\n", os.Getenv("PUBSUB_EMULATOR_HOST"))
		}

		t := agent.Client.Topic(topic)

		res := t.Publish(ctx, &pubsub.Message{Data: []byte(msg)})
		log.Print(res)

		// if err != nil {
		// 	log.Println("COULD NOT PUBLISH MESSAGE", err)
		// } else {
		// 	// log.Printf("Published a message with a message id: %s\n", msgIDs[0])
		// }
		return
	}
}

func (agent *Agent) shouldPublish() bool {
	return agent.ProjectID != ""
}
