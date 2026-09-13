import(
	"slices"
)

func isAnagram(s string, t string) bool {
	// converting into runes
	ss := []rune(s)
	tt := []rune(t)
	//sortig runes
	slices.Sort(ss)
	slices.Sort(tt)
	// convertig back to string
	s= string(ss)
	t= string(tt)

if s==t {
	return true
}
return false
}
