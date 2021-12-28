package routes

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/examples/helloworld/controllers"
	"github.com/qbhy/goal/session"
)

func V1Routes(router contracts.Router) {
	v1 := router.Group("/", session.StartSession)

	v1.Get("/", controllers.HelloWorld)
}
