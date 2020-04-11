package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#ks!"

	//возвращает значение string.NewReplacer настроенное для замены всех # на о
	replacer := strings.NewReplacer("#", "o")

	//вызывает метод Replace для strings.NewReplacer и передает ему строку
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
