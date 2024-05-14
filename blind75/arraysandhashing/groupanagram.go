package arraysandhashing

func groupAnagrams(strs []string) [][]string {
  anagramMap := make(map[[26]int][]string)
  for _, val := range strs {
    counts := [26]int{}

    for _, char := range val {
      counts[char - 'a']++
    }
    anagramMap[counts] = append(anagramMap[counts], val)
  }

  result := make([][]string, len(anagramMap))
  idx := 0
  for _, v := range anagramMap {
    result[idx] = v
    idx++
  }

  return result
}
