package leetGo

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	begin := 0
	for begin != slow {
		begin = nums[begin]
		slow = nums[slow]
	}
	return begin
}
