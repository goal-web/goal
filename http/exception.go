package http

import (
	"github.com/qbhy/goal/contracts"
)

type HttpException struct {
	Exception contracts.Exception
	Request   contracts.HttpRequest
}

func (this HttpException) Error() string {
	return this.Exception.Error()
}

func (this HttpException) Fields() contracts.Fields {
	return contracts.Fields{
		"method": this.Request.Request().Method,
		"path":   this.Request.Path(),
		"query":  this.Request.QueryParams(),
		"fields": this.Request.Fields(),
	}
}
