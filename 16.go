func threeSumClosest(nums []int, target int) int {
	ret := 0
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			res := nums[i] + nums[l] + nums[r]
			if res == target {
				return target
			}
			if abs(res - target) < abs(ret - target) {
				ret = res
			}
			if res > target {
				r--
			}
			if res < target {
				l++
			}
		}
	}
	return ret
}

func abs(c int) int {
	if c >= 0 {
		return c
	} else {
		return -c
	}
}