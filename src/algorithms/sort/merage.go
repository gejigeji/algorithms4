package sort

import (
	"algorithms/types"
)

func Merge(c types.Sortable) {
	var aux = make(types.Sortable, c.Len())
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
	
	var sort func(types.Sortable, int, int, chan bool)
	var finish = make(chan bool)
	defer close(finish)
	sort = func(c types.Sortable, lo, hi int, wait chan bool) {
		defer func(){wait<-true}()
		if hi<=lo {
			return
		}
		var mid = lo+(hi-lo)/2
		var low = make(chan bool, 1)
		defer close(low)
		var high = make(chan bool, 1)
		defer close(high)
		go sort(c, lo, mid, low)
		go sort(c, mid+1, hi, high)
		<-low
		<-high
		merge(c, lo, mid, hi)
	}

	go sort(c, 0, c.Len()-1, finish)
	<-finish
}

func MergeBU(c types.Sortable){
	var aux = make(types.Sortable, c.Len())
	var merge = func(c types.Sortable, lo, mid, hi int, wait chan int) {
		defer func(){
			wait<-1
		}()
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
	var sort = func(c types.Sortable){
		var l = c.Len()
		for sz:=1;sz<l;sz=sz+sz {
			var count=0
			var wait = make(chan int, 1024)
			for lo:=0;lo < l-sz; lo+=sz+sz {
				count++
				go merge(c, lo, lo+sz-1, min(lo+sz+sz-1, l-1), wait)
			}
			for {
				count -= <-wait
				if count == 0 {
					close(wait)
					break
				}
			}
		}
	}
	sort(c)
}

func min(x, y int) int {
	if x<y {
		return x
	} else {
		return y
	}
}
