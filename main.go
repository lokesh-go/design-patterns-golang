package main

import (
	"fmt"

	factory "github.com/lokesh-go/design-patterns-golang/factory"
)

func main() {
	var input string
	fmt.Print("Enter database name: ")
	fmt.Scan(&input)

	db := factory.Database(input)
	fmt.Println(db.GetClient())
}
