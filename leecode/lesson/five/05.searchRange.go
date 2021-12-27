package five
/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。

进阶：
	你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？ // 二分查找

示例 1：
	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]

示例2：
	输入：nums = [5,7,7,8,8,10], target = 6
	输出：[-1,-1]

示例 3：
	输入：nums = [], target = 0
	输出：[-1,-1]


提示：
	0 <= nums.length <= 105
	-10^9<= nums[i]<= 10^9
	nums是一个非递减数组
	-10^9<= target<= 10^9
*/
// 这种方式会超时
func searchRange(nums []int, target int) []int {
	lenN := len(nums)
	low := 0
	high := lenN - 1

	res := []int{-1, -1}
	for low <= high {
		mid := low + (high - low) / 2 // 避免数组越界
		if nums[mid] == target {
			end := mid + 1
			start := mid - 1
			for ; start >= 0; start--  {
				if nums[end] != target {
					break
				}
			}
			for ; end <= high; end++ {
				if nums[end] != target {
					break
				}
			}
			// 判断首位元素
			if nums[start] == target {
				res[0] = start
			} else {
				res[0] = start+1
			}
			if nums[end] == target {
				res[1] = end
			} else  {
				res[1] = end - 1
			}
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
