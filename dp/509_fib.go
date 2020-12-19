package dp
func fib(N int) int {
	if N == 0{
		return 0
	}
	if N == 1 || N == 2 {
		return  1
	}
	prev,curr := 1,1
	for i:=3;i<=N;i++{
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
