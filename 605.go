func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	for i := 0; i < l; i += 2 {
		if flowerbed[i] == 0 {
			if i == l-1 || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
	}
	return n <= 0
}