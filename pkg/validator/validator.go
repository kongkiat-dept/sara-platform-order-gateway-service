package validator

import (
	"time"

	validators "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(inf interface{}) error
}

type validator struct {
	validator *validators.Validate
}

var (
	dateFormat = "2006-01-02 15:04:05"
)

func New() Validator {
	v := validators.New()
	dateValidator(v)
	return &validator{
		validator: v,
	}
}

func (v *validator) ValidateStruct(inf interface{}) error {

	return v.validator.Struct(inf)
}

//  date format validator ...
func dateValidator(v *validators.Validate) {
	v.RegisterValidation("dateFormat", func(fl validators.FieldLevel) bool {
		if fl != nil {
			d := fl.Field().String()
			_, err := time.Parse(dateFormat, d)
			return err == nil
		} else {
			return true
		}
	})
}
