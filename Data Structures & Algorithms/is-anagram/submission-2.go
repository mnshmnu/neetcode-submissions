func isAnagram(s string, t string) bool {
	counts := make(map[rune]int)

	for _, char := range(s) {
		counts[char]++
	}

	for _, char := range(t) {
		counts[char]--
		if counts[char] == -1 {
			return false
		}
	}

	for _, count := range(counts) {
		if count > 0 {
			return false
		}
	}

	return true
}
