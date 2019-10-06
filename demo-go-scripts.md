# [2.0] TechCon Demo

> TO: make audiance trust Golang to be used in their business
## Go comunity

### The Popular
- Google search trends compare with Javascript: [Error 429 (Too Many Requests)](https://trends.google.com/trends/explore?cat=31&q=%2Fm%2F09gbxjr,%2Fm%2F07sbkfb,%2Fm%2F01t6b,%2Fm%2F07sbkfb)
- Git star/ PR/ Push: [GitHut 2.0](https://madnight.github.io/githut/#/pushes/2019/2)
- Top Popular Language - PYPL : [PYPL PopularitY of Programming Language index](http://pypl.github.io/PYPL.html)
- Top languages over time: [Projects \| The State of the Octoverse](https://octoverse.github.com/projects#languages)

### The wanted and Paid 
- Top Paid: [Stack Overflow Developer Survey 2019](https://insights.stackoverflow.com/survey/2019#top-paying-technologies)
- Most Love, Wanted Language: [Stack Overflow Developer Survey 2019](https://insights.stackoverflow.com/survey/2019#most-loved-dreaded-and-wanted)

---
## Go is Influenced by Limbo, Oberon, C, Alef, Pascal, Modula, Python
- [Programming Languages Influence Network \| Exploring Data](https://exploring-data.com/vis/programming-languages-influence-network/#Go)
  
---
## Why Go is easy to learn

### Easy to use documents
- [The Tour of Go](https://tour.golang.org) (online/ offline)
  - Showing several method and channel to learn go effectivelly
  - Provide simple, clear and straight forwad documents beside a brower IDE environment for practise imediatly (like code academey) lead you go accoss all go language feature/ ability/ component.
  
  Demo:
  - 1. Go to [Pakage demo](https://tour.golang.org/basics/1)
    ```go
    package main

    import (
      "fmt"
      "math/rand"
    )

    func main() {
      for i := 0; i < 5; i++ {
        fmt.Println("My favorite number is", rand.Intn(100))
      }
    }
    ```
  - 2. Magic button
    ![d060106b.png](:storage\f9dee581-3294-4d6b-94e5-ef536c303502\d060106b.png)

 

## Static Type
- [Go Types Document](https://golang.org/ref/spec#Types)

### Basic Types
- string
- bool
- numeric : 
  - int, int8, uint8, int16, uint16...
  - float32, float64

[Demo Basic Type](https://tour.golang.org/basics/11)

### Composite Types
- [pointer](https://tour.golang.org/moretypes/1)
- [struct](https://tour.golang.org/moretypes/5)
- [function type](https://tour.golang.org/moretypes/24) (First-Class type)
- Container types: [Array](https://tour.golang.org/moretypes/6), slice type, [map type](https://tour.golang.org/moretypes/19)
```go
package main

import "fmt"

var m = map[string]int {"one" : 1, "two" : 2, "three" : 3}

func main() {
	fmt.Println(m["two"])
}
```
- Channel Type: 
```go
TBD
```
- Interface Type
```TDB
TDB
```

### Function is First-Class Type
- This means the language supports passing functions as arguments to other functions, returning them as the values from other functions, and assigning them to variables or storing them in data structures


## Feel like Dynamic Type Language

### Short variable declarations   
```go
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
``` 