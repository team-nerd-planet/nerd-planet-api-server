package util

import (
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateBody[T any](c *gin.Context) (*T, error) {
	var input T

	if err := c.ShouldBind(&input); err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, err
	}

	return &input, nil
}

func ValidateQuery[T any](c *gin.Context) (*T, error) {
	var input T

	if err := c.ShouldBindQuery(&input); err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, err
	}

	return &input, nil
}

func ValidateInt64Param(c *gin.Context, key string) (*int64, error) {
	param := c.Param(key)
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, err
	}

	return &id, nil
}
