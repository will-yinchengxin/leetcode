package four

/*
实现pow(x,n)，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

示例 1：
	输入：x = 2.00000, n = 10
	输出：1024.00000

示例 2：
	输入：x = 2.10000, n = 3
	输出：9.26100

示例 3：
	输入：x = 2.00000, n = -2
	输出：0.25000
	解释：2-2 = 1/2 * 2 = 1/4 = 0.25
*/
// 递归实现
func myPow(x float64, n int) float64 {
	// 指数可为正可为负
	if (n >= 0) {
		return myPow_r(x, n)
	} else {
		// < 0 的处理有些模糊
		return 1/(myPow_r(x, -1*(n+1))*x)
	}
}

func myPow_r(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	halfPow := myPow_r(x, n/2)
	if n % 2 == 1 {
		return halfPow * halfPow * x
	} else {
		return halfPow * halfPow
	}
}