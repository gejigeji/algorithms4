package main

import (
	"fmt"
	atypes "algorithms/types"
	"math/rand"
)

func main(){
	var c = new(atypes.BinSearchTree)
	var length = 100
	for i:= 0;i < length; i++{
		var key = atypes.Integer(rand.Intn(1000))
		var val = atypes.Val(i)
		c.Put(key, val)
	}
	c.Show()
	//fmt.Println(c.Get(atypes.Integer(10)))
	fmt.Println(c.Floor(atypes.Integer(100)))
	fmt.Println(c.Min())
}



