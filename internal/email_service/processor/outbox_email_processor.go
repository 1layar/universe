package processor

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/1layar/universe/internal/email_service/app/appconfig"
	"github.com/1layar/universe/internal/email_service/service"
	"github.com/hibiken/asynq"
)

func RegHandleOutBoxProcess(mux *asynq.ServeMux, emailManager *service.EmailManager, appconfig *appconfig.Config) {
	mux.Handle("email:delivery", asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		log.Info().Msg("start processing email delivery task processed")
		if err := emailManager.ProcessEmailDeliveryTask(ctx, task); err != nil {
			log.Err(err).Msg("could not process email delivery task")
			return fmt.Errorf("could not process email delivery task: %w", err)
		}

		log.Info().Msg("email delivery task processed")
		return nil
	}))
}
