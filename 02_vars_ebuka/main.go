package main

import "fmt"

func main() {

	var age = 20
	var isHot = true

	name := "Ebuka"
	size := 2.0
	name, email := "Ebuka", "ebulamicheal@gmail.com"

	fmt.Println(name, age, isHot, size, email)

	fmt.Printf("%T\n", size)
}
