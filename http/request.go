package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type Request struct {
	echo.Context
}

func (this Request) Get(key string) (value interface{}) {
	if value = this.Context.Get(key); value != nil {
		return value
	}
	if value = utils.StringOr(
		this.Param(key),
		this.QueryParams().Get(key),
	); value != "" {
		return value
	}
	return this.FormValue(key)
}

func (this Request) All() contracts.Fields {
	data := make(contracts.Fields)

	for key, query := range this.QueryParams() {
		data[key] = query
	}
	for _, paramName := range this.ParamNames() {
		data[paramName] = this.Param(paramName)
	}
	if form, existsForm := this.FormParams(); existsForm == nil {
		for key, values := range form {
			data[key] = values
		}
	}
	if multiForm, existsForm := this.MultipartForm(); existsForm == nil {
		for key, values := range multiForm.Value {
			data[key] = values
		}
		for key, values := range multiForm.File {
			data[key] = values
		}
	}

	return data
}
