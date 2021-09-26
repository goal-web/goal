package routes

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
)

var (
	MethodTypeError = errors.New("http method type unknown")
)

type group struct {
	prefix      string
	middlewares []interface{}
	routes      []contracts.Route
}

func NewGroup(prefix string, middlewares ...interface{}) contracts.RouteGroup {
	return &group{
		prefix:      prefix,
		routes:      make([]contracts.Route, 0),
		middlewares: middlewares,
	}
}

// AddRoute 添加一条路由
func (group *group) AddRoute(route contracts.Route) contracts.RouteGroup {
	group.routes = append(group.routes, route)

	return group
}

// Add 添加路由，method 只允许字符串或者字符串数组
func (group *group) Add(method interface{}, path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	methods := make([]string, 0)
	switch r := method.(type) {
	case string:
		methods = []string{r}
	case []string:
		methods = r
	default:
		panic(MethodTypeError)
	}
	group.AddRoute(&route{
		method:      methods,
		path:        group.prefix + path,
		middlewares: middlewares,
		handler:     handler,
	})

	return group
}

func (group *group) Get(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.GET, path, handler, middlewares...)
}

func (group *group) Post(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.POST, path, handler, middlewares...)
}

func (group *group) Delete(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.DELETE, path, handler, middlewares...)
}

func (group *group) Put(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.PUT, path, handler, middlewares...)
}

func (group *group) Trace(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.TRACE, path, handler, middlewares...)
}

func (group *group) Patch(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.PATCH, path, handler, middlewares...)
}

func (group *group) Options(path string, handler interface{}, middlewares ...interface{}) contracts.RouteGroup {
	return group.Add(echo.OPTIONS, path, handler, middlewares...)
}

func (group *group) Middlewares() []interface{} {
	return group.middlewares
}

func (group *group) Routes() []contracts.Route {
	return group.routes
}
