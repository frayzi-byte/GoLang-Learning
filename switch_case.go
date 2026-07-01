package main

import "fmt"

func main() {
	var a int
	a = 30
	switch a {

	case 30:
		fmt.Println("A = 30")
		fallthrough
	case 20:
		fmt.Println("A = 20")
		fallthrough
	case 10:
		fmt.Println("A = 10")
		fallthrough
	case 5, 6, 7:
		fmt.Println("A = 5/6/7")
		fallthrough
	default:
		fmt.Println("/A/ is undefined")

	}
}
