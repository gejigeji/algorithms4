package main

import (
	"fmt"
	asort "algorithms/sort"
	"math/rand"
)

type ints []int

func (c ints) Less(i, j int) bool{
	return c[i]<c[j]
}

func (c ints) Len() int {
	return len(c)
}

func (c ints) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main(){
	var c = make(ints, 10)
	for i:=0; i<10; i++ {
		c[i] = rand.Int()
	}
	fmt.Printf("before:%v\n", c)
	asort.Insertion(c)
	fmt.Printf("after:%v\n", c)
}
