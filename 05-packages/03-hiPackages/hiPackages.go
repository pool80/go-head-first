package main

//импортируем созданный пакет
import (
	"greeting"
	"greeting/deutsch"
)

func main() {
	//при вызове пакета нужно указать имя пакета и точку
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()
}
