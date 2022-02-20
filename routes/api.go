package routes

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/http/controllers"
)

func ApiRoutes(router contracts.Router) {

	router.Get("/queue", controllers.DemoJob)

	router.Get("/", controllers.HelloWorld)
	//router.Get("/", controllers.HelloWorld, ratelimiter.Middleware(100))
	router.Post("/login", controllers.LoginExample)

	router.Get("/myself", controllers.GetCurrentUser, auth.Guard("jwt"))

	authRouter := router.Group("", auth.Guard("jwt"))
	authRouter.Get("/myself", controllers.GetCurrentUser, auth.Guard("jwt"))

	router.Post("/mail", controllers.SendEmail)
}
