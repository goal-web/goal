package routes

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/auth/gate"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/http/controllers"
	"github.com/goal-web/goal/app/models"
	"github.com/goal-web/session"
)

func ApiRoutes(router contracts.Router) {

	router.Get("/queue", controllers.DemoJob)

	router.Post("/article", controllers.CreateArticle, gate.Authorize("create", models.ArticleModel))

	router.Get("/", controllers.HelloWorld)
	//router.Get("/", controllers.HelloWorld, ratelimiter.Middleware(100))
	router.Get("/counter", controllers.Counter, session.StartSession)
	router.Get("/db", controllers.DatabaseQuery)
	router.Get("/redis", controllers.RedisExample)
	router.Get("/users", controllers.GetUsers)
	router.Post("/login", controllers.LoginExample)

	router.Get("/myself", controllers.GetCurrentUser, auth.Guard("jwt"))

	authRouter := router.Group("", auth.Guard("jwt"))
	authRouter.Get("/myself", controllers.GetCurrentUser, auth.Guard("jwt"))

	router.Post("/mail", controllers.SendEmail)
}
