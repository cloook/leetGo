func minPatches(nums []int, n int) int {
	total,i,res := 1,0,0
	for total <= n {
		if i < len(nums) && total >= nums[i] {
            total+=nums[i]
			i++
		} else {
			res++
			total *= 2
		}
	}
	return res
}