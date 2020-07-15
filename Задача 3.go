
package main

import (
	"fmt"
)
func main() {
	a := 42
	b := 153
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

