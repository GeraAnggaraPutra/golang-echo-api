package entity

import (
	"github.com/go-playground/validator/v10"
)

type Request struct {
	NAME    string `json:"name" form:"name" validate:"required"`
	AGE     int    `json:"age" form:"age" validate:"required,gte=1,lte=100"`
	ADDRESS string `json:"address" form:"address" validate:"required"`
}

func (r *Request) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
