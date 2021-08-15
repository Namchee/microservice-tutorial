package mq

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Namchee/microservice-tutorial/post/endpoints"
	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/nsqio/go-nsq"
)

type MessageHandler struct {
	Endpoints *endpoints.PostEndpoints
}

func (mh *MessageHandler) HandleMessage(message *nsq.Message) error {
	var msg *entity.Message

	if err := json.Unmarshal(message.Body, &msg); err != nil {
		return err
	}

	user, err := strconv.ParseInt(msg.Content, 10, 64)

	if err != nil {
		return err
	}

	_, err = mh.Endpoints.DeletePostByUser(context.Background(), int(user))

	if err != nil {
		return err
	}

	return nil
}
