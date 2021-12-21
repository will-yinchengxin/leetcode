package four

/*
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0, F(1)= 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/
const mod = 1000000007
var helpSeli []int

func fib(n int) int {
	helpSeli = make([]int, n+1)
	return fib_r(n)
}

func fib_r(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if helpSeli[n] != 0 {
		return helpSeli[n]
	}
	helpSeli[n] = (fib_r(n-1) + fib_r(n-2)) % mod
	return helpSeli[n]
}