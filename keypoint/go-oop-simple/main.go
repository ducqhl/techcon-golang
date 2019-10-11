package main

import "fmt"

// Cat struct
type Cat struct {
	name  string
	speak string
	move  string
}

// Move a Cat
func (cat Cat) Move() {
	fmt.Println(cat.move)
}

func main() {
	cat := Cat{"Leo", "Meow", "Move"}
	cat.Move()
}
