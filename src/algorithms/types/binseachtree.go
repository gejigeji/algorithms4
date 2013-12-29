package types

import "fmt"

type BinSearchTree struct{
	root *Node
}

type Key Compare

type Val int

type Node struct{
	key Key
	val Val
	left, right *Node
	N int
}

func (bst *BinSearchTree) Size() int{
	if bst.root != nil{
		return size(bst.root)
	}else{
		return 0
	}
}

func size(node *Node) int{
	if node == nil {
		return 0
	}
	return node.N
}

func (bst *BinSearchTree) Get(key Key) Val{
	if bst.root != nil{
		return get(bst.root, key)
	}else{
		panic("There is no node in this binsearchtree")
	}
}

func get(node *Node, key Key) Val{
	var cmp int = key.Compare(node.key)
	if cmp < 0 {
		if node.left != nil {
			return get(node.left, key)
		}else{
			var message = fmt.Sprintf("Key %v is not in the binsearchtree", key)
			panic(message)
		}
	}else if cmp > 0 {
		if node.right != nil {
			return get(node.right, key)
		}else{
			var message = fmt.Sprintf("Key %v is not in the binsearchtree", key)
			panic(message)
		}
	}else {
		return node.val
	}
}

func (bst *BinSearchTree) Put(key Key, val Val) {
	var node = bst.root
	if node == nil{
		bst.root = &Node{key, val, nil, nil, 1}
	}else{
		put(node, key, val)
	}
}

func put(node *Node, key Key, val Val) {
	var cmp = key.Compare(node.key)
	if cmp < 0 {
		var son = node.left
		if son == nil{
			node.left = &Node{key, val, nil, nil, 1}
		}else{
			put(son, key, val)
		}
	}else if cmp > 0 {
		var son = node.right
		if son == nil{
			node.right = &Node{key, val, nil, nil, 1}
		}else{
			put(son, key, val)
		}
	}else{
		node.val = val
	}
	node.N = size(node.left) + size(node.right) + 1
}

func (bst *BinSearchTree) Show(){
	var node = bst.root
	if node == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		show(node)
	}
}

func show(node *Node){
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
	if node.left == nil{
	}else{
		show(node.left)
	}
	if node.right == nil{
	}else{
		show(node.right)
	}
}

func (bst *BinSearchTree) Min() Key{
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		return min(bst.root).key
	}
}

func min(node *Node) *Node{
	if node.left == nil{
		return node
	}else{
		return min(node.left)
	}
}

func (bst *BinSearchTree) Floor(key Key) Key{
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		return floor(bst.root, key).key
	}
}

func floor(node *Node, key Key) *Node{
	var cmp = key.Compare(node.key)
	if cmp == 0 {
		return node
	}else if cmp < 0 {
		if node.left == nil{
			var message = fmt.Sprintf("There isn't any key not bigger than %v", key)
			panic(message)
		}else{
			return floor(node.left, key)
		}
	}else{
		if node.right == nil{
			return node
		}else if key.Compare(node.right.key) < 0{
			return node
		}else{
			return floor(node.right, key)
		}
	}
}

func (bst *BinSearchTree) Select(k int) Key{
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		return selectNode(bst.root, k).key
	}
}

func selectNode(node *Node, k int) *Node{
	if node.left == nil{
		if 0 > k {
			panic("Out of range")
		}else if 0 < k{
			if node.right == nil{
				panic("Out of range")
			}else{
				return selectNode(node.right, k-1)
			}
		}else{
			return node
		}
	}
	var t = size(node.left)
	if t > k{
		return selectNode(node.left, k)
	}else if t < k{
		if node.right == nil{
			panic("Out of range")
		}
		return selectNode(node.right, k-t-1)
	}else{
		return node
	}
}
