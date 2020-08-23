package main

import "fmt"

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func main() {
	//preorderTraversal(nil)
	fmt.Print(buildTreePreIn([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
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

// Maximum depth of Binary Tree

var depth int

func DepthTraverse(root *TreeNode, level int) {
	if root != nil {
		if level > depth {
			depth = level
		}
		if root.Left != nil {
			DepthTraverse(root.Left, level+1)
		}
		if root.Right != nil {
			DepthTraverse(root.Right, level+1)
		}
	}
}


func maxDepth(root *TreeNode) int {
	depth = 0
	DepthTraverse(root, 1)
	return depth
}

// Path Sum using DFS

func hasPathSum(root *TreeNode, sum int) bool {
	right := false
	left := false
	if root == nil {
		return false
	}
	isLeaf := root.Left == nil && root.Right == nil
	currentVal := sum - root.Val
	if currentVal  == 0 && isLeaf {
		return true
	}
	if root.Left != nil {
		left = hasPathSum(root.Left, currentVal)
	}
	if root.Right != nil {
		right = hasPathSum(root.Right, currentVal)
	}
	return left || right
}

// Merge InOrder and PostOrder array to Tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}
	//fmt.Println("============================")
	//fmt.Println("inorder", inorder)
	//fmt.Println("postorder", postorder)

	val := postorder[len(postorder) - 1]
	rootPos := 0
	for i, v := range inorder {
		if v == val {
			rootPos = i
			break
		}
	}
	//fmt.Println(rootPos)
	//fmt.Println("============================"
	return &TreeNode{
		Val:   val,
		Left:  buildTree(inorder[:rootPos], postorder[:rootPos]),
		Right: buildTree(inorder[rootPos+1:], postorder[rootPos:len(postorder) - 1]),
	}
}

// Merge InOrder and PostOrder array to Tree

func buildTreePreIn(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 && len(preorder) == 0 {
		return nil
	}
	//fmt.Println("============================")
	//fmt.Println("inorder", inorder)
	//fmt.Println("preorder", preorder)

	val := preorder[0]
	rootPos := 0
	for i, v := range inorder {
		if v == val {
			rootPos = i
			break
		}
	}
	//fmt.Println(rootPos)
	//fmt.Println("============================")
	return &TreeNode{
		Val:   val,
		Left:  buildTree(preorder[1:rootPos+1], inorder[:rootPos]),
		Right: buildTree(preorder[rootPos+1:], inorder[rootPos+1:]),
	}
}