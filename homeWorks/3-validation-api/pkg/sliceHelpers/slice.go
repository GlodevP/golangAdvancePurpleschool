package slicehelpers

func ContainsInStringSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func RemoveInStringSlice(s []string, e string) []string {
	for i, a := range s {
		if a == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
