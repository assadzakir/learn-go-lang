package main

import (
	"fmt"
	"reflect"
)

/*
	Variables

	-Statically typed

	-Use var keyword if declaring at package level

	-Name must start with "_" or a letter

	-Go will initialize with a zero value

	-Package level variable are global

	-Shorthand declare and initialize ":=" only works inside of function

	-Variables declare at the function level must be used

*/
var (

	/* 	---------------------------------------
	   	 name, course string
	   	 module float64
	   	----------------------------------------
	   	 name, course, module = "Assad", "GO", 2.2
	   	---------------------------------------- */
	name   = "Assad"
	course = "Go"
	module = 2.2
)

func main() {
	// fmt.Println("Name is set to ", name)
	// fmt.Println("Module is set to ", module)
	fmt.Println("Name is ", reflect.TypeOf(name))
	fmt.Println("Module is ", reflect.TypeOf(module))
}
