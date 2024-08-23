package mux

import (
	"github.com/hibiken/asynq"
)

func Create(authM *asynq.Server) *asynq.ServeMux {
	mux := asynq.NewServeMux()

	return mux
}
