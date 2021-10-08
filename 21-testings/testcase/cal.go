package testcase

func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n1, n2 int) int {
	return n1 - n2
}
