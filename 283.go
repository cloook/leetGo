func moveZeroes(nums []int) {
	c, n := 0, len(nums)
	for i, v := range nums {
		if v == 0 {
			c++
		} else {
			nums[i-c] = v
		}
	}
	for i := 0; i < c; i++ {
		if n-i-1 >= 0 {
			nums[n-i-1] = 0
		}
	}
}