package providers

import (
	"github.com/goal-web/contracts"
)

type App struct {
	path string
}

func NewApp(path string) contracts.ServiceProvider {
	return &App{path}
}

func (app App) Register(instance contracts.Application) {

}

func (app App) Start() error {
	return nil
}

func (app App) Stop() {
}
