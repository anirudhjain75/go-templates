package main

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func main() {
	preorderTraversal(nil)
}

// Traverse through binary tree
// Pre-Order Traversal
var preOrderRes []int

func preOrderRecursiveTraverse(root *TreeNode) {
	if root != nil {
		preOrderRes = append(preOrderRes, root.Val)
		if root.Left != nil {
			preOrderRecursiveTraverse(root.Left)
		}
		if root.Right != nil {
			preOrderRecursiveTraverse(root.Right)
		}
	}
}

func preorderTraversal(root *TreeNode) []int {
	preOrderRes = make([]int, 0)
	preOrderRecursiveTraverse(root)
	return preOrderRes
}

// InOrder Traversal

var InOrderRes []int

func inOrderRecursiveTraverse(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			inOrderRecursiveTraverse(root.Left)
		}
		InOrderRes = append(InOrderRes, root.Val)
		if root.Right != nil {
			inOrderRecursiveTraverse(root.Right)
		}
	}
}

func inorderTraversal(root *TreeNode) []int {
	InOrderRes = make([]int, 0)
	inOrderRecursiveTraverse(root)
	return InOrderRes
}

// PostOrder Traversal

var PostOrderRes []int

func PostOrderRecursiveTraverse(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			PostOrderRecursiveTraverse(root.Left)
		}
		if root.Right != nil {
			PostOrderRecursiveTraverse(root.Right)
		}
		PostOrderRes = append(PostOrderRes, root.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	PostOrderRes = make([]int, 0)
	PostOrderRecursiveTraverse(root)
	return PostOrderRes
}

// Breadth First Search

var BFSMatrix [][]int

func levelOrder(root *TreeNode) [][]int {
	BFSMatrix = make([][]int, 0)
	if root != nil {
		Queue := make([]*TreeNode, 0)
		Queue = append(Queue, root)
		NextLevel := make([]*TreeNode, 0)
		ElemsInLevel := make([]int, 0)
		for len(Queue) != 0 {
			CurrentNode := Queue[0]
			ElemsInLevel = append(ElemsInLevel, CurrentNode.Val)
			Queue = Queue[1:]
			if CurrentNode.Left != nil {
				NextLevel = append(NextLevel, CurrentNode.Left)
			}
			if CurrentNode.Right != nil {
				NextLevel = append(NextLevel, CurrentNode.Right)
			}
			if len(Queue) == 0 && len(NextLevel) > 0 {
				Queue = NextLevel
				NextLevel = []*TreeNode{}
				BFSMatrix = append(BFSMatrix, ElemsInLevel)
				ElemsInLevel = []int{}
			}
		}
		if len(ElemsInLevel) > 0 {
			BFSMatrix = append(BFSMatrix, ElemsInLevel)
		}
	}
	return BFSMatrix
}