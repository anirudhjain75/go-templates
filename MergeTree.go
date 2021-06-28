package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// both root are nil
	// Either of root is nil
	// None of root is nil
	res := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil || root2 == nil {
		if root1 == nil && root2 != nil{
			return root2
		} else if root2 == nil && root1 != nil{
			return root1
		}
	} else {
		res.Val = root1.Val+root2.Val
		res.Left = mergeTrees(root1.Left, root2.Left)
		res.Right = mergeTrees(root1.Right, root2.Right)
	}

	return res
}
