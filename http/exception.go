package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
)

type HttpException struct {
	Exception contracts.Exception
	Context   echo.Context
}

func (h HttpException) Error() string {
	return h.Exception.Error()
}

func (h HttpException) Fields() contracts.Fields {
	return h.Exception.Fields()
}
