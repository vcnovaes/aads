package datastuctures

type MinStack[T comparable] struct {
	stack []node[T]
}

func (ms MinStack[T]) Top() node[T] {
	return ms.stack[len(ms.stack)-1]
}
func (ms *MinStack[T]) Push(elt T) {
	if len(ms.stack) == 0 {
		ms.stack = append(ms.stack, node[T]{value: elt, min_pos: 0})
	}
}

func (ms *MinStack[T]) createNode(elt T) node[T] {
	var min int
	for i := len(ms.stack) - 1; i > 0; i-- {
		if elt < ms.stack[i].value {
			i = ms.stack[i].min_pos
		} else {
			min = ms.stack[i].min_pos
		}
	}
	return node[T]{value: elt, min_pos: min}
}

type node[T comparable] struct {
	value   T
	min_pos int
}
