package contracts

type Route interface {
	Middlewares() []interface{}
	Method() []string
	Path() string
	Handler() interface{}
}

type RouteGroup interface {
	Get(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Post(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Delete(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Put(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Patch(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Options(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Trace(path string, handler interface{}, middlewares ...interface{}) RouteGroup
	Middlewares() []interface{}
	Routes() []Route
}

type Router interface {
	Get(path string, handler interface{}, middlewares ...interface{})
	Post(path string, handler interface{}, middlewares ...interface{})
	Delete(path string, handler interface{}, middlewares ...interface{})
	Put(path string, handler interface{}, middlewares ...interface{})
	Patch(path string, handler interface{}, middlewares ...interface{})
	Options(path string, handler interface{}, middlewares ...interface{})
	Trace(path string, handler interface{}, middlewares ...interface{})
	Group(prefix string, middlewares ...interface{}) RouteGroup
	Start(string) error
}
