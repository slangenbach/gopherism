package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var strategy string
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	cardSum := card1Value + card2Value

	switch {
	case card1 == "ace" && card2 == "ace":
		strategy = "P"
	case cardSum == 21:
		if dealerCardValue < 10 {
			strategy = "W"
		} else {
			strategy = "S"
		}
	case cardSum >= 17 && cardSum <= 20:
		strategy = "S"
	case cardSum >= 12 && cardSum <= 16:
		if dealerCardValue >= 7 {
			strategy = "H"
		} else {
			strategy = "S"
		}

	case cardSum <= 11:
		strategy = "H"
	}
	return strategy
}
