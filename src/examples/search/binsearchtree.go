package main

import (
	"algorithms/types"
	"fmt"
	"math/rand"
)

func main(){
	var c = new(types.BinSearchTree)
	var length = 10
	for i:= 0;i < length; i++{
		var key = types.Integer(rand.Intn(1000))
		var val = types.Val(i)
		c.Put(key, val)
	}
	c.Show()
	//fmt.Println(c.Get(types.Integer(10)))
	fmt.Println(c.Floor(types.Integer(100)))
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
	fmt.Println(c.Rank(types.Integer(1000)))

	c.Delete(types.Integer(318))
	c.Show()
	fmt.Println(c.Min())

	c.DeleteMin()
	c.Show()
	fmt.Println(c.Min())
	fmt.Println(c.Max())

	var q = c.Keys(c.Min(), c.Max())

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())


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



