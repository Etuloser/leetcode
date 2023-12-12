package findtheindexofthefirstoccurrenceinastring

func solution1(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
outer:
	for i := 0; i+n <= m; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
