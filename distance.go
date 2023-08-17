// At lot of the code heavily inspired/stolen from https://github.com/lmas/Damerau-Levenshtein and https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
package main

func min(n ...int) int {
	m := n[0]
	for _, i := range n {
		if m > i {
			m = i
		}
	}
	return m
}

func calculateDistance(a string, b string) int {
	aLength, bLength := len(a), len(b)

	//make a matrix aLength x bLength - adding two columns for the edge and empty string
	matrix := make([][]int, aLength+2)

	for col := range matrix {
		matrix[col] = make([]int, bLength+2)
	}

	/*
	* What is happening here?
	* If you follow:
	* https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
	* we need to have an exhaustive list of the alphabet used in the two strings
	* so if we concat them using +
	* and then loop over the "key/value" pair, we can create a map of the alphabet
	* range over a string returns int, rune as described in https://go.dev/ref/spec#For_range
	* Runes https://go.dev/ref/spec#Rune_literals
	 */

	da := make(map[rune]int)
	for _, r := range a + b {
		da[r] = 0
	}

	/*
	* Here we set the initial grid such as
	*     .  "" f  o  o
	*  .  0  6  6  6  6
	* ""  6  0  1  2  3
	*  b  6  1  0  0  0
	*  a  6  2  0  0  0
	*  r  6  3  0  0  0
	* can watch https://www.youtube.com/watch?v=MiqoA-yF-0M to see why
	*
	* As far as I can undertand, the "6" border is to rule out those operations because of their high cost,
	* whilst ensuring the operations don't fall out of the limits of the matrix indices
	 */

	maxDist := aLength + bLength
	for i := 0; i <= aLength; i++ {
		matrix[i+1][1] = i
		matrix[i+1][0] = maxDist
	}

	for j := 0; j <= bLength; j++ {
		matrix[1][j+1] = j
		matrix[0][j+1] = maxDist
	}

	// then we fill in the matrix
	for i := 1; i <= aLength; i++ {
		db := 0
		for j := 1; j <= bLength; j++ {

			k := da[rune(b[j-1])]
			l := db

			c := 1
			if a[i-1] == b[j-1] {
				c = 0
				db = j
			}

			matrix[i+1][j+1] = min(
				matrix[i][j]+c,   //substitution
				matrix[i+1][j]+1, //insertion
				matrix[i][j+1]+1, //deletion
				matrix[k][l]+(i-k-1)+1+(j-l-1), // transposition
			)

		}
		da[rune(a[i-1])] = i
	}
	return matrix[aLength+1][bLength+1]
}
