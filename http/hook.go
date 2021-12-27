package http

import "github.com/qbhy/goal/contracts"

type RequestBefore struct {
	request contracts.HttpRequest
}

func (this *RequestBefore) Event() string {
	return "REQUEST_BEFORE"
}

func (this *RequestBefore) Sync() bool {
	return true
}

type RequestAfter struct {
	request contracts.HttpRequest
}

func (this *RequestAfter) Event() string {
	return "REQUEST_AFTER"
}

func (this *RequestAfter) Sync() bool {
	return true
}
