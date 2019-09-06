package main

import (
	"fmt"
)

// SingleObject 节点数据
type SingleObject interface{}

// SingleNode 单链表节点
type SingleNode struct {
	Data SingleObject
	Next *SingleNode
}

// SingleList 单链表
type SingleList struct {
	Head *SingleNode
	Tail *SingleNode
	Size uint
}

// Init 初始化
func (list *SingleList) Init() {
	list.Size = 0
	list.Head = nil
	list.Tail = nil
}

// Append 添加节点到链表尾部
func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1
		return true
	}

	tail := list.Tail
	tail.Next = node
	list.Tail = node
	list.Size += 1
	return true
}

// Insert 插入节点到指定位置
func (list *SingleList) Insert(index uint, node *SingleNode) bool {
	if node == nil {
		return false
	}

	if index > list.Size {
		return false
	}

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Size += 1
		return true
	}

	var i uint
	nowNode := list.Head
	for i = 0; i < index; i++ {
		nowNode = nowNode.Next
	}
	// 顺序不要弄错
	next := nowNode.Next
	nowNode.Next = node
	node.Next = next
	list.Size += 1
	return true
}

// Delete 删除指定位置的节点
func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return false
	}

	if index == 0 {
		head := list.Head.Next
		list.Head = head
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size -= 1
		return true
	}

	var i uint
	nowNode := list.Head
	for i = 1; i < index; i++ {
		nowNode = nowNode.Next
	}
	next := nowNode.Next
	nowNode.Next = next.Next
	if index == list.Size-1 {
		list.Tail = nowNode
	}
	list.Size -= 1
	return true
}

// Get 获取指定位置的节点
func (list *SingleList) Get(index uint) *SingleNode {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return nil
	}

	if index == 0 {
		return list.Head
	}

	// 如果没有var i uint，则需要for i := 0; (uint)i < index; i++
	var i uint
	nowNode := list.Head
	for i = 0; i < index; i++ {
		nowNode = nowNode.Next
	}
	return nowNode
}

// Display 输出链表
func (list *SingleList) Display() {
	if list == nil {
		fmt.Println("this single list is nil")
	}

	fmt.Printf("this single list size is %d \n", list.Size)

	var i uint
	nowNode := list.Head
	for i = 0; i < list.Size; i++ {
		fmt.Printf("No %3d data is %v\n", i+1, nowNode.Data)
		nowNode = nowNode.Next
	}
}

func main() {
	singleList := new(SingleList)

	// 初始化
	singleList.Init()

	// 添加数据元素
	// error: singleList.Append(&{1, nil}) syntax error: unexpected {, expecting expression
	singleList.Append(&SingleNode{1, nil})
	singleList.Append(&SingleNode{2, nil})
	singleList.Append(&SingleNode{3, nil})
	singleList.Append(&SingleNode{5, nil})
	singleList.Append(&SingleNode{6, nil})

	// 打印链表
	singleList.Display()

	// 插入数据元素
	singleList.Insert(4, &SingleNode{4, nil})

	// 打印链表
	singleList.Display()

	// 获取数据元素
	singleList.Get(0)
	singleList.Get(5)

	// 删除数据元素
	singleList.Delete(2)

	// 打印链表
	singleList.Display()
}
