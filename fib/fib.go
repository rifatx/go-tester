package fib

func FibR(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n + FibR(n-1)
}

func FibI(n int) int {
	r := 0

	for n > 0 {
		r = r + n
		n--
	}

	return r
}
