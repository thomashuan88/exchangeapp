package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Validator is a validation middleware
type Validator struct {
	validate *validator.Validate
}

// NewValidator returns a new Validator instance
func NewValidator() binding.StructValidator {
	return &Validator{
		validate: validator.New(),
	}
}

// Validate validates a struct
func (v *Validator) Validate(obj interface{}) error {
	return v.validate.Struct(obj)
}

func (v *Validator) Engine() interface{} {
	return v.validate
}

func (v *Validator) ValidateStruct(structure interface{}) error {
	return v.validate.Struct(structure)
}
