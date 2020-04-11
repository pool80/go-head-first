package main

import "fmt"

func main() {
	var originalCount int = 10
	fmt.Println("I started with", originalCount, "apples.")

	var eatenCount int = 4
	eatenCount = 7
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left")
}
