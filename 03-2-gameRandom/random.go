package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//получаем рекущую дату и время в формате целого числа
	seconds := time.Now().Unix()
	//функция генератора случайных чисел
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I`ve chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it")
	fmt.Println(target)
}
