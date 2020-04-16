package main

import (
	"fmt"
)

//ширина и высота стены передаются параметрами
func paintNeeded(width float64, height float64) {

	//ширина умножается на высоту
	area := width * height

	//выводим расход краски
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(77.235, 5.6)
}
