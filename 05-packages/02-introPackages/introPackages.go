package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Farenheit:")
	farenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (farenheit - 32) * 5 / 9
	fmt.Printf("%0.2f deegres Celsius\n", celsius)
}
