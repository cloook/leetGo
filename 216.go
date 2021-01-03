func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	dfs(k, n, 1, c, &res)
	return res
}

func dfs(k, target, index int, c []int, res *[][]int) {
	if target == 0 {
		if len(c) == k {
			*res = append(*res, append([]int(nil), c...))
		}
		return
	}
	for i := index; i < 10; i++ {
		if target >= i {
			c = append(c, i)
			dfs(k, target-i, i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}