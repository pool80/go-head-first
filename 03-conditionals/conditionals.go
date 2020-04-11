package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//запросить у пользователя значение
	fmt.Print("Enter a grade: ")

	//создать буферизиованное средство чтения
	reader := bufio.NewReader(os.Stdin)

	//возвращает введенный пользователем текст до enter
	//метод возвращает 2 значения,
	//для 2 переменной используется пустой идентификатор как заполнитель для кода ошибки
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
