package main

import (
	"fmt"
	//"io/ioutil"
)
//赋值必须使用，赋值不算使用
func grade(score int)string{
	g :=""
	switch  {
	case score>100 ||score<0:
		panic(fmt.Sprintf("Wrong score:%d",score))
	case score<60:
		g  = "F"
	case score<70:
		g = "D"
	case score<80:
		g = "C"
	case score<90:
		g ="b"
	case score<=100:
		g ="a"
	}
	return g;
}
func main() {
	const filename = "aa.txt"
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Printf("%s\n", err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	fmt.Println(
	grade(60),
	grade(70),
	grade(100),
	grade(40),
	//grade(-1),
	)
}
