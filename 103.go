/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**************DFS 深度优先********************/

 func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	dfs(root, 0, &ans)
	return ans
}

func dfs(root *TreeNode, depth int, ans *[][]int) {
	// dfs终止条件
	if root == nil {
		return
	}
	// 当前层还没有集合，先创建集合
	if depth >= len(*ans) {
		*ans = append(*ans, []int{})
	}
	if depth & 1 == 0 {
		// 顺序插入
		(*ans)[depth] = append((*ans)[depth], root.Val)
	} else {
		// 倒序插入
		(*ans)[depth] = append([]int{root.Val}, (*ans)[depth]...)
	}
	dfs(root.Left, depth +1, ans)
	dfs(root.Right, depth +1, ans)
}


/**********BFS 广度优先***********/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res[][]int
	if root == nil {
		return res
	}
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	isLeftStart := true
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int,l)
		for i := 0; i< l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, nodee.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeftStart {
				ans[i] = node.Val
			} else {
				ans[l-i-1] = node.Val
			}
		}
		res = append(res, ans)
		isLeftStart = !isLeftStart
		queue = queue[l:]
	}
}