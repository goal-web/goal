package exceptions

import (
	"qbhy/contracts"
	"qbhy/events"
	"qbhy/logs"
	"qbhy/http"
)

type ExceptionHandler struct{}

func (handler ExceptionHandler) Handle(exception contracts.Exception) {
	switch e := exception.(type) {
	case events.EventException:
		logs.WithException(e).Info("事件报错啦")
	case http.HttpException:
		logs.WithException(e).Error("控制器报错啦")
		_ = e.Context.String(500, e.Error())
	default:
		logs.WithException(e).Info("默认异常")
	}
}
