package http

import "github.com/qbhy/goal/contracts"

type route struct {
	method      []string
	path        string
	middlewares []interface{}
	handler     contracts.MagicalFunc
}

func (route *route) Middlewares() []interface{} {
	return route.middlewares
}

func (route *route) Method() []string {
	return route.method
}

func (route *route) Path() string {
	return route.path
}

func (route *route) Handler() contracts.MagicalFunc {
	return route.handler
}
