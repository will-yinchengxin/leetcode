package seven

/*
一. 二叉树
	题型套路：
		- LCA（生命周期评估）最近公共祖先

		- 二叉树 单 / 双 / 循环链表

		- 按照遍历结果反向构建二叉树

		- 二叉树上的最长路径和

二. 字符串匹配算法


三. 堆
	- 堆必须是一个完全二叉树
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
				for i/2 > 0 && h.arr[i] > h.arr[i/2] { // 自下网上堆化
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
			- 排序， 基于堆的排序数据
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