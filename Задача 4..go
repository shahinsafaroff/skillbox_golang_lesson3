//------------Задача 4. Злостные вредители -- Первая часть-----

package main

import "fmt"

func main() {

	var startHeight, growth, ate, targetDays, height int
	fmt.Println("вводите высоту рассад: ")
	fmt.Scan(&startHeight)
	fmt.Println("вводите эжедневный рост сорта: ")
	fmt.Scan(&growth)
	fmt.Println("вводите эжедневный урон роста из за термитов: ")
	fmt.Scan(&ate)
	//fmt.Scan(&targetHeight)
	fmt.Println("вводите нужное количество дней для выращивание: ")
	fmt.Scan(&targetDays)

	height = startHeight + (growth - ate)*(targetDays - 1) + growth/2

	fmt.Println("Полученный итоговый рост бамбука: ", height)
}
