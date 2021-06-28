package main

type StackImplemented []byte

func (s StackImplemented) Push(v byte) StackImplemented {
	return append(s, v)
}

func (s StackImplemented) Pop() (StackImplemented, byte) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return  s[:l-1], s[l-1]
}

func (s StackImplemented) Peek() byte {

	return  s[len(s)-1]
}

func isValid(s string) bool {

	sByte := []byte(s)
	stack := make(StackImplemented, 0, len(s))
	if len(sByte) == 0 {
		return true
	}

	for _, v := range sByte {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		if stack.Peek() == '[' && v == ']' {
			stack, _ = stack.Pop()
		} else if stack.Peek() == '(' && v == ')' {
			stack, _ = stack.Pop()
		} else if stack.Peek() == '{' && v == '}' {
			stack, _ = stack.Pop()
		} else {
			if v == '[' || v == '{' || v == '(' {
				stack.Push(v)
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
