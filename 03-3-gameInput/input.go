package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I`ve chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it")
	fmt.Println(target)

	//читаем с клавиатуры
	reader := bufio.NewReader(os.Stdin)

	//запросить число
	fmt.Print("Make a guess: ")
	//прочитать данные с клавиатуры
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//удаление символа новой строки
	input = strings.TrimSpace(input)
	//входная строка преобразуется в целое число
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
}
