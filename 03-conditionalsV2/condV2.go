package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//запросить у пользователя значение
	fmt.Print("Enter a grade: ")

	//создать буферизиованное средство чтения
	reader := bufio.NewReader(os.Stdin)

	//возвращает введенный пользователем текст до enter
	//метод возвращает 2 значения,
	//для 2 переменной возвращенный код ошибки сохраняется в переменной
	input, err := reader.ReadString('\n')

	//сообщить об ошибке и остановить программу
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
