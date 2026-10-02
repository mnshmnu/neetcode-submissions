type Freq struct {
	Key int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	out := make([]int, 0, k)

	for _, num := range nums {
		freqMap[num] += 1
	}

	freqs := make([]Freq,0,len(nums))
	for key, v := range freqMap {
		freqs = append(freqs, Freq{Key: key, Value: v})
	}

	sort.Slice(freqs, func(a, b int) bool {
		return freqs[a].Value > freqs[b].Value
	})


	for i:=0;i<k;i++ {
		out = append(out, freqs[i].Key)
	}

	return out
}
