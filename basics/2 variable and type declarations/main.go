package main

import "fmt"

func main() {
	// *** VARIABLE DECLARATIONS ***
	// var name type = expression
	var number int = 6
	// we can either omit type or expression
	// but not both
	// if no expression is assigned, a zero
	// value is assigned which for ints is 0,
	// for bools is False
	var number2 int
	var number3 = 10

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)

	// multiple vars of same type can be handled this way
	var num1, num2, num3 int
	fmt.Println(num1, num2, num3)

	// to assign multiple varibales of different types
	// in single line, you have to use respective expressions
	var num4, num5 = 100, "Nine Thousand"
	fmt.Println(num4, num5)

	// short declarations
	// name := expression
	bool1 := true
	fmt.Println(bool1)

	// *** POINTERS ***
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p) // accesses the contents of the pointer
	// for long declaration
	var x2 int = 111
	var p2 *int = &x2
	fmt.Println(p2, *p2)

	// *** TYPE DECLARATIONS ***
	type fahrenheit float32
	type celsius float32

	var f fahrenheit = 32.0
	var c celsius = 0.0

	// f = c gives type error because they are of different types
	f = fahrenheit(c)

	fmt.Println(f, c)
}
