# Golang - Go Higher Demo

### Object Oriented Language Y/N
- Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. 
- The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general.

```go
package main

import "fmt"

type Animal interface {
    sound()
    move()
}

type Dog struct {}
type Cat struct {}

func (dog Dog) Animal {
    
}
```
