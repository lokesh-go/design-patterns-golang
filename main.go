package main

import (
	"fmt"
	"reflect"
	"sync"

	abstractfactory "github.com/lokesh-go/design-patterns-golang/abstractfactory"
	factory "github.com/lokesh-go/design-patterns-golang/factory"
	singleton "github.com/lokesh-go/design-patterns-golang/singleton"
)

func main() {
	//factoryExample()
	//singletonExample()
	abstractFactoryExample()
}

func factoryExample() {
	var input string
	fmt.Print("Enter database name: ")
	fmt.Scan(&input)

	// Calls factory design patterns
	db := factory.New().Database(input)
	fmt.Println(db.GetClient())
}

func singletonExample() {
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singleton.GetInstance()
		}()
	}
	wg.Wait()
}

func abstractFactoryExample() {
	dbFactory := abstractfactory.AbstractFactory("filesystem")
	fmt.Println("ABS: ", reflect.TypeOf(dbFactory))
}
