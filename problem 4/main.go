package main

import "fmt"

func main() {
	fmt.Print(sum_to_n_c(5))
}

//using formula for sum up to N
//time complexity: O(1)
//space complexity: O(1)
func sum_to_n_a(n int) int {
	return (n + 1) * n / 2
}

//using recursion
//time complexity: O(n)
//space complexity: O(n)
func sum_to_n_b(n int) int {
	if (n == 0) {
		return 0
	} else {
		return n + sum_to_n_b(n - 1)
	}
}

//using for loop
//time complexity: O(n)
//space complexity: O(1)
func sum_to_n_c(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}