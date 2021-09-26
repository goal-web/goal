package routes

type route struct {
	method      []string
	path        string
	middlewares []interface{}
	handler     interface{}
}

func (route *route) Middlewares() []interface{} {
	return route.middlewares
}

func (route *route) Method() []string {
	return route.method
}

func (route *route) Path() string {
	return route.path
}

func (route *route) Handler() interface{} {
	return route.handler
}
