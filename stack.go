package stack

type Stack struct {
	size int
	items []interface{}
}

func New() *Stack {
	return &Stack{0, nil}
}

func (stack *Stack) Push(item interface{}) {
	stack.items = append(stack.items, item)
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	var item interface{} = nil
	if stack.size > 0 {
		item = stack.items[stack.size-1]
		stack.items[stack.size-1] = nil
		stack.size--
	}
	return item
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Peek() interface{} {
	var item interface{} = nil
	if stack.size > 0 {
		item = stack.items[stack.size-1]
	}
	return item
}
