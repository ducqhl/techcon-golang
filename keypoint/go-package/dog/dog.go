package dog

import (
	"fmt"
)

// Dog object
type Dog struct{}

// Sound as a Dog
func (dog Dog) Sound() { fmt.Println("Woaw Woaw") }

// Move as a Dog
func (dog Dog) Move()  { fmt.Println("Dog Walk") }