package raindrops

import "fmt"

// Converts number into string representing factors
func Convert(number int) string {
	var result string
	factors := []int{3, 5, 7}
	rules := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for _, factor := range factors {
		if number%factor == 0 {
			result += rules[factor]
		}
	}

	if len(result) != 0 {
		return result
	} else {
		return fmt.Sprint(number)
	}

}
