package main

import "fmt"

func main() {
	cycle(1, 2, 3)
	f := multiply
	fmt.Println(f(10, 10))
	a := display
	a("SUCCESS")

}

func cycle(numbers ...int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum =", sum)
}

func multiply(x int, y int) int { return x * y }

func display(message string) {
	fmt.Println(message)
}
