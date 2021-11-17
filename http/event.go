package http

import (
	"github.com/qbhy/goal/contracts"
)

var (
	HTTP_SERVE_CLOSED = contracts.EventName("HTTP_SERVE_CLOSED")
)

type HttpServeClosed struct {
}

func (this *HttpServeClosed) Name() contracts.EventName {
	return HTTP_SERVE_CLOSED
}
