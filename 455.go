func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	n,m := len(g),len(s)
    for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && g[i] > s[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return
}