package levenshtein

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

func Levenshtein(stringA string, stringB string) int {
	lenA := len(stringA) + 1
	lenB := len(stringB) + 1
	matrix := make([][]int, lenB)
	for i := 0; i < lenA; i++ {
		matrix[i] = make([]int, lenB)
	}

	for x := 0; x < lenA; x++ {
		matrix[x][0] = x
	}

	for y := 0; y < lenB; y++ {
		matrix[0][y] = y
	}

	for x := 1; x < lenA; x++ {
		for y := 1; y < lenB; y++ {
			if stringA[x-1] == stringB[y-1] {
				matrix[x][y] = min(
					matrix[x-1][y]+1,
					matrix[x-1][y-1],
					matrix[x][y-1]+1)
			} else {
				matrix[x][y] = min(
					matrix[x-1][y]+1,
					matrix[x-1][y-1]+1,
					matrix[x][y-1]+1)
			}
		}
	}
	return (matrix[lenA-1][lenB-1])
}
