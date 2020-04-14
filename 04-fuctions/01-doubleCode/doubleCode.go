package main

import (
	"fmt"
)

func main() {
	width := 4.2
	height := 3.0
	area := width * height
	fmt.Printf("%.2f litters needed\n", area/10)

	width = 5.2
	height = 3.5
	area = width * height
	fmt.Printf("%.2f litters needed\n", area/10)
}
