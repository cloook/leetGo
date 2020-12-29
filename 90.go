func subsetsWithDup(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(nums)
	for k := 0; k <= len(nums); k++ {
		dfs(nums, k, 0, c, &res)
	}
	return res
}

func dfs(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		*res = append(*res, append([]int(nil), c...))
		return
	}

	for i := start; i < len(nums)-(k-len(c)) + 1; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		c = append(c, nums[i])
		dfs(nums,k,i+1,c,res)
		c = c[:len(c)-1]
	}
	return
}