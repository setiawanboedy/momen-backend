package handler

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)

func ErrorValidationHandler(err error) []string {
	var errorList []string
	var jsErr *json.UnmarshalTypeError
	if errors.As(err, &jsErr) {
		errorList = append(errorList, "something wrong with input")
	} else {
		for _, e := range err.(validator.ValidationErrors) {
			errorList = append(errorList, e.Error())
		}
	}

	return errorList
}
