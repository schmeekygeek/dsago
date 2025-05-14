package arrays

func BinarySearch(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
			mid = (left + right) / 2
		} else if nums[mid] > target {
			right = mid - 1
			mid = (left + right) / 2
		}
	}
	return -1
}
