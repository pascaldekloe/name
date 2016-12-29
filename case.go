// Package name implements naming conventions like camel case and snake case.
package name

import "unicode"

// SnakeCase returns the snake case version of word sequence s.
// The input can be camel case or just a bunch of words.
// Upper case abbreviations are preserved. Use strings.ToLower and
// strings.ToUpper to enforce a letter case.
func SnakeCase(s string) string {
	return Delimit(s, '_')
}

// Delimit returns word sequence s delimited with sep.
// The input can be camel case or just a bunch of words.
// Upper case abbreviations are preserved. Use strings.ToLower and
// strings.ToUpper to enforce a letter case.
func Delimit(s string, sep rune) string {
	out := make([]rune, 0, len(s)+5)

	for _, r := range s {
		switch {
		case !unicode.IsLetter(r):
			if !unicode.IsNumber(r) {
				if i := len(out); i != 0 && out[i-1] != sep {
					out = append(out, sep)
				}
				continue
			}

		case !unicode.IsUpper(r):
			if i := len(out) - 1; i >= 0 {
				if last := out[i]; unicode.IsUpper(last) {
					out = out[:i]
					if i > 0 && out[i-1] != sep {
						out = append(out, sep)
					}
					out = append(out, unicode.ToLower(last))
				}
			}

		default: // upper case letter
			if last := len(out) - 1; last >= 0 && unicode.IsLower(out[last]) {
				out = append(out, sep)
			}

		}
		out = append(out, r)
	}

	if len(out) == 0 {
		return ""
	}

	// trim tailing separator
	if i := len(out) - 1; out[i] == sep {
		out = out[:i]
	}

	return string(out)
}
