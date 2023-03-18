package factory

import (
	cachePkg "github.com/lokesh-go/design-patterns-golang/factory/pkg/cache"
	mongodbPkg "github.com/lokesh-go/design-patterns-golang/factory/pkg/mongodb"
)

// methods ...
type methods interface {
	GetClient() string
}

type db struct{}

func new() methods {
	return &db{}
}

// GetClient: default response
func (d *db) GetClient() (res string) {
	return "Invalid Input"
}

// Database ...
func Database(database string) methods {
	switch database {
	case "mongodb":
		{
			return mongodbPkg.New()
		}
	case "cache":
		{
			return cachePkg.New()
		}
	default:
		{
			return new()
		}
	}
}
