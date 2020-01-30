package exam

// TwoSum q
func TwoSum(nums []int, target int) []int {
	l := len(nums)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
