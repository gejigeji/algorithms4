package sort

import (
	"algorithms/types"
)

func Select(a types.Sortable, k int) types.Less{
	lo := 0
	hi := a.Len()-1
	for hi > lo {
		j := partition(a, lo, hi)
		if j == k {
			return a[k]
		} else if j > k {
			hi = j - 1
		} else if j < k {
			lo = j + 1
		}
	}
	return a[k]
}
