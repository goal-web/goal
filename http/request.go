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

func NewRequest(ctx echo.Context) contracts.HttpRequest {
	return &Request{
		Context: ctx,
		fields:  nil,
	}
}

func (this *Request) GetString(key string) string {
	return utils.GetStringField(this.Fields(), key)
}

func (this *Request) GetInt64(key string) int64 {
	return utils.GetInt64Field(this.Fields(), key)
}

func (this *Request) GetInt(key string) int {
	return utils.GetIntField(this.Fields(), key)
}

func (this *Request) GetFloat64(key string) float64 {
	return utils.GetFloat64Field(this.Fields(), key)
}

func (this *Request) GetFloat(key string) float32 {
	return utils.GetFloatField(this.Fields(), key)
}

func (this *Request) GetBool(key string) bool {
	return utils.GetBoolField(this.Fields(), key)
}

func (this *Request) GetFields(key string) contracts.Fields {
	if field, isTypeRight := this.Fields()[key].(contracts.Fields); isTypeRight {
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

func (this *Request) Fields() contracts.Fields {
	if this.fields != nil {
		return this.fields
	}
	data := make(contracts.Fields)

	for key, query := range this.QueryParams() {
		if len(query) == 1 {
			data[key] = query[0]
		} else {
			data[key] = query
		}
	}
	for _, paramName := range this.ParamNames() {
		data[paramName] = this.Param(paramName)
	}
	if form, existsForm := this.FormParams(); existsForm == nil {
		for key, values := range form {
			if len(values) == 1 {
				data[key] = values[0]
			} else {
				data[key] = values
			}
		}
	}
	if multiForm, existsForm := this.MultipartForm(); existsForm == nil {
		for key, values := range multiForm.Value {
			if len(values) == 1 {
				data[key] = values[0]
			} else {
				data[key] = values
			}
		}
		for key, values := range multiForm.File {
			if len(values) == 1 {
				data[key] = values[0]
			} else {
				data[key] = values
			}
		}
	}

	return data
}
