package abstractfactory

import factory "github.com/lokesh-go/design-patterns-golang/factory"

// AbstractFactory ...
func AbstractFactory(fact string) interface{} {
	switch fact {
	case "database":
		{
			return factory.New().Database
		}
	case "filesystem":
		{
			return factory.New().FileSystem
		}
	default:
		{
			return nil
		}
	}
}
