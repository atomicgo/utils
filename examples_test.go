package utils_test

import (
	"atomicgo.dev/utils"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func ExampleToJSON() {
	var person = Person{"John Doe", 42}

	fmt.Println(utils.ToJSON(person))
}
