**Golang PubSub Client for Google's PubSub Service**

This is just a demo project, something we're using internally to manage our PubSub messages.

There are no tests yet, since it's just a testing package.

**How to use**

Use the publisher, subscriber or both. It's up to you.

First include the project:

```
import ("github.com/cucumber-tony/pubsub/publisher")
```

Then set it up:

```
agent := publisher.NewAgent()
agent.ProjectID = *ProjectID
```

Then publish a message:

```
slcD := []string{"apple", "peach", "pear"}
msg, _ := json.Marshal(slcD)
agent.Publish(msg, "my-topic")
```

You should see something like this:

```
2016/10/08 18:57:44 Published a message with a message id: 27
```

We recommend using the PubSub emulator for testing and development purposes. After installing and running this, export the emulator host as so:

```
export PUBSUB_EMULATOR_HOST=localhost:8602
```

And then run again. All should be well.

**Authenticating**

If you're using this in GCE, you probably don't need any auth credentials. However, if you're using outside or DO need them, you can include by setting the following ENV variable:

```
if *Creds != "" {
  os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", *Creds)
}
```

Where *Creds* is the path to the JSON key, generated from your GCE portal.
