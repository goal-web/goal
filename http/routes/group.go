package routes

import (
	"errors"
	"github.com/labstack/echo/v4"
)

var (
	MethodTypeError = errors.New("http method type unknown")
)

type Group struct {
	prefix      string
	middlewares []echo.MiddlewareFunc
	routes      []Route
}

func NewGroup(prefix string, middlewares ...echo.MiddlewareFunc) *Group {
	return &Group{
		prefix:      prefix,
		routes:      make([]Route, 0),
		middlewares: middlewares,
	}
}

// AddRoute 添加一条路由
func (group *Group) AddRoute(route Route) *Group {
	group.routes = append(group.routes, route)

	return group
}

// Add 添加路由，method 只允许字符串或者字符串数组
func (group *Group) Add(method interface{}, path string, handler HttpHandler, middlewares ...echo.MiddlewareFunc) *Group {
	methods := make([]string, 0)
	switch r := method.(type) {
	case string:
		methods = []string{r}
	case []string:
		methods = r
	default:
		panic(MethodTypeError)
	}
	group.AddRoute(Route{
		method:      methods,
		path:        group.prefix + path,
		middlewares: middlewares,
		handler:     handler,
	})

	return group
}

func (group *Group) Get(path string, handler HttpHandler, middlewares ...echo.MiddlewareFunc) *Group {
	return group.Add(GET, path, handler, middlewares...)
}

func (group *Group) Post(path string, handler HttpHandler, middlewares ...echo.MiddlewareFunc) *Group {
	return group.Add(POST, path, handler, middlewares...)
}

func (group *Group) Delete(path string, handler HttpHandler, middlewares ...echo.MiddlewareFunc) *Group {
	return group.Add(DELETE, path, handler, middlewares...)
}

func (group *Group) Put(path string, handler HttpHandler, middlewares ...echo.MiddlewareFunc) *Group {
	return group.Add(PUT, path, handler, middlewares...)
}
