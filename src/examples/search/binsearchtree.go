package main

import (
	"fmt"
	atypes "algorithms/types"
	"math/rand"
)

func main(){
	var c = new(atypes.BinSearchTree)
	var length = 10
	for i:= 0;i < length; i++{
		var key = atypes.Integer(rand.Intn(1000))
		var val = atypes.Val(i)
		c.Put(key, val)
	}
	c.Show()
	//fmt.Println(c.Get(atypes.Integer(10)))
	fmt.Println(c.Floor(atypes.Integer(100)))
	fmt.Println(c.Select(0))
	fmt.Println(c.Select(1))
	fmt.Println(c.Select(2))
	fmt.Println(c.Select(3))
	fmt.Println(c.Select(4))
	fmt.Println(c.Select(5))
	fmt.Println(c.Select(6))
	fmt.Println(c.Select(7))
	fmt.Println(c.Select(8))
	fmt.Println(c.Min())
	fmt.Println(c.Rank(atypes.Integer(1000)))

	c.Delete(atypes.Integer(318))
	c.Show()
	fmt.Println(c.Min())

	c.DeleteMin()
	c.Show()
	fmt.Println(c.Min())
/*
	c.DeleteMin()
	c.Show()
	fmt.Println(c.Min())

	c.DeleteMin()
	c.Show()
	fmt.Println(c.Min())

	c.DeleteMin()
	c.Show()
	fmt.Println(c.Min())

	c.DeleteMin()
	c.Show()
	fmt.Println(c.Min())
	*/
}



