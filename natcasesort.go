package natcasesort

import (
	"unicode"
)

type Sort []string

func (s Sort) Len() int      { return len(s) }
func (s Sort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Sort) Less(i, j int) bool {
	iRunes := []rune(s[i])
	jRunes := []rune(s[j])

	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		// the lowercase runes are the same, so compare the original
		if ir != jr {
			return ir < jr
		}
	}
	return false
}
