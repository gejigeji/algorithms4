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

	var sort func(types.Sortable, int, int, chan bool)
	var finish = make(chan bool)
	sort = func(c types.Sortable, lo, hi int, finish chan bool){
		defer func(){finish<-true}()
		if hi<=lo {
			return
		}
		var low = make(chan bool)
		var high = make(chan bool)
		var j = partition(c, lo, hi)
		go sort(c, lo, j-1, low)
		go sort(c, j+1, hi, high)
		<-low
		<-high
	}
	go sort(c, 0, c.Len()-1, finish)
	<-finish
}
