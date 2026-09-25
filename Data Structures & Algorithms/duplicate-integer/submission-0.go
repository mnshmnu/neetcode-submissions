func hasDuplicate(nums []int) bool {
    uniqueNumbers := make(map[int]bool, len(nums))

    for _, num := range nums {
        _, ok := uniqueNumbers[num]
        if ok {
            return true
        }
        uniqueNumbers[num] = true
    }

    return false
}
