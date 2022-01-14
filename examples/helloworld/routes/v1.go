package routes

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/examples/helloworld/app/controllers"
	"github.com/qbhy/goal/session"
)

func V1Routes(router contracts.Router) {
	router.Static("/", "public")

	v1 := router.Group("", session.StartSession)

	v1.Get("/", controllers.HelloWorld)
	v1.Get("/counter", controllers.Counter)
	v1.Get("/db", controllers.DatabaseQuery)
	v1.Get("/redis", controllers.RedisExample)
}
