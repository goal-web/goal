package routes

import "github.com/labstack/echo/v4"

var (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type HttpHandler = func(echo.Context) interface{}

type Route struct {
	method      []string
	path        string
	middlewares []echo.MiddlewareFunc
	handler     HttpHandler
}

func (route Route) Handle(ctx echo.Context) interface{} {
	return route.handler(ctx)
}
