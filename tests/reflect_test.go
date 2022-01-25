package tests

import (
	"fmt"
	"github.com/goal-web/supports/utils"
	"github.com/labstack/echo/v4"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var ctx *echo.Context
	ctxInstance := echo.New().NewContext(nil, nil)
	fmt.Println(utils.GetTypeKey(reflect.TypeOf(ctx)))
	fmt.Println(utils.GetTypeKey(reflect.TypeOf(ctxInstance)))
}

func TestParseTag(t *testing.T) {
	rawTag := `di:"true,false"`

	fmt.Println(utils.ParseStructTag(reflect.StructTag(rawTag))["di"])
}
