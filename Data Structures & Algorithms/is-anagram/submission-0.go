func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqMapS, freqMapT := createFreqMap(s), createFreqMap(t)

	if len(freqMapS) != len(freqMapT) {
		return false
	}

	for _, char := range(s) {
		charFreqS, okS := freqMapS[char]
		charFreqT, okT := freqMapT[char]

		if okS && okT {
			if charFreqS == charFreqT {
				continue
			}
			return false
		}
		return false
	}

	return true
}

func createFreqMap(s string)( freqMap map[rune]int) {
	freqMap = make(map[rune]int)
	for _, char := range(s) {
		_, ok := freqMap[char]
		if ok {
			freqMap[char]++
			continue
		}
		freqMap[char] = 1
	}
	return freqMap
}
