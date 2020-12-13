func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	if n < 3 {
		return res
	}
	// 固定第一个数
	for i := 0; i < n; i++ {
		// 找第二个和第三个，使用双指针
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n -1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
						left++
				}
				for left < right && nums[right-1] == nums[right] {
						right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
		
	}
	return res
}