package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	abc()
	def(40, 50, 60)
}

func abc() {
	a := 10
	b := 20
	c := 30

	fmt.Println(a, b, c)
}

func def(d int, e int, f int) {
	fmt.Println(d, e, f)
}
