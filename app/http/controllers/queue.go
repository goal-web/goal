package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/jobs"
	"github.com/golang-module/carbon/v2"
	"time"
)

func DemoJob(queue contracts.Queue, request contracts.HttpRequest) any {
	var err = queue.Push(jobs.NewDemo(request.GetString("info")))
	if err != nil {
		return contracts.Fields{
			"error": err.Error(),
		}
	}

	err = queue.Later(
		time.Now().Add(time.Second*time.Duration(request.GetInt("delay"))),
		jobs.NewDemo("delay+"+request.GetString("info")),
	)
	if err != nil {
		return contracts.Fields{
			"error": err.Error(),
		}
	}

	return contracts.Fields{
		"now": carbon.Now().String(),
	}
}
