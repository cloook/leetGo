func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	q,s := 0,0
	for q < len(nums) {
		// 快慢指针相等
		if nums[q] == nums[s] {
			q++
		} else {
            s++
			nums[s] = nums[q]
		}
	}
	return s+1
}