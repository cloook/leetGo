import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, v := range nums {
		if v <= 0 {
			nums[i] = n + 1
		}
	}
	for _, v := range nums {
		if abs(v) <= n && nums[abs(v)-1] > 0 {
			nums[abs(v)-1] = -nums[abs(v)-1]
		}
	}

	for i, v := range nums {
		fmt.Print(v)
		if v > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}