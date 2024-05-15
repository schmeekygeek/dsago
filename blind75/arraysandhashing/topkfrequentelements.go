package arraysandhashing

func topKFrequent(nums []int, k int) []int {
  freqmap := make(map[int]int)
  freq := make([][]int, len(nums) + 1)
  result := []int{}

  for _, val := range nums {
    freqmap[val]++
  }
  
  for key, val := range freqmap {
    freq[val] = append(freq[val], key)
  }

  count := k
  for i := len(freq) - 1; i >= 0; i-- {
    if count == 0 {
      break
    }
    for _, val := range freq[i] {
      result = append(result, val)
      count--
    }
  }

  return result
}
