package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	abstractfactory "github.com/lokesh-go/design-patterns-golang/abstractfactory"
	adapter "github.com/lokesh-go/design-patterns-golang/adapter.go"
	builder "github.com/lokesh-go/design-patterns-golang/builder"
	facad "github.com/lokesh-go/design-patterns-golang/facad"
	factory "github.com/lokesh-go/design-patterns-golang/factory"
	singleton "github.com/lokesh-go/design-patterns-golang/singleton"
)

func main() {
	//factoryExample()
	//singletonExample()
	//abstractFactoryExample()
	//builderExample()
	//adapterExample()
	facadeExample()
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

func builderExample() {
	mcaStudentBuilder := builder.NewStudentBuilder("mca")
	mcaStudentBuilder.SetCollege()
	mcaStudentBuilder.SetName()
	mcaStudentBuilder.SetRollNo()
	mcaStudentBuilder.SetStream()
	fmt.Println(mcaStudentBuilder)
}

func adapterExample() {
	adapter.AdapterExample()
}

func facadeExample() {
	newAcc := facad.NewPersonAccount("abc", 123)
	err := newAcc.CreditAmount("abc", 123, 1000)
	if err != nil {
		log.Fatal(err)
	}
	newAcc.Status()
}
