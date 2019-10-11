package main

import "fmt"

// Animal Interface
type Animal interface {
	sound()
	move()
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}

	dog.sound()
	dog.move()

	fmt.Println()

	cat.sound()
	cat.move()
}

/******************************
* Animal Like Structs
*******************************/

// Cat struct
type Cat struct{}

// Dog struct
type Dog struct{}

/**********************************
*  Implement Animal Interface
***********************************/

func (dog Dog) sound() { fmt.Println("Woaw Woaw") }
func (dog Dog) move()  { fmt.Println("Dog Walk") }

func (cat Cat) sound() { fmt.Println("Mew Mew") }
func (cat Cat) move()  { fmt.Println("Cat walk") }
