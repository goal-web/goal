package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/container"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"strings"
)

var (
	ignoreError = errors.New("忽略该错误") // 用于中间件直接返回响应

	// magical functions
	exceptionHandler = container.NewMagicalFunc(func(handler contracts.ExceptionHandler, exception contracts.Exception) {
		handler.Handle(exception)
	})
)

func New(container contracts.Application) contracts.Router {
	return &router{
		app:    container,
		events: container.Get("events").(contracts.EventDispatcher),
		echo:   echo.New(),
		routes: make([]contracts.Route, 0),
		groups: make([]contracts.RouteGroup, 0),
	}
}

type router struct {
	events contracts.EventDispatcher
	app    contracts.Application
	echo   *echo.Echo
	groups []contracts.RouteGroup
	routes []contracts.Route
}

func (this *router) errHandler(err error, ctx echo.Context) {
	request, isRequest := ctx.(contracts.HttpRequest)
	if !isRequest {
		request = NewRequest(ctx)
	}
	if ignoreError == err {
		return
	}
	var httpException HttpException
	switch rawErr := err.(type) {
	case HttpException:
		httpException = rawErr
	default:
		httpException = HttpException{
			Exception: exceptions.ResolveException(err),
			Request:   request,
		}
	}

	// 调用容器内的异常处理器
	this.app.Call(exceptionHandler, httpException)
}

func (this *router) Group(prefix string, middlewares ...interface{}) contracts.RouteGroup {
	groupInstance := NewGroup(prefix, middlewares...)

	this.groups = append(this.groups, groupInstance)

	return groupInstance
}

func (this *router) Close() error {
	return this.echo.Close()
}
func (this *router) Static(path, directory string) {
	if strings.HasPrefix(directory, "/") {
		directory = this.app.Get("path").(string) + "/" + directory
	}
	this.echo.Static(path, directory)
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
	this.echo.Use(this.resolveMiddlewares(middleware)...)
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
		handler:     container.NewMagicalFunc(handler),
	})
}

// start 启动 server
func (this *router) Start(address string) error {
	// recovery 。 这里为了 contracts 不依赖 echo ，要求 Request 必须继承自 echo.Context !!!
	this.Use(
		func(request *Request, next echo.HandlerFunc) (result error) {
			defer func() {
				if err := recover(); err != nil {
					this.errHandler(exceptions.ResolveException(err), request)
					result = ignoreError
				}
			}()

			// 触发钩子
			this.events.Dispatch(&RequestBefore{request})
			return next(request)
		},
	)

	this.mountRoutes(this.routes)

	for _, routeGroup := range this.groups {
		this.mountRoutes(routeGroup.Routes(), routeGroup.Middlewares()...)
	}

	this.echo.HTTPErrorHandler = this.errHandler
	this.echo.Debug = this.app.Debug()

	return this.echo.Start(address)
}

// mountRoutes 装配路由
func (this *router) mountRoutes(routes []contracts.Route, middlewares ...interface{}) {
	for _, routeItem := range routes {
		(func(routeInstance contracts.Route) {
			this.echo.Match(routeInstance.Method(), routeInstance.Path(), func(context echo.Context) error {
				request := &Request{Context: context}
				results := this.app.Call(routeInstance.Handler(), request)
				if len(results) > 0 {
					if result, isErr := results[0].(error); isErr {
						return result
					}
					this.events.Dispatch(&RequestAfter{request})
					HandleResponse(results[0], request)
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
					request := NewRequest(context)
					rawResult := this.app.Call(container.NewMagicalFunc(middlewareItem), request, next)[0]
					switch result := rawResult.(type) {
					case error:
						return result
					default:
						this.events.Dispatch(&RequestAfter{request})
						HandleResponse(result, context)
						return ignoreError
					}
				}
			})
		})(middlewareItem)
	}
	return middlewares
}
