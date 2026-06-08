package main

import "fmt"

func main() {
	// transactions := [3]int{5, 10, -7}
	// banks := [2]string{"Приватбанк", "ПУМБ"}
	// fmt.Println(transactions)
	// fmt.Println(banks)
	transactions := [4]int{1, 2, 3, 4}
	banks := [2]string{}
	fmt.Println(transactions[1])
	banks[0] = "Приватбанк"
	banks[1] = "ПУМБ"
	fmt.Println(banks)
	patrial := transactions[1:4]
	fmt.Println(patrial)

}
