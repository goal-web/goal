package http

import (
	"errors"
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"github.com/labstack/echo/v4"
	"strings"
)

var (
	ignoreError = errors.New("忽略该错误") // 用于中间件直接返回响应

	// magical functions
	exceptionHandler = container.NewMagicalFunc(func(handler contracts.ExceptionHandler, exception Exception) interface{} {
		return handler.Handle(exception)
	})
)

func New(container contracts.Application) contracts.Router {
	return &router{
		app:       container,
		events:    container.Get("events").(contracts.EventDispatcher),
		echo:      echo.New(),
		routes:    make([]contracts.Route, 0),
		groups:    make([]contracts.RouteGroup, 0),
		functions: make([]contracts.MagicalFunc, 0),
	}
}

type router struct {
	events    contracts.EventDispatcher
	app       contracts.Application
	echo      *echo.Echo
	groups    []contracts.RouteGroup
	routes    []contracts.Route
	functions []contracts.MagicalFunc
}

func (this *router) errHandler(err interface{}, request contracts.HttpRequest) (result interface{}) {
	if ignoreError == err || err == nil {
		return nil
	}
	var httpException Exception
	switch rawErr := err.(type) {
	case Exception:
		httpException = rawErr
	default:
		httpException = Exception{
			error:   exceptions.ResolveException(err),
			Request: request,
		}
	}

	// 调用容器内的异常处理器
	return this.app.StaticCall(exceptionHandler, httpException)
}

func (this *router) Group(prefix string, middlewares ...interface{}) contracts.RouteGroup {
	groupInstance := NewGroup(prefix, middlewares...)

	this.groups = append(this.groups, groupInstance)

	return groupInstance
}

func (this *router) Close() error {
	return this.echo.Close()
}

// addFunc 添加一个静态方法并且返回对应的索引
func (this *router) addFunc(handler interface{}) int {
	id := len(this.functions)
	if magicalFunc, ok := handler.(contracts.MagicalFunc); ok {
		this.functions = append(this.functions, magicalFunc)
	} else {
		this.functions = append(this.functions, container.NewMagicalFunc(handler))
	}
	return id
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

// Start 启动 httpserver
func (this *router) Start(address string) error {
	// recovery 。 这里为了 contracts 不依赖 echo ，要求 Request 必须继承自 echo.Context !!!
	this.Use(
		func(request *Request, next echo.HandlerFunc) (result error) {
			defer func() {
				if res := this.errHandler(recover(), request); res != nil {
					HandleResponse(res, request)
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

	this.echo.HTTPErrorHandler = func(err error, context echo.Context) {
		this.app.StaticCall(exceptionHandler, Exception{error: err, Request: NewRequest(context)})
	}
	this.echo.Debug = this.app.Debug()

	return this.echo.Start(address)
}

// mountRoutes 装配路由
func (this *router) mountRoutes(routes []contracts.Route, middlewares ...interface{}) {
	for _, routeItem := range routes {
		(func(routeInstance contracts.Route) {
			this.echo.Match(routeInstance.Method(), routeInstance.Path(), func(context echo.Context) error {
				// 包装 request
				request := NewRequest(context)
				// 调用控制器方法
				results := this.app.StaticCall(routeInstance.Handler(), request)
				this.events.Dispatch(&RequestAfter{request})
				// 若有返回值则处理返回并且不继续往下执行
				if len(results) > 0 {
					if result, isErr := results[0].(error); isErr {
						return result
					}
					HandleResponse(results[0], request)
					return ignoreError
				}
				return nil
			}, this.resolveMiddlewares(append(middlewares, routeInstance.Middlewares()...))...)
		})(routeItem)
	}
}

// resolveMiddlewares 装配中间件
func (this *router) resolveMiddlewares(interfaceMiddlewares []interface{}) []echo.MiddlewareFunc {
	middlewares := make([]echo.MiddlewareFunc, 0)

	for _, middlewareItem := range interfaceMiddlewares {
		if middleware, isEchoMiddleware := middlewareItem.(echo.MiddlewareFunc); isEchoMiddleware {
			middlewares = append(middlewares, middleware)
			continue
		}

		// 如果不是 echo 原生的中间件，包装成 echo 中间件并且静态化
		(func(middleware interface{}) {
			id := this.addFunc(middleware) // 静态化
			middlewares = append(middlewares, func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(context echo.Context) (err error) {
					request := NewRequest(context)
					rawResult := this.app.StaticCall(this.functions[id], request, next)[0]
					switch result := rawResult.(type) {
					case contracts.HttpResponse: // 如果中间件直接返回 response 实例，则直接响应
						logs.WithError(result.Response(request)).Error("response err from middleware")
						this.events.Dispatch(&RequestAfter{request})
						return ignoreError
					}
					return
				}
			})
		})(middlewareItem)
	}

	return middlewares
}
