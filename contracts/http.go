package contracts

import "github.com/labstack/echo/v4"

type HttpResponse interface {
	Status() int
	Response(ctx echo.Context) error
}

type HttpRequest interface {
	echo.Context
	All() Fields
}
