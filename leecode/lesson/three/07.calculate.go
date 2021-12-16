package three

import (
	"strings"
)

/*
给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。

表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格。 整数除法仅保留整数部分。

示例1:
	输入: "3+2*2"
	输出: 7

示例 2:
	输入: " 3/2 "
	输出: 1

示例 3:
	输入: " 3+5 / 2 "
	输出: 5
*/
func Calculate(s string) int {
	s = strings.Trim(s, " ")
	lenS := len(s)

	nums := []int{}
	symbol := []byte{}
	
	for i := 0; i < lenS; i++ {
		if s[i] == ' ' {
			continue
		} else if judgeNum(s[i]) {
			res := 0
			for judgeNum(s[i]) {
				res = int(s[i] - '0') + res * 10
				i++
			}
			nums = append(nums, res)
		} else {
			for len(symbol) > 0 && judgeSymbol(s[i], symbol[len(symbol) - 1]) {
				num1 := nums[len(nums) - 1]
				num2 := nums[len(nums) - 2]
				symbols := symbol[len(symbol) - 1]
				val := cal(symbols, num1, num2)
				symbol = symbol[:len(symbol) - 1]
				nums = nums[:len(nums) - 2]
				nums = append(nums, val)
			}
			symbol = append(symbol, s[i])
		}
	}
	for len(symbol) > 0 {
		num1 := nums[len(nums) - 1]
		num2 := nums[len(nums) - 2]
		symbols := symbol[len(symbol) - 1]
		val := cal(symbols, num1, num2)
		symbol = symbol[:len(symbol) - 1]
		nums = nums[:len(nums) - 2]
		nums = append(nums, val)
	}
	return nums[len(nums) - 1]
}
func cal(sym byte, a, b int) int {
	switch sym {
	case '+': return a + b
	case '-': return a - b
	case '*': return a * b
	case '/': return a / b
	default: return -1
	}
}
func judgeNum(num byte) bool {
	if num >= '0' && num <= '9' {
		return true
	}
	return false
}
// a 为新加入, b 为原始
func judgeSymbol(a, b byte) bool {
	if (a == '*' || a == '/') && (b == '+' || b == '-') {
		return true
	}
	return false
}

//func Calculate(s string) int {
//	s = strings.Trim(s, " ")
//	lenS := len(s)
//	// fmt.Println('+', '-', '*', '/', '0') // 43 45 42 47 48
//	// 两个栈, 一个装符号, 一个装数字, 符号位, 运算级别高的继续存储, 运算优先级别相同或者级别较高的, 直接计算
//	var num []int
//	var symbol []byte
//	for i := 0; i < lenS; {
//		if s[i] == ' ' {
//			i++
//			continue
//		} else if (judgeNum(s[i])) {
//			// 数字处理, 可能为多位数字, 如: 32 234 ....
//			result := 0
//			for i < lenS && judgeNum(s[i]) {
//				result = int((s[i] - '0')) + result * 10
//				i++
//			}
//			num = append(num, result)
//		} else {
//			// 符号处理, 符号栈为空, 或者新添加符号优先级高于末尾符号优先级 直接插入, 否则进行计算
//			if len(symbol) == 0 || judgeSymbol(s[i], symbol[len(symbol) - 1]){
//				symbol = append(symbol, s[i])
//			} else {
//				// 符号栈中有数值, 新添加的符号位为同级别符号位, 或低级别符号位
//				for len(symbol) > 0 && !judgeSymbol(s[i], symbol[len(symbol) - 1]) {
//					num1 := num[len(num)-1]
//					num2 := num[len(num)-2]
//					num = num[:len(num)-2]
//					symb := symbol[len(symbol)-1]
//					symbol = symbol[:len(symbol)-1]
//					// 传入数值的顺序应注意, 前置位为 被减数 / 被除数
//					res := cal(symb, num2, num1)
//					num = append(num, res)
//				}
//				symbol = append(symbol, s[i])
//			}
//			i++
//		}
//	}
//	for len(symbol) > 0 {
//		num1 := num[len(num)-1]
//		num2 := num[len(num)-2]
//		num = num[:len(num)-2]
//		symb := symbol[len(symbol)-1]
//		symbol = symbol[:len(symbol)-1]
//		res := cal(symb, num2, num1)
//		num = append(num, res)
//	}
//	return num[len(num)-1]
//}
//
//func judgeNum(a byte) bool {
//	return a >= '0' && a <= '9'
//}
//
//func judgeSymbol(a, b byte) bool {
//	if (a == '*' || a == '/') && (b == '+' || b == '-') {
//		return true
//	}
//	return false
//}
//
//func cal(a byte, num1, num2 int) int {
//	switch a {
//	case '+': return num1 + num2
//	case '-': return num1 - num2
//	case '*': return num1 * num2
//	case '/': return num1 / num2
//	}
//	return -1
//}