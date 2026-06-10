package main

import "fmt"

func main() {
	// transactions := [3]int{5, 10, -7}
	// transactions2 := [3]int{}
	// banks := [2]string{"Приватбанк", "ПУМБ"}
	// banks2 := [2]string{}
	// fmt.Println(transactions)
	// fmt.Println(banks)
	transactions := [5]int{1, 2, 3, 4, 5}
	transactionsNew := transactions
	transactionsNew[0] = 30

	fmt.Println(transactions)
	fmt.Println(transactionsNew)

}
