package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
)

type HttpException struct {
	Exception contracts.Exception
	Context   echo.Context
}

func (this HttpException) Error() string {
	return this.Exception.Error()
}

func (this HttpException) Fields() contracts.Fields {
	return contracts.Fields{
		"method": this.Context.Request().Method,
		"path":   this.Context.Path(),
		"query":  this.Context.QueryParams(),
		"fields": NewRequest(this.Context).Fields(),
	}
}
