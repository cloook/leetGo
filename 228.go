import "strconv"

func summaryRanges(nums []int) []string {
	l, r, n := 0, 0, len(nums)
	if n == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	res := []string{}
	for r < n-1 {
		r++
		if nums[r]-nums[r-1] > 1 {
			if r-l == 1 {
				res = append(res, strconv.Itoa(nums[l]))
			} else {
				res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r-1]))
			}
			l = r
		}
		if r == n-1 {
			if nums[r]-nums[r-1] != 1 {
				res = append(res, strconv.Itoa(nums[r]))
			} else {
				res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
			}
		}
	}
	return res
}