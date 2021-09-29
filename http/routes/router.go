package routes

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/http"
)

var (
	ignoreError = errors.New("忽略该错误") // 用于中间件直接返回响应
)

func New(container contracts.Container) contracts.Router {
	return &router{
		app:    container,
		e:      echo.New(),
		routes: make([]contracts.Route, 0),
		groups: make([]contracts.RouteGroup, 0),
	}
}

type router struct {
	app    contracts.Container
	e      *echo.Echo
	groups []contracts.RouteGroup
	routes []contracts.Route
}

func (this *router) errHandler(err error, context echo.Context) {
	if ignoreError == err {
		return
	}
	var httpException http.HttpException
	switch rawErr := err.(type) {
	case http.HttpException:
		httpException = rawErr
	default:
		httpException = http.HttpException{
			Exception: exceptions.ResolveException(err),
			Context:   context,
		}
	}

	go func() {
		// 调用容器内的异常处理器
		this.app.Call(func(handler contracts.ExceptionHandler, exception http.HttpException) {
			handler.Handle(exception)
		}, httpException)
	}()
}

func (this *router) Group(prefix string, middlewares ...interface{}) contracts.RouteGroup {
	groupInstance := NewGroup(prefix, middlewares...)

	this.groups = append(this.groups, groupInstance)

	return groupInstance
}

func (this *router) Get(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.GET, path, handler, middlewares...)
}

func (this *router) Post(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.POST, path, handler, middlewares...)
}

func (this *router) Delete(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.DELETE, path, handler, middlewares...)
}

func (this *router) Put(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.PUT, path, handler, middlewares...)
}

func (this *router) Patch(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.PATCH, path, handler, middlewares...)
}

func (this *router) Options(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.OPTIONS, path, handler, middlewares...)
}

func (this *router) Trace(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.TRACE, path, handler, middlewares...)
}

func (this *router) Use(middleware ...interface{}) {
	this.e.Use(this.resolveMiddlewares(middleware)...)
}

func (this *router) Add(method interface{}, path string, handler interface{}, middlewares ...interface{}) {
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
		middlewares: middlewares,
		handler:     handler,
	})
}

// Start 启动 server
func (this *router) Start(address string) error {
	this.mountRoutes(this.routes)

	for _, routeGroup := range this.groups {
		this.mountRoutes(routeGroup.Routes(), routeGroup.Middlewares()...)
	}

	// recovery
	this.Use(func(request http.Request, next echo.HandlerFunc) (result error) {
		defer func() {
			if err := recover(); err != nil {
				this.errHandler(exceptions.ResolveException(err), request)
				result = ignoreError
			}
		}()
		return next(request)
	})

	this.e.HTTPErrorHandler = this.errHandler

	return this.e.Start(address)
}

// mountRoutes 装配路由
func (this *router) mountRoutes(routes []contracts.Route, middlewares ...interface{}) {
	for _, routeItem := range routes {
		(func(routeInstance contracts.Route) {
			this.e.Match(routeInstance.Method(), routeInstance.Path(), func(context echo.Context) error {
				request := http.Request{Context: context}
				results := this.app.Call(routeInstance.Handler(), request)
				if len(results) > 0 {
					if result, isErr := results[0].(error); isErr {
						return result
					}
					http.HandleResponse(results[0], request)
					return ignoreError
				}
				return nil
			}, this.resolveMiddlewares(append(middlewares, routeInstance.Middlewares()...))...)
		})(routeItem)
	}
}

func (this *router) resolveMiddlewares(interfaceMiddlewares []interface{}) []echo.MiddlewareFunc {
	middlewares := make([]echo.MiddlewareFunc, 0)

	for _, middlewareItem := range interfaceMiddlewares {
		if middleware, isEchoMiddleware := middlewareItem.(echo.MiddlewareFunc); isEchoMiddleware {
			middlewares = append(middlewares, middleware)
			continue
		}
		(func(middleware interface{}) {
			middlewares = append(middlewares, func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(context echo.Context) (err error) {
					rawResult := this.app.Call(middlewareItem, http.Request{context}, next)[0]
					switch result := rawResult.(type) {
					case error:
						return result
					default:
						http.HandleResponse(result, context)
						return ignoreError
					}
				}
			})
		})(middlewareItem)
	}
	return middlewares
}
