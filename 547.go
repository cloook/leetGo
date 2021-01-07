func findCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	for i,v := range vis {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				vis[form] = true
				for to, conn := range isConnected[form] {
					if connn == 1 && !vis[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return

}