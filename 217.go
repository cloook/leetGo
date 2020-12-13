func containsDuplicate(nums []int) bool {
	dMap := make(map[int]int)
	for i,v := range nums {
		_, ok := dMap[v]
		if (ok) {
			return true
		} else {
			dMap[v] = 1
		}
	}
	return flase
}