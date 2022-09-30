package levenshtein

import (
	"bytes"
	"math"
	"strings"
	"unicode/utf8"
)

/* get the minimum of three values */
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

/*
/* calculate the Levenshtein distance between two strings
*/
func DistanceOld(strA string, strB string) int {
	/* Check which of the input strings is longer
	in order to determine the size of the matrix */

	var stringA = ""
	var stringB = ""
	if len(strA) > len(strB) {
		stringA = strB
		stringB = strA
	} else {
		stringA = strA
		stringB = strB
	}

	lenA := len(stringA) + 1
	lenB := len(stringB) + 1
	distanceMatrix := make([][]int, lenB)

	/* Initialize the distance matrix */
	for i := 0; i < lenA; i++ {
		distanceMatrix[i] = make([]int, lenB)
	}
	/* source prefixes can be transformed into empty string by
	dropping all characters*/
	for x := 0; x < lenA; x++ {
		distanceMatrix[x][0] = x
	}
	/* target prefixes can be reached from empty source prefix
	by inserting every character */
	for y := 0; y < lenB; y++ {
		distanceMatrix[0][y] = y
	}

	/* calculate for all x,y the distanceMatrix[x][y] which holds the
	Levenshtein distance between stringA[0..x-1] and stringB[0..y-1] */
	for x := 1; x < lenA; x++ {
		for y := 1; y < lenB; y++ {
			if stringA[x-1] == stringB[y-1] {
				distanceMatrix[x][y] = min(
					distanceMatrix[x-1][y]+1, // deletion
					distanceMatrix[x-1][y-1], // insertion
					distanceMatrix[x][y-1]+1) // substitution
			} else {
				distanceMatrix[x][y] = min(
					distanceMatrix[x-1][y]+1,
					distanceMatrix[x-1][y-1]+1,
					distanceMatrix[x][y-1]+1)
			}
		}
	}
	return (distanceMatrix[lenA-1][lenB-1])
}

func compareStrings(s1, s2 string) bool {
	if len(s1) == len(s2) && strings.EqualFold(s1, s2) {
		return true
	}
	return false
}

// Levenshtein distance
// https://en.wikipedia.org/wiki/Levenshtein_distance
func Distance(strA string, strB string) int {
	// Check which of the input strings is longer
	// in order to determine the size of the matrix
	var stringA = ""
	var stringB = ""
	if len(strA) > len(strB) {
		stringA = strB
		stringB = strA
	} else {
		stringA = strA
		stringB = strB
	}

	lenA := len(stringA) + 1
	lenB := len(stringB) + 1
	distanceMatrix := make([]int, lenA)

	// Initialize the distance matrix
	for x := 0; x < lenA; x++ {
		distanceMatrix[x] = x
	}

	// calculate for all x,y the distanceMatrix[x][y] which holds the
	// Levenshtein distance between stringA[0..x-1] and stringB[0..y-1]
	for y := 1; y < lenB; y++ {
		previousDiagonal := distanceMatrix[0]
		distanceMatrix[0] = y
		for x := 1; x < lenA; x++ {
			oldDiagonal := distanceMatrix[x]
			if stringA[x-1] == stringB[y-1] {
				distanceMatrix[x] = previousDiagonal
			} else {
				distanceMatrix[x] = min(
					distanceMatrix[x-1]+1, // deletion
					previousDiagonal+1,    // insertion
					distanceMatrix[x]+1)   // substitution
			}
			previousDiagonal = oldDiagonal
		}
	}
	return (distanceMatrix[lenA-1])
}

func sort(s1, s2 string) (shorter, longer string) {
	if utf8.RuneCountInString(s1) < utf8.RuneCountInString(s2) {
		return s1, s2
	}
	return s2, s1
}

// Calculate calculates Jaro-Winkler distance of two strings.
// The function lowercases its parameters.
// Taken from https://github.com/jhvst/go-jaro-winkler-distance
// https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func DistanceJW(s1, s2 string) float64 {

	s1, s2 = sort(strings.ToLower(s1), strings.ToLower(s2))

	// A sliding window for match search.
	window := uint(math.Floor(
		math.Max(
			float64(utf8.RuneCountInString(s1)),
			float64(utf8.RuneCountInString(s2)),
		)/2,
	) - 1)

	runes1 := bytes.Runes([]byte(s1))
	runes2 := bytes.Runes([]byte(s2))

	var m uint = 0
	var matches []bool
	for i := 0; i < len(runes1); i++ {
		match := false
		if runes1[i] == runes2[i] {
			m++
			match = true
		}
		matches = append(matches, match)
	}

	if m == 0 {
		return 0.0
	}

	var t float64 = 0
	slider := runes2[0:window]
	for i := 0; i < len(runes1); i++ {

		if matches[i] {
			continue
		}

		idx := strings.Index(string(slider), string(runes1[i]))
		if idx != -1 && !matches[idx] {
			t += 0.5
			matches[idx] = true
		}

		start := uint(math.Max(
			0,
			float64(i-int(window)),
		))
		end := uint(math.Min(
			float64(i+int(window)),
			float64(len(runes1)),
		))

		slider_new := runes2[int(start):int(end)]
		if len(slider_new) >= int(window) {
			slider = slider_new
		}
	}

	var term1, term2, term3 float64
	term1 = float64(m) / float64(len(runes1))
	term2 = float64(m) / float64(len(runes2))
	term3 = (float64(uint(float64(m) - t))) / float64(m)

	var simj float64
	simj = (term1 + term2 + term3) / 3

	var p float64 = 0.1
	var l uint = 0
	var common_prefix uint = uint(math.Min(4.0, float64(len(s1))))
	for i := 0; i < len(s1[0:common_prefix]); i++ {
		if s1[0:common_prefix][i] == s2[0:common_prefix][i] {
			l++
		}
	}

	return simj + float64(l)*p*(1-simj)
}
