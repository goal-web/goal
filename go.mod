module github.com/goal-web/goal

go 1.17

require (
	github.com/goal-web/container v0.1.4
	github.com/goal-web/contracts v0.1.24
	github.com/goal-web/pipeline v0.1.4
	github.com/goal-web/supports v0.1.12
	github.com/labstack/echo/v4 v4.5.0
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/apex/log v1.9.0 // indirect
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
