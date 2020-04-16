package main

import "fmt"

//объявляем что функция возвращает float
func paintNeeded(width, height float64) float64 {

	area := width * height

	//функция возвращает расход краски
	return area / 10.0
	//выводим расход краски
	//fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	//объявляем переменные для хранения расхода краски
	var amount, total float64
	//вызываем paintNeeded и сохраняем возвращаемое значение
	amount = paintNeeded(4.2, 3.0)
	//выводим расход краски для первой стены
	fmt.Printf("%.2f liters needed\n", amount)
	//прибавляем расход для текушей стены
	total += amount
	//повторяем действиядля второй стены
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	//выводим общий расход по всем стенам
	fmt.Printf("%.2f liters needed\n", total)
}
