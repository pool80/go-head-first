package main

import (
	"fmt"
	"time" //импортируем пакет time
)

func main() {
	//метод time.Now возвращает значение time.Time
	var now time.Time = time.Now()

	//у значений time.Time есть метод Year, который возвращает текущий год
	var year int = now.Year()
	fmt.Println(year)

}
