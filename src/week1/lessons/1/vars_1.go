package main

import "fmt"

func main() {

	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 = 1

	// пропуск типа
	var num2 = 20

	fmt.Println(num0, num1, num2)

	// короткое объявление переменной
	num := 30
	// только для новых переменных
	// var := 31

	num += 1
	fmt.Println("+=", num)

	num++
	fmt.Println("++", num)

	// camelCase - принятый стиль
	userInsex := 10
	// under_score - не принято
	user_index := 10
	fmt.Println(userInsex, user_index)

	// объявление нескольких переменных
	var weight, height = 11, 21

	// короткое объявление
	// хотя-бы одна меременная должна быть новой
	weight, age := 12, 22

	fmt.Println(weight, height, age)
}
