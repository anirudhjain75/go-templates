package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var xV int
var yV int
var xL int
var yL int
var xP int
var yP int

func traverseThroughBinaryTree(root *TreeNode, level int) int {
	var left int
	var right int
	if root.Left != nil {
		left = traverseThroughBinaryTree(root.Left, level + 1)
	}
	if root.Right != nil {
		right = traverseThroughBinaryTree(root.Right, level + 1)
	}

	if left == xV || right == xV {
		xP = root.Val
		xL = level
	}

	if left == yV || right == yV {
		yP = root.Val
		yL = left
	}

	return root.Val
}

func isCousins(root *TreeNode, x int, y int) bool {
	xV = x
	yV = y
	traverseThroughBinaryTree(root, 0)

	if xL == 0 || yL == 0 {
		return false
	}

	if xP == yP {
		return false
	}

	return xL == yL
}

func validate(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}
	left := validate(root.Left, min, &root.Val)
	right := validate(root.Right, &root.Val, max)

	res := true

	if root.Left != nil && root.Left.Val >= root.Val  {
		res = false
	}
	if root.Right != nil && root.Right.Val <= root.Val && root.Right.Val > *min {
		res = false
	}

	checker := false
	if root.Val > *min && root.Val < *max {
		checker = true
	}

	return res && left && right && checker
}

func isValidBST(root *TreeNode) bool {

	// Validate by length

	min := -math.MaxInt64
	max := math.MaxInt64

	return validate(root, &min, & max)
}

// Using Breadth First Search Logic
func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	res = append(res, []int{})

	currentLevel := make([]*TreeNode, 0)
	nextLevel := make([]*TreeNode, 0)

	currentLevel = append(currentLevel, root)
	leftToRight := true;

	currentLevelNumber := 0

	for len(currentLevel) != 0 {
		temp := currentLevel[len(currentLevel) - 1]
		res[currentLevelNumber] = append(res[currentLevelNumber], temp.Val)
		currentLevel = currentLevel[:len(currentLevel) - 1]
		if leftToRight {
			if temp.Left != nil {
				nextLevel = append(nextLevel, temp.Left)
			}
			if temp.Right != nil {
				nextLevel = append(nextLevel, temp.Right)
			}
		} else {
			if temp.Right != nil {
				nextLevel = append(nextLevel, temp.Right)
			}
			if temp.Left != nil {
				nextLevel = append(nextLevel, temp.Left)
			}
		}

		if len(currentLevel) == 0 {
			leftToRight = !leftToRight;
			currentLevel = append(currentLevel, nextLevel...)
			res = append(res, []int{})
			currentLevelNumber++
			nextLevel = make([]*TreeNode, 0)
		}
	}
	return res[:len(res)-1]
}
