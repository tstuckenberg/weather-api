package weather

import (
	"../domain"
)

//Options struct
type Options struct {
	BasePath string
	// Database *gorm.DB
	// Database *sql.DB
}

// NewResource create new resource
func NewResource(options *Options) *Resource {

	// database := options.Database

	// if database == nil {
	// 	panic("weather.Database is required")
	// }

	// u := &Resource{options, nil, database}
	u := &Resource{options, nil}
	u.generateRoutes(options.BasePath)
	return u
}

// Resource type
type Resource struct {
	options *Options
	routes  *domain.Routes
	// Database *gorm.DB
	// Database *sql.DB
}

// Routes method returns routes
func (resource *Resource) Routes() *domain.Routes {
	return resource.routes
}
