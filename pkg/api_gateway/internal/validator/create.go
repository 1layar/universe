package validator

import "github.com/1layar/universe/pkg/shared/dto"

func CreateValidator() *dto.XValidator {
	return dto.NewValidator()
}
