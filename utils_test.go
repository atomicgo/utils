package utils_test

import (
	"fmt"

	"atomicgo.dev/utils"
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
	prettyJSON, _ := utils.PrettyJSON(json)
	fmt.Println(prettyJSON)

	// Output:
	// {
	//   "Name": "John Doe",
	//   "Age": 42
	// }
}
