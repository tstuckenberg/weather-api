package server

import (
	"fmt"

	"../domain"
	"github.com/gin-gonic/gin"
)

// Router type for server
type Router struct{ *gin.Engine }

// NewRouter Returns a new Router object
func NewRouter() *Router {

	router := gin.New()
	return &Router{router}

}

//AddRoutes adds routes to router
func (router *Router) AddRoutes(routes *domain.Routes) *Router {
	if routes == nil {
		return router
	}
	for _, route := range *routes {

		router.Handle(route.Method, route.Path, route.Handler)
	}
	return router
}

//AddResources adds resources to router
func (router *Router) AddResources(resources ...domain.IResource) *Router {
	for _, resource := range resources {
		if resource.Routes() == nil {
			// server/router instantiation error
			// its safe to throw panic here
			// panic(errors.New(fmt.Sprintf("Routes definition missing: %v", resource)))
			panic(fmt.Errorf("Routes definition missing: %v", resource))

		}
		router.AddRoutes(resource.Routes())
	}
	return router
}
