package sort

import (
	"algorithms/types"
)

func Merge(c types.Sortable) {
	var aux = make(types.Sortable, c.Len())
	copy(aux, c)
	var merge = func(c types.Sortable, lo, mid, hi int) {
		var i=lo
		var j=mid+1
		for k:=lo; k<=hi; k++ {
			aux[k] = c[k]
		}
		for k := lo; k<=hi; k++ {
			if i>mid {
				c[k] = aux[j]
				j++
			} else if j> hi {
				c[k] = aux[i]
				i++
			} else if aux.Less(j, i) {
				c[k] = aux[j]
				j++
			} else {
				c[k] = aux[i]
				i++
			}
		}
	}
	
	var sort func(types.Sortable, int, int)
	sort = func(c types.Sortable, lo, hi int) {
		if hi<=lo {
			return
		}
		var mid = lo+(hi-lo)/2
		sort(c, lo, mid)
		sort(c, mid+1, hi)
		merge(c, lo, mid, hi)
	}


	sort(c, 0, c.Len()-1)
}


