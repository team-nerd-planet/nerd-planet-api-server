package util

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

var validate *validator.Validate

func ValidateBody[T any](c iris.Context) (*T, error) {
	var input T

	if err := c.ReadJSON(&input); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	return &input, nil
}

func ValidateQuery[T any](c iris.Context) (*T, error) {
	var input T

	if err := c.ReadQuery(&input); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	return &input, nil
}

func ValidateForm[T any](c iris.Context) (*T, error) {
	var input T

	if err := c.ReadForm(&input); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	return &input, nil
}
