package sort

import (
	"algorithms/types"
)

func Quick3way(c types.Sortable){
	var sort func(types.Sortable, int, int, chan bool)
	var finish = make(chan bool)
	sort = func(c types.Sortable, lo, hi int, finish chan bool){
		defer func(){finish<-true}()
		if hi<=lo {
			return
		}
		var lt = lo
		var i = lo+1
		var gt = hi
		var v = c[lo]

		for i<=gt {
			var cmp = c[i].Compare(v)
			if cmp < 0 {
				c.Swap(lt, i)
				lt++
				i++
			} else if cmp > 0 {
				c.Swap(i, gt)
				gt --
			} else {
				i ++
			}
		}

		var low = make(chan bool)
		var high = make(chan bool)
		go sort(c, lo, lt-1, low)
		go sort(c, gt+1, hi, high)
		<-low
		<-high
	}
	go sort(c, 0, c.Len()-1, finish)
	<-finish
}
