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

	score = 50 + 10
	fmt.Println(score)

	score = score + 10
	fmt.Println(score)
}

///ВАЖНО! В GoLang очень строгая типизация, и переменная записанная как текст не может быть изменена в число
///Пример:
///text:= "Game Over"
///text=30
///ЭТО ОШИБКА
