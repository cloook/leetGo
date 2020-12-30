func lastStoneWeight(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }
    sort.Ints(stones)
    newWeight := stones[len(stones)-1] - stones[len(stones)-2]
    stones = append(stones[:len(stones)-2], newWeight)
    return lastStoneWeight(stones)
}