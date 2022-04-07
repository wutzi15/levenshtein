package levenshtein_test

import (
	"testing"

	"github.com/wutzi15/levenshtein"
)

func TestKitten(t *testing.T) {
	lev := levenshtein.Levenshtein("kitten", "sitting")
	if lev != 3 {
		t.Errorf("Levenshtein distance should be 3, but is %d", lev)
	}
}

func TestText(t *testing.T) {
	lev := levenshtein.Levenshtein("Text", "Test")
	if lev != 1 {
		t.Errorf("Levenshtein distance should be 1, but is %d", lev)
	}
}

func TestEqual(t *testing.T) {
	lev := levenshtein.Levenshtein("Test", "Test")
	if lev != 0 {
		t.Errorf("Levenshtein distance should be 0, but is %d", lev)
	}
}

func TestEmpty(t *testing.T) {
	lev := levenshtein.Levenshtein("", "")
	if lev != 0 {
		t.Errorf("Levenshtein distance should be 0, but is %d", lev)
	}
}

func TestAGrtb(t *testing.T) {
	lev := levenshtein.Levenshtein("Saturday", "Sunday")
	if lev != 3 {
		t.Errorf("Levenshtein distance should be 0, but is %d", lev)
	}
}
