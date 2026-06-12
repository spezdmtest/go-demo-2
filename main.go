package main

import "fmt"

func main() {
	// transactions := [3]int{5, 10, -7}
	// transactions2 := [3]int{}
	// banks := [2]string{"Приватбанк", "ПУМБ"}
	// banks2 := [2]string{}
	// fmt.Println(transactions)
	// fmt.Println(banks)
	// transactions := [6]int{1, 2, 3, 4, 5, 6}
	// transactionsNew := transactions
	// transactionsNew[0] = 30
	// fmt.Println(transactions)
	// fmt.Println(transactionsNew)
	// patrial := transactions[1:4]
	// patrial := transactions[:4]
	// transacrionPatrial := transactions[1:5]
	// transacrionNewPatrial := transacrionPatrial[:1]
	// transacrionNewPatrial[0] = 30
	// // fmt.Println(transacrionPatrial)
	// fmt.Println(transacrionNewPatrial)
	// fmt.Println(len(transacrionPatrial), cap(transacrionPatrial))
	// fmt.Println(len(transacrionNewPatrial), cap(transacrionNewPatrial))
	// fmt.Println(newTransactions)
	// fmt.Println(transactions)
	// slice := transactions[:2]
	// slice[0] = 100

	transactions := []int{0, 20, 35}
	temp := transactions
	transactions = append(transactions, 100)
	fmt.Println(temp)
	fmt.Println(transactions)

}
