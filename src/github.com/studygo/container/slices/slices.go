package main

import "fmt"

func main(){
	arr :=[...]int {0,1,2,3,4,5,6}
	//slice是一个数据结构，是对array的view
	fmt.Println("arr[2:6]",arr[2:6])
	fmt.Println("arr[2:]",arr[2:])
	fmt.Println("arr[:6]",arr[:6])
	fmt.Println("arr[:]",arr[:])
	fmt.Println("arr[2:6]",arr[2:6])
//slice本身是没有值的，它对应就是与原先的内容，也就是指针



arr1 :=[...]int {1,2,3,4,5,6,7,8,9}
s1 :=arr1[2:6]
s2 :=s1[3:7]
s2 =append(s2,10)
s2 = append(s2,30)
s2 = append(s2,30)
s2 = append(s2,30)
	s2 = append(s2,30)
	s2 = append(s2,30)

fmt.Println(s1,s2)
}
