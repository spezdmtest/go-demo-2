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

	// transactions := []int{0, 20, 35}
	// temp := transactions
	// transactions = append(transactions, 100)
	// fmt.Println(temp)
	// fmt.Println(transactions)
	//-------------------------------------------------------------

	// у циклі запитуємо введення транзакцій: -10, 10, 40.5
	// додаємо кожну транзакцію до масиву транзакцій
	// виводимо масив транзакцій
	// вивести суму баланса в консоль

	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, tr2...)
	// fmt.Println(tr1)

	// for index, value := range tr1 {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range tr1 {
	// 	fmt.Println(value)
	// }

	// for index := range tr1 {
	// 	fmt.Println(tr1[index])
	// }

	tr := make([]string, 0, 2)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1")
	tr = append(tr, "2")
	fmt.Println(tr)

	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f$", balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введіть транзакцію (n для завершення): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
