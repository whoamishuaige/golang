package basic

//数据类型的使用包括定义转换枚举常量等
import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数外的定义方式，无法使用：=的方式，定义的变量也可以不使用
//go语言函数外的定义和函数内部使用有所不同
var (
	a   = 1
	b   = 2
	c   = 3
	d   = "def"
	dad = "1ada"
)

//欧拉公式验证
func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

//类型转换，不自动，手动
func triangle() {
	var a, b int = 300000000, 400000000
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//const的使用
func consts() {
	const filename, a, b = "abc.txt", 1, 2
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}

//枚举类型
func enums() {
	const (
		cpp = iota
		_
		php
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, php, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

//不定义内容自动分配内容
//犯错点1：printf才会打印格式，println不能，话说为啥没有万能打印方法
func varZeroValue() {
	var a int
	var b string
	fmt.Printf("%d,%q\n", a, b)
}

//测试一下这个插件的效果，看看是什么情况，感觉好奇妙好叼啊
//定义变量类型和初值
func varInitValue() {
	var a int = 1
	var b string = "def"
	fmt.Printf("%d,%q\n", a, b)
}

//自动判断数据类型
//var a = 1等价于 a:=1

func varAutoType() {
	a := 1
	b := "string"
	fmt.Println(a, b)
	var s, d, e, f = 1, 2, true, "dfd"
	fmt.Println(s, d, e, f)
	j, q, u, y := 3, 6, 7, 1
	fmt.Println(j, q, u, y)
}
func main() {
	fmt.Println("HELLO,WORLD")
	varZeroValue()
	varInitValue()
	varAutoType()
	euler()
	triangle()
	consts()
	enums()
}
