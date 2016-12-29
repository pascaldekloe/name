// Package name implements naming conventions like camel case and snake case.
package name

import "unicode"

// SnakeCase returns the snake case version of word sequence s.
// The input can be camel case or just a bunch of words.
// Upper case abbreviations are preserved. Use strings.ToLower when
// all lower case is required.
func SnakeCase(s string) string {
	return Sep(s, '_')
}

// Sep returns a token separated version of word sequence s.
// The input can be camel case or just a bunch of words.
// Upper case abbreviations are preserved. Use strings.ToLower when
// all lower case is required.
func Sep(s string, sep rune) string {
	in := []rune(s)
	out := make([]rune, 0, len(in)+5)

	wordStart := func() {
		if len(out) != 0 && out[len(out)-1] != sep {
			out = append(out, sep)
		}
	}

	for i, r := range s {
		if !unicode.IsLetter(r) {
			if unicode.IsNumber(r) {
				out = append(out, r)
			} else {
				wordStart()
			}
			continue
		}
		// letter

		if !unicode.IsUpper(r) {
			out = append(out, r)
			continue
		}
		// upper case

		if i+1 == len(in) {
			if i > 0 && !unicode.IsUpper(in[i-1]) && unicode.IsLetter(in[i-1]) {
				out = append(out, sep)
			}
			out = append(out, r)
			break
		}
		// not last

		if !unicode.IsUpper(in[i+1]) && unicode.IsLetter(in[i+1]) {
			wordStart()
			out = append(out, unicode.ToLower(r))
			continue
		}
		// followed by another upper

		if i > 0 && !unicode.IsUpper(in[i-1]) && unicode.IsLetter(in[i-1]) {
			wordStart()
		}
		out = append(out, r)
	}

	if len(out) == 0 {
		return ""
	}

	// trim tailing separator
	if out[len(out)-1] == sep {
		out = out[:len(out)-1]
	}

	return string(out)
}
