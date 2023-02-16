package card

import (
	"bank/pkg/bank/types"
)

// Лимит снятия с карты 20 000сомони
const withdrawLimit = 20_000_00

// Withdraw функсия по снятию средств из карты
func Withdraw(card types.Card, amount types.Money) types.Card {

	// Проверяем если сумма меньше или равно 0 будет возврашена само карта
	if amount <= 0 {
		return card
	}

	// Проверяем если сумма больше лимита будет возврашена само карта
	if amount > withdrawLimit {
		return card
	}

	// Проверяем если карта не активна будет возврашена само карта
	if !card.Active {
		return card
	}

	// Проверяем если баланс карты меньше суммы будет возврашена само карта
	if card.Balance < amount {
		return card
	}

	// В этом случае из баланса карты будет снят сумма
	card.Balance -= amount
	return card
}

// Лимит зачисления на карту 50 000сомони
const depositMoneyLimit = 50_000_00

func Deposit(card *types.Card, amount types.Money) {

	// Проверяем если карта не активна будет возврашена само карта
	if !card.Active {
		return
	}

	// Проверяем если сумма больше лимита будет возврашена само карта
	if amount > depositMoneyLimit {
		return
	}

	// В этом случае в баланс карты будет зачислен сумма
	card.Balance += amount
	return
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	// Лимит зачисления на карту бонус 5 000сомони
	const bonusLimit = 5_000_00

	bonus := ((card.MinBalance * types.Money(percent)) / 100) * types.Money(daysInMonth) / types.Money(daysInYear)

	// Проверяем если карта не активна будет возврашена само карта
	if !card.Active {
		return
	}

	// На карте сумма больше 0
	if card.Balance <= 0 {
		return
	}

	// Бонус не может быть выше 5 000сомони если выше вернет 5 000сомони
	if bonus >= bonusLimit {
		bonus = bonusLimit
	}
	// В этом случае в баланс карты будет зачислен бонус
	card.Balance += bonus
	return
}

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются. 
func Total(cards []types.Card) types.Money {
	summ := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <= 0 {
			continue
		}

		summ += card.Balance
	}
	return summ
}


//func PaymentSources(cards []types.Card) []types.PaymentSource  {
	
//}