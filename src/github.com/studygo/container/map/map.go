package main

import "fmt"

func main(){
	m := map[string] string{
		"name":"zhangsan",
		"age":"18",
		"sex":"male",
	}
	m2 :=make(map[string]int)
	var m3 = make(map[string]int)
	fmt.Println(m,m2,m3)

	for k,v :=range m{
		fmt.Println(k,v)
	}
	name := m["names"]
	fmt.Printf("%q\n",name)

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}
