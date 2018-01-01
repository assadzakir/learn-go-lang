package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 2.2
	b := 2

	fmt.Println("a is type", reflect.TypeOf(a), " and b is type ", reflect.TypeOf(b))

	// c := a + b  error invalid operation: a + b (mismatched types float64 and int)
	c := int(a) + b

	fmt.Println("c value", c)
	fmt.Println("c is type", reflect.TypeOf(c))

}
