package main

import "fmt"

func main() {
	printing()
	fmt.Print("\n")
	printing_2(6, 7)
	fmt.Print("\n")
	fmt.Print("Success")
}

func printing() {
	m := 5
	var n int
	n = 4

	fmt.Print("m=", m, " ", "n=", n)
}

func printing_2(a int, b int) {
	fmt.Print(a, " ", b)
}
