package validator

import "github.com/1layar/universe/src/shared/dto"

func CreateValidator() *dto.XValidator {
	return dto.NewValidator()
}
