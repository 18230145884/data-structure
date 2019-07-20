package main

import (
	"errors"
	"fmt"
)

type ElemType int  //元素类型
const MAXSIZE = 20 //申请内存大小
//定义数据元素：结构体作为数据元素，内容包括存放数据的数组、数据的长度
type SqList struct {
	data   [MAXSIZE]ElemType
	length int
}

//获取数据元素
func GetElem(s SqList, i int) (err error, res ElemType) {
	//查找的线性表为空   下标越界（小于0或者大于长度值）
	if s.length == 0 || i < 0 || i > s.length {
		//return errors.New("查找失败")
		err = errors.New("查找失败")
	}
	//*e = s.data[i]
	res = s.data[i]
	return
}

// 插入数据
// 传插入数据的时候既可以直接传值，也可以传指针
/*
func ListInsert(s *SqList, i int, e ElemType) error {
	if s.length == MAXSIZE {
		return errors.New("线性表已满，不能插入数据")
	}
	if i < 0 || i > s.length {
		return errors.New("插入的位置不正确")
	}
	//i位开始后移一位
	for j := s.length; j >= i; j-- {
		s.data[j+1] = s.data[j]
	}
	s.data[i] = e
	s.length++
	return nil
}
*/
func ListInsert(s *SqList, i int, e *ElemType) error {
	if s.length == MAXSIZE {
		return errors.New("线性表已满，不能插入数据")
	}
	if i < 0 || i > s.length {
		return errors.New("插入的位置不正确")
	}
	//i位开始后移一位
	for j := s.length; j >= i; j-- {
		s.data[j+1] = s.data[j]
	}
	s.data[i] = *e
	s.length++
	return nil
}

// 删除数据
func ListDelete(s *SqList, i int) (err error) {
	if s.length == 0 {
		// return errors.New("线性表为空")
		err = errors.New("线性表为空")
	}
	if i < 0 || i > s.length-1 {
		// return errors.New("删除的位置不正确")
		err = errors.New("删除的位置不正确")
	}

	// *e=s.data[i]
	for j := i; j < s.length-1; j++ {
		s.data[j] = s.data[j+1]
	}

	// 清除数据
	s.data[s.length-1] = 0
	s.length--
	return nil
}

func main() {
	// 添加数据
	var s SqList
	var result ElemType = 10
	ListInsert(&s, 0, &result)
	result = 20
	ListInsert(&s, 1, &result)
	result = 30
	ListInsert(&s, 2, &result)
	result = 40
	err := ListInsert(&s, 3, &result)
	fmt.Println("添加数据")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	// 获取数据元素
	_, res := GetElem(s, 1)
	fmt.Println(res)

	// 删除数据元素
	err = ListDelete(&s, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
