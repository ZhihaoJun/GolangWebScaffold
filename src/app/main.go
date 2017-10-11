package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/boj/redistore"
	"github.com/labstack/echo"
	sessions "github.com/utahta/echo-sessions"
)

type App struct {
	addr      string
	secretKey string

	e          *echo.Echo
	redisStore *redistore.RediStore
}

func (a *App) Hello(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func (a *App) SessionCount(c echo.Context) error {
	countStr := sessions.GetRaw(c, "count").(string)
	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 1
	}
	sessions.Set(c, "count", count)
	if err := sessions.Save(c); err != nil {
		return c.String(http.StatusInternalServerError, "Session save failed")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Page has been visited %d times", count))
}

func (a *App) loadEnv() {
	// ATTENTION: you should use your own prefix
	a.addr = ":1323"
	if v, ok := os.LookupEnv("APP_ADDR"); ok {
		a.addr = v
	}

	// ATTENTION: this is default secre key
	// you should change to a random one
	a.secretKey = "secret-key"
	if v, ok := os.LookupEnv("APP_SESSION_SECRET"); ok {
		a.secretKey = v
	}
}

func (a *App) Setup() {
	a.loadEnv()
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte(a.secretKey))
	if err != nil {
		panic(err)
	}
	a.redisStore = store

	a.e = echo.New()
	// ATTENTION: you should use your own session id prefix
	a.e.Use(sessions.Sessions("APP_SESS", store))

	// setup routers
	a.e.GET("/", a.Hello)
	a.e.GET("/count", a.SessionCount)
}

func (a *App) Run() {
	log.Println(a.e.Start(a.addr))
}

func (a *App) Close() {
	a.redisStore.Close()
}

func main() {
	a := &App{}
	a.Setup()
	defer a.Close()
	a.Run()
}
