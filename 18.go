func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ret := [][]int{}
	for i := 0; i < n-3; i++ {
		// 固定第一个数
		for k := i +1; k < n-2; k++ {
			// 固定第二个数
			sum2 := nums[i] + nums[k]
			l := k+1
			r := k+2
			for l < r {
				sum := sum2 + nums[l] + nums[r]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[k], nums[l], nums[r]})
				}
				if sum < target {
					l++
				}
				if sum > target {
					r--
				}
			}
		} 
	}
}