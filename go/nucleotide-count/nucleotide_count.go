package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	hist := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nuc := range d {
		if _, ok := hist[nuc]; !ok {
			return nil, fmt.Errorf("%v is not a valid nuc", nuc)
		}
		hist[nuc] += 1
	}
	return hist, nil

}
