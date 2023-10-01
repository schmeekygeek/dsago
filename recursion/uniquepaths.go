package recursion

func uniquePaths(n, m int) int {
  if n == 1 || m == 1 {
    return 1
  }
  return uniquePaths(n, m-1) + uniquePaths(n-1, m)
}
