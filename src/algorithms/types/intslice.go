// package types for easy use algorithms code.
// the Ints just like buildin package sort.IntSlice
package types

import (
	"math/rand"
)

type Ints []int

func (c Ints) Less(i, j int) bool{
	return c[i]<c[j]
}

func (c Ints) Len() int {
	return len(c)
}

func (c Ints) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func RandIntsn(size, top int) Ints {
	var ret = make(Ints, size)
	for i:=0; i<size; i++ {
		ret[i] = rand.Intn(top)
	}
	return ret
}

func RandInts(size int) Ints {
	var ret = make(Ints, size)
	for i:=0; i<size; i++ {
		ret[i] = rand.Int()
	}
	return ret
}

func PRandIntsn(size, top int) Ints {
	var ret = make(Ints, size)
	var ch = make(chan int, size)
	defer close(ch)
	for i:=0; i<size; i++ {
		go func(){
			ch <- rand.Intn(top)
		}()
	}
	for i:=0; i<size; i++ {
		ret[i] = <-ch
	}
	return ret
}

func PRandInts(size int) Ints {
	var ret = make(Ints, size)
	var ch = make(chan int, size)
	defer close(ch)
	for i:=0; i<size; i++ {
		go func(){
			ch <- rand.Int()
		}()
	}
	for i:=0; i<size; i++ {
		ret[i] = <-ch
	}
	return ret
}
