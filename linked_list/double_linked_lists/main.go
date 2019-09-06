package main

import (
	"errors"
	"fmt"
)

// DoubleObject 存储数据元素
type DoubleObject interface{}

// DoubleNode 双链表节点
type DoubleNode struct {
	Data DoubleObject
	Prev *DoubleNode
	Next *DoubleNode
}

// DoubleList 双链表
type DoubleList struct {
	Head   *DoubleNode
	Tail   *DoubleNode
	Length int
}

// DoubleListInit 双链表初始化
func (l *DoubleList) DoubleListInit() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

// DoubleListInsert 向链表任意位置插入节点
func (l *DoubleList) DoubleListInsert(location int, node *DoubleNode) (err error) {
	if location < 0 || location > l.Length {
		err = errors.New("location error")
		return err
	}

	if node == nil {
		err = errors.New("Node error")
		return err
	}

	if location == 0 {
		head := l.Head
		l.Head = node
		head.Prev = node
		node.Prev = nil
		node.Next = head
	} else {
		nowNode, _ := l.DoubleListGet(location)
		node.Prev = nowNode
		node.Next = nowNode.Next
		nowNode.Next = node
		nowNode.Next.Prev = node
	}

	l.Length++

	return nil
}

// DoubleListLastInsert 向链表最后插入节点
func (l *DoubleList) DoubleListLastInsert(node *DoubleNode) (err error) {
	if node == nil {
		err = errors.New("Node error")
		return err
	}

	if l.Length == 0 {
		l.Head = node
		l.Tail = node
		node.Prev = nil
		node.Next = nil
	} else {
		nowNode := l.Tail
		nowNode.Next = node
		node.Prev = nowNode
		node.Next = nil
		l.Tail = node
	}

	l.Length++

	return nil
}

// DoubleListDelete 删除节点
func (l *DoubleList) DoubleListDelete(location int) (err error) {
	if location < 0 || location > l.Length {
		err = errors.New("Location error")
		return err
	}

	nowNode := l.Head
	for i := 0; i <= location; i++ {
		nowNode = nowNode.Next
	}
	prev := nowNode.Prev
	next := nowNode.Next
	nowNode.Prev = nil
	nowNode.Next = nil
	prev.Next = next
	next.Prev = prev

	l.Length--

	return nil
}

// DoubleListGet 查询节点
func (l *DoubleList) DoubleListGet(location int) (node *DoubleNode, err error) {
	if location < 0 || location > l.Length {
		err = errors.New("Location error")
		return nil, err
	}

	nowNode := l.Head
	for i := 1; i < location; i++ {
		nowNode = nowNode.Next
	}

	node = nowNode

	return node, nil
}

// DoubleListDisplay 打印链表
func (l *DoubleList) DoubleListDisplay() {
	nowNode := l.Head
	for i := 0; i < l.Length; i++ {
		fmt.Printf("NO %d data is %d \n", i+1, nowNode.Data)
		nowNode = nowNode.Next
	}
}

func main() {
	doubleList := new(DoubleList)

	// 初始化链表
	doubleList.DoubleListInit()

	// 向链表末尾插入节点
	doubleList.DoubleListLastInsert(&DoubleNode{0, nil, nil})
	doubleList.DoubleListLastInsert(&DoubleNode{1, nil, nil})
	doubleList.DoubleListLastInsert(&DoubleNode{2, nil, nil})
	doubleList.DoubleListLastInsert(&DoubleNode{3, nil, nil})
	doubleList.DoubleListLastInsert(&DoubleNode{4, nil, nil})
	doubleList.DoubleListLastInsert(&DoubleNode{6, nil, nil})

	// 打印链表
	doubleList.DoubleListDisplay()

	// 向链表中插入节点
	doubleList.DoubleListInsert(5, &DoubleNode{5, nil, nil})

	// 打印链表
	doubleList.DoubleListDisplay()

	// 获取节点
	doubleList.DoubleListGet(2)

	// 删除节点
	doubleList.DoubleListDelete(3)

	// 打印链表
	doubleList.DoubleListDisplay()
}
