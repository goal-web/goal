package routes

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/examples/helloworld/app/controllers"
	"github.com/qbhy/goal/session"
)

func ApiRoutes(router contracts.Router) {

	v1 := router.Group("", session.StartSession)

	v1.Get("/", controllers.HelloWorld)
	v1.Get("/counter", controllers.Counter)
	v1.Get("/db", controllers.DatabaseQuery)
	v1.Get("/redis", controllers.RedisExample)
}
