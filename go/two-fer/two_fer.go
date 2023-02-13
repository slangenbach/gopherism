// Functions for sharing stuff
package twofer

import "fmt"

// Returns a sentence specifying whom to share with
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	} else {
		return "One for you, one for me."
	}
}
