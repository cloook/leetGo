func wordPattern(pattern string, s string) bool {
	wMap := make(map[string]byte)
	pMap := make(map[byte]string)
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, w := range words {
		p := pattern[i]
		if wMap[w] > 0 && wMap[w] != p || pMap[p] != '' && pMap != w {
			return false
		}
		wMap[w] = p
		pMap[p] = w
	}
	return true
}