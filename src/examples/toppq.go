package main

import (
	"fmt"
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
	var length = 100
	var c = atypes.NewMaxPQ(20)
	for i:=0; i<length; i++ {
		var data = atypes.Integer(rand.Intn(10000000))
		c.Insert(data)
	}
	fmt.Println("Min Data Filter by the MaxPQ")
	for {
		data := c.DelMax()
		if data == nil{
			break
		}
		fmt.Printf("Max Data In The Stack:%v\n", data)
	}
}
