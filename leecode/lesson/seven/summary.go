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