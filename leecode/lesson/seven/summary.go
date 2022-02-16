package seven

/*
一. 二叉树
	题型套路：
		- LCA（生命周期评估）最近公共祖先

		- 二叉树 单 / 双 / 循环链表

		- 按照遍历结果反向构建二叉树

		- 二叉树上的最长路径和

二. 字符串匹配算法（非重点）
	单模式串匹配算法（在主串中查找一个模式串）
		- BF算法（*）
			暴力匹配算法，朴素字符串匹配算法：如果模式串长度为 m，主串长度为 n，那么在主串中就会有 n-m+1 个长度为 m 的字串
			我们只需要暴力的对比这 n-m+1 个子串与模式串，就算可以找到主串与模式串匹配的子串

			func bf(arrModel []int, lenModel int, arrMain []int, lenMain int) int {
				for i := 0; i < lenMain - lenModel; i++ {
					j := 0
					for j < lenModel {
						if arrMain[i+j] != arrModel[j] {
							break
						}
						j++
					}
					if j == lenModel {
						return i
					}
				}
				return -1
			}
			最好时间复杂度：O（n） 主串: abcd....fghe 模式串: zabc
			最坏时间复杂度：O（n） 主串: aaaa....aaaa 模式串: aaab

		- RK算法（*）
			通过哈希算法对主串中的 n-m+1 个子串分别求哈希值，然后逐个与模式串的哈希值比较，如果某个子串的哈希值与模式串的哈希值相等
			那么就说明这个子串和模式串匹配了（通过第一个哈希值可以进一步推导出下一个子串的哈希值）
		- BM算法
		- KMP算法
			BM和KMP的核心算法是，当模式串中的某个字符不匹配时，能够跳过一些肯定不会匹配的情况，将模式串往后多滑动几位

	多模式串匹配算法（在主串中查找多个模式串）
		falskdfjkladsfjfadjslkfj 中 查找 fal jkl adj，一把梭得到多个结果

		- Trie树（*）
			- 字符串查找（完全匹配，就是普通的查找）
			- 多模式串匹配（敏感词过滤）
			- 前缀匹配（搜索引擎，输入法，浏览器）
		- AC自动机

三. 堆（非重点）
	- 堆必须是一个 完全二叉树
	- 堆里每个节点的值必须 大于等于（小于等于）其子树中每个节点的值， 也就是堆中每个节点的值都 大于等于（小于等于）其左右子节点的值

	如果堆中每个节点的值都 大于等于 其子树中每个节点的值， 我们成为 大顶堆
	如果堆中每个节点的值都 小于等于 其子树中每个节点的值， 我们成为 小顶堆

	堆是完全二叉树， 所以其适合使用数组来存储(非完全二叉树使用数组存储会产生很多空洞， 适合使用链表来存储

					12
				   /   \
				  9     6
				/  \   /  \
			   7    8  3   5
			  /
			 4

	root = 1， 左：2i， 右：2i + 1

	堆上的操作
		- 堆中插入数据
			将新的数据插入到末尾， 然后执行自下而上的堆化 // 时间复杂度 O(h)
			type Heap struct {
				arr   []int // 数组， 从下标 1 开始存储
				n     int	// 可以存储的最大数据个数
				count int	// 已经存储的数据个数
			}

			func initHeap(cap int) *Heap {
				return &Heap{
					arr: make([]int, 0 , cap+1),
					n: cap,
					count: 0,
				}
			}
			func (h *Heap) inset(data int) {
				if h.count >= h.n {
					return // 满了
				}
				h.count++
				h.arr[h.count] = data
				i := h.count
				for i/2 > 0 && h.arr[i] > h.arr[i/2] { // 自下向上堆化
					swap(h.arr, i, i/2)
				}
			}
			func swap(arr []int, i, j int) {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}

		- 取堆顶元素
			func top() {
				if count == 0 {
					return min_value/max_value
				}
				return a[1]
			}
			大顶堆取最大元素， 小顶堆取最小元素

		- 删除堆顶元素
			把最后一个节点放置堆顶，进行自上而下的堆化方式 // O(h)

		- 更新元素值
			如果更新之后的值变小了，就进行自上而下的堆化
			如果更新之后的值变大了，就进行自下而上的堆化

	堆排序：
		包含两个大步骤：
			- 建堆， 将数组中的数据原地组织成一个堆
				存储是按层进行存储的， 然后基于 2i 和 2i + 1 这种方式求下标即可
					- 从小到大
					- 从大到小

				1）方式一：从前往后处理每个元素， 对每个元素执行 自下而上 的堆化 （O(nlog2n)）
					可以主要看叶子节点，数组中的存储，【前半段 | 后半段】，假设有 n 个节点， 那么前半段就是 0 ~ n/2， 后半段就是 （n/2 + 1） ~ n
					前半段一定是非叶子节点，后半段一定为叶子节点（假设为非叶子，那么一定存在 （n/2+1）* 2 的子节点， 即 n+2，即超出了 n， 所以后半段一定为叶子节点）
					那么加入的节点就看新增的节点要与多少个节点进行比较，与层级有关系，假设为 h 层，那么就是 叶子节点为 k 个，h 为 log2_n,那么求解出 k 即可
					那么 k 也就是 n/2, 也就是叶子节点个数为 n/2，最终结果也就是 >= n/2*log2_n, O(nlog2n)
				2）方式二：从后往前处理每个元素，对每个元素执行 自上而下 的堆化（O(n)）

				func buildHeap(arr []int, arrCount int) {
					for i := arrCount/2; i >= 1 ; i++ { // 非叶子节点 1 ~ n/2, i 最小为1（从 1 开始计数）
						heapify(arr, arrCount, i)
					}
				}
				func heapify(arr []int, arrCount, i int) {
					for {
						maxPos := i
						if i*2 <= arrCount && arr[i] < arr[i*2] {
							maxPos = i*2
						}
						if i*2+1 <= arrCount && arr[maxPos] < arr[i*2+1] {
							maxPos = i*2+1
						}
						if maxPos == i {
							break
						}
						swap(arr, i, maxPos)
						i = maxPos
					}
				}
				func swap(arr []int, i, maxPos int) {
					arr[i], arr[maxPos] = arr[maxPos], arr[i]
				}

			- 排序， 基于堆的排序数据(n/2^n/2 <= n! <= n^n)
				func sort(arr []int, arrCount int) {
					buildHeap(arr, arrCount)
					k := arrCount
					for k > 1 {
						swap(arr, 1, k)
						k--
						heapify(a, k, 1)
					}
				}

				- 原地排序
				- 非稳定排序
				- 时间复杂度 O（nlogn）
				- 空间复杂度 O（1）

		题型：
			- 优先级队列

			- TOP K
				维护一个大小为 K 的 小顶堆，当有数据添加到集合中时
				如果堆中的数据个数小于 K, 我们直接将数据插入小顶堆
				如果队中的数据个数等于K， 我们就拿新添加的数据与堆顶元素比较
					- 如果新添加的数据大于堆顶元素，我们就把堆顶元素删除，并将这个新添加的数据插入到堆中
					- 如果新添加的数据小于堆顶元素，则不做任何处理
				也就是说小顶堆中一直维护当前数据集合中的 TOP K
				每次询问当前数据的 TOP K 操作就变得非常高效，直接输出小顶堆中的元素即可 （小顶堆维护TOP K.png）

			- 求中位数
				一直维护具有一下特征的两个堆
					- 一个 大顶堆， 一个 小顶堆
					- 每个堆中的元素个数接近 n/2, 如果 n 是偶数，两个堆中的数据个数都是 n/2, 如果 n 是奇数，大顶堆有 2n+1 个元素
						小顶堆中有 n/2 个元素
					- 大顶堆中的元素都小于小顶堆中的元素
				那么 大顶堆 中的 堆顶元素 就是 中位数

				那么再插入元素的过程中， 怎样维护两个堆继续满足上述特征呢？
					如果新数据小于等于大顶堆的堆顶元素，就将这个新数据插入到大顶堆，否则，插入到小顶堆
					不过，此时就有可能出现两个堆中的数据个数不符合前面的约定，我们通过一个堆不停的将堆顶元素移动到另一个堆，通过这样的调整，让
					两个堆中的数据个数满足前面的约定

			- 求百分位数
				一直维护具有一下特征的两个堆
				假设当前数据个数是n， 大顶堆中保存前 99%*n 个数据，小顶堆中保存后 1%*n 个数据，大顶堆堆顶元素就是我们要找的 99百分位 数据

				当插入新数据时，我们根据这个数据跟大顶堆堆顶的元素的大小关系，决定将其插入到那个堆中，如果新数据小于等于大顶堆堆顶元素，那么就将其插入大顶堆
				否则就将其插入小顶堆

				为了保持大顶堆中的数据占 99%，小顶堆中的数据占 1%， 我们需要重新计算大顶堆和小顶堆中的数据个数，看是否还符合 99：1 ，如果不符合
				我们就将其中一个堆中的数据移动到另一个堆中，直到满足这个比例为止
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}