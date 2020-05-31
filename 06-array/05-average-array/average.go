package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("C:\\Users\\maxim\\go\\src\\HeadFirst\\06-array\\04-text-file\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	//индекс элеменов игнорируется, перебор всех чисел в массиве
	for _, number := range numbers {
		//текущее число прибавляется к сумме
		sum += number
	}
	//получить длинну массива в виде значения int и преобразовать к типу float64
	sampleCount := float64(len(numbers))
	//чтобы вычислить среднее, делим сумму значений на длинну массива
	fmt.Println("Average: %0.2f\n", sum/sampleCount)
}
