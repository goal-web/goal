module github.com/goal-web/goal

go 1.17

require (
	github.com/goal-web/contracts v0.1.24
	github.com/goal-web/http v0.1.0
	github.com/goal-web/pipeline v0.1.5
	github.com/goal-web/supports v0.1.12
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
)

require (
	github.com/apex/log v1.9.0 // indirect
	github.com/goal-web/container v0.1.4 // indirect
	github.com/labstack/echo/v4 v4.6.3 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
)

replace (
	github.com/goal-web/application => ../application
	github.com/goal-web/container => ../container
	github.com/goal-web/contracts => ../contracts
	github.com/goal-web/database => ../database
	github.com/goal-web/supports => ../supports
)
