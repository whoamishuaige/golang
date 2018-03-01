package main

import "fmt"
func eval(a,b int,op string)(result int,err error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_ :=div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsuported operation:%s",op)
	}

}
func sum(numbers ... int) int{
		s :=0
		for i:=range  numbers{
			s += numbers[i]
		}
		return s
}
func div(a,b int ) (q,r int) {
	return a/b,a%b
}
func swap2 (a,b *int ){
	*a,*b= *b,*a
}
func swap (a,b int)(int ,int){
	return b,a
}
func main()  {
	q, r := div(13, 3)
	fmt.Println(q,r)
	if result, err := eval(3, 4, "/");err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
	c := 1
	d := 4
	a,b :=swap(1,2)
	swap2(&c,&d)
	fmt.Println(a,b,
	sum(1,2,3,4,5),c,d)
}

