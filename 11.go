func maxArea(height []int) int {
	max,start,end := 0,0,len(height)-1
	for start < end {
		width := end - start
		if height[start] > height[end] {
			if height[end] * width > max {
				max = height[end] * width
			}
			end--
		} else {
			if height[start] * width > max {
				max = height[start] * width
			}
			start++
		}
		return max
	}
}