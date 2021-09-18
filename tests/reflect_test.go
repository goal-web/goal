package tests

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/qbhy/goal/utils"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var ctx *echo.Context
	ctxInstance := echo.New().NewContext(nil,nil)
	fmt.Println(utils.GetTypeKey(reflect.TypeOf(ctx)))
	fmt.Println(utils.GetTypeKey(reflect.TypeOf(ctxInstance)))
}
