func isAnagram(s string, t string) bool {
		// check if both are of samme lengtj
	if len(s) != len(t) {
		return false
	}

	// create united Hash
	UH := make(map[byte]int8, len(s))

	// Filling up hash with nums of characters contaning
	for i := range s {
		UH[s[i]]++
		UH[t[i]]--
	}
	for k := range UH {
		if UH[k] != 0 {
			return false
		}
	}

	return true
}
