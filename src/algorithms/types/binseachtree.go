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

//Return size of the binsearchtree
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

//Return the value of node with Key key
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

//Put node with Key key and Val val in the binsearchtree
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

//Show all node in the binsearchtree
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

//Return node with minimum Key
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

//Return node with maximum Key
func (bst *BinSearchTree) Max() Key{
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		return max(bst.root).key
	}
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

//Return *Node containing key of rank k
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

//Return rank of Key key
func (bst *BinSearchTree) Rank(key Key) int{
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		return rank(bst.root, key)
	}
}

//Return number of nodes with key less than key in the subtree rooted at node.
func rank(node *Node, key Key) int{
	var cmp = key.Compare(node.key)
	if cmp < 0 {
		if node.left == nil{
			return 0
		}else{
			return rank(node.left, key)
		}
	}else if cmp > 0{
		if node.left == nil{
			return 1 + rank(node.right, key)
		}else if node.right == nil{
			return 1 + size(node.left)
		}else{
			return 1 + size(node.left) + rank(node.right, key)
		}
	}else if node.left == nil{
		return 0
	}else{
		return size(node.left)
	}
}

//Delete node with minimum Key
func (bst *BinSearchTree) DeleteMin() {
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		if bst.root.left == nil && bst.root.right == nil{
			bst.root = nil
		}else{
			bst.root = deleteMin(bst.root)
		}
	}
}

func deleteMin(node *Node) *Node{
	if node.left == nil{
		return node.right
	}
	if node.left.left == nil && node.left.right == nil{
		node.left = nil
		node.sizeSum()
	}else{
		node.left = deleteMin(node.left)
		node.sizeSum()
	}
	return node
}

//Calculate the size of the subtree rooted at node
func (node *Node) sizeSum(){
	if node.left == nil && node.right == nil{
		node.N = 1
	}else if node.left == nil{
		node.N = size(node.right) + 1
	}else{
		node.N = size(node.left) + size(node.right) + 1
	}
}

//Delete node with Key key
func (bst *BinSearchTree) Delete(key Key) {
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}else{
		if bst.root.left == nil && bst.root.right == nil{
			if bst.root.key == key{
				bst.root = nil
			}else{
				var message = fmt.Sprintf("Node containing key %v is not exist.", key)
				panic(message)
			}
		}else{
			bst.root = deleteNode(bst.root, key)
		}
	}
}

func deleteNode(node *Node, key Key) *Node{
	var cmp = key.Compare(node.key)
	if cmp < 0{
		if node.left == nil{
			var message = fmt.Sprintf("Node containing key %v is not exist.", key)
			panic(message)
		}else{
			if node.left.left == nil && node.left.right == nil{
				if node.left.key == key{
					node.left = nil
					node.sizeSum()
					return node
				}else{
					var message = fmt.Sprintf("Node containing key %v is not exist.", key)
					panic(message)
				}
			}else{
				node.left = deleteNode(node.left, key)
			}
		}
	}else if cmp > 0 {
		if node.right == nil{
			var message = fmt.Sprintf("Node containing key %v is not exist.", key)
			panic(message)
		}else{
			if node.right.left == nil && node.right.right == nil{
				if node.right.key == key{
					node.right = nil
					node.sizeSum()
					return node
				}else{
					var message = fmt.Sprintf("Node containing key %v is not exist.", key)
					panic(message)
				}
			}else{
				node.right = deleteNode(node.right, key)
			}
		}
	}else{
		var nodeTmp *Node = node
		var nodeMinTmp = min(nodeTmp.right)
		node.key = nodeMinTmp.key
		node.val = nodeMinTmp.val
		//*node = *nodeMinTmp
		node.left = nodeTmp.left
		if nodeTmp.left == nil && nodeTmp.right == nil{
			node.right = nil
		}else{
			node.right = deleteMin(nodeTmp.right)
		}
	}
	node.sizeSum()
	return node
}

//Golang has no simple override method
//func (bst *BinSearchTree) Keys() Queue{
//	return keys(bst.min(), bst.max())
//}

//Return Queue of node with Key between Key lo and Key hi
func (bst *BinSearchTree) Keys(lo Key, hi Key) *Queue{
	if bst.root == nil{
		var message = fmt.Sprintf("There isn't any node in this binsearchtree")
		panic(message)
	}
	var q = new(Queue)
	keys(bst.root, q, lo, hi)
	return q
}

func keys(node *Node, q *Queue, lo Key, hi Key){
	var cmplo = lo.Compare(node.key)
	var cmphi = hi.Compare(node.key)
	if cmplo < 0 && node.left != nil{
		keys(node.left, q, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0{
		q.Push(node.key)
	}
	if cmphi > 0 && node.right != nil{
		keys(node.right, q, lo, hi)
	}
}


