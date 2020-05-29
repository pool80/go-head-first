package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//файл данных открывается для чтения
	file, err := os.Open("C:\\Users\\maxim\\go\\src\\HeadFirst\\06-array\\04-text-file\\data.txt")
	//если при открытии файла произошла ошибка, сообщить о ней и выйти
	if err != nil {
		log.Fatal(err)
	}
	//для файла создается новое значение Scanner
	scanner := bufio.NewScanner(file)
	//цикл выполняется до тех пор пока не будет достигнуть конец файла, а scanner.Scan не вернет конец файла
	//читает строку из файла
	for scanner.Scan() {
		//выводит строку
		fmt.Println(scanner.Text())
	}
	//закрывает файл для освобождения ресурсов
	err = file.Close()
	//если при завершении файла произошла ошибка сообщить о ней и завершить программу
	if err != nil {
		log.Fatal(err)
	}
	//если при сканировании файла произошла ощибка, сообщить о ней и завершить работу
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
