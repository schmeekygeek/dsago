package recursion

func reverse(s string) string {
  if len(s) == 1 {
    return s
  }

  return reverse(string(s[1:])) + string(s[0])
}
