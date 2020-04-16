package main

import "fmt"

var metersPerLiter float64

//ширина и высота стены передаются параметрами
func paintNeeded(width, height float64) float64 {

	//ширина умножается на высоту
	area := width * height

	//возвращаем значение переменной
	return area / metersPerLiter
	//выводим расход краски
	//fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	metersPerLiter = 10
	fmt.Printf("%.2f\n", paintNeeded(4.2, 3.0))
	fmt.Printf("%.2f\n", paintNeeded(77.235, 5.6))
}
