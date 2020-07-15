//----------Задача 2. Симулятор маршрутки-----------------

package main

import "fmt"

func main() {
	stations := 4
	passengersInEachStation := 3
	paymentPerPerson := 20
	revenue := stations * passengersInEachStation * paymentPerPerson
	driverSalary := revenue / 4
	gasoline := revenue / 5
	taxes := revenue / 5
	maintenance := revenue / 5
	netProfit := revenue - driverSalary - gasoline - taxes - maintenance

	fmt.Println("Всего заработали: ", revenue)
	fmt.Println("Зарплата водителя: ", driverSalary)
	fmt.Println("Расходы на топливо: ", gasoline)
	fmt.Println("Налоги: ", taxes)
	fmt.Println("Расходы на ремонт машины: ", maintenance)
	fmt.Println("Итого доход: ", netProfit)
}
