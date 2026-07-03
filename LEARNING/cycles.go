package main

import "fmt"

func main() {

	nums := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for i, value := range nums {
		fmt.Println(i, value)
	}
}
