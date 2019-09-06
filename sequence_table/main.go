package main

import (
	"errors"
	"fmt"
)

// ElemType 元素类型
type ElemType int

// MAXSIZE 申请内存大小
const MAXSIZE = 20

// SeqList 定义顺序表
type SeqList struct {
	data   [MAXSIZE]ElemType
	length int
}

// SeqListInit 初始化顺序表
func (s *SeqList) SeqListInit() {
	for i := 0; i < MAXSIZE; i++ {
		s.data[i] = 0
	}
	s.length = 0
}

// SeqListInsert 插入数据
func (s *SeqList) SeqListInsert(i int, e ElemType) (err error) {
	if s.length == MAXSIZE {
		err = errors.New("The sequence table is full and cannot insert data")
		return err
	}

	if i < 0 || i > s.length {
		err = errors.New("Insert location is error")
		return err
	}

	for j := s.length; j >= i; j-- {
		s.data[j+1] = s.data[j]
	}
	s.data[i] = e
	s.length++
	return nil
}

// SeqListDelete 删除数据元素
func (s *SeqList) SeqListDelete(i int) (err error) {
	if s.length == 0 {
		err = errors.New("The sequence table is empty")
		return err
	}

	if i < 0 || i > s.length-1 {
		err = errors.New("Delete location is error")
		return err
	}

	for j := i; j < s.length-1; j++ {
		s.data[j] = s.data[j+1]
	}

	s.data[s.length-1] = 0
	s.length--
	return nil
}

// SeqListGetElem 获取数据元素
func (s *SeqList) SeqListGetElem(i int) (err error, res ElemType) {
	if s.length == 0 || i < 0 || i > s.length {
		err = errors.New("Find error")
	}

	res = s.data[i]
	return nil, res
}

// SeqListDisplay 打印顺序表
func (s *SeqList) SeqListDisplay() {
	for i := 0; i < s.length; i++ {
		fmt.Printf("NO %d data is: %d \n", i+1, s.data[i])
	}
}

func main() {
	seqList := new(SeqList)

	// 初始化
	seqList.SeqListInit()

	// 添加数据元素
	seqList.SeqListInsert(0, 0)
	seqList.SeqListInsert(1, 1)
	seqList.SeqListInsert(2, 2)
	seqList.SeqListInsert(3, 3)
	seqList.SeqListInsert(4, 4)
	seqList.SeqListInsert(5, 5)

	// 打印顺序表
	seqList.SeqListDisplay()

	// 获取数据元素
	seqList.SeqListGetElem(2)

	// 删除数据元素
	seqList.SeqListDelete(4)

	// 打印顺序表
	seqList.SeqListDisplay()
}
