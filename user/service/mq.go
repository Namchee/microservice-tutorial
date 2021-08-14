package service

import (
	"context"
	"encoding/json"

	"github.com/Namchee/microservice-tutorial/user/entity"
	nsq "github.com/nsqio/go-nsq"
)

const (
	topic = "post"
)

type nsqPublisher struct {
	client *nsq.Producer
}

func NewNSQPublisher(client *nsq.Producer) PublisherService {
	return &nsqPublisher{client}
}

func (pub *nsqPublisher) Publish(ctx context.Context, msg *entity.Message) error {
	str, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	err = pub.client.Publish(topic, str)

	if err != nil {
		return err
	}

	return nil
}
