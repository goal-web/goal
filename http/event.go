package http

type HttpServeClosed struct {
}

func (this *HttpServeClosed) Event() string {
	return "HTTP_SERVE_CLOSED"
}
