package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"./server"

	"./weather"
)

//App definition
type App struct {
	Router *server.Router
}

//Initialize App
func (a *App) Initialize() {

	a.Router = server.NewRouter()

	a.Router.Use(gin.Logger())

	a.initializeRoutes()

}

//Run App
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

//Initialize routes
func (a *App) initializeRoutes() {

	weatherResource := weather.NewResource(&weather.Options{})

	a.Router.AddResources(weatherResource)

}
