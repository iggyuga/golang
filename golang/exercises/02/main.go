package main

import (
	"fmt"
)

func main() {
	//anonymous function expression
	myFunc := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	test, err := myFunc2(4)
	test2, _ := myFunc2(8)

	fmt.Println(myFunc(3))
	fmt.Println(myFunc2(4))
	fmt.Println(test, err)
	fmt.Println(test2)

}

func myFunc2(n int) (int, bool) {
	return n / 2, n%2 == 0
}
