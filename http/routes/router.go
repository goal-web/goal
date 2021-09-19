package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/http"
)

func New(container contracts.Container) Router {
	return Router{
		app:    container,
		e:      echo.New(),
		routes: make([]Route, 0),
		groups: make([]*Group, 0),
	}
}

type Router struct {
	app    contracts.Container
	e      *echo.Echo
	groups []*Group
	routes []Route
}

func (router *Router) Group(prefix string, middlewares ...echo.MiddlewareFunc) *Group {
	group := NewGroup(prefix, middlewares...)

	router.groups = append(router.groups, group)

	return group
}

func (router *Router) Get(path string, handler interface{}, middlewares ...echo.MiddlewareFunc) {
	router.Add(GET, path, handler, middlewares...)
}

func (router *Router) Post(path string, handler interface{}, middlewares ...echo.MiddlewareFunc) {
	router.Add(POST, path, handler, middlewares...)
}

func (router *Router) Delete(path string, handler interface{}, middlewares ...echo.MiddlewareFunc) {
	router.Add(DELETE, path, handler, middlewares...)
}

func (router *Router) Put(path string, handler interface{}, middlewares ...echo.MiddlewareFunc) {
	router.Add(PUT, path, handler, middlewares...)
}

func (router *Router) Use(middleware ...echo.MiddlewareFunc) {
	router.e.Use(middleware...)
}

func (router *Router) mountMiddleware(middlewares []echo.MiddlewareFunc) []echo.MiddlewareFunc {
	mountedMiddlewares := make([]echo.MiddlewareFunc, 0)
	for _, middleware := range middlewares {
		(func(middleware echo.MiddlewareFunc) {
			mountedMiddlewares = append(mountedMiddlewares, func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
				return router.app.Call(middleware, handlerFunc)[0].(echo.HandlerFunc)
			})
		})(middleware)
	}
	return mountedMiddlewares
}

func (router *Router) Add(method interface{}, path string, handler interface{}, middlewares ...echo.MiddlewareFunc) {
	methods := make([]string, 0)
	switch v := method.(type) {
	case string:
		methods = []string{v}
	case []string:
		methods = v
	}
	router.routes = append(router.routes, Route{
		method:      methods,
		path:        path,
		middlewares: middlewares,
		handler:     handler,
	})
}

// Start 启动 server
func (router *Router) Start(address string) error {
	router.mountRoutes(router.routes)

	for _, group := range router.groups {
		router.mountRoutes(group.routes, group.middlewares...)
	}

	return router.e.Start(address)
}

// mountRoutes 装配路由
func (router *Router) mountRoutes(routes []Route, middlewares ...echo.MiddlewareFunc) {
	for _, route := range routes {
		(func(route Route) {
			router.e.Match(route.method, route.path, func(context echo.Context) error {
				defer func() {
					if err := recover(); err != nil {
						exceptions.Handle(http.HttpException{
							Exception: exceptions.ResolveException(err),
							Context:   context,
						})
					}
				}()
				request := http.Request{Context: context}
				results := router.app.Call(route.handler, request)
				if len(results) > 0 {
					http.HandleResponse(results[0], request)
				}
				return nil
			}, append(middlewares, route.middlewares...)...)
		})(route)
	}
}
