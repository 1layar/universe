package queue

import (
	"github.com/hibiken/asynq"

	"github.com/1layar/universe/internal/email_service/app/appconfig"
)

func NewClient(config *appconfig.Config) *asynq.Client {
	client := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     config.RedisAddr,
			Username: config.RedisUsername,
			Password: config.RedisPassword,
		},
	)

	return client
}
