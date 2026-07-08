package main

import "fmt"

func main() {
	first_pointer()
	second_pointer()
}

func first_pointer() {
	var x int = 4
	var p *int
	p = &x
	fmt.Println(*p)
	*p = 45
	fmt.Println(*p)
}

func second_pointer() {
	var x int = 4
	var p *int
	p = &x
	fmt.Println("Adress : ", p)
	fmt.Println("Value : ", *p)
}
