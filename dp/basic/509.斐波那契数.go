package basic

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		cur := 0
		next := 1
		for i := 2; i < n+1; i++ {
			cur, next = next, cur+next
		}
		return next
	}
}
