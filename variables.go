package main

import "fmt"

func main() {
	text := "Game Over"
	score := 111

	fmt.Println(score)
	fmt.Println(text)

	text = "Game Paused"
	score = 50

	fmt.Println(score)
	fmt.Println(text)

	print_user("Andrew", 35)
}

func print_user(name string, age int) {
	fmt.Println(name, age)
}

///ВАЖНО! В GoLang очень строгая типизация, и переменная записанная как текст не может быть изменена в число
///Пример
/*
	text:= "Game Over"
	text=30
*/
///ЭТО ОШИБКА
