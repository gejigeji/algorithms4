package main

import (
	"fmt"
	asort "algorithms/sort"
	atypes "algorithms/types"
	"math/rand"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	var c = make(atypes.Sortable, 10)
	for i:=0; i<10; i++ {
		var data = atypes.Integer(rand.Intn(300))
		c[i] = data
	}
	fmt.Printf("before:%v\n", c)
	asort.Merge(c)
	fmt.Printf("after:%v\n", c)
}
