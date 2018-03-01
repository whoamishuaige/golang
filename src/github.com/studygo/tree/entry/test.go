package main

import (
	"fmt"

	"github.com/studygo/tree"
	//"github.com/studygo/basic1/basic"
)
//import ""
//使用组合
func main()  {
var root tree.Node
root =tree.Node{Value:3}
root.Left = &tree.Node{}
root.Right = &tree.Node{5,nil,nil}
root.Right.Left = &tree.Node{}
root.Left.Right=tree.CreateNode(3)
fmt.Println(root)
fmt.Println(root.Right)
fmt.Println(root.Left)
//root.Trav()
root.SetValue(200);
root.Trav()
//basic.Test();
fmt.Println()
//myTreeNode{&root}.PostOrder();
myRoot := myTreeNode{&root}
myRoot.PostOrder()
//myTreeNode{&root}.PostOrder()
fmt.Println();
}
type  myTreeNode struct {
	node * tree.Node
}
func (myNode *myTreeNode)PostOrder(){
	if myNode==nil||myNode.node==nil {
		//fmt.Println("mynode is nil")
		return
	}
	left :=myTreeNode{myNode.node.Left}
	left.PostOrder()
	myNode.node.Print()
	//myTreeNode{myNode.node.Left}.PostOrder()
	//myTreeNode{myNode.node.Right}.PostOrder()
	right :=myTreeNode{myNode.node.Right}
	right.PostOrder()

}