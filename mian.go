package main

import (
	"errors"
	"fmt"
)

// element type 
type ElemTypeInt int
// apply for memory
const maxSize = 20
// define data element, include array and array length
type seqList struct {
	array 		[maxSize] ElemTypeInt
	arrayLength int
}

// get data element
func getData (seq seqList, i int) (err error, res ElemTypeInt) {
	// judge if sequence is empty or if subscript crossover
	if seq.length == 0 || i < 0 || i > s.length {
		err = errors.New("find fail")
	} 

	res = seq.array[i]
	
	return nil, res
}

// insert data element
func insertData (seq seqList, i int, e ElemTypeInt) (err error) {
	// judge if sequence is full
	if seq.
}





















