package lengthoflastword

func solution1(s string) (ans int) {
	index := len(s) - 1

	for s[index] == ' ' {
		index--
	}
	for index > 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
