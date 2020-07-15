//------------Задача 4. Злостные вредители -- Вторая часть-----

package main

import "fmt"

func main() {

	var startHeight, growth, ate, days, targetHeight int
	fmt.Println("вводите высоту рассад: ")
	fmt.Scan(&startHeight)
	fmt.Println("вводите эжедневный рост сорта: ")
	fmt.Scan(&growth)
	fmt.Println("вводите эжедневный урон роста из за термитов: ")
	fmt.Scan(&ate)
	fmt.Println("вводите нужное высоту для вырубки созревщего бамбука")
	fmt.Scan(&targetHeight)

	days = (targetHeight - startHeight - growth) / (growth - ate) + 2

	fmt.Println("сколько полных дней понадобится бамбуку, чтобы его можно было срубить и продать? ", days)

}
