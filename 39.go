// 官方答案
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

// 自己尝试剪枝
func combinationSum(candidates []int, target int) (ans [][]int) {
	// 排序是剪枝的前提
	sort.Ints(candidates)
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		// 回溯的终止
		if target <= 0 {
			if target == 0 {
				// push解
				ans = append(ans, append([]int(nil), comb...))
			}
			return
		}
		// 生成子节点，开始搜索
		for i := idx; i < len(candidates); i++ {
			// 剪枝，大于target的没必要搜索了
			if target < candidates[i] {
				break
			}
			// push节点元素
			comb = append(comb, candidates[i])
			// 继续生成子节点
			dfs(target-candidates[i], i)
			// 终止后剪枝，向上一层继续搜索
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}