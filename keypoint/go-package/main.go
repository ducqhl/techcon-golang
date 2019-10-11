package main

import (
	"fmt"
	"./animal"
	"./dog"
)

func main() {
	var d animal.Animal = dog.Dog{}
	var c animal.Animal = Cat{}

	d.Sound()
	d.Move()

	fmt.Println()

	c.Sound()
	c.Move()
}

/******************************
* Animal Like Structs
*******************************/

// Dog object
type Dog struct{}

// Cat struct
type Cat struct{}

/**********************************
*  Implement Animal Interface
***********************************/

func (cat Cat) Sound() { fmt.Println("Mew Mew") }
func (cat Cat) Move()  { fmt.Println("Cat walk") }
