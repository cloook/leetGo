func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	top,left,bottom,right := 0,0,n-1,n-1
	count,sum := 0, n * n
	res := make([][]int,n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for count < sum {
		i,j := top,left
		for j <= right && count < sum {
			res[i][j] = count + 1
            count++
			j++
		}
		i,j = top + 1, right
		for i <= bottom && count < sum {
			res[i][j] = count + 1
            count++
			i++
		}
		i,j = bottom, right -1
		for j >= left && count < sum {
			res[i][j] = count + 1
            count++
			j--
		}
		i,j = bottom -1, left
		for i > top && count < sum {
			res[i][j] = count + 1
            count++
			i--
		}
		// 进入下一层
		top,left,bottom,right = top+1,left+1,bottom-1,right-1
	}
	return res
}