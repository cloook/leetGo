func rotate(nums []int, k int)  {
	n := len(nums)
	for i:= 0;i < k; i++ {
		temp := nums[n-1]
		for j := n-1;j >0;j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}