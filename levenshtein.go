package levenshtein

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
