package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	apps := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}
	numRunes := utf8.RuneCountInString(log)

	if numRunes > 0 {
		for _, char := range log {
			app, exists := apps[char]
			if exists {
				return app
			}
		}
	}
	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newChar string

	for _, char := range log {
		if char == oldRune {
			newChar += string(newRune)
		} else {
			newChar += string(char)
		}
	}
	return newChar
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
