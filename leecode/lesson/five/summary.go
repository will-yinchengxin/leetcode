package five
/*
一. 二分查找:
	编写二分查找的正确姿势:
		- 查找区间永远是闭区间 [low, high]
		- 循环条件永远是 low <= high
		- 对于 low == high 的情况, 必要的时候特殊处理, 在 while 内部补充退出条件
		- 返回值永远是 mid, 而不要是 low 或者 high
		- low 或者 high 的更新条件永远是 low=mid+1 和 high=mid-1
		- 对于非确定性查找, 使用前后探测法, 来确定搜索区间  !!!
		- 先处理命中情况, 再处理在左右半部分查找的情况

	非确定性查找:
		- 第一个, 最后一个相等的
		- 第一个大于等于的,最后一个小于等于的
		- 循环数组寻找最小值
		- 寻找峰值

	模板:
		func bSearch(a []int, len int, val int) int {
			low := 0
			high := len-1
			for low <= high {
				mid := low + (high - low) / 2 // 避免数组越界
				if a[mid] == val {
					return mid
				} else if a[mid] < val {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			return -1
		}

二. 哈希表:
	实现方式:
		- 链表法 (常用)
		- 开放寻址法

	为什么要动态扩容？
		在不可预知将来有多少数据会插入到哈希表的情况下，只能给哈希表预先设置一个起始大小。当不停往哈希表中插入数据，哈希表中数据会越来越多，
		不管是哪种哈希冲突解决方案，哈希表的性能都会随着装入数据的个数而降低。当哈希表性能下降到一定程度之后，就触发扩容。

	怎么判断性能下降到一定程度呢？
		装载因子=数据个数/槽的个数
		对于基于 开放寻址法 解决哈希冲突 的 哈希表，装载因子要小于1；对于基于链表法解决哈希冲突的哈希表，装载因子可以大于1。
		当装载因子大于某个阈值时，我们就要对哈希表进行扩容。

	怎么实现动态扩容？
		哈希表的扩容，类似数组的扩容。申请一个更大的哈希表（比如2倍大小），将原哈希表中的数据搬移到新哈希表。
		!!! 对于比较大的哈希表，比如1GB大小，数据重新计算哈希值和搬移比较耗时。

	如何避免集中扩容？
		为了解决集中扩容耗时过多的问题，我们将扩容操作穿插在多次插入操作的过程中完成, 分批次完成(类似均摊)
		此时的就需要维护 原哈希表 和 新哈希表, 查询操作就需要在两个哈希表中进行, 但是综合考量, 这种方式比 集中扩容 有明显优势

	哈希表的应用: 位图, 布隆过滤器
		有1千万个整数，范围 0-10 亿，如何快速判定某个数据是否出现在这1千万个数据中？

		能否使用 位图 存储下来?
			设置位图大小为 4 千万, 通过哈希函数，让哈希落在4千万范围内，比如 f(x)=x%n。其中，x表示数据，n表示位图的天小（n=4千万）。

		减少存储空间带来的问题:
			哈希冲突

		哈希冲突导致误判：
			- 判定为存在——>有可能不存在（误判）例如：3123和1123哈希值相同，存储了3123，没有存储1123，查询1123也会返回true
			- 判定为不存在——>肯定不存在（不会误判）例如：查询3123返回false，那就3123就一定不存储

		存在冲突和误判的位图适用的场景：
			把这种存在冲突误判的位图，叫做布隆过滤器比如在访问数据库查询数据前，先访问内存中的位图，如果经过位图判定数据不存在，
			就不需要继续访问数据库，这样就能减少数据库操作。

		布隆过滤器降低冲突概率的方法：
			用多个二进制位率表示一个数据 a、b 经过 3 个 哈希函数 得到的值都相等的概率肯定要小很多
*/
