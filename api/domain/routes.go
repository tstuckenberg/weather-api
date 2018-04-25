package domain

import "github.com/gin-gonic/gin"

// Route type
// Note that DefaultVersion must exists in RouteHandlers map
// See routes.go for examples
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// Routes type
type Routes []Route

// Append Returns a new slice of Routes
func (r *Routes) Append(routes ...*Routes) Routes {
	res := Routes{}
	// copy current route
	for _, route := range *r {
		res = append(res, route)
	}
	for _, _routes := range routes {
		for _, route := range *_routes {
			res = append(res, route)
		}
	}
	return res
}
