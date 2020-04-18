package main

import "fmt"

func main() {
	amount := 6
	//передает аргумент функции
	double(amount)
	//выводит исходное значение
	fmt.Println(amount)
}

//парамеьр настраивает копию аргумента
func double(number int) {
	//изменяет скопированное, а не исходное значение
	number *= 2
}
