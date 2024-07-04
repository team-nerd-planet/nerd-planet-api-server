package util

import (
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
)

func NewError(ctx iris.Context, status int, errors ...any) {
	description := make([]string, 0)

	for _, err := range errors {
		switch e := err.(type) {
		case error:
			description = append(description, e.Error())
		case string:
			description = append(description, e)
		}
	}

	ctx.StopWithJSON(status, HTTPError{
		Code:        status,
		Message:     http.StatusText(status),
		Description: strings.Join(description, ", "),
	})
}

type HTTPError struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description" binding:"omitempty"`
}
