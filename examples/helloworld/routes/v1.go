package routes

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/examples/helloworld/controllers"
)

func V1Routes(router contracts.Router) {
	router.Get("/", controllers.HelloWorld)
}
