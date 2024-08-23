package watermill

import (
	"github.com/1layar/universe/src/api_gateway/internal/app/appconfig"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

func CreateSub(config *appconfig.Config) *amqp.Subscriber {
	amqpConfig := amqp.NewDurableQueueConfig(config.AmqpUrl)
	subscriber, err := amqp.NewSubscriber(
		// This config is based on this example: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
		// It works as a simple queue.
		//
		// If you want to implement a Pub/Sub style service instead, check
		// https://watermill.io/pubsubs/amqp/#amqp-consumer-groups
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	return subscriber
}

func CreatePub(config *appconfig.Config) *amqp.Publisher {
	amqpConfig := amqp.NewDurableQueueConfig(config.AmqpUrl)
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	return publisher
}
