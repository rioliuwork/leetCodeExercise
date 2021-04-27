package exercise

/**

给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。



示例 1：


输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32
示例 2：


输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23


提示：

树中节点数目在范围 [1, 2 * 104] 内
1 <= Node.val <= 105
1 <= low <= high <= 105
所有 Node.val 互不相同
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var sum int
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	if root.Left != nil {
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Right != nil {
		sum += rangeSumBST(root.Right, low, high)
	}
	return sum
}

/**
官方解答：
方法一：深度优先搜索
思路

按深度优先搜索的顺序计算范围和。记当前子树根节点为 \textit{root}root，分以下四种情况讨论：

1：root 节点为空

返回 00。

2：root 节点的值大于 high

由于二叉搜索树右子树上所有节点的值均大于根节点的值，即均大于 high，故无需考虑右子树，返回左子树的范围和。

3：root 节点的值小于 low

由于二叉搜索树左子树上所有节点的值均小于根节点的值，即均小于 low，故无需考虑左子树，返回右子树的范围和。

4：root 节点的值在 [low,high] 范围内

此时应返回 root 节点的值、左子树的范围和、右子树的范围和这三者之和。
*/

func rangeSumBST1(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST1(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST1(root.Right, low, high)
	}
	return root.Val + rangeSumBST1(root.Left, low, high) + rangeSumBST1(root.Right, low, high)
}

/**
官方解答:
方法二：广度优先搜索
思路

使用广度优先搜索的方法，用一个队列 qq 存储需要计算的节点。每次取出队首节点时，若节点为空则跳过该节点，否则按方法一中给出的大小关系来决定加入队列的子节点。
*/
func rangeSumBST2(root *TreeNode, low, high int) (sum int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			return 0
		}
		if node.Val > high {
			q = append(q, node.Left)
		} else if node.Val < low {
			q = append(q, node.Right)
		} else {
			sum += node.Val
			q = append(q, node.Left, node.Right)
		}
	}
	return
}
