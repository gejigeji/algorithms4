package sort

import (
	"algorithms/types"
)

func partition(a types.Sortable, lo, hi int) int {
	i := lo
	j := hi+1
	v := lo
	for {
		for i++;a.Less(i, v);i++ {
			if i == hi {
				break
			}
		}
		for j--;a.Less(v, j);j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a.Swap(i, j)
	}
	a.Swap(lo, j)
	return j
}

func quick(a types.Sortable, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	quick(a, lo, j-1)
	quick(a, j+1, hi)
}

func Quick(a types.Sortable) {
	quick(a, 0, a.Len()-1)
}

/*
func QuickBU(a types.Sortable){
	var finish = make(chan bool,1)
	defer close(finish)
	var sort func(types.Sortable, int, int, chan bool)
	sort = func(a types.Sortable, lo, hi int, wait chan bool){
		defer func(){wait <- true}()
		if hi <= lo {
			return
		}
		j := partition(a, lo, hi)
		var low = make(chan bool,1)
		defer close(low)
		var high = make(chan bool,1)
		defer close(high)
		go sort(a, lo, j-1, low)
		go sort(a, j+1, hi, high)
		<-low
		<-high
	}
	go sort(a, 0, a.Len()-1, finish)
	<-finish
}
*/

func QuickBU(a types.Sortable) {
	var finish = make(chan bool,1)
	defer close(finish)
	quickbu(a, 0, a.Len()-1, finish)
	<-finish
}
func quickbu(a types.Sortable, lo, hi int,wait chan bool){
	defer func(){wait<-true}()
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	var low = make(chan bool,1)
	defer close(low)
	var high = make(chan bool,1)
	defer close(high)
	go quickbu(a, lo, j-1, low)
	go quickbu(a, j+1, hi, high)
	<-low
	<-high
}
