package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//удаляем символ новой строки из входной строки
	input = strings.TrimSpace(input)
	//прелбразование значения во float64
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	//сравниваем с float64
	if grade >= 60 {
		status := "passing"
	} else {
		status := "failing"
	}

}
