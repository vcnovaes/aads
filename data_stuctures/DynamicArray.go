package datastuctures

type DynamicArray[T interface{}] struct {
	array []T
	cur   uint
	size  uint
}

func (dynamic_arr *DynamicArray[T]) rearrange(new_size uint) {
	var new_arr = make([]T, new_size)
	for idx, elt := range dynamic_arr.array {
		new_arr[idx] = elt
	}
	dynamic_arr.array = new_arr
	dynamic_arr.size = new_size
}

func (dynamic_arr *DynamicArray[T]) Push(el T) {
	if len(dynamic_arr.array) == int(dynamic_arr.size) {
		dynamic_arr.rearrange(dynamic_arr.size * 2)
	}
	dynamic_arr.array[dynamic_arr.cur] = el
	dynamic_arr.cur++
}
