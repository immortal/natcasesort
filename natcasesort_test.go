package natcasesort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {

	var samples = []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"v", "u", "a", "p", "z", "P", "w"},
			[]string{"a", "P", "p", "u", "v", "w", "z"},
		},
		{
			[]string{"z", "x", "y", "p", "Z", "X", "Z"},
			[]string{"p", "X", "x", "y", "Z", "Z", "z"},
		},
		{
			[]string{"e", "x", "y", "p", "E", "X", "E"},
			[]string{"E", "E", "e", "p", "X", "x", "y"},
		},
		{
			[]string{"9", "a", "1", "f", "8", "A", "40", "e", "z", "7", "Z", "v", "13"},
			[]string{"1", "7", "8", "9", "13", "40", "A", "a", "e", "f", "v", "Z", "z"},
		},
	}
	for _, s := range samples {
		for k, v := range s.expected {
			sort.Sort(Sort(s.in))
			if s.in[k] != v {
				t.Errorf("Expecting %q Got: %q", v, s.in[k])
			}
		}
	}
}

func TestOut(t *testing.T) {
	in := []string{
		"z",
		"p",
		"14",
		"40",
		"30",
		"P",
		"b",
		"f",
		"g",
		"1",
		"x",
		"v",
		"S",
		"s",
		"13",
		"9",
		"7",
		"a",
		"e",
		"E",
		"zz",
		"A",
		"aa"}
	sort.Sort(Sort(in))
	t.Log(in)
}
