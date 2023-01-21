package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExits := units[unit]
	_, itemExists := bill[item]
	if !unitExits {
		return false
	}

	if itemExists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, itemExists := bill[item]
	unitValue, unitExists := units[unit]
	if !itemExists || !unitExists || itemValue-unitValue < 0 {
		return false
	} else {
		if itemValue-unitValue == 0 {
			delete(bill, item)
		} else {
			bill[item] -= unitValue
		}
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}
	return itemValue, true
}
