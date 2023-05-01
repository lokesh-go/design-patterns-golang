package factory

import (
	cachePkg "github.com/lokesh-go/design-patterns-golang/factory/pkg/cache"
	jsonPkg "github.com/lokesh-go/design-patterns-golang/factory/pkg/file/json"
	txtPkg "github.com/lokesh-go/design-patterns-golang/factory/pkg/file/txt"
	mongodbPkg "github.com/lokesh-go/design-patterns-golang/factory/pkg/mongodb"
)

type factory struct{}

// Factories ...
type Factories interface {
	Database(string) dbMethods
	FileSystem(string) fileMethods
}

// New ...
func New() Factories {
	return &factory{}
}

// Database ...
func (f *factory) Database(database string) dbMethods {
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
			return newdb()
		}
	}
}

// FileSystem ...
func (f *factory) FileSystem(extension string) fileMethods {
	switch extension {
	case "json":
		{
			return jsonPkg.New()
		}
	case "txt":
		{
			return txtPkg.New()
		}
	default:
		{
			return newfile()
		}
	}
}
