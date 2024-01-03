package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"middleware/internal/app/MW"
	"middleware/internal/app/Service"
	"middleware/internal/app/endpoint"
)

type App struct {
	e    *endpoint.Endpoint
	s    *Service.Service
	echo interface{}
}

func New() (*App, error) {
	a := &App{}
	a.s = Service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(MW.RoleCheck)

	a.echo.GET("/status", a.e.Status)
	return a, nil

}

func (a *App) Run() error {
	fmt.Println("Server running")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
