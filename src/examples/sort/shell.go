package main

import (
	"fmt"
	asort "algorithms/sort"
	"algorithms/types"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	var c = datastruct.PRandIntsn(100, 1000)
	fmt.Printf("before:%v\n", c)
	asort.Shell(c)
	fmt.Printf("after:%v\n", c)
}
