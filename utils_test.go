package utils_test

import (
	"atomicgo.dev/utils"
	"fmt"
)

func ExampleTernary() {
	fmt.Println(utils.Ternary(true, "a", "b"))
	fmt.Println(utils.Ternary(false, "a", "b"))

	// Output:
	// a
	// b
}

func ExamplePrettyJSON() {
	person := Person{Name: "John Doe", Age: 42}
	json, _ := utils.ToJSON(person)
	prettyJson, _ := utils.PrettyJSON(json)
	fmt.Println(prettyJson)

	// Output:
	// {
	//   "Name": "John Doe",
	//   "Age": 42
	// }
}
