package main

import (
	"fmt"
	asort "algorithms/sort"
	"algorithms/types"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	var c = types.PRandIntsn(100, 1000)
	fmt.Printf("before:%v\n", c)
	asort.Shell(c)
	fmt.Printf("after:%v\n", c)


	//my type Sortable test
	var a0,a1,a2,a3,a4 types.Float = 1,3,1.3,5.7,0
	var a5 types.Integer = 3
	var a = types.Sortable{a0,a1,a2,a3,a4,a5}
	fmt.Printf("before:%v\n", a)
	asort.MergeBU(a)
	fmt.Printf("after:%v\n", a)
	//end my test

}
