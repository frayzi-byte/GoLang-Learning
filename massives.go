package main

import "fmt"

func main() {
	mass()
	mass_2([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	mass_3()
}

func mass() {
	var numbers [10]int = [10]int{1, 2}
	fmt.Println(numbers)
	numbers[0] = 30
	fmt.Println("/", numbers[0], "/", "NUMBERS[0]")
}

func mass_2(numbers_2 [10]int) {
	fmt.Println(numbers_2)
	fmt.Println("/", numbers_2[5], "/", "NUMBERS_2[5]")
}

func mass_3() {
	var numbers_3 [3][3]int
	numbers_3 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(numbers_3)
}
