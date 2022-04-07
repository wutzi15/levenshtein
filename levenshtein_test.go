package levenshtein_test

import (
	"testing"

	"github.com/wutzi15/levenshtein"
)

func TestTable(t *testing.T) {
	tests := []struct {
		name   string
		inputA string
		inputB string
		want   int
	}{
		{name: "AGrtB", inputA: "Saturday", inputB: "Sunday", want: 3},
		{name: "kitten", inputA: "kitten", inputB: "sitting", want: 3},
		{name: "text", inputA: "Text", inputB: "Test", want: 1},
		{name: "equal", inputA: "Text", inputB: "Text", want: 0},
		{name: "empty", inputA: "", inputB: "", want: 0},
	}
	for i, tc := range tests {
		got := levenshtein.Distance(tc.inputA, tc.inputB)
		if got != tc.want {
			t.Errorf("Test %d (%s): Levenshtein distance should be %d, but is %d", i, tc.name, tc.want, got)
		}
	}
}
