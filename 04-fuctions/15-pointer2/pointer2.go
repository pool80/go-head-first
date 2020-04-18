package main

import "fmt"

func main() {
	amount := 6
	//вместо значения переменной передается указатель на нее
	double(&amount)
	fmt.Println(amount)
}

//вместо значения int получает указатель
func double(number *int) {
	//обновляет значение на которое ссылается указатель
	*number *= 2
}
