package sort

import (
	"algorithms/types"
)

func Quick3Way(a types.Sortable){
	quick3Way(a, 0, a.Len()-1)
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

func Quick3WayBU(a types.Sortable){
	var finish = make(chan bool, 1)
	defer close(finish)
	quick3WayBU(a, 0, a.Len()-1, finish)
	<- finish
}

func quick3WayBU(a types.Sortable, lo, hi int, wait chan bool){
	defer func() {wait <- true}()
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
	var low = make(chan bool,1)
	defer close(low)
	var high = make(chan bool,1)
	defer close(high)
	go quick3WayBU(a, lo, lt-1, low)
	go quick3WayBU(a, gt+1, hi, high)
	<-low
	<-high
}
