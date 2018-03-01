package main

import "fmt"


//1:数组传递以传值调用
//2:修改可以通过指针调用


func printArray(arr *[5] int){
		arr[0] = 100;
		for i,v := range arr{
			fmt.Println(i,v)
		}
}
func main(){
	var arr1 [5] int;
	var arr2 = [5] int {1,2,3,4,5}
	arr3 := [3]int{4,5,6}
	arr4 :=[...]int{10,23,12,34,12}
	var arr5 [5][4] int
	fmt.Println(arr1);
	fmt.Println(arr2);
	fmt.Println(arr3);
	fmt.Println(arr4)
	fmt.Println(arr5);
	printArray(&arr1);
	fmt.Println(arr1)
}