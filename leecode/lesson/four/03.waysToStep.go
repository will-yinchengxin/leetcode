package four
/*
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

示例1:
	 输入：n = 3
	 输出：4
	 说明: 有四种走法

示例2:
	 输入：n = 5
	 输出：13

提示:
	n范围在[1, 1000000]之间
*/
// 递归 + 备忘录
func waysToStep(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, ok := hashMap[n]; ok {
		return val
	}
	hashMap[n] = (waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)) % mod
	return hashMap[n]
}

// 非递归实现
func wayToStepAno(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	helpSel := make([]int, n + 1)
	helpSel[1] = 1
	helpSel[2] = 2
	helpSel[3] = 4
	for i := 4; i < n ; i++ {
		// 计算公式 ???
		helpSel[i] = (helpSel[i-1]%mod + helpSel[i-2]%mod + helpSel[i-3])%mod
	}
	return helpSeli[n]
}