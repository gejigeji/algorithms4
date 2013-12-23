package main

import (
	"fmt"
	asort "algorithms/sort"
	atypes "algorithms/types"
	"math/rand"
)

func main(){
	var c = make(atypes.Sortable, 10)
	for i:=0; i<10; i++ {
		var data = atypes.Integer(rand.Intn(100))
		c[i] = data
	}
	fmt.Printf("before:%v\n", c)
	asort.Merge(c)
	fmt.Printf("after:%v\n", c)
}
