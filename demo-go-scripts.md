# [2.0] TechCon Demo

> TO: make audiance trust Golang to be used in their business
## Go Compile
Hello World Example
```bash
go build keypoint/hello-world/main.go
./main
```
## Go OOP
The way golang approad OOP is different from other language as Java or C#.
```bash
go run ./keypoint/go-oop/main.go
```

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
- [Channel Type](https://tour.golang.org/concurrency/3)

### Function is First-Class Type
- This means the language supports 
  - passing functions as arguments 
  - returning them as the values from other function
  - assigning them to variables or storing them in data structures


## Go Package
```bash
go run ./keypoint/go-package/src/main.go
```
## Go Library
[go-search](https://go-search.org/)

### Short variable declarations   
```go
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
``` 

## Go Concurrency
[Concurrency Example](https://tour.golang.org/concurrency/1)
