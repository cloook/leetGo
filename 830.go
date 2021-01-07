func largeGroupPositions(s string) [][]int {
	l := 0
	n := len(s)
	res := [][]int{}
	for r := 0; r < n; r++ {
		if s[r] != s[l] || r == n-1 {
            if r == n-1 && s[r] == s[l] && r-l >= 2 {
                res = append(res, []int{l, r})
            } else if r-l >= 3  {
				res = append(res, []int{l, r-1})
			}
			l = r
		}
	}
	return res
}