func hasDuplicate(nums []int) bool {
    count := make(map[int]bool)
	for _,v := range nums {
		if count[v] {
			return true
		}
		count[v] = true
	}
	return false
}
