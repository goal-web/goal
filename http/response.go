package http

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/labstack/echo/v4"
	"go/types"
	"os"
)

var (
	FileTypeError = errors.New("文件参数类型错误")
)

// HandleResponse 处理控制器函数的响应
func HandleResponse(response interface{}, ctx echo.Context) {
	switch res := response.(type) {
	case error, contracts.Exception:
		panic(res)
	case string:
		logs.WithError(ctx.String(200, res)).Debug("response error")
	case contracts.HttpResponse:
		logs.WithError(res.Response(ctx)).Debug("response error")
	case types.Nil:
		return
	default:
		logs.WithError(ctx.JSON(200, res)).Debug("response json error")
	}

}

type Response struct {
	status   int
	Json     interface{}
	String   string
	FilePath string
	File     *os.File
}

func StringResponse(str string, code ...int) Response {
	status := 200
	if len(code) > 0 {
		status = code[0]
	}
	return Response{
		status: status,
		String: str,
	}
}

func JsonResponse(json interface{}, code ...int) Response {
	status := 200
	if len(code) > 0 {
		status = code[0]
	}
	return Response{
		status: status,
		Json:   json,
	}
}

// FileResponse 响应文件
func FileResponse(file interface{}) Response {
	switch f := file.(type) {
	case *os.File:
		return Response{File: f}
	case string:
		return Response{FilePath: f}
	default:
		panic(FileTypeError)
	}
}

func (res Response) Status() int {
	return res.status
}

func (res Response) Response(ctx contracts.Context) error {
	if res.Json != nil {
		return ctx.JSON(res.Status(), res.Json)
	}
	if res.FilePath != "" {
		return ctx.File(res.FilePath)
	}
	if res.File != nil {
		return ctx.File(res.File.Name())
	}

	return ctx.String(res.Status(), res.String)
}
