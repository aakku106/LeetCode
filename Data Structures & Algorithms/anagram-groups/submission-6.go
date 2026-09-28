func groupAnagrams(strs []string) [][]string{
	if len(strs) == 0 || len(strs) > 10000 {
		return [][]string{}
	}
	final := make([][]string, 0, len(strs)/2)
	notAnargamGroup := strs

	for len(notAnargamGroup) != 0 {
		if len(notAnargamGroup) == 1 {
			final = append(final, notAnargamGroup)
			break
		}
		first := notAnargamGroup[0]
		newNotAnargamsList := make([]string, 0, len(notAnargamGroup)/2)
		anagramGroup := make([]string, 0, len(strs)/2)
		for _, v := range notAnargamGroup {
			if v == first {
				anagramGroup = append(anagramGroup, v)
				continue
			}
			if ok := isAnagrams(first, v); ok {
				anagramGroup = append(anagramGroup, v)
			} else {
				newNotAnargamsList = append(newNotAnargamsList, v)
			}
		}
		final = append(final, anagramGroup)
		notAnargamGroup = newNotAnargamsList
	}

	return final
}

// Function to check Anagrams, takes two strings and return true if anagram false else wise
func isAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//create a hash table
	Table := make(map[byte]int8, len(s))
	// fill the hash table
	for i := range s {
		Table[s[i]]++
		Table[t[i]]--
	}
	// check if 0, cause repeat=0
	for key := range Table {
		if Table[key] != 0 {
			return false
		}
	}
	return true
}