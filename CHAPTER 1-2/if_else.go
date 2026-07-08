package main

import "fmt"

func main() {

	a := 6
	b := 8
	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is greater than b")
	}

	more_ifls(1, 2)
	more_ifls(2, 1)
	more_ifls(1, 1)

}

func more_ifls(a int, b int) {
	if a < b {
		fmt.Println("A < B")
	} else if a == b {
		fmt.Println("A = B")
	} else {
		fmt.Println("A > B")
	}
}
