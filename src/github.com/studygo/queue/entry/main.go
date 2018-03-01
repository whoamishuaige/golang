package main

import (
	"github.com/studygo/queue"
	"fmt"
)

func main()  {
		q :=queue.Queue{1}
		q.Push(2)
		q.Push(2)
		fmt.Println(q.Pop())
		fmt.Println(q.Pop())
		fmt.Println(q.IsEmpty());
		fmt.Println(q.Pop())
		fmt.Println(q.IsEmpty());

}
