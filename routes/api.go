package routes

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/controllers"
)

func Api(router contracts.HttpRouter) {

	router.Get("/wrk/{name}", func(request contracts.HttpRequest) any {
		return request.Param("name")
	})

	router.Get("/", controllers.HelloWorld)
	router.Get("/view", func(views contracts.Views) any {
		return views.Render("view.html")
	})
}
