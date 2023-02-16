package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1000,
			Amount: 45,
		},
		{
			ID:     1050,
			Amount: 600,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)

	// Output: 1050
}
