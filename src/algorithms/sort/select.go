package sort

import ("algorithms/types")

func Select(c types.Sortable, k int) types.Compare{
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
	var lo=0
	var hi = c.Len()-1
	for hi > lo {
		var j = partition(c, lo, hi)
		if j==k {
			return c[k]
		} else if j>k {
			hi = j-1
		} else if j<k {
			lo = j+1
		}
	}
	return c[k]
}
