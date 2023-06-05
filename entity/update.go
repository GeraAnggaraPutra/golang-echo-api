package entity

import (
	"github.com/go-playground/validator/v10"
)

type Update struct {
	NAME    string `json:"name" form:"name"`
	AGE     int    `json:"age" form:"age" validate:"gte=1,lte=100"`
	ADDRESS string `json:"address" form:"address"`
}


func (r *Update) ValidateUpdate() error {
	validate := validator.New()
	return validate.Struct(r)
}