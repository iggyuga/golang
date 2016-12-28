package main

import "fmt"

func main() {

	fmt.Println(myFunc(3))

}

func myFunc(a int) (int, bool) {
	if a%2 == 0 {
		return (a / 2), true
	}
	return (a / 2), false

}
