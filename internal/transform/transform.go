package transform

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// transformFunc is a func that transforms the elements of parts and returns
// the transformed elements in a slice.
type transformFunc func(parts []string) []string

// Do returns a slice of parts transformed with style s.
func Do(parts []string, s Style) ([]string, error) {
	tf, ok := styleFuncs[s]
	if !ok {
		return nil, fmt.Errorf("%q is not a valid transformation", string(s))
	}

	return tf(parts), nil
}

// noop just returns the parts.
func noop(parts []string) []string {
	return parts
}

// alternate returns "alternating WORD case" parts.
func alternate(parts []string) []string {
	for i, p := range parts {
		if i%2 == 0 {
			parts[i] = strings.ToLower(p)
		} else {
			parts[i] = strings.ToUpper(p)
		}
	}

	return parts
}

// capitalise returns "Capitalise First Letter" parts.
func capitalise(parts []string) []string {
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return parts
}

// invert returns "cAPITALISE eVERY lETTER eXCEPT tHE fIRST" parts.
func invert(parts []string) []string {
	for i := range parts {
		// Strings are immutable so we need to cast to byte slice to be able
		// to modify specific characters.
		pw := []byte(parts[i])
		for j, w := range pw {
			if j == 0 {
				continue
			}

			pw[j] = byte(unicode.ToTitle(rune(w)))
		}
		parts[i] = string(pw)
	}

	return parts
}

// lower returns "lower case" parts.
func lower(parts []string) []string {
	for i, p := range parts {
		parts[i] = strings.ToLower(p)
	}

	return parts
}

// upper returns "UPPER CASE" parts.
func upper(parts []string) []string {
	for i, p := range parts {
		parts[i] = strings.ToUpper(p)
	}

	return parts
}

// random returns "EVERY word randomly CAPITALISED or NOT" parts.
func random(parts []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, p := range parts {
		if r.Int()%2 == 0 {
			parts[i] = strings.ToUpper(p)
		} else {
			parts[i] = strings.ToLower(p)
		}
	}

	return parts
}
