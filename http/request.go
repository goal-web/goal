package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type Request struct {
	echo.Context
	fields contracts.Fields
}

func (this *Request) GetString(key string) string {
	return utils.GetStringField(this.All(), key)
}

func (this *Request) GetInt64(key string) int64 {
	return utils.GetInt64Field(this.All(), key)
}

func (this *Request) GetInt(key string) int {
	return utils.GetIntField(this.All(), key)
}

func (this *Request) GetFloat64(key string) float64 {
	return utils.GetFloat64Field(this.All(), key)
}

func (this *Request) GetFloat(key string) float32 {
	return utils.GetFloatField(this.All(), key)
}

func (this *Request) GetBool(key string) bool {
	return utils.GetBoolField(this.All(), key)
}

func (this *Request) GetFields(key string) contracts.Fields {
	if field, isTypeRight := this.All()[key].(contracts.Fields); isTypeRight {
		return field
	}
	return nil
}

func (this *Request) Get(key string) (value interface{}) {
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

func (this *Request) All() contracts.Fields {
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
