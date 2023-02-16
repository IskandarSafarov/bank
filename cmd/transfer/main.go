package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	cd := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  20000,
		MinBalance: 100,
		Currency: "TJS",
		Color:    "White",
		Name:     "Alif",
		Active:   true,
	}

	newCard := card.Withdraw(cd, 10000)

	fmt.Println(newCard.Balance)
}