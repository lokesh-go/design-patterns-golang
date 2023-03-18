package main

import (
	"fmt"
	"sync"

	factory "github.com/lokesh-go/design-patterns-golang/factory"
	singleton "github.com/lokesh-go/design-patterns-golang/singleton"
)

func main() {
	factoryExample()
	singletonExample()
}

func factoryExample() {
	var input string
	fmt.Print("Enter database name: ")
	fmt.Scan(&input)

	// Calls factory design patterns
	db := factory.Database(input)
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
