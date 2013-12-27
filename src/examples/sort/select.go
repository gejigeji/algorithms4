package main

import (
	"fmt"
	asort "algorithms/sort"
	atypes "algorithms/types"
	"math/rand"
	"runtime"
	"runtime/pprof"
	"flag"
	"log"
	"os"
)

func main(){
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	if *cpuprofile==""{
		run()
	} else {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		run()
	}
}

func run(){
	var length = 10
	var c = make(atypes.Sortable, length)
	for i:=0; i<length; i++ {
		var data = atypes.Integer(rand.Intn(1000000000))
		c[i] = data
	}
	fmt.Printf("before:%v\n", c)
	var mink atypes.Less
	var k int = 9
	mink = asort.Select(c,k)
	asort.Quick(c)
	fmt.Printf("after:%v\n", c)
	fmt.Printf("%v smallest:%v\n", k, mink)
}

