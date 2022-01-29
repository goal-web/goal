package http

import (
	"github.com/goal-web/contracts"
)

type Exception struct {
	error
	Request contracts.HttpRequest
}

func (this Exception) Fields() contracts.Fields {
	return contracts.Fields{
		"method": this.Request.Request().Method,
		"path":   this.Request.Path(),
		"query":  this.Request.QueryParams(),
		"fields": this.Request.Fields(),
	}
}
