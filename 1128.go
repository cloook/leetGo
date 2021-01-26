package leetGo

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	cnt := [100]int{}
	for _,v := range dominoes {
		if v[0] < v[1] {
			v[0],v[1] = v[1],v[0]
		}
		k := v[0] * 10 + v[1]
		ans += cnt[k]
		cnt[k]++
	}
	return
}
