package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/jobs"
	"github.com/golang-module/carbon/v2"
	"time"
)

func DemoJob(queue contracts.Queue, request contracts.HttpRequest) string {
	var err = queue.Push(jobs.NewDemo(request.GetString("info")))
	if err != nil {
		return err.Error()
	}

	err = queue.Later(
		time.Now().Add(time.Second*time.Duration(request.GetInt("delay"))),
		jobs.NewDemo("delay+"+request.GetString("info")),
	)
	if err != nil {
		return err.Error()
	}

	return carbon.Now().String()
}
