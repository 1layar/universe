package dto

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/iancoleman/strcase"
)

type (
	XValidator struct {
		validator *validator.Validate
		uni       *ut.UniversalTranslator
	}

	ErrorResponse struct {
		Error       bool   `json:"error"`
		FailedField string `json:"failed_field"`
		Tag         string `json:"tag"`
		Value       any    `json:"value"`
		Message     string `json:"message"`
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	GlobalHandlerResp[T any] struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}
)

func NewValidator() *XValidator {
	en := en.New()
	validate := validator.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	return &XValidator{
		validator: validator.New(),
		uni:       uni,
	}
}

func (v XValidator) msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "gte", "min":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	case "exist":
		param := fe.Param()
		checkBang := param[0] == '!'
		label := "not exist"

		if checkBang {
			label = "exist"
			param = param[1:]
		}

		return fmt.Sprintf("%s is %s", param, label)
	}

	trans, _ := v.uni.GetTranslator("en")
	return fe.Translate(trans) // default error
}

func (v XValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) {
	v.validator.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = strcase.ToSnake(err.Field()) // Export struct field name
			elem.Tag = err.Tag()                            // Export struct tag
			elem.Value = err.Value()                        // Export field value
			elem.Error = true
			elem.Message = v.msgForTag(err)

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
