package strand

import (
	"fmt"
	"strings"
)

func ToRNA(dna string) string {
	mapping := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	rna := strings.Builder{}

	for _, nuc := range dna {
		rna.WriteRune(mapping[nuc])
	}

	return fmt.Sprint(rna.String())
}
