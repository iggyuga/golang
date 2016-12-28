package main

import (
	"fmt"
)

func main() {

	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())

}

func foo(data ...int) string {
	return "hello from foo"
}
