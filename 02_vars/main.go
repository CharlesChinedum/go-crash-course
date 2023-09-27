package main

import "fmt"

func main() {
	// fmt.Println("Hello World! Let's GO!!")
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Leo"
	var age int32 = 21
	const isCool = true

	// Shorthand
	// name := "Leo"
	// size := 1.3

	// Multiple vars
	name, email, number := "Leo", "charleschinedum2@gmail.com", 12

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", number)
}
