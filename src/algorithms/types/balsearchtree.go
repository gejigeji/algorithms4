package types

import "fmt"

type Color bool

var RED Color = true
var BLACK Color = false

type BalSearchTree struct{
	root *NodeRB
}

func NewBalST() *BalSearchTree{
	return &BalSearchTree{root:nil}
}

type NodeRB struct{
	key Key
	val Val
	left, right *NodeRB
	N int
	color Color
}

func isRed(node *NodeRB) bool{
	if node == nil {
		return false
	}
	return node.color == RED
}

func (bst *BalSearchTree) Size() int{
	if bst.root != nil{
		return size(bst.root)
	}else{
		return 0
	}
}

func rotateLeft(h *NodeRB) *NodeRB{
	var x = h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func rotateRight(h *NodeRB) *NodeRB{
	var x = h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func flipColors(h *NodeRB) {
	h.color = RED
	h.left.color = BLACK
	h.right.color = BLACK
}

func (bst *BalSearchTree) Put(key Key, val Val){
	bst.root = putRB(bst.root, key, val)
	bst.root.color = BLACK
}


func putRB(o *NodeRB, key Key, val Val) *NodeRB{
		if o ==  nil{
			return &NodeRB{key, val, nil, nil, 1, RED}
		}
		var cmp = key.Compare(o.key)
		if cmp < 0 {
			o.left = putRB(o.left, key, val)
		}else if cmp > 0 {
			o.right = putRB(o.right, key, val)
		}else {
			o.val = val
		}
		if isRed(o.right) && !isRed(o.left) {
			o = rotateLeft(o)
		}
		if isRed(o.left.left) && isRed(o.left) {
			o = rotateRight(o)
		}
		if isRed(o.left) && isRed(o.right){
			flipColors(o)
		}
		o.N = size(o.left) + size(o.right) + 1
		return o
}

func (bst *BalSearchTree) Show(){
	showRB(bst.root)
}

//Show all node in the binsearchtree
func showRB(node *NodeRB){
	if node != nil{
		fmt.Println("        node:  color: ",node.color,"  key: ",node.key, " value: ", node.val, " size: ",node.N)
		if node.left != nil{
			fmt.Println(" left  child:  color: ",node.left.color,"  key: ",node.left.key, " value: ", node.left.val, " size: ", node.left.N)
		}else{
			fmt.Println(" left  child:  nil")
		}
		if node.right != nil{
			fmt.Println(" right child:  color: ",node.right.color,"  key: ",node.right.key, " value: ", node.right.val, " size: ", node.right.N, "\n")
		}else{
			fmt.Println(" right child:  nil\n")
		}
		showRB(node.left)
		showRB(node.right)
	}
}
