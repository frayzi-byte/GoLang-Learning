package main

import "fmt"

func main() {
	var num int ///num := ... (int)
	num = 30
	fmt.Println(num)
	calc()
	compare(3, 3)
}

func calc() {
	a := 10
	b := 20
	clcpl := a + b
	clcmn := a - b
	clcmt := a * b
	clcdv := a / b
	fmt.Println(clcpl)
	fmt.Println(clcmn)
	fmt.Println(clcmt)
	fmt.Println(clcdv)
}

func compare(a int, b int) {
	var c bool = a == b
	fmt.Println(c)
}
