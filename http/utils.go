package http

import (
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
)

func convertToMiddlewares(middlewares ...interface{}) (results []contracts.MagicalFunc) {
	for _, middleware := range middlewares {
		magicalFunc, isMiddleware := middleware.(contracts.MagicalFunc)
		if !isMiddleware {
			magicalFunc = container.NewMagicalFunc(middleware)
		}
		if magicalFunc.NumOut() != 1 {
			panic(MiddlewareError)
		}
		results = append(results, magicalFunc)
	}
	return
}
