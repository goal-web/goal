package session

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/pipeline"
)

func StartSession(session contracts.Session, request contracts.HttpRequest, next pipeline.Pipe) interface{} {
	session.Start()
	return next(request)
}
