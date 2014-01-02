package types

type Color bool

var RED Color = true
var BLACK Color = false

type BalSearchTree struct{
	root *NodeRB
}

type NodeRB struct{
	key Key
	val Val
	left, right *NodeRB
	N int
	color Color
}

func isRed(node NodeRB) bool{
	if node.color == RED{
		return true
	}else{
		return false
	}
}

func (bst *BalSearchTree) Size() int{
	if bst.root != nil{
		return size(bst.root)
	}else{
		return 0
	}
}
