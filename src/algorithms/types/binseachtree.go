package types

import "fmt"

type BinSearchTree struct{
	root *Node
}

func NewBinST() *BinSearchTree{
	return &BinSearchTree{root:nil}
}

type Key Compare

type Val int

type Node struct{
	key Key
	val Val
	left, right *Node
	N int
}

func (bst *BinSearchTree) isEmpty() bool{
	if bst.root == nil{
		return true
	}else{
		return false
	}
}

func (bst *BinSearchTree) check(){
	if bst.isEmpty(){
		var message = fmt.Sprintf("There isn't any node in binsearchtree %v\n",bst)
		panic(message)
	}
}


//Return size of the binsearchtree
func (bst *BinSearchTree) Size() int{
	if !bst.isEmpty(){
		return size(bst.root)
	}else{
		return 0
	}
}


func size(other interface{}) int{
	switch o := other.(type){
	case *NodeRB:
		if o == nil{
			return 0 
		}
		return o.N
	case *Node:
		if o == nil{
			return 0 
		}
		return o.N
	default:
		panic("Can not use method size()")
	}
}
/*
func size(node *Node) int{
	if node == nil {
		return 0
	}
	return node.N
}*/

//Return the value of node with Key key
func (bst *BinSearchTree) Get(key Key) Val{
	bst.check()
	return *get(bst.root, key)
}

func get(node *Node, key Key) *Val{
	if node == nil{
		return nil
	}
	var cmp int = key.Compare(node.key)
	if cmp < 0{
		return get(node.left, key)
	}else if cmp > 0{
		return get(node.right, key)
	}else{
		return get(node.right, key)
	}
}

//Put node with Key key and Val val in the binsearchtree
func (bst *BinSearchTree) Put(key Key, val Val) {
	bst.root = put(bst.root, key, val)
}

func put(node *Node, key Key, val Val) *Node{
	if node == nil{
		return &Node{key, val, nil, nil, 1}
	}
	var cmp = key.Compare(node.key)
	if cmp < 0 {
		node.left = put(node.left, key, val)
	}else if cmp > 0{
		node.right = put(node.right, key, val)
	}else {
		node.val = val
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}

//Show all node in the binsearchtree
func (bst *BinSearchTree) Show(){
	bst.check()
	show(bst.root)
}

func show(node *Node){
	if node != nil{
		fmt.Println("        node:  key: ",node.key, " value: ", node.val, " size: ",node.N)
		if node.left != nil{
			fmt.Println(" left  child:  key: ",node.left.key, " value: ", node.left.val, " size: ", node.left.N)
		}else{
			fmt.Println(" left  child:  nil")
		}
		if node.right != nil{
			fmt.Println(" right child:  key: ",node.right.key, " value: ", node.right.val, " size: ", node.right.N, "\n")
		}else{
			fmt.Println(" right child:  nil\n")
		}
		show(node.left)
		show(node.right)
	}
}

//Return node with minimum Key
func (bst *BinSearchTree) Min() Key{
	bst.check()
	return min(bst.root).key
}

func min(node *Node) *Node{
	if node.left == nil{
		return node
	}else{
		return min(node.left)
	}
}

//Return node with maximum Key
func (bst *BinSearchTree) Max() Key{
	bst.check()
	return max(bst.root).key
}

func max(node *Node) *Node{
	if node.right == nil{
		return node
	}else{
		return max(node.right)
	}
}


//Return node with Key not bigger than key
func (bst *BinSearchTree) Floor(key Key) Key{
	bst.check()
	return floor(bst.root, key).key
}

func floor(node *Node, key Key) *Node{
	if node == nil {
		return nil
	}
	var cmp = key.Compare(node.key)
	if cmp == 0 {
		return node
	}else if cmp < 0 {
		return floor(node.left, key)
	}
	var t = floor(node.right, key)
	if t != nil {
		return t
	}else{
		return node
	}
}

func (bst *BinSearchTree) Select(k int) Key{
	bst.check()
	return selectNode(bst.root, k).key
}

//Return *Node containing key of rank k
func selectNode(node *Node, k int) *Node{
	if node == nil {
		return nil
	}
	var t = size(node.left)
	if t > k{
		return selectNode(node.left, k)
	}else if t < k{
		return selectNode(node.right, k-t-1)
	}else{
		return node
	}
}

//Return rank of Key key
func (bst *BinSearchTree) Rank(key Key) int{
	bst.check()
	return *rank(bst.root, key)
}

//Return number of nodes with key less than key in the subtree rooted at node.
func rank(node *Node, key Key) *int{
	if node == nil{
		var tmp = 0
		return &tmp
	}
	var cmp = key.Compare(node.key)
	if cmp < 0 {
		return rank(node.left, key)
	}else if cmp > 0{
		var tmp = 1 + size(node.left) + *rank(node.right, key)
		return &tmp
	}else{
		var tmp = size(node.left)
		return &tmp
	}
}

//Delete node with minimum Key
func (bst *BinSearchTree) DeleteMin(){
	bst.check()
	bst.root = deleteMin(bst.root)
}

func deleteMin(node *Node) *Node{
	if node.left == nil{
		return node.right
	}
	node.left = deleteMin(node.left)
	node.N = size(node.left) + size(node.right) + 1
	return node
}

//Delete node with Key key
func (bst *BinSearchTree) Delete(key Key){
	bst.check()
	bst.root = deleteNode(bst.root, key)
}

func deleteNode(node *Node, key Key) *Node{
	if node == nil {
		return nil
	}
	var cmp = key.Compare(node.key)
	if cmp < 0 {
		node.left = deleteNode(node.left, key)
	}else if cmp > 0 {
		node.right = deleteNode(node.right, key)
	}else{
		if node.right == nil{
			return node.left
		}
		if node.left == nil{
			return node.right
		}
		var t = node
		var minNode = min(t.right)
		node.key = minNode.key
		node.val = minNode.val
		node.right = deleteMin(t.right)
		node.left = t.left
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}

//Golang has no simple override method
//func (bst *BinSearchTree) Keys() Queue{
//	return keys(bst.min(), bst.max())
//}

//Return Queue of node with Key between Key lo and Key hi
func (bst *BinSearchTree) Keys(lo Key, hi Key) *Queue{
	bst.check()
	var q = new(Queue)
	keys(bst.root, q, lo, hi)
	return q
}

func keys(node *Node, q *Queue, lo Key, hi Key){
	if node == nil{
		return
	}
	var cmplo = lo.Compare(node.key)
	var cmphi = hi.Compare(node.key)
	if cmplo < 0{
		keys(node.left, q, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0{
		q.Push(node.key)
	}
	if cmphi > 0 {
		keys(node.right, q, lo, hi)
	}
}


