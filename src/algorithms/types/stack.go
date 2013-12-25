package types

import ("fmt")

type BinTree []Compare
//FIXME: 没有处理越界
func NewBinTree(max int) BinTree{
	var b = make(BinTree, max+1)
	b[0]=Integer(0)
	return b
}

func (b BinTree)Size()int{
	var ret, ok = b[0].(Integer)
	if ok {
		return int(ret)
	} else {
		var message = fmt.Sprintf("It is't a size in Buffer[0]. It is %v", 
			b[0])
		panic(message)
	}
}

func (b BinTree)IsEmpty() bool{
	return b.Size()==0
}

func (b BinTree)Insert(item Compare) int{
	b.incSize()
	b[b.Size()] = item
	return b.Swim(b.Size())
}

func (b BinTree)DelMax() Compare{
	if b.IsEmpty() {
		return nil
	}
	var max = b[1]
	b.Swap(1, b.Size())
	b.recSize()
	b[b.Size()+1] = nil
	b.Sink(1)
	return max
}

func (b BinTree)Len() int {
	return len(b)-1
}

func (b BinTree)Less(i, j int) bool {
	return b[i].Less(b[j])
}

func (b BinTree)Swap(i, j int){
	b[i], b[j] = b[j], b[i]
}

func (b BinTree)Swim(k int) int{
	for k>1 && b.Less(k/2, k) {
		b.Swap(k/2, k)
		k=k/2
	}
	return k
}

func (b BinTree)Sink(k int)int{
	for 2*k <= b.Size(){
		var j = 2*k
		if j<b.Size()&&b.Less(j, j+1){
			j++
		}
		if (!b.Less(k, j)) {
			break
		}
		b.Swap(k, j)
		k=j
	}
	return k
}

func (b BinTree)incSize(){
	var data, ok = b[0].(Integer)
	if ok {
		data++
		b[0] = data
	}else{
		var message = fmt.Sprintf("It is't a size in Buffer[0]. It is %v", 
			b[0])
		panic(message)
	}
}

func (b BinTree)recSize(){
	var data, ok = b[0].(Integer)
	if ok {
		data--
		b[0] = data
	}else{
		var message = fmt.Sprintf("It is't a size in Buffer[0]. It is %v", 
			b[0])
		panic(message)
	}
}

type MaxPQ struct {
	*BinTree
}

func NewMaxPQ(size int) MaxPQ {
	var b = NewBinTree(size)
	return MaxPQ{&b}
}

func (m MaxPQ) IsEmpty() bool{
	return m.BinTree.IsEmpty()
}

func (m MaxPQ) Size() int {
	return m.BinTree.Size()
}

func (m *MaxPQ) Insert(item Compare) {
	if m.BinTree.Size()==m.BinTree.Len() {
		m.BinTree.DelMax()
	}
	m.BinTree.Insert(item)
}

func (m *MaxPQ) DelMax() Compare {
	return m.BinTree.DelMax()
}
