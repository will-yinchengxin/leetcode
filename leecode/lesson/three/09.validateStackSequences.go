package three

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：
	输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
	输出：true
	解释：我们可以按以下顺序执行：
	push(1), push(2), push(3), push(4), pop() -> 4,
	push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

示例 2：
	输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
	输出：false
	解释：1 不能在 2 之前弹出。

长度 < 1000
元素大小 < 1000
*/
func ValidateStackSequences(pushed []int, popped []int) bool {
	lenPush := len(pushed)
	lenPop := len(popped)
	if lenPop != lenPush || lenPop == 0 || lenPush == 0 {
		return true
	}
	// 辅助栈用来记录
	// 如果 push 和 pop 的首元素相等, 那么 pop 弹出首元素
	// 不相等, 则将 push 中的元素放置到临时栈 stack 中
	// pop 再次与 临时栈中的结果对比
 	stack := []int{}
	for i := 0; i < lenPush; i++ {
		if pushed[i] != popped[0] {
			stack = append(stack, pushed[i])
		} else {
			popped = popped[1:]
		}
		for len(stack) > 0 && popped[0] == stack[len(stack) - 1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}

	for len(popped) > 0 {
		return false
	}
	return true
}
