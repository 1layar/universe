package rule

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/dto"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
	"github.com/ThreeDotsLabs/watermill/message"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

type UserExistRule struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
	Validator  *dto.XValidator
}

func (h UserExistRule) GetTopic(tag string) string {
	return "controller.ExistCommand." + tag
}

func (h UserExistRule) getBackendConfig(suflix string) requestreply.PubSubBackendConfig {
	topic := h.GetTopic(suflix)
	backendConfig := requestreply.PubSubBackendConfig{
		Publisher: h.Publisher,
		SubscriberConstructor: func(subscriberContext requestreply.PubSubBackendSubscribeParams) (message.Subscriber, error) {
			return h.Subscriber, nil
		},
		GenerateSubscribeTopic: func(subscriberContext requestreply.PubSubBackendSubscribeParams) (string, error) {

			return topic, nil
		},
		GeneratePublishTopic: func(subscriberContext requestreply.PubSubBackendPublishParams) (string, error) {

			return topic, nil
		},
		ModifyNotificationMessage: func(msg *message.Message, params requestreply.PubSubBackendOnCommandProcessedParams) error {
			// to make it deterministic
			msg.UUID = uuid.NewString()

			return params.HandleErr
		},
	}
	backendConfig.AckCommandErrors = true

	return backendConfig
}

func (h *UserExistRule) GetUsernameBackend() *requestreply.PubSubBackend[command.GetUsernameExistsResult] {
	backendConfig := h.getBackendConfig("username")
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetUsernameExistsResult]{},
	)

	if err != nil {
		panic(err)
	}

	return backend
}

func (h *UserExistRule) GetEmailBackend() *requestreply.PubSubBackend[command.GetEmailExistsResult] {
	backendConfig := h.getBackendConfig("email")
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetEmailExistsResult]{},
	)

	if err != nil {
		panic(err)
	}

	return backend
}

func (h *UserExistRule) GetSKUBackend() *requestreply.PubSubBackend[command.GetSKUExistsResult] {
	backendConfig := h.getBackendConfig("SKU")
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetSKUExistsResult]{},
	)

	if err != nil {
		panic(err)
	}

	return backend
}

// Validate implements IRule.
func (r *UserExistRule) Validate(fl validator.FieldLevel) bool {
	option := fl.Param()
	group := ""
	groupVal := ""
	queries := strings.Split(option, ":")
	checkBang := option[0] == '!'

	if len(queries) == 2 {
		option = queries[0]

		group = queries[1]
		rfValue := fl.Parent().FieldByName(group)

		// get rvalue
		switch rfValue.Kind() {
		case reflect.String:
			groupVal = rfValue.String()
		case reflect.Int:
			groupVal = strconv.Itoa(int(rfValue.Int()))
		}
	}

	if checkBang {
		option = option[1:]
	}

	switch option {
	case "username":
		usernameExistCmd := command.GetUsernameExistsCommand{
			Username: fl.Field().String(),
		}

		if group != "" {
			usernameExistCmd.Field = map[string]string{}
			usernameExistCmd.Field[strings.ToLower(group)] = groupVal
		}

		replyCh, cancel, err := requestreply.SendWithReplies(
			context.Background(),
			r.CommandBus,
			r.GetUsernameBackend(),
			usernameExistCmd,
		)
		if err != nil {
			panic(err)
		}

		defer cancel()

		result := <-replyCh

		if result.Error != nil {
			panic(result.Error)
		}

		if checkBang {
			return !result.HandlerResult.Exists
		}

		return result.HandlerResult.Exists
	case "email":
		emailExistCmd := command.GetEmailExistsCommand{
			Email: fl.Field().String(),
		}

		if group != "" {
			emailExistCmd.Field = map[string]string{}
			emailExistCmd.Field[strings.ToLower(group)] = groupVal
		}

		replyCh, cancel, err := requestreply.SendWithReplies(
			context.Background(),
			r.CommandBus,
			r.GetEmailBackend(),
			emailExistCmd,
		)
		if err != nil {
			panic(err)
		}

		defer cancel()

		result := <-replyCh

		if result.Error != nil {
			panic(result.Error)
		}

		if checkBang {
			return !result.HandlerResult.Exists
		}

		return result.HandlerResult.Exists
	case "sku":
		skuExistCmd := command.GetSKUExistsCommand{
			SKU: fl.Field().String(),
		}

		if group != "" {
			skuExistCmd.Field = map[string]string{}
			skuExistCmd.Field[strings.ToLower(group)] = groupVal
		}

		replyCh, cancel, err := requestreply.SendWithReplies(
			context.Background(),
			r.CommandBus,
			r.GetSKUBackend(),
			skuExistCmd,
		)
		if err != nil {
			panic(err)
		}

		defer cancel()

		result := <-replyCh

		if result.Error != nil {
			panic(result.Error)
		}

		if checkBang {
			return !result.HandlerResult.Exists
		}

		return result.HandlerResult.Exists
	default:
		return false
	}
}

func RegisterUserExistRule(r UserExistRule, v *dto.XValidator) IRule {
	v.RegisterValidation("exist", r.Validate)

	return &r
}
