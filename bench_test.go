package levenshtein_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/wutzi15/levenshtein"
)

func getString() string {
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
func BenchmarkDistance(b *testing.B) {
	// read file to string
	text := getString()
	// run benchmark
	for i := 0; i < b.N; i++ {
		levenshtein.Distance(text, text)
	}
}
func BenchmarkDistanceFast(b *testing.B) {
	// read file to string
	text := getString()
	// run benchmark
	for i := 0; i < b.N; i++ {
		levenshtein.DistanceFast(text, text)
	}
}

func BenchmarkCalc(b *testing.B) {
	// read file to string
	text := getString()
	// run benchmark
	for i := 0; i < b.N; i++ {
		levenshtein.Calculate(text, text)
	}
}
