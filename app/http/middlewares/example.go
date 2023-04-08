package middlewares

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
)

func Example(request contracts.HttpRequest, next contracts.Pipe) any {
	fmt.Println("controller before")

	result := next(request)

	fmt.Println("controller after")
	return http.JsonResponse(contracts.Fields{
		"result": result,
	}, 200)
}
