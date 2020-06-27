package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	copied := new(Node)
	copyMover := copied

	// Index of Original
	m := make(map[*Node]int)
	// Index of Copy
	copyIndex := make([]*Node, 0)

	// Random Index
	randomIndex := make([]*Node, 0)

	index := 0

	for head != nil {
		copyMover.Val = head.Val
		if head.Next != nil {
			copyMover.Next = new(Node)
		}

		m[head] = index
		index++

		randomIndex = append(randomIndex, head.Random)
		copyIndex = append(copyIndex, copyMover)
		copyMover = copyMover.Next
		head = head.Next
	}

	copyMover = copied

	index = 0

	for copyMover != nil {


		if randomIndex[index] == nil {
			copyMover.Random = nil
		} else {
			v, _ := m[randomIndex[index]]
			copyMover.Random = copyIndex[v]
		}
		index++

		copyMover = copyMover.Next
	}

	return copied
}