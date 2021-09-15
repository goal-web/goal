package http

import (
	"github.com/labstack/echo/v4"
	"qbhy/contracts"
)

type HttpException struct {
	exception contracts.Exception
	Method    string
	Path      string
	Context echo.Context
}

func (h HttpException) Error() string {
	return h.exception.Error()
}

func (h HttpException) Fields() contracts.Fields {
	return h.exception.Fields()
}

