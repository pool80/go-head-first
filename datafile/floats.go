package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//GetFloats читает значение float64 из каждой строки файла
//функция возвращает массив чисел и любую обнаруженную ошибку
func GetFloats(fileName string) ([3]float64, error) {
	//объявление возвращаемого массива
	var numbers [3]float64
	//открываем файл с переданным именем
	file, err := os.Open(fileName)
	//если при открытии файла произошла ошибка, вернуть ее
	if err != nil {
		return numbers, err
	}
	//переменная для хранения индекса, по которому должно проводиться присваивание
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//строка прочитанная из файла преобразуется в float64
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		//если при открытии файла произошла ошибка, вернуть ее
		if err != nil {
			return numbers, err
		}
		//переход к следующему индексу массива
		i++
	}
	err = file.Close()
	//если при закрытии файла произошла ошибка, вернуть ее
	if err != nil {
		return numbers, err
	}
	//если при сканировании файла произошла ошибка, вернуть ее
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	//если программ дошла до этой точки, значит ошибок не было,
	//поэтому программа возвращает массив чисел и значение ошибки nil
	return numbers, nil
}
