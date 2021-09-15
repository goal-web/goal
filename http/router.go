package http

import (
	"github.com/labstack/echo/v4"
	"goal/contracts"
	"goal/exceptions"
	"goal/logs"
)

func New() Router {
	return Router{
		e: echo.New(),
	}
}

type Router struct {
	e *echo.Echo
}

type HttpHandler = func(echo.Context) interface{}

func (router Router) Get(path string, handler HttpHandler) {
	router.Add("GET", path, handler)
}

func (router Router) Post(path string, handler HttpHandler) {
	router.Add("POST", path, handler)
}

func (router Router) Delete(path string, handler HttpHandler) {
	router.Add("DELETE", path, handler)
}

func (router Router) Put(path string, handler HttpHandler) {
	router.Add("PUT", path, handler)
}

func (router Router) Use(middleware ...echo.MiddlewareFunc) {
	router.e.Use(middleware...)
}

func (router Router) Add(method, path string, handler HttpHandler) {
	router.e.Add(method, path, func(context echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				exceptions.Handle(HttpException{
					exception: exceptions.ResolveException(err),
					Method:    method,
					Path:      path,
					Context:   context,
				})
			}
		}()
		switch res := handler(context).(type) {
		case error, contracts.Exception:
			exceptions.Handle(HttpException{
				exception: exceptions.ResolveException(res),
				Method:    method,
				Path:      path,
				Context:   context,
			})
		case string:
			logs.WithError(context.String(200, res)).Debug("response error")
		}

		return nil
	})
}

func (router Router) Start(address string) error {
	return router.e.Start(address)
}
