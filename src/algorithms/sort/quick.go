package sort

import (
	"algorithms/types"
)

func Quick(c types.Sortable){
	var partition = func(c types.Sortable, lo, hi int) int{
		var i=lo
		var j=hi+1
		var v = c[lo]
		for{
			for {
				i++
				if !c[i].Less(v) {
					break
				}
				if i==hi {
					break
				}
			}
			for {
				j--
				if !v.Less(c[j]) {
					break
				}
				if j==lo {
					break
				}
			}
			if i>=j {
				break
			}
			c.Swap(i, j)
		}
		c.Swap(lo, j)
		return j
	}

	var sort func(types.Sortable, int, int)
	sort = func(c types.Sortable, lo, hi int){
		if hi<=lo {
			return
		}
		var j = partition(c, lo, hi)
		sort(c, lo, j-1)
		sort(c, j+1, hi)
	}
	sort(c, 0, c.Len()-1)
}
