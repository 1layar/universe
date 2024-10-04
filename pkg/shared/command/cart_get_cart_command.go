package command

import "github.com/1layar/universe/pkg/shared/dto"

type (
	CartGetCartCommand struct {
		SessionId string
	}

	CartGetCartCommandResult struct {
		Cart *dto.CartResp
	}
)
