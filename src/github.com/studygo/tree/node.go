package tree
//使用组合
import (
	"fmt"
	//"text/template/parse"
)

type Node struct {
	Value int
	Left,Right *Node
}
func (node *Node) SetValue(value int ){
	node.Value = value
}
func (node Node)Print(){
	fmt.Print(node.Value," ");
}
func CreateNode(value int) *Node {
	return &Node{Value:value}
}
func (node *Node) Trav() {
	if node== nil{
		return
	}
	node.Left.Trav()
	node.Print()
	node.Right.Trav()
}

