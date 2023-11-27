<h1 align="center">AtomicGo | utils</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Futils&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/utils/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/utils?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/utils" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/utils/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/utils" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/utils?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/utils">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-6-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/utils" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/utils?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/utils#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/utils</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# utils

```go
import "atomicgo.dev/utils"
```

Package utils is a collection of useful, quickly accessible utility functions.

## Index

- [func AppendToFile\[T string | \[\]byte\]\(path string, content T\) error](<#AppendToFile>)
- [func DownloadFile\(url, path string\) error](<#DownloadFile>)
- [func Fetch\(url string\) \(string, error\)](<#Fetch>)
- [func FileExists\(path string\) bool](<#FileExists>)
- [func PrettyJSON\(inputJSON string, indent ...string\) \(string, error\)](<#PrettyJSON>)
- [func ReadFile\(path string\) \(string, error\)](<#ReadFile>)
- [func Ternary\[T any\]\(condition bool, a, b T\) T](<#Ternary>)
- [func ToInt\[T string | constraints.Number\]\(v T\) int](<#ToInt>)
- [func ToJSON\(v any\) \(string, error\)](<#ToJSON>)
- [func ToPrettyJSON\(v any, indent ...string\) \(string, error\)](<#ToPrettyJSON>)
- [func ToString\(v any\) string](<#ToString>)
- [func WriteFile\[T string | \[\]byte\]\(path string, content T\) error](<#WriteFile>)


<a name="AppendToFile"></a>
## func [AppendToFile](<https://github.com/atomicgo/utils/blob/main/file.go#L27>)

```go
func AppendToFile[T string | []byte](path string, content T) error
```

AppendToFile appends the given content to the given file. Accepts a string or a byte slice.

<a name="DownloadFile"></a>
## func [DownloadFile](<https://github.com/atomicgo/utils/blob/main/file.go#L49>)

```go
func DownloadFile(url, path string) error
```

DownloadFile downloads the given URL to the given path. If the file already exists, it will be overwritten.

<a name="Fetch"></a>
## func [Fetch](<https://github.com/atomicgo/utils/blob/main/utils.go#L34>)

```go
func Fetch(url string) (string, error)
```

Fetch returns the body of a GET request to the given URL.

<a name="FileExists"></a>
## func [FileExists](<https://github.com/atomicgo/utils/blob/main/file.go#L42>)

```go
func FileExists(path string) bool
```

FileExists returns true if the given file exists.

<a name="PrettyJSON"></a>
## func [PrettyJSON](<https://github.com/atomicgo/utils/blob/main/utils.go#L20>)

```go
func PrettyJSON(inputJSON string, indent ...string) (string, error)
```

PrettyJSON returns a pretty\-printed JSON string. If indent is not provided, it defaults to " " \(two spaces\).

<details><summary>Example</summary>
<p>



```go
person := Person{Name: "John Doe", Age: 42}
json, _ := utils.ToJSON(person)
prettyJson, _ := utils.PrettyJSON(json)
fmt.Println(prettyJson)

// Output:
// {
//   "Name": "John Doe",
//   "Age": 42
// }
```

#### Output

```
{
  "Name": "John Doe",
  "Age": 42
}
```

</p>
</details>

<a name="ReadFile"></a>
## func [ReadFile](<https://github.com/atomicgo/utils/blob/main/file.go#L11>)

```go
func ReadFile(path string) (string, error)
```

ReadFile reads the given file and returns its content.

<a name="Ternary"></a>
## func [Ternary](<https://github.com/atomicgo/utils/blob/main/utils.go#L11>)

```go
func Ternary[T any](condition bool, a, b T) T
```

Ternary is a ternary operator. It returns a if the condition is true, otherwise it returns b.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.Ternary(true, "a", "b"))
	fmt.Println(utils.Ternary(false, "a", "b"))

}
```

#### Output

```
a
b
```

</p>
</details>

<a name="ToInt"></a>
## func [ToInt](<https://github.com/atomicgo/utils/blob/main/to.go#L39>)

```go
func ToInt[T string | constraints.Number](v T) int
```

ToInt converts the given value to an int. If the value is a float, it will be rounded to the nearest integer. \(Rounds up if the decimal is 0.5 or higher\)

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.ToInt(1337))
	fmt.Println(utils.ToInt(1337.4))
	fmt.Println(utils.ToInt(1337.5))
	fmt.Println(utils.ToInt(1337.6))
	fmt.Println(utils.ToInt("1337"))
	fmt.Println(utils.ToInt("1337.4"))
	fmt.Println(utils.ToInt("1337.5"))
	fmt.Println(utils.ToInt("1337.6"))

}
```

#### Output

```
1337
1337
1338
1338
1337
1337
1338
1338
```

</p>
</details>

<a name="ToJSON"></a>
## func [ToJSON](<https://github.com/atomicgo/utils/blob/main/to.go#L12>)

```go
func ToJSON(v any) (string, error)
```

ToJSON converts the given value to a JSON string.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/utils"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var person = Person{"John Doe", 42}

	json, _ := utils.ToJSON(person)
	fmt.Println(json)

}
```

#### Output

```
{"Name":"John Doe","Age":42}
```

</p>
</details>

<a name="ToPrettyJSON"></a>
## func [ToPrettyJSON](<https://github.com/atomicgo/utils/blob/main/to.go#L20>)

```go
func ToPrettyJSON(v any, indent ...string) (string, error)
```



<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/utils"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John Doe", Age: 42}
	prettyJson, _ := utils.ToPrettyJSON(person)
	fmt.Println(prettyJson)

}
```

#### Output

```
{
  "Name": "John Doe",
  "Age": 42
}
```

</p>
</details>

<a name="ToString"></a>
## func [ToString](<https://github.com/atomicgo/utils/blob/main/to.go#L33>)

```go
func ToString(v any) string
```

ToString converts the given value to a string.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"atomicgo.dev/utils"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"John Doe", 42}

	fmt.Println(utils.ToString(person))

}
```

#### Output

```
{John Doe 42}
```

</p>
</details>

<a name="WriteFile"></a>
## func [WriteFile](<https://github.com/atomicgo/utils/blob/main/file.go#L21>)

```go
func WriteFile[T string | []byte](path string, content T) error
```

WriteFile writes the given content to the given file. Accepts a string or a byte slice.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
