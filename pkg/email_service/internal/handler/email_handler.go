package handler

import (
	"github.com/goccy/go-json"
	"go.uber.org/fx"

	"github.com/1layar/universe/pkg/email_service/internal/dto"
	"github.com/1layar/universe/pkg/email_service/internal/service"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/event"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rs/zerolog/log"
)

type EmailHandler struct {
	fx.In
	Manager    *service.EmailManager
	Subscriber *amqp.Subscriber
}

func (e EmailHandler) RegEmailHandler(messages <-chan *message.Message) {
	for msg := range messages {
		id := msg.UUID
		payload := msg.Payload

		// convert payload into map[string]any
		var payloadMap event.UserRegisterEvent

		err := json.Unmarshal(payload, &payloadMap)
		if err != nil {
			msg.Nack()
			continue
		}

		event, err := e.Manager.Compose(msg.Context(), dto.Compose{
			ID:       id,
			Event:    constant.AUTH_REGISTER_EVENT,
			Agent:    "mailtrap_sandbox",
			Email:    payloadMap.Email,
			Template: "Account Verification",
			Payload: map[string]any{
				"Email":       payloadMap.Email,
				"ConfirmLink": payloadMap.ConfirmLink,
				"Name":        payloadMap.Name,
				"ConfirmCode": payloadMap.ConfirmCode,
			},
		})

		if err != nil {
			msg.Nack()
			continue
		}

		info, err := e.Manager.Send(event.Message.ID)

		if err != nil {
			msg.Nack()
			continue
		}

		log.Info().Msgf("Email sent: %s", info.ID)

		msg.Ack()
	}

}
