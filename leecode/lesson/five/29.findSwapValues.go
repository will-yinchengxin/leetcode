package five

/*
给定两个整数数组，请交换一对数值（每个数组中取一个数值），使得两个数组所有元素的和相等。

返回一个数组，第一个元素是第一个数组中要交换的元素，第二个元素是第二个数组中要交换的元素。
若有多个答案，返回任意一个均可。若无满足条件的数值，返回空数组。

示例:
	输入: array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
	输出: [1, 3]

示例:
	输入: array1 = [1, 2, 3], array2 = [4, 5, 6]
	输出: []

提示：
	1 <= array1.length, array2.length <= 100000
*/
func findSwapValues(array1 []int, array2 []int) []int {
	/*
		1) 连个数组的和一定为偶数, 奇数情况下此题目不成立
		2) 获取 两个数组 的 和 与 sum/2 的差值
		3) 应用 hashMap 找出各自的差值即可
	*/
	len1 := len(array1)
	len2 := len(array2)

	hash1 := make(map[int]struct{})
	sum1 := 0
	for i := 0; i < len1; i++ {
		if _, ok := hash1[array1[i]]; !ok {
			hash1[array1[i]] = struct{}{}
		}
		sum1 += array1[i]
	}

	hash2 := make(map[int]struct{})
	sum2 := 0
	for i := 0; i < len2; i++ {
		if _, ok := hash2[array2[i]]; !ok {
			hash2[array2[i]] = struct{}{}
		}
		sum2 += array2[i]
	}

	sumAll := sum1 + sum2
	if !judgeEvenOrOdd(sumAll) {
		return []int{}
	}

	diff := sum1 - sum2

	res := []int{}
	for i:=0 ; i < len2; i++{
		if _ ,ok :=hash1[diff - array2[i]]; ok {
			res = append(res, array2[i])
			res = append(res, array2[i] - diff)
		}
	}
	return res
}
func judgeEvenOrOdd(num int) bool {
	if num % 2 == 1 {
		return false
	}
	return true
}
