package recursion

func nthtriangular(n int) int {
	if n == 1 {
		return 1
	}
	return n + nthtriangular(n-1)
}
