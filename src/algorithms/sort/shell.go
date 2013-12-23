package sort

import "sort"

func Shell(c sort.Interface) {
	var l = c.Len()
	var h = 1
	for h < l/3 {
		h=3*h+1
	}
	for h >= 1 {
		for i:=h; i<l; i++ {
			for j:=i; j>=h&&c.Less(j, j-h); j-=h {
				c.Swap(j, j-h)
			}
		}
		h = h/3
	}
}
