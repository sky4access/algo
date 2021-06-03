package array

func fib(m map[int]int64, n int) int64 {
	if num, ok := m[n]; ok {
		return num
	}
	if n < 2 {
		return int64(n)
	}
	m[n] = fib(m, n-1) + fib(m, n-2)
	return m[n]
}

func fibMaster(n int) int64 {
	cache := make(map[int]int64)
	return fib(cache, n)
}
