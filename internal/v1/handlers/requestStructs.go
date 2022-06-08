package handlers

import (
	validator "github.com/go-ozzo/ozzo-validation/v4"
)

type MathopsRequest struct {
	NumOne *float64 `json:"num_one"` // encode-decode JSON using as keys this name tags
	NumTwo *float64 `json:"num_two"`
}

func (r MathopsRequest) Validate() error {
	return validator.ValidateStruct(&r,
		validator.Field(&r.NumOne, validator.NotNil),
		validator.Field(&r.NumTwo, validator.NotNil),
	)
}
