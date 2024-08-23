package handler

import (
	"context"

	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/dto"
	"github.com/1layar/universe/src/shared/event"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
)

func RegAsignAccess(c AuthHandler) {
	messages, err := c.Subscriber.Subscribe(context.Background(), constant.ADD_PRODUCT_EVENT)
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range messages {
			id := msg.UUID
			payload := msg.Payload

			var payloadMap event.ProductAddEvent

			err := json.Unmarshal(payload, &payloadMap)

			if err != nil {
				msg.Nack()
				continue
			}

			err = c.AuthorizeService.Grant(dto.PermissionDto{
				UserId:      payloadMap.UserID,
				Role:        "user",
				Subject:     "product",
				SubjectID:   payloadMap.ID,
				Permissions: []string{"writer", "reader", "owner"},
			})

			if err != nil {
				msg.Nack()
				continue
			}

			log.Info().Str("id", id).Interface("payload", payloadMap).Msg("Received message")
			msg.Ack()
		}
	}()
}
