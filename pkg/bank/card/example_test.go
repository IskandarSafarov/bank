package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// ExampleWithdraw_positive Успешное снятие баланса
func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

// ExampleWithdraw_noMoney Попытка сниятия при нулевом балансе
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000)
	fmt.Println(result.Balance)
	// Output: 0
}

// ExampleWithdraw_inactive Попытка сниятия баланса из неактивной карты
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

// ExampleWithdraw_limit Попытка снятие баланса больше лимита(20 000сомони)
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

// ---------------//
// ExampleDeposit_positive Успешное зачисление стредств
func ExampleDeposit_positive() {
	card := types.Card{Balance: 10_000_00, Active: true}
	Deposit(&card, 5_000_00)
	fmt.Println(card.Balance)
	// Output: 1500000
}

// ExampleDeposit_inactive Зачисление стредств в неактивную карту
func ExampleDeposit_inactive() {
	card := types.Card{Balance: -10_000_00, Active: false}
	Deposit(&card, 50_000_00)
	fmt.Println(card.Balance)
	// Output: -1000000
}

// ExampleDeposit_limit Успешное зачисление стредств до 50 000сомони
func ExampleDeposit_limit() {
	card := types.Card{Balance: 0, Active: true}
	Deposit(&card, 50_000_00)
	fmt.Println(card.Balance)
	// Output: 5000000
}

// ExampleDeposit_active Успешное зачисление средств до 50 000сомони при минусовом балансе на карте
func ExampleDeposit_active() {
	card := types.Card{Balance: -10_000_00, Active: true}
	Deposit(&card, 50_000_00)
	fmt.Println(card.Balance)
	// Output: 4000000
}

//--------------------//
// ExampleAddBonus_positive Нормальное зачисление средств
func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000_00, Active: true, MinBalance: 1000000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
}

// ExampleAddBonus_noactive Попытка зачислить на неактивную карту
func ExampleAddBonus_noactive() {
	card := types.Card{Balance: 10_000_00, Active: false, MinBalance: 10000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1000000
}

// ExampleAddBonus_negativeBalance Попытка зачислить зачислить на карту с отрицательным балансом
func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -10_000_00, Active: true, MinBalance: 10000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -1000000
}

// ExampleAddBonus_moreBonus Попытка зачислить бонус более 5_000 сомони
func ExampleAddBonus_moreBonus() {
	card := types.Card{Balance: 1000000, Active: true, MinBalance: 100_000_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1500000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	//Output: 3000000

}

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			PAN:     "5058",
		},
		{
			Balance: 2_000_00,
			Active:  true,
			PAN:     "good",
		},
	}
	paySource := PaymentSources(cards)
	fmt.Println(len(paySource))
	// Output: 2
}
