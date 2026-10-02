
func groupAnagrams(strs []string) [][]string {
	flMap := make(map[[26]int][]string)
	res := make([][]string, 0, len(strs))

	for _, str := range strs {
		fl := getFreqListFromString(str)

		if items, ok := flMap[fl]; ok {
			items = append(items, str)
			flMap[fl] = items
			continue
		}

		flMap[fl] = []string{str}
	}

	for _, value := range flMap {
		res = append(res, value)
	}

	return res
}

func getFreqListFromString(s string) (out [26]int) {
	freqString := [26]int{}

	for _, char := range s {
		alphInd := char - 'a'
		freqString[alphInd] += 1
	}

	return freqString
}
