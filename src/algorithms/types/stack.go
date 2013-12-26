package types

import ("fmt")

type BinTree struct{
	tree *Sortable
	judge func(Compare, Compare)bool
}
//FIXME: 没有处理越界
func NewBinTree(max int, j func(Compare, Compare)bool) BinTree{
	var tree = make(Sortable, max+1)
	tree[0]=Integer(0)
	return BinTree{&tree, j}
}

func (b BinTree)Size()int{
	var tree = *b.tree
	var ret, ok = tree[0].(Integer)
	if ok {
		return int(ret)
	} else {
		var message = fmt.Sprintf("It is't a size in Buffer[0]. It is %v", 
			tree[0])
		panic(message)
	}
}

func (b BinTree)IsEmpty() bool{
	return b.Size()==0
}

func (b BinTree)Insert(item Compare) int{
	b.incSize()
	var tree = *b.tree
	tree[b.Size()] = item
	return b.Swim(b.Size())
}

func (b BinTree)DelMax() Compare{
	if b.IsEmpty() {
		return nil
	}
	var tree = *b.tree
	var max = tree[1]
	b.tree.Swap(1, b.Size())
	b.recSize()
	tree[b.Size()+1] = nil
	b.Sink(1)
	return max
}

func (b BinTree)Len() int {
	return b.tree.Len()-1
}

func (b BinTree)Swim(k int) int{
	var tree = *b.tree
	for k>1 && b.judge(tree[k/2], tree[k]) {
		b.tree.Swap(k/2, k)
		k=k/2
	}
	return k
}

func (b BinTree)Sink(k int)int{
	var tree = *b.tree
	for 2*k <= b.Size(){
		var j = 2*k
		if j<b.Size()&&b.judge(tree[j], tree[j+1]){
			j++
		}
		if (!b.judge(tree[k], tree[j])) {
			break
		}
		b.tree.Swap(k, j)
		k=j
	}
	return k
}

func (b BinTree)incSize(){
	var tree = *b.tree
	var data, ok = tree[0].(Integer)
	if ok {
		data++
		tree[0] = data
	}else{
		var message = fmt.Sprintf("It is't a size in Buffer[0]. It is %v", 
			tree[0])
		panic(message)
	}
}

func (b BinTree)recSize(){
	var tree = *b.tree
	var data, ok = tree[0].(Integer)
	if ok {
		data--
		tree[0] = data
	}else{
		var message = fmt.Sprintf("It is't a size in Buffer[0]. It is %v", 
			tree[0])
		panic(message)
	}
}

type TopPQ struct {
	*BinTree
}

func NewMaxPQ(size int) TopPQ {
	var b = NewBinTree(size, func(x, y Compare)bool{return x.Less(y)})
	return TopPQ{&b}
}

func NewMinPQ(size int) TopPQ {
	var b = NewBinTree(size, func(x, y Compare)bool{return y.Less(x)})
	return TopPQ{&b}
}

func (m TopPQ) IsEmpty() bool{
	return m.BinTree.IsEmpty()
}

func (m TopPQ) Size() int {
	return m.BinTree.Size()
}

func (m *TopPQ) Insert(item Compare) {
	if m.BinTree.Size()==m.BinTree.Len() {
		m.BinTree.DelMax()
	}
	m.BinTree.Insert(item)
}

func (m *TopPQ) DelMax() Compare {
	return m.BinTree.DelMax()
}
