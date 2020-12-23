func firstUniqChar(s string) int {
	dict := [26]int{}
	for _,v := range s {
		ch := v - 'a'
		dict[ch]++
	}
	for i,v := range s {
		if dict[v-'a'] == 1 {
			return i
		}
	}
	return -1
}