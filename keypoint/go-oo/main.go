package main

import "fmt"

func main() {
	cat := Cat{"Leo", "Meow", "Move"}
	cat.Move()
}

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
