package leetcode

func TwoSum(nums []int, target int) []int {

	var Result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				Result = []int{i, j}
				break
			}
		}
	}
	return Result
}
