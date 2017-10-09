package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
)

type App struct {
	addr string

	e *echo.Echo
}

func (a *App) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (a *App) Setup() {
	a.addr = ":1323"
	if v, ok := os.LookupEnv("APP_ADDR"); ok {
		a.addr = v
	}
	a.e = echo.New()

	// setup routers
	a.e.GET("/", a.HelloWorld)
}

func (a *App) Run() {
	log.Println(a.e.Start(a.addr))
}

func (a *App) Close() {

}

func main() {
	a := &App{}
	a.Setup()
	defer a.Close()
	a.Run()
}
