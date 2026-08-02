package app

import "log"

type App struct {
	container *Container
}

func NewApp(container *Container) *App {
	return &App{
		container: container,
	}
}

func (a *App) Run() error {
	go func() {
		if err := a.container.GRPCServer.Start("50051"); err != nil {
			log.Fatal(err)
		}
	}()

	//a.container.GRPCServer.Stop()
	return nil
}
