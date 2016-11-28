package subscriber

import (
	"log"
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

type Agent struct {
	Verbose      bool
	ProjectID    string
	Subscription string
	Env          string
}

func NewAgent() *Agent {
	agent := &Agent{
		Verbose: true,
		Env:     "development",
	}
	return agent

}

func (agent *Agent) Subscribe() *pubsub.MessageIterator {

	if agent.Env == "development" || agent.Env == "production" {

		ctx := context.Background()
		client, err := pubsub.NewClient(PubSubCtx, agent.ProjectID)

		if err != nil {
			log.Fatal(err)
		}

		subscription = client.Subscription(agent.Subscription)

		it, err := subscription.Pull(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return it
	}
	return nil
}
