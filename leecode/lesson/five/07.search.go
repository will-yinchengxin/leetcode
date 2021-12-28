package five
/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。

示例 1：
	输入：nums = [4,5,6,7,0,1,2], target = 0
	输出：4

示例2：
	输入：nums = [4,5,6,7,0,1,2], target = 3
	输出：-1

示例 3：
	输入：nums = [1], target = 0
	输出：-1
*/
// 数组的特征是 先增后降, 问题的关键是找到中间点
func searchOne(nums []int, target int) int {
	/*
		1) 正常逻辑是 nums[mid] > target 向左边收缩查询
			那么考虑向右的情况(右边为递增序列), nums[right] < nums[mid] && target <= nums[right]
		2) 正常逻辑是 nums[mid] < target 向右边收缩查询
			那么考虑向左的情况(左边为递增序列), nums[left] < nums[mid] && target >= nums[left]
	*/
	lenN := len(nums)
	low := 0
	high := lenN - 1

	tag := -1
	for low <= high {
		mid := low + (high - low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[high] < nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] > nums[mid] && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return  tag
}