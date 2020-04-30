package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	//если при чтении ввода произошла ошибка, она будет возвращена функцией return 0, err
	if err != nil {
		//если при чтении ввода произошла ошибка, она будет возвращена функцией
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	//возвращает ошибки при преобразовании строки в float64
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("Enter a grade:")
	grade, err := getFloat()
	// если фуекция вернула ошибку, программа сообщает об этом и завершается
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
