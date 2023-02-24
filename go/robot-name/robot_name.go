package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

// Registry for names of active robots
var nameRegistry = make(map[string]bool)

// Max number of names available
const allNames = 26 * 26 * 1000
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generate random name consisting of two letters and three digits
func GenerateName() string {
	rand.Seed(time.Now().UnixNano())

	genLetter := func() string {
		return string(letters[rand.Intn(len(letters))])
	}

	return fmt.Sprintf("%s%s%03d", genLetter(), genLetter(), rand.Intn(1000))

}

// Check if name exist in registry
func isNameUsed(name string) bool {
	return nameRegistry[name]
}

// Assign a random, unused name to a robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	for {
		name := GenerateName()

		if !isNameUsed(name) {
			nameRegistry[name] = true
			r.name = name
			break
		}

		if len(nameRegistry) == allNames {
			return "", fmt.Errorf("no more names available")
		}
	}

	return r.name, nil

}

// Reset robot and clear its name
func (r *Robot) Reset() {
	nameRegistry[r.name] = false
	r.name = ""
}
