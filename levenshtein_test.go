package levenshtein_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/wutzi15/levenshtein"
)

func GetString() string {
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text string
	for scanner.Scan() {
		text = text + scanner.Text()
	}
	return text
}

func TestTable(t *testing.T) {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	text := GetString()
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
		{name: "index", inputA: text, inputB: text, want: 0},
	}
	for i, tc := range tests {
		got := levenshtein.Distance(tc.inputA, tc.inputB)
		if got != tc.want {
			t.Errorf("Test %d (%s): Levenshtein distance should be %d, but is %d", i, tc.name, tc.want, got)
		}
	}
}

func floatEQ(a, b float64) bool {
	return (a - b) < 0.00001
}
func TestTableJW(t *testing.T) {
	text := GetString()
	tests := []struct {
		name   string
		inputA string
		inputB string
		want   float64
	}{
		{name: "AGrtB", inputA: "Saturday", inputB: "Sunday", want: 0.187500},
		{name: "kitten", inputA: "kitten", inputB: "sitting", want: 0.822222},
		{name: "text", inputA: "Text", inputB: "Test", want: 0.883333},
		{name: "equal", inputA: "Text", inputB: "Text", want: 1.0},
		{name: "empty", inputA: "", inputB: "", want: 1.0},
		{name: "index", inputA: text, inputB: text, want: 1.0},
	}
	for i, tc := range tests {
		got := levenshtein.DistanceJW(tc.inputA, tc.inputB)
		if !floatEQ(got, tc.want) {
			t.Errorf("Test %d (%s): Jaro-Winkler distance distance should be %f, but is %f", i, tc.name, tc.want, got)
		}
	}
}
