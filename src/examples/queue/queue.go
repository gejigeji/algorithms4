package main

import (
	"algorithms/types"
	"fmt"
)

func main(){
	var q = new(types.Queue)

	q.Push(types.Integer(0))
	q.Push(types.Integer(1))
	q.Push(types.Integer(0))
	q.Push(types.Integer(8))
	q.Push(types.Integer(9))
	q.Push(types.Integer(0))
	q.Push(types.Integer(4))
	q.Push(types.Integer(5))

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
