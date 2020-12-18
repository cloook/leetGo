func combinationSum2(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target <= 0 {
			if target == 0 {
				ret = append(ret, append([]int(nil),comb...))
			}
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			comb = append(comb, candidates[i])
			dfs(target - candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}