module github.com/goal-web/goal

go 1.17

require (
	github.com/goal-web/contracts v0.1.27
	github.com/goal-web/database v0.0.0-00010101000000-000000000000
	github.com/goal-web/supports v0.1.12
	github.com/golang-jwt/jwt v3.2.2+incompatible
)

require (
	github.com/apex/log v1.9.0 // indirect
	github.com/goal-web/application v0.1.0 // indirect
	github.com/goal-web/container v0.1.4 // indirect
	github.com/goal-web/querybuilder v0.1.12 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/qbhy/parallel v1.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/goal-web/application => ../application
	github.com/goal-web/container => ../container
	github.com/goal-web/contracts => ../contracts
	github.com/goal-web/database => ../database
	github.com/goal-web/session => ../session
	github.com/goal-web/http => ../http
	github.com/goal-web/supports => ../supports
)
