package main

import (
	"fmt"
	//импортируем пакет math/rand
	"math/rand"
)

func main() {
	//вызываем функцию rand.Intn,  которая генерирует целое число
	target := rand.Intn(100) + 1
	fmt.Println(target)
}
