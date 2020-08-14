package LinkListDemo

import "fmt"

type LinkListObject interface {

}

type Node struct {
	Data LinkListObject //链表数据
	Next *Node
}

type LinkList struct {
	headNode *Node
}


func (link *LinkList) IsEmpty() bool {
	if link.headNode == nil{
		return true
	}else {
		return false
	}

}

func (link *LinkList) length() int  {
	cur := link.headNode
	count := 0
	for cur != nil{
		count ++
		cur = cur.Next
	}
	return count
}


func (link *LinkList) Add(data LinkListObject)  *Node{
	node := &Node{Data: data}
	node.Next = link.headNode
	link.headNode = node
	return node
}

func (link *LinkList)ShowList()  {
	if !link.IsEmpty(){
		cur := link.headNode
		for {
			fmt.Printf("%v aaa\n",cur.Data)
			if cur.Next != nil{
				cur = cur.Next
			}else {
				break
			}

		}
	}
}




