func twoSum(nums []int, target int) []int {
    response := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++{
			if nums[i] + nums[j] == target && i != j{
				response = append(response, i, j)
			}
		}
	}
	return response[:2]
}


