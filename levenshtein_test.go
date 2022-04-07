package levenshtein_test

import (
	"testing"

	"github.com/wutzi15/levenshtein"
)

func TestKitten(t *testing.T) {
	lev := levenshtein.Distance("kitten", "sitting")
	if lev != 3 {
		t.Errorf("Levenshtein distance should be 3, but is %d", lev)
	}
}

func TestText(t *testing.T) {
	lev := levenshtein.Distance("Text", "Test")
	if lev != 1 {
		t.Errorf("Levenshtein distance should be 1, but is %d", lev)
	}
}

func TestEqual(t *testing.T) {
	lev := levenshtein.Distance("Test", "Test")
	if lev != 0 {
		t.Errorf("Levenshtein distance should be 0, but is %d", lev)
	}
}

func TestEmpty(t *testing.T) {
	lev := levenshtein.Distance("", "")
	if lev != 0 {
		t.Errorf("Levenshtein distance should be 0, but is %d", lev)
	}
}

func TestAGrtb(t *testing.T) {
	lev := levenshtein.Distance("Saturday", "Sunday")
	if lev != 3 {
		t.Errorf("Levenshtein distance should be 0, but is %d", lev)
	}
}
