package transport

import (
	"context"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
	"github.com/ThreeDotsLabs/watermill/message"
)

func GetTopic(topic string) string {
	return "controller." + topic
}

func GetBackendConfig(cmd string, publisher *amqp.Publisher, subscriber *amqp.Subscriber) requestreply.PubSubBackendConfig {
	topic := GetTopic(cmd)
	backendConfig := requestreply.PubSubBackendConfig{
		Publisher:        publisher,
		AckCommandErrors: true,
		SubscriberConstructor: func(subscriberContext requestreply.PubSubBackendSubscribeParams) (message.Subscriber, error) {
			return subscriber, nil
		},
		GenerateSubscribeTopic: func(subscriberContext requestreply.PubSubBackendSubscribeParams) (string, error) {

			return topic, nil
		},
		GeneratePublishTopic: func(subscriberContext requestreply.PubSubBackendPublishParams) (string, error) {

			return topic, nil
		},
	}

	return backendConfig
}

func GetBackend[T any](cmd string, publisher *amqp.Publisher, subscriber *amqp.Subscriber) (*requestreply.PubSubBackend[T], error) {
	backendConfig := GetBackendConfig(cmd, publisher, subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[T]{},
	)

	if err != nil {
		return nil, err
	}

	return backend, nil
}

func GetReqRep[T any, P any](cmd string, publisher *amqp.Publisher, subscriber *amqp.Subscriber, ctx context.Context, bus *cqrs.CommandBus, param P) (<-chan requestreply.Reply[T], func(), error) {
	backend, err := GetBackend[T](cmd, publisher, subscriber)

	if err != nil {
		return nil, nil, err
	}

	replyCh, cancel, err := requestreply.SendWithReplies[T](
		ctx,
		bus,
		backend,
		param,
	)

	if err != nil {
		return nil, nil, err
	}

	return replyCh, cancel, nil
}
