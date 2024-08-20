package validator

import "github.com/1layar/merasa/backend/src/shared/dto"

func CreateValidator() *dto.XValidator {
	return dto.NewValidator()
}
