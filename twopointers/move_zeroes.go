package twopointers

func MoveZeroes(nums []int) []int {
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
		}
	}
	return nums
}
