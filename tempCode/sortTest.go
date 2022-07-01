package main

import (
	"fmt"
	"sort"
)

type MyStructType struct {
	Id   uint
	Name string
	Date string
}

type SortBy []MyStructType

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	// 基于ID顺序进行排序
	return a[i].Id < a[j].Id
}

func main() {
	// 如下：为自定义结构体自定义排序规则的例子
	// 创建一个结构体数组，并添加数据
	var myStructs []MyStructType
	myStructs = append(myStructs, MyStructType{Id: 1, Name: "a", Date: "2022-03-17"})
	myStructs = append(myStructs, MyStructType{Id: 3, Name: "b", Date: "2022-03-18"})
	myStructs = append(myStructs, MyStructType{Id: 2, Name: "c", Date: "2019-03-17"})
	myStructs = append(myStructs, MyStructType{Id: 5, Name: "d", Date: "1992-03-17"})
	myStructs = append(myStructs, MyStructType{Id: 4, Name: "e", Date: "2022-07-01"})

	// 输出结构体
	for _, v := range myStructs {
		fmt.Println(v)
	}

	// 对结构体数组进行排序
	sort.Sort(SortBy(myStructs))

	// 再次输出结构体
	for _, v := range myStructs {
		fmt.Println(v)
	}
}
