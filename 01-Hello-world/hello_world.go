// The Go Language is case sensitive.

/* 
 package main
-Executable Program
-Entry point 
*/
package main

import (
	"fmt"
	"runtime"
)
/* 
func main()
-Entry point
-Only for executables
-can't remain "main"
-takes no arguments
-No return values 
*/
func main() {
	fmt.Println("Hello from", runtime.GOOS)
}


