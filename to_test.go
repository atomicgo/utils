package utils_test

import (
	"fmt"

	"atomicgo.dev/utils"
)

type Person struct {
	Name string
	Age  int
}

func ExampleToJSON() {
	person := Person{"John Doe", 42}

	json, _ := utils.ToJSON(person)
	fmt.Println(json)

	// Output:
	// {"Name":"John Doe","Age":42}
}

func ExampleToInt() {
	fmt.Println(utils.ToInt(1337))
	fmt.Println(utils.ToInt(1337.4))
	fmt.Println(utils.ToInt(1337.5))
	fmt.Println(utils.ToInt(1337.6))
	fmt.Println(utils.ToInt("1337"))
	fmt.Println(utils.ToInt("1337.4"))
	fmt.Println(utils.ToInt("1337.5"))
	fmt.Println(utils.ToInt("1337.6"))

	// Output:
	// 1337
	// 1337
	// 1338
	// 1338
	// 1337
	// 1337
	// 1338
	// 1338
}

func ExampleToString() {
	person := Person{"John Doe", 42}

	fmt.Println(utils.ToString(person))

	// Output:
	// {John Doe 42}
}

func ExampleToPrettyJSON() {
	person := Person{Name: "John Doe", Age: 42}
	prettyJSON, _ := utils.ToPrettyJSON(person)
	fmt.Println(prettyJSON)

	// Output:
	// {
	//   "Name": "John Doe",
	//   "Age": 42
	// }
}
