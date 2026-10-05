func groupAnagrams(strs []string) [][]string {
	g := make(map[[26]byte][]string)
	for _, v := range strs {
		var c [26]byte
		for i := 0; i < len(v); i++ {
			c[v[i]-'a']++
		}
		g[c] = append(g[c], v)
	}
	f := make([][]string, 0, len(g))
	for _, v := range g {
		f = append(f, v)
	}
	return f
}
