package session

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/http"
)

type RequestBeforeListener struct {
}

func (this *RequestBeforeListener) Handle(event contracts.Event) {
}

type RequestAfterListener struct {
}

// Handle 如果开启了 session 那么请求结束时保存 session
func (this *RequestAfterListener) Handle(event contracts.Event) {
	if requestAfterEvent, ok := event.(*http.RequestAfter); ok {
		if session, isSession := requestAfterEvent.Request().Get("session").(contracts.Session); isSession {
			if session.IsStarted() {
				session.Save()
			}
		}
	}
}
