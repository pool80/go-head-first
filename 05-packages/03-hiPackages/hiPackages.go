package main

//импортируем созданный пакет
import "greeting"

func main()  {
	//при вызове пакета нужно указать имя пакета и точку
	greeting.Hello()
	greeting.Hi()
}