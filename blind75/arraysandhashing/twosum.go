package arraysandhashing

func twoSum(nums []int, target int) []int {
  // val: index
  values := make(map[int]int)
  for i := range nums {
    if val, found := values[target - nums[i]]; found {
      return []int{i, val}
    }
    values[nums[i]] = i
  }
  return []int{}
}
