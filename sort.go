// Sort

/*
	기본형 정렬
	- sort.Ints(x []int)
	- sort.Strings(x []string)

	커스텀 정렬
	1) Using sort.Sort(data sort.Interface)
	- type alias
	- sort.Interface : Len , Less , Swap Implements

	2) Using sort.Slice(x interface{}, less func(i, j int) bool)
	 - sort.Slice(x []T, func(i, j int) bool{
		return x[i].data < x[j].data
	})
*/

package main

import "sort"

type CustomInterface []Custom
type Custom struct {
	Id    int
	Name  string
	Email string
}

func (c CustomInterface) Len() int {
	return len(c)
}

func (c CustomInterface) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CustomInterface) Less(i, j int) bool {
	return c[i].Id < c[j].Id
}

func SortExample() {
	strs := []string{"c", "d", "z", "a", "y"}
	sort.Strings(strs)

	ints := []int{1, 10, 5, 3, 7, 2}
	sort.Ints(ints)

	custom := []Custom{{3, "Lee", "iddbxk@naver.com"},
		{1, "Kim", "iddbxk@daum.net"},
		{2, "Park", "iddbxk@gmail.com"},
	}

	sort.Sort(CustomInterface(custom))
	sort.Slice(custom, func(i, j int) bool {
		return custom[i].Id < custom[j].Id
	})
}
