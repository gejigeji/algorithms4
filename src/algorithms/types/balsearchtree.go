package types

import "fmt"

type Color bool

var RED Color = true
var BLACK Color = false

type BalSearchTree struct {
	root *NodeRB
}

func NewBalST() *BalSearchTree {
	return &BalSearchTree{root: nil}
}

type NodeRB struct {
	key                 Key
	val                 Val
	left, right, father *NodeRB
	N                   int
	color               Color
}

func (bst *BalSearchTree) isEmpty() bool {
	if bst.root == nil {
		return true
	} else {
		return false
	}
}

func (bst *BalSearchTree) check() {
	if bst.isEmpty() {
		var message = fmt.Sprintf("There isn't any node in binsearchtree %v\n", bst)
		panic(message)
	}
}

func isRed(node *NodeRB) bool {
	if node == nil {
		return false
	}
	return node.color == RED
}

func (bst *BalSearchTree) Size() int {
	if bst.root != nil {
		return size(bst.root)
	} else {
		return 0
	}
}

func rotateLeft(h *NodeRB) *NodeRB {
	var x = h.right
	x.father = h.father
	h.right = x.left
	h.father = x
	if x.left != nil {
		x.left.father = h
	}
	x.left = h
	x.color = h.color
	h.color = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func rotateLeftPure(h *NodeRB) *NodeRB {
	var x = h.right
	x.father = h.father
	h.right = x.left
	h.father = x
	if x.left != nil {
		x.left.father = h
	}
	x.left = h
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func rotateRight(h *NodeRB) *NodeRB {
	var x = h.left
	x.father = h.father
	if x.right != nil {
		x.right.father = h
	}
	h.left = x.right
	h.father = x
	x.right = h
	x.color = h.color
	h.color = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func rotateRightPure(h *NodeRB) *NodeRB {
	var x = h.left
	x.father = h.father
	if x.right != nil {
		x.right.father = h
	}
	h.left = x.right
	h.father = x
	x.right = h
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func flipColors(h *NodeRB) {
	h.color = RED
	h.left.color = BLACK
	h.right.color = BLACK
}

func (bst *BalSearchTree) Put(key Key, val Val) {
	bst.root = putRB(bst.root, key, val)
	bst.root.color = BLACK
}

func putRB(o *NodeRB, key Key, val Val) *NodeRB {
	if o == nil {
		return &NodeRB{key, val, nil, nil, nil, 1, RED}
	}
	var cmp = key.Compare(o.key)
	if cmp < 0 {
		o.left = putRB(o.left, key, val)
		o.left.father = o
	} else if cmp > 0 {
		o.right = putRB(o.right, key, val)
		o.right.father = o
	} else {
		o.val = val
	}
	if isRed(o.right) && !isRed(o.left) {
		o = rotateLeft(o)
	}
	if o.left != nil {
		if isRed(o.left.left) && isRed(o.left) {
			o = rotateRight(o)
		}
	}
	if isRed(o.left) && isRed(o.right) {
		flipColors(o)
	}
	o.N = size(o.left) + size(o.right) + 1
	return o
}

func (bst *BalSearchTree) Show() {
	bst.check()
	showRB(bst.root)
}

//Show all node in the binsearchtree
func showRB(node *NodeRB) {
	if node != nil {
		fmt.Println("        node:  color: ", node.color, "  key: ", node.key, " value: ", node.val, " size: ", node.N)
		if node.father != nil {
			fmt.Println(" father     :  color: ", node.father.color, "  key: ", node.father.key, " value: ", node.father.val, " size: ", node.father.N)
		} else {
			fmt.Println(" father     :  nil")
		}
		if node.left != nil {
			fmt.Println(" left  child:  color: ", node.left.color, "  key: ", node.left.key, " value: ", node.left.val, " size: ", node.left.N)
		} else {
			fmt.Println(" left  child:  nil")
		}
		if node.right != nil {
			fmt.Println(" right child:  color: ", node.right.color, "  key: ", node.right.key, " value: ", node.right.val, " size: ", node.right.N, "\n")
		} else {
			fmt.Println(" right child:  nil\n")
		}
		showRB(node.left)
		showRB(node.right)
	}
}

//Return the value of node with Key key
func (bst *BalSearchTree) Get(key Key) Val {
	bst.check()
	return *getRB(bst.root, key)
}

func getRB(node *NodeRB, key Key) *Val {
	if node == nil {
		return nil
	}
	var cmp int = key.Compare(node.key)
	if cmp < 0 {
		return getRB(node.left, key)
	} else if cmp > 0 {
		return getRB(node.right, key)
	} else {
		return getRB(node.right, key)
	}
}

//Return node with minimum Key
func (bst *BalSearchTree) Min() Key {
	bst.check()
	return minRB(bst.root).key
}

func minRB(node *NodeRB) *NodeRB {
	if node.left == nil {
		return node
	} else {
		return minRB(node.left)
	}
}

//Return node with maximum Key
func (bst *BalSearchTree) Max() Key {
	bst.check()
	return maxRB(bst.root).key
}

func maxRB(node *NodeRB) *NodeRB {
	if node.right == nil {
		return node
	} else {
		return maxRB(node.right)
	}
}

func (P *NodeRB) anotherChild(N *NodeRB) *NodeRB {
	if P.left == N {
		return P.right
	} else {
		return P.left
	}
}

func (P *NodeRB) isLeftSon() bool {
	if P.father.left == P {
		return true
	} else {
		return false
	}
}

var tmp *NodeRB

func (bst *BalSearchTree) DeleteMin() {
	bst.check()
	var minNode = minRB(bst.root)
	deleteMinRB(bst, minNode)
}

func deleteMinRB(bst *BalSearchTree, minNode *NodeRB) {
	if minNode.right != nil {
		if minNode.father != nil {
			minNode.right.father = minNode.father
			if minNode.isLeftSon() {
				minNode.father.left = minNode.right
			} else {
				minNode.father.right = minNode.right
			}
		} else {
			minNode.right.father = nil
			bst.root = minNode.right
		}
		balance(minNode.right, minNode.color, minNode.father, bst)
	} else {
		if minNode.father != nil {
			if minNode.isLeftSon() {
				minNode.father.left = nil
			} else {
				minNode.father.right = nil
			}
		} else {
			bst.root = nil
		}
		balance(nil, minNode.color, minNode.father, bst)
	}
	return
}

func balance(N *NodeRB, col Color, P *NodeRB, bst *BalSearchTree) {
	//0.0
	if P == nil {
		N.color = BLACK
		return
	}

	//0.1
	if col == RED {
		return
	}

	//0.2
	if col == BLACK && isRed(N) {
		N.color = BLACK
		return
	}

	var S = P.anotherChild(N)
	var SL = S.left
	var SR = S.right
	//1
	if !isRed(N) && !isRed(P) && !isRed(S) && !isRed(S.left) && !isRed(S.right) {
		fmt.Println("1 used")
		S.color = RED
		balance(P, BLACK, P.father, bst)
		return
	}

	//2
	if isRed(S) {
		fmt.Println("2 used")
		if P.father != nil {
			tmp = P.father
			if P.isLeftSon() {
				tmp.left = rotateLeftPure(P)
			} else {
				tmp.right = rotateLeftPure(P)
			}
		} else {
			bst.root = rotateLeftPure(P)
		}
		P.color, S.color = S.color, P.color
		S = P.anotherChild(N)
		SL = S.left
		SR = S.right
	}

	//2.1
	if !isRed(S) && !isRed(S.left) && !isRed(S.right) {
		fmt.Println("2.1 used")
		P.color, S.color = S.color, P.color
		return
	}

	//2.2.1
	if !isRed(S) && !isRed(S.right) && isRed(S.left) {
		fmt.Println("2.2.1 used")
		if S.father != nil {
			tmp = S.father
			if S.isLeftSon() {
				tmp.left = rotateRightPure(S)
			} else {
				tmp.right = rotateRightPure(S)
			}
		} else {
			bst.root = rotateRightPure(S)
		}
		SL.color, S.color = S.color, SL.color
		fmt.Println("p.right.key", P.right.key)
		fmt.Println("s.key", S.key)
		S = SL
		SL = S.left
		SR = S.right
	}

	//2.2.2
	if !isRed(S) && isRed(S.right) {
		fmt.Println("2.2.2 used")
		if P.father != nil {
			tmp = P.father
			if P.isLeftSon() {
				tmp.left = rotateLeftPure(P)
			} else {
				tmp.right = rotateLeftPure(P)
			}
		} else {
			bst.root = rotateLeftPure(P)
		}
		P.color, S.color = S.color, P.color
		SR.color = BLACK
		return
	}
}

func (bst *BalSearchTree) GetNodeRB(key Key) *NodeRB {
	bst.check()
	return getNodeRB(bst.root, key)
}

func getNodeRB(node *NodeRB, key Key) *NodeRB {
	if node == nil {
		return nil
	}
	var cmp int = key.Compare(node.key)
	if cmp < 0 {
		return getNodeRB(node.left, key)
	} else if cmp > 0 {
		return getNodeRB(node.right, key)
	} else {
		return node
	}
}

/*
func (bst *BalSearchTree) Delete(key Key){
	bst.check()
	bst.root = deleteNodeRB(bst.root, key)
}

func deleteNodeRB(node *NodeRB, key Key) *Node{
	if node == nil{
		return nil
	}
	var targetNode = getNode
	var cmp = key.Compare(node.key)
	if cmp < 0 {
		node.left = deleteNodeRB(node.left, key)
	}else if cmp > 0 {
		node.right = deleteNodeRB(node.right, key)
	}else{
		if node.right == nil{
			return node.left
		} node.left == nil{
			return node.right
		}
		var t = node
		var minNode = minRB(t.right)
		node.key = minNode.key
		node.val = minNode.val
		node.right = deleteMinRB(t.right)
		node.left = t.left
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}
*/
