package main

import (
	"fmt"
	"math"
)

func main() {

	// *** INTEGERS ***
	var x = 1
	var x32 int32 = 1
	var x64 int64 = 1
	fmt.Println(x, x32, x64)
	// to actually print the types of ints
	fmt.Printf("%T, %T, %T \n", x, x32, x64)

	// math operations
	y := 100
	fmt.Println(
		y-x,
		y+x,
		y/x,
		y%x,
		y == x,
		y < x,
		y > x,
	)

	// *** FLOATING POINT NUMBERS ***
	pi := 3.1416
	var pi32 float32 = 3.1416
	fmt.Printf("%T, %T \n", pi, pi32)

	// scientific notation
	sci := 1e100
	fmt.Println(sci)

	// *** COMPLEX NUMBERS ***
	comp := complex(1, 2)
	fmt.Println(comp, "complex number")

	// *** MATH ***
	z := 3.123456789
	fmt.Println(math.Ceil(z), "Ceil")
	fmt.Println(math.Min(pi, z), "Min")
	fmt.Println(math.Max(pi, z), "Max")
	fmt.Println(math.Abs(z), "Abs")
	fmt.Println(math.Sqrt(z), "Sqrt")
	fmt.Println(math.Pow(z, 100), "Pow")
	fmt.Println(math.Floor(z), "Floor")

}
