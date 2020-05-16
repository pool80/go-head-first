package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade:")
	grade, err := keyboard.GetFloat()
	// если функция вернула ошибку, программа сообщает об этом и завершается
	if err != nil {
		log.Fatal(err)
	}

	var status string
	//сравниваем с float64
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
