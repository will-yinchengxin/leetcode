package six
/*
二叉树:
	- 二叉树(BT)
		概念:
			根节点, 父节点, 子节点, 左子节点, 右子节点, 叶子节点, 兄弟节点, 子树, 左子树, 右子树, 高度, 深度, 层
			二叉树(n叉树), 满二叉树, 完全二叉树

			完全二叉树: 一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，
						如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，则这棵二叉树称为完全二叉树

		二叉树的存储:
			- 基于指针存储:
				type BT_Node struct {
					val interface
					left *BT_Node
					right *BT_Node
				}
			- 基于数组存储:
				根节点存储下标为 1 的位置, 节点之间父子关系根据数组下标计算获得
				例: 如果数组下标为 i
					- 下标为 i*2的位置存储 左子节点
					- 下标为 i*2+1 的位置存储 右子节点
					- 下标为 i/2 的位置存储的就是他的父节点
			- 基于数组存储为什么从下标为 1 的位置开始存储?
				因为通过浪费一个节点可以简化计算公式, 如果从 0 开始计算公式就变成了 左: i*2+1 右: i*2+2
				这样公式是复杂的

		二叉树的遍历(二叉树的遍历类似于链表的遍历):
			- 前序遍历: 根->左子树->右子树
			- 中序遍历: 左子树->根->右子树
			- 后续遍历: 左子树->右子树->根

			前序遍历:
				func preOrder(root Node) {
					if root == nil {
						return // 终止条件, 也就是看那种情况的最小子问题满足条件
					}
					fmt.Println(root.data) // 可能为其他操作,这里使用打印代替
					preOrder(root.left)
					preOrder(root.right)
				}
			时间复杂度: O(n)    空间复杂度:O(h), n元素格式, h二叉树的高度

			时间复杂度通过递归树来分析: 每个节点 递 和 归 的过程都是常量级别的 C, 总体也就是 O(nC) 化简为 O(n)

			空间复杂度与树的高度 h 有关,这里 h 是从1开始计数的,
			那么一个节点为 n 的二叉树的高度是多少呢?
				max(h) = n //线性排列
				min(h) = log2(n) // 完全二叉树
				树的高度 h 介于二者之间

			那么树的高度为 h 的二叉树有多少个节点呢?
				min(n) = h // 树的高度,也就是 线性排列 的树
				max(n) = 2^h - 1 // 满二叉树
				节点个数 n 介于二者之间

			前/中/后序遍历 的逻辑是一致的,只是操作的时机(例如打印)不同, 时间/空间 复杂度都是一致的

			tips: 树和子树的处理逻辑一般都是相同的, 一般采用递归来实现, 通过分解为最小子问题来求解父问题

		二叉树遍历总结:
			前/中/后 可以看作树上的 深度优先遍历, 也属于回溯算法, 根据节点的不同分为 前/中/后序遍历
				func preOrder(root Node) {
					if root == nil {
						return
					}
					fmt.Println(root.data) // 前序遍历
					preOrder(root.left)
					preOrder(root.right)
				}
				func middleOrder(root Node) {
					if root == nil {
						return
					}
					middleOrder(root.left)
					fmt.Println(root.data) // 中序遍历
					middleOrder(root.right)
				}
				func lastOrder(root Node) {
					if root == nil {
						return
					}
					lastOrder(root.left)
					lastOrder(root.right)
					fmt.Println(root.data) // 后序遍历
				}


	- 二叉查找树(BST)
		是一种特殊的二叉树, 支持快速的 查找, 插入, 删除 数据

		对于 二叉查找树(BST) 任意一个节点
			- 左子树每个节点的值都小于这个节点
			- 右子树每个节点的值都大于这个节点

		二叉查找树(BST) 查找操作:
			- 先取根节点, 如果 root == target, 直接返回
			- 如果 root > target, 在左子树中继续查找
			- 如果 root < target, 在右子树中继续查找
			//------------------------递归实现------------------------------
			type Node struct {
				val int
				left *Node
				right *Node
			}
			func constructor(val int) Node {
				return Node {
					val: val,
				}
			}
			func findR(root *Node, target int) *Node {
				if root == nil {
					return nil
				}
				if root.val == target {
					return root
				}
				if root.val > target {
					return findR(root.left, target)
				} else {
					return findR(root.right, target)
				}
			}
			时间复杂度: O(h) 空间复杂度:O(h)
			//----------------------------非递归实现-----------------------------
			func find(root *Node, target int) *Node {
				p := root
				for p != nil {
					if target < p.val {
						p = p.left
					} else if target > p.val {
						p = p.right
					} else {
						return p
					}
				}
				return nil
			}
			时间复杂度: O(h) 空间复杂度:O(1)


		二叉查找树(BST) 插入操作:
			- 插入的元素比当前的节点值小
				- 当前节点左子树为空, 数据插入 左子节点位置
				- 当前节点左子树不为空, 继续遍历左子树, 寻找插入的位置
			- 插入的元素比当前的节点值大
				- 当前节点右子树为空, 数据插入 右子节点位置
				- 当前节点右子树不为空, 继续遍历右子树, 寻找插入的位置
			//------------------------递归实现------------------------------
			func insert(root *Node, target int) {
				if root == nil {
					return
				}
				insertR(root, target)
			}
			func insertR(root *Node, target int) {
				if target > root.val {
					if root.right == nil {
						root.right = &Node{
							val: target,
						}
					} else {
						insertR(root.right, target)
					}
				} else {
					if root.left == nil {
						root.left = &Node{
							val: target,
						}
					} else {
						insertR(root.left, target)
					}
				}
			}
			时间复杂度: O(h) 空间复杂度:O(h)
			//-------------------------非递归实现----------------------------
			func insertR(root *Node, target int) {
				if root == nil {
					return
				}
				p := root
				for p != nil {
					if target > p.val {
						if p.right == nil {
							p.right = &Node{
								val: target,
							}
							return
						}
						p = p.right
					} else {
						if p.left == nil {
							p.left = &Node{
								val: target,
							}
							return
						}
						p = p.left
					}
				}
			}
			时间复杂度: O(h) 空间复杂度:O(1)

		二叉查找树(BST) 删除操作:
			- 要删除节点没有子节点(简单)
				只需要直接将父节点中指向要删除节点的指针置为 nil 即可
			- 要删除节点只有一个子节点(中等)

			- 要删除节点有两个子节点(复杂)


	- 平衡二叉查找树(BBST) // 经典: 红黑树

*/
type Node struct {
	val int
	left *Node
	right *Node
}
func insertR(root *Node, target int) {
	if root == nil {
		return
	}
	p := root
	for p != nil {
		if target > p.val {
			if p.right == nil {
				p.right = &Node{
					val: target,
				}
				return
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = &Node{
					val: target,
				}
				return
			}
			p = p.left
		}
	}
}
