package subscriber

import (
	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"log"
	"sync"
)

var (
	subscription *pubsub.Subscription
	countMu      sync.Mutex
	count        int
	PubSubClient *pubsub.Client
	PubSubCtx    context.Context
)

type Agent struct {
	Verbose      bool
	ProjectID    string
	Subscription string
	Env          string
	Client       *pubsub.Client
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

func (agent *Agent) Subscribe() *pubsub.MessageIterator {

	if agent.Env == "development" || agent.Env == "production" {

		ctx := context.Background()

		subscription = agent.Client.Subscription(agent.Subscription)

		it, err := subscription.Pull(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return it
	}
	return nil
}
