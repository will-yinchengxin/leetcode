package five
/*
符合下列属性的数组 arr 称为 山脉数组 ：
arr.length >= 3
存在 i（0 < i< arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。

示例 1：
	输入：arr = [0,1,0]
	输出：1

示例 2：
	输入：arr = [0,2,1,0]
	输出：1

示例 3：
	输入：arr = [0,10,5,2]
	输出：1

示例 4：
	输入：arr = [3,4,5,1]
	输出：2

示例 5：
	输入：arr = [24,69,100,99,79,78,67,36,26,19]
	输出：2

提示：
	3 <= arr.length <= 104
	0 <= arr[i] <= 106
	题目数据保证 arr 是一个山脉数组
*/
// 查找数组中最大元素, 使用二分查找
func peakIndexInMountainArray(arr []int) int {
	// [3,4,5,1]
	// 4 5 1 2 3
	lenA := len(arr)
	low := 0
	high := lenA-1

	for low <= high {
		mid := low + (high - low)/2

		if low == high {
			return mid
		}
		if (mid != 0 && arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1]) || (mid == 0 && arr[mid] > arr[high]) {
			return mid
		} else if arr[mid] > arr[high] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}