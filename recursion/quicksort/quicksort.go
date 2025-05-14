package quicksort

func QuickSort(nums []int) {
	Sort(nums, 0, len(nums))
}

func Partition(nums []int, high int, low int) int {
	pivot := nums[0]
	index := low - 1
	for i := low; i < high; i++ {
		index++
		// temp := nums[i]

	}
	return pivot
}

func Sort(nums []int, high int, low int) {

}
