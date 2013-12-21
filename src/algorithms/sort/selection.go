package sort

import "sort"

func Selection(c sort.Interface) {
	var l = c.Len()
	for i:=0; i< l; i++{
		var min = i;
		for j:=i+1; j < l; j++ {
			if c.Less(j, min) {
				min = j
			}
		}
		c.Swap(i, min)
	}
}

