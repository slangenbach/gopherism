package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32

	if balance < 0.0 {
		interest = 3.213
	} else if balance >= 0.0 && balance < 1000.0 {
		interest = 0.5
	} else if balance >= 1000.0 && balance < 5000.0 {
		interest = 1.621
	} else {
		interest = 2.475
	}

	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance) * (float32(balance) / 100))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var interimBalance float64
	var years int

	for years = 0; balance < targetBalance; years++ {
		interimBalance = AnnualBalanceUpdate(balance)
		balance = interimBalance
	}
	return years
}
