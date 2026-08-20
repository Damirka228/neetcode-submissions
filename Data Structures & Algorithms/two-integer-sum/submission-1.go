func twoSum(nums []int, target int) []int {
    response := []int{}
	for i := range nums {
		for j := range nums {
			if nums[i] + nums[j] == target && i != j{
				response = append(response, i, j)
			}
		}
	}
	return response[:2]
}


