package sort

import "sort"

func Insertion(c sort.Interface){
	var l = c.Len()
	for i:=1; i<l; i++ {
		for j:=i; j>0&&c.Less(j, j-1); j-- {
			c.Swap(j, j-1)
		}
	}
}
