package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"gitlab.com/7chip/little-bird/backend/micro/config/gcp"
)

type Client struct {
	client *pubsub.Client
}

// NewClient creates a new micro pubsub client wrapper
// It will help you to connect to topics and subscriptions, and panic if they do not exist
func NewClient(cfg *gcp.Config) (*Client, error) {
	client, err := pubsub.NewClient(context.Background(), cfg.ProjectID)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

// MustLoadTopic will return the requested pubsub topic, or panic if tht topic does not exist
func (c *Client) MustLoadTopic(t string) *pubsub.Topic {
	topic := c.client.Topic(t)
	ok, err := topic.Exists(context.Background())
	if err != nil {
		panic(fmt.Errorf("pubsub: topic exists call failed: %v", err))
	}
	if !ok {
		panic(fmt.Errorf("pubsub: topic '%s' does not exists", t))
	}
	return topic
}

// MustLoadSubscriber will return the requested subscription, or panic if that subscription does not exist
func (c *Client) MustLoadSubscriber(s string) *pubsub.Subscription {
	sub := c.client.Subscription(s)
	ok, err := sub.Exists(context.Background())
	if err != nil {
		panic(fmt.Errorf("pubsub: subscription exists call failed: %v", err))
	}
	if !ok {
		panic(fmt.Errorf("pubsub: subscription '%s' does not exists", s))
	}
	return sub
}

// Close the pubsub client wrapper, and the underlying pubsub connection
func (c *Client) Close() error {
	return c.client.Close()
}
