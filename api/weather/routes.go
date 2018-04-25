package weather

import (
	"strings"

	"../domain"
)

const defaultBasePath = "/v1/weather"

func (r *Resource) generateRoutes(basePath string) *domain.Routes {

	if basePath == "" {
		basePath = defaultBasePath
	}

	var baseRoutes = domain.Routes{
		domain.Route{
			Handler: r.HandleGetWeather,
			Method:  "GET",
			Path:    "/v1/weather",
		},
	}

	routes := domain.Routes{}

	for _, route := range baseRoutes {
		r := domain.Route{
			Handler: route.Handler,
			Method:  route.Method,
			Path:    strings.Replace(route.Path, defaultBasePath, basePath, -1),
		}
		routes = routes.Append(&domain.Routes{r})
	}

	r.routes = &routes
	return r.routes

}
