package app

import (
	"fmt"
	"log"
	"middleware/internal/app/endpoint"
	"middleware/internal/app/mw"
	"middleware/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {

	app := &App{}

	app.s = service.New()

	app.e = endpoint.New(app.s)

	app.echo = echo.New()

	app.echo.Use(mw.RoleCheck)

	app.echo.GET("status", app.e.Status)

	return app, nil

}

func (app *App) Run() error {

	fmt.Println("Server started")

	err := app.echo.Start(":8080")

	if err != nil {

		log.Fatal(err)
	}

	return nil
}
