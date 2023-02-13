package raindrops

import "fmt"

// Determins if number is a factor
func isFactor(number int, factor int) bool {
	return number%factor == 0
}

// Converts number into string representing factors
func Convert(number int) string {
	var factorSum int
	factors := []int{3, 5, 7}
	rules := map[int]string{
		3:  "Pling",
		5:  "Plang",
		7:  "Plong",
		8:  "PlingPlang",
		10: "PlingPlong",
		12: "PlangPlong",
		15: "PlingPlangPlong",
	}

	for _, factor := range factors {
		if isFactor(number, factor) {
			factorSum += factor
		}
	}

	res, ok := rules[factorSum]
	if ok {
		return res
	} else {
		return fmt.Sprint(number)
	}

}
