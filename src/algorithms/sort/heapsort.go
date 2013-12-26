package sort

import (
	"algorithms/types"
)

func siftDown(data types.Sortable, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1){
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func HeapSort(data types.Sortable, a, b int) {
	first := a
	lo := 0
	hi := b-a

	for i := (hi - 1)/2; i >= 0; i--{
		siftDown(data, i, hi, first)
	}

	for i := hi -1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

