package main

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil;
	}
	return helperArrToBST(nums, 0, len(nums) - 1)
}

func helperArrToBST (nums []int, low int, high int)  *TreeNode {
	if low > high { // Done
		return nil
	}
	mid := low + (high - low) / 2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = helperArrToBST(nums, low, mid - 1)
	node.Right = helperArrToBST(nums, mid + 1, high)
	return node
}