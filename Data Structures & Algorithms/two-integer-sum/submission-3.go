func twoSum(nums []int, target int) []int {
    numbers := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := numbers[complement]; ok {
			return []int{j, i}
		}
		
		numbers[num] = i
	}

	return []int{}
}
