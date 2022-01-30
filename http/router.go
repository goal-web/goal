package http

import (
	"errors"
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
	"github.com/goal-web/pipeline"
	"github.com/labstack/echo/v4"
	"strings"
)

var (
	ignoreError     = errors.New("忽略该错误")            // 用于中间件直接返回响应
	MiddlewareError = errors.New("middleware error") // 中间件必须有一个返回值

	// magical middlewares
	exceptionHandler = container.NewMagicalFunc(func(handler contracts.ExceptionHandler, exception Exception) interface{} {
		return handler.Handle(exception)
	})
)

func New(container contracts.Application) contracts.Router {
	router := &Router{
		app:         container,
		events:      container.Get("events").(contracts.EventDispatcher),
		echo:        echo.New(),
		routes:      make([]contracts.Route, 0),
		groups:      make([]contracts.RouteGroup, 0),
		middlewares: make([]contracts.MagicalFunc, 0),
	}

	router.Use(router.recovery)

	return router
}

type Router struct {
	events contracts.EventDispatcher
	app    contracts.Application
	echo   *echo.Echo
	groups []contracts.RouteGroup
	routes []contracts.Route

	// 全局中间件
	middlewares []contracts.MagicalFunc
}

func (this *Router) Group(prefix string, middlewares ...interface{}) contracts.RouteGroup {
	groupInstance := NewGroup(prefix, middlewares...)

	this.groups = append(this.groups, groupInstance)

	return groupInstance
}

func (this *Router) Close() error {
	return this.echo.Close()
}

func (this *Router) Static(path, directory string) {
	if strings.HasPrefix(directory, "/") {
		directory = this.app.Get("path").(string) + "/" + directory
	}
	this.echo.Static(path, directory)
}

func (this *Router) Get(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.GET, path, handler, middlewares...)
}

func (this *Router) Post(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.POST, path, handler, middlewares...)
}

func (this *Router) Delete(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.DELETE, path, handler, middlewares...)
}

func (this *Router) Put(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.PUT, path, handler, middlewares...)
}

func (this *Router) Patch(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.PATCH, path, handler, middlewares...)
}

func (this *Router) Options(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.OPTIONS, path, handler, middlewares...)
}

func (this *Router) Trace(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.TRACE, path, handler, middlewares...)
}

func (this *Router) Use(middlewares ...interface{}) {
	for _, middleware := range middlewares {
		if magicalFunc, ok := middleware.(contracts.MagicalFunc); ok {
			this.middlewares = append(this.middlewares, magicalFunc)
		} else {
			this.middlewares = append(this.middlewares, container.NewMagicalFunc(middleware))
		}
	}
}

func (this *Router) Add(method interface{}, path string, handler interface{}, middlewares ...interface{}) {
	methods := make([]string, 0)
	switch v := method.(type) {
	case string:
		methods = []string{v}
	case []string:
		methods = v
	default:
		panic(errors.New("method 只能接收 string 或者 []string"))
	}
	this.routes = append(this.routes, &route{
		method:      methods,
		path:        path,
		middlewares: convertToMiddlewares(middlewares...),
		handler:     container.NewMagicalFunc(handler),
	})
}

// Start 启动 httpserver
func (this *Router) Start(address string) error {

	this.mountRoutes(this.routes)

	for _, routeGroup := range this.groups {
		this.mountRoutes(routeGroup.Routes(), routeGroup.Middlewares()...)
	}

	this.echo.HTTPErrorHandler = func(err error, context echo.Context) {
		this.app.StaticCall(exceptionHandler, Exception{error: err, Request: NewRequest(context)})
	}
	this.echo.Debug = this.app.Debug()

	return this.echo.Start(address)
}

// mountRoutes 装配路由
func (this *Router) mountRoutes(routes []contracts.Route, middlewares ...contracts.MagicalFunc) {
	for _, routeItem := range routes {
		(func(routeInstance contracts.Route) {
			this.echo.Match(routeInstance.Method(), routeInstance.Path(), func(context echo.Context) error {
				request := NewRequest(context)
				defer func() {
					this.events.Dispatch(&RequestAfter{request})
				}()

				result := pipeline.Static(this.app).SendStatic(request).
					ThroughStatic(
						this.middlewares...,
					).
					ThroughStatic(
						append(middlewares, routeInstance.Middlewares()...)...,
					).
					ThenStatic(routeInstance.Handler())

				this.events.Dispatch(&ResponseBefore{request})
				HandleResponse(result, request)

				return nil
			})
		})(routeItem)
	}
}
