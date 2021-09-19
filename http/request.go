package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/utils"
)

type Request struct {
	echo.Context
}

func (this Request) Get(key string) (value interface{}) {
	if value = utils.StringOr(
		this.Param(key),
		this.QueryParams().Get(key),
	); value != "" {
		return value
	}
	return utils.NotNil(this.Context.Get(key), this.FormValue(key))
}
