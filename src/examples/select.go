package main

import (
	"fmt"
	atypes "algorithms/types"
	asort "algorithms/sort"
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
	var length = 100
	var c = make(atypes.Sortable, length)
	for i:=0; i<length; i++ {
		var data = atypes.Integer(rand.Intn(10000000))
		c[i]=data
	}
	fmt.Printf("select top n from the collection %v\n", c)
	for i:=length/20; i<length/10; i++{
		data := asort.Select(c, i)
		fmt.Printf("Select %v from the collection got:%v\n", i, data)
	}
}
