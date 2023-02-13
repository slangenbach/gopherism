package hamming

import "errors"

// Calculate Hamming distane between two strings
func Distance(a, b string) (int, error) {
	var count int

	if len(a) != len(b) {
		return 0, errors.New("strings need to be of equal length")
	} else {
		for idx := range a {
			if a[idx] != b[idx] {
				count += 1
			}
		}
		return count, nil
	}
}
