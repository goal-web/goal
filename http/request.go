package http

import (
	"net/http"
)

type Request struct {
	rawRequest *http.Request
}
