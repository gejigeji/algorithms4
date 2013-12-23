package main

import (
	"fmt"
	asort "algorithms/sort"
	atypes "algorithms/types"
	"math/rand"
	"runtime"
)

func main(){
	var length = 10
	runtime.GOMAXPROCS(runtime.NumCPU())
	var c = make(atypes.Sortable, length)
	for i:=0; i<length; i++ {
		var data = atypes.Integer(rand.Intn(1000))
		c[i] = data
	}
	fmt.Printf("before:%v\n", c)
	asort.MergeBU(c)
	fmt.Printf("after:%v\n", c)
}
