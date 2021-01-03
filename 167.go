import "sort"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	n := len(numbers)
	res := []int{}
	for i := 0; i < n; i++ {
		if numbers[i] > target {
			break
		}
		m[target-numbers[i]] = i
	}
	for i := 0; i < n; i++ {
		k := numbers[i]
		if _, ok := m[k]; ok {
			res = append(res, i+1)
			res = append(res, m[numbers[i]]+1)
			sort.Ints(res)
			return res
		}
	}
	return res
}