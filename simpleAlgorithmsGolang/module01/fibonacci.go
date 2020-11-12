package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	var curr int
	fib_2 :=  0
	fib_1 :=  1
	for i:=2; i<=n;i++ {
		curr = fib_2 + fib_1
		fib_2 = fib_1
		fib_1 = curr
	}
	return curr
}
