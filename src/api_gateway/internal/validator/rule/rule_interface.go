package rule

import "github.com/go-playground/validator/v10"

type IRule interface {
	Validate(fl validator.FieldLevel) bool
}
