package command

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
)

func CreateCommandProcessor(router *message.Router, sub *amqp.Subscriber) *cqrs.CommandProcessor {
	logger := watermill.NewStdLogger(false, false)
	cqrsMarshaler := cqrs.JSONMarshaler{}

	commandProcessor, err := cqrs.NewCommandProcessorWithConfig(
		router,
		cqrs.CommandProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
				// we are using queue RabbitMQ config, so we need to have topic per command type
				return params.CommandName, nil
			},
			SubscriberConstructor: func(params cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				// we can reuse subscriber, because all commands have separated topics
				return sub, nil
			},
			OnHandle: func(params cqrs.CommandProcessorOnHandleParams) error {
				start := time.Now()

				err := params.Handler.Handle(params.Message.Context(), params.Command)

				logger.Info("Command handled", watermill.LogFields{
					"command_name": params.CommandName,
					"duration":     time.Since(start),
					"err":          err,
				})

				return err
			},
			Marshaler:                cqrsMarshaler,
			Logger:                   logger,
			AckCommandHandlingErrors: true,
		},
	)
	if err != nil {
		panic(err)
	}

	return commandProcessor

}
