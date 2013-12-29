package types

//import "fmt"

type NodeQ struct{
	key Key
	next *NodeQ
}

type Queue struct{
	head  *NodeQ
	tail *NodeQ
	count int
}

func NewQueue() *Queue{
	return &Queue{count:0}
}

func (q *Queue) Push(key Key){
	var node = &NodeQ{key: key}
	if q.count == 0 {
		q.head = node
		q.tail = node
	}else{
		q.tail.next = node
		q.tail = node
	}
	q.count += 1
}

func (q *Queue) Pop() Key{
	var k Key
	if q.count == 0{
		panic("No node in queue")
	}else if q.count == 1{
		k = q.head.key
		q.head = nil
		q.tail = nil
		q.count -= 1
	}else{
		k = q.head.key
		q.head = q.head.next
	}
	return k
}

