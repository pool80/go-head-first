package main

import (
	"fmt"
)

func main() {
	//функция вставляет значение в форматную строку, но вместо вывода или возвращения строки
	//она возвращает значение ошибки
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error())
	fmt.Println(err)
	//
}
