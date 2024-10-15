package queue

import (
	"github.com/hibiken/asynq"

	"github.com/1layar/universe/internal/email_service/app/appconfig"
)

func NewServer(config *appconfig.Config) *asynq.Server {
	// Initialize Asynq client
	redisConnOpt := asynq.RedisClientOpt{
		Addr:     config.RedisAddr,
		Username: config.RedisUsername,
		Password: config.RedisPassword,
	}
	asynqServer := asynq.NewServer(redisConnOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	return asynqServer
}
