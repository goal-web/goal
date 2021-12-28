package session

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
)

func StartSession(session contracts.Session, request contracts.HttpRequest, next echo.HandlerFunc) error {
	session.Start()
	return next(request.(echo.Context))
}
