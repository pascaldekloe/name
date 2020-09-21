// Package name implements various naming conventions. The two categories are
// delimiter-separated and letter case-separated words. Each of the formatting
// functions support both techniques for input, without any context.
package name

import (
	"strings"
	"unicode"
)

// CamelCase returns the medial capitals form of the words in s.
// Words consist of Unicode letters and/or numbers in any order.
// Upper case sequences [abbreviations] are preserved.
//
// Argument upper forces the letter case for the first rune.
// Use true for UpperCamelCase, a.k.a. PascalCase.
// Use false for lowerCamelCase, a.k.a. dromedaryCase.
//
// BUG(pascaldekloe): Abbreviations at the beginning of a name
// may look odd in lowerCamelCase, i.e., "tCPConn".
//
// BUG(pascaldekloe): CamelCase concatenates abbreviations by
// design, i.e., "DB-API" becomes "DBAPI".
func CamelCase(s string, upper bool) string {
	var b strings.Builder
	b.Grow(len(s))

	for i, r := range s {
		if i == 0 {
			if upper {
				b.WriteRune(unicode.ToUpper(r))
			} else {
				b.WriteRune(unicode.ToLower(r))
			}
			upper = false
			continue
		}

		switch {
		case unicode.IsLetter(r):
			if upper {
				r = unicode.ToUpper(r)
			}

			fallthrough
		case unicode.IsNumber(r):
			upper = false
			b.WriteRune(r)

		default:
			upper = true
		}
	}

	return b.String()
}

// SnakeCase returns Delimit(s, '_'), a.k.a. the snake_case.
func SnakeCase(s string) string {
	return Delimit(s, '_')
}

// DotSeparated returns Delimit(s, '.'), a.k.a. the dot notation.
func DotSeparated(s string) string {
	return Delimit(s, '.')
}

// Delimit returns the words in s delimited with separator sep.
// Words consist of Unicode letters and/or numbers in any order.
// Upper case sequences [abbreviations] are preserved.
// Use strings.ToLower or ToUpper to enforce one letter case.
func Delimit(s string, sep rune) string {
	var b strings.Builder
	b.Grow(len(s) + len(s)/4)

	var last rune // previous rune; pending write
	sepDist := 1  // distance between a sep and the current rune r
	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			if unicode.IsLower(last) {
				if b.Len() == 0 {
					last = unicode.ToUpper(last)
				} else {
					b.WriteRune(last)
					last = sep
					sepDist = 1
				}
			}

		case unicode.IsLetter(r): // lower-case
			if unicode.IsUpper(last) {
				if sepDist > 2 {
					b.WriteRune(sep)
				}
				last = unicode.ToLower(last)
			}

		case !unicode.IsNumber(r):
			if last == 0 || last == sep {
				continue
			}
			r = sep
			sepDist = 0
		}

		if last != 0 {
			b.WriteRune(last)
		}
		last = r
		sepDist++
	}

	if last != 0 && last != sep {
		if b.Len() == 0 {
			last = unicode.ToUpper(last)
		}
		b.WriteRune(last)
	}

	return b.String()
}
