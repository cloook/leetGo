func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		// 如果没有在栈中出现过，进栈，考略是否前移
		if !inStack[ch-'a'] {
			// 当前元素比栈顶小，往前翻
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				// 后面不会再出现了，留下栈顶元素，切段
				if left[last] == 0 {
					break
				}
				// 后面还有，把栈顶释放了
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}