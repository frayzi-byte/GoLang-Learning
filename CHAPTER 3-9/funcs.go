package main

import "fmt"

func main() {
	g := 30
	add(1, 2, 1.23, 3.45, 6.56)
	add2(g)
}

func add(x, y int, a, b, c float32) {
	var d = a + b + c
	var z = x + y
	fmt.Println("x + y =", z)
	fmt.Println("a + b + c =", d)
}

func add2(t int) {
	fmt.Println("t before :", t)
	t = t + 20
	fmt.Printf("t after :", t)
}
