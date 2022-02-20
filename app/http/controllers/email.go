package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/email"
)

func SendEmail(request contracts.HttpRequest, mailer contracts.Mailer) string {
	mail := email.New("测试邮件", email.Text(request.GetString("content"))).SetTo(request.GetString("to"))

	if request.GetString("queue") != "" {
		mail.Queue(request.GetString("queue"))
	}

	var err = mailer.Send(mail)

	if err != nil {
		return err.Error()
	}

	return "ok"
}
