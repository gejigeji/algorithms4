package sort

import (
	"algorithms/types"
)

func Quick3Way(a types.Sortable){
	quick(a, 0, a.Len()-1)
}

func quick3Way(a types.Sortable, lo, hi int){
	if  hi <= lo{
		return
	}
	lt := lo
	i := lo+1
	gt := hi
	v := a[lo]
	for i <= gt{
		cmp := a[i].Compare(v)
		if cmp < 0{
			a.Swap(lt,i)
			lt++
			i++
		}else if cmp > 0{
			a.Swap(i,gt)
			gt--
		}else{
			i++
		}
	}
	quick3Way(a, lo, lt-1)
	quick3Way(a, gt+1, hi)
}
