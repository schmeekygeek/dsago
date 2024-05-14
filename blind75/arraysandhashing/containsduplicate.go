package arraysandhashing

func containsDuplicate(nums []int) bool {
  if len(nums) <= 1 {
    return false
  }

  values := make(map[int]struct{})
  for _, val := range nums {
    if _, ok := values[val]; ok {
      return true
    }
    values[val] = struct{}{}
  }
  return false
}
