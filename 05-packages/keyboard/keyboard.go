package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	//если при чтении ввода произошла ошибка, она будет возвращена функцией return 0, err
	if err != nil {
		//если при чтении ввода произошла ошибка, она будет возвращена функцией
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	//возвращает ошибки при преобразовании строки в float64
	if err != nil {
		return 0, err
	}
	return number, nil
}
