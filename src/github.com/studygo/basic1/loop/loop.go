package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"

)

func convertToBin(n int) string {
		result := ""
	for ;n>0 ;n/=2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
		return result;
}
func printFileContent(filename string)  {
	file,err := os.Open(filename);
	if err!=nil{
		panic(err)
	}
	scanner :=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func forever(){
	for   {
		fmt.Println("dadada")
	}
}
func main()  {
	fmt.Printf("%q%q\n",
		convertToBin(0), convertToBin(3),
	)
	printFileContent("aa.txt")
	//forever()
}
