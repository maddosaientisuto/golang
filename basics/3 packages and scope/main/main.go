package main

import (
	"fmt"
)

// 3 levels of scope in go
// block --> package --> universe
// *** PACKAGE ***
var package_lvl_num int = 99

// it can be called in any function since it's declared
// at package level

func main() {

	// *** BLOCK ***
	// means within curly brackets

	var blocknum int = 10
	// we can access it from within same block(curly braces)
	fmt.Println(blocknum)
	// introduce another block
	{
		var otherblocknum int = 100
		// we can access it within this block only
		fmt.Println(otherblocknum, "called within same block")
	}
	// we can't access otherblocknum outside of that block
	// fmt.Println(otherblocknum)

	// the package level var can be accessed from here
	fmt.Println(package_lvl_num, "called from main")

	// good thing about package level vars is that they can be accessed
	// from other files too if they have same package name
	// fmt.Println(other.Othernum, "from other file")

	nonMain()
}

func nonMain() {
	fmt.Println(package_lvl_num, "called from nonMain function")
}
