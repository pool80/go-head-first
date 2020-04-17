package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("height can`t be negative")
	//fmt.Println(err.Error())
	//log.Fatal(err.Error())

	//функции fmt и log проверяют содержит ли передающееся значение метод Error,
	//если содержит выводят возвращаемое значение Error
	fmt.Println(err)
	log.Fatal(err)
}
