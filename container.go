package main

import (
	"fmt"
)

//Interfaces for containers and iterators
type Container interface {
	begin() Iterator
	append(v interface{})
}
type Iterator interface {
	Next() Iterator
	IsEnd() bool
	Deref() interface{}
}

//Vector structure
type Vector []interface{}

//Vector methods
func (v *Vector) append(val interface{}) {
	*v = append(*v, val)
}
func (thisV *Vector) begin() Iterator {
	return &vectorIterator{
		v:   *thisV,
		ndx: 0,
	}
}

//Vector Iterator structure
type vectorIterator struct {
	v   Vector
	ndx int
}

//Vector Iterator methods
func (vIt *vectorIterator) Next() Iterator {
	vIt.ndx += 1
	return vIt
}
func (vIt *vectorIterator) IsEnd() bool {
	if vIt.ndx >= len(vIt.v) {
		//fmt.Printf("ndx: %d\nsize: %d\n", vIt.ndx, vIt.v.size())
		return true
	}
	return false
}
func (vIt *vectorIterator) Deref() interface{} {
	return vIt.v[vIt.ndx]
}

//List Structure
type List struct {
	curNodeVal interface{}
	nxtNodePtr *List
}

//List Methods
func (l *List) append(val interface{}) {

	node := l
	prevNode := node
	for true {
		if node != nil && node.curNodeVal != nil {
			//fmt.Println("node was not nil, it holds: ", node.curNodeVal)
			prevNode = node
			node = node.nxtNodePtr
			//fmt.Println("updated node:", node)
		} else {
			break
		}
	}
	//fmt.Println("node reached nil value. Now trying to insert", val)
	if node == nil {
		var newListItem List
		newListItem.curNodeVal = val
		newListItem.nxtNodePtr = nil
		prevNode.nxtNodePtr = &newListItem
	} else {
		node.curNodeVal = val
	}
}
func (thisL *List) begin() Iterator {
	return &listIterator{
		l: thisL,
	}
}

//List Iterator structure
type listIterator struct {
	l *List
}

//List Iterator Methods
func (lIt *listIterator) Next() Iterator {
	lIt.l = lIt.l.nxtNodePtr
	return lIt
}
func (lIt *listIterator) IsEnd() bool {
	//fmt.Println("lIt.l: ", lIt)
	if lIt.l == nil {
		return true
	}
	return false
}
func (lIt *listIterator) Deref() interface{} {
	return lIt.l.curNodeVal
}

//Functions that work on Containers
func SumInt(c Container) int {
	//Sum all the ints in the container
	intSum := 0
	for it := c.begin(); it.IsEnd() != true; it.Next() {
		if val, ok := it.Deref().(int); ok {
			//fmt.Println("currently adding to sum: ", val)
			intSum += val
		}
	}
	return intSum
}
func SumFloat64(c Container) float64 {
	//Sum all the float64s in the container
	float64Sum := 0.0
	for it := c.begin(); it.IsEnd() != true; it.Next() {
		if val, ok := it.Deref().(float64); ok {
			float64Sum += val
		}
	}
	return float64Sum
}

func main() {
	fmt.Println("Welcome to CS341 Project 3!")
	fmt.Println("✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ Testing Vector ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶")
	var myVecIt vectorIterator
	fmt.Print("✶ Before appending anything myVecIt:         ")
	fmt.Println(myVecIt.v, "<- Empty Vector")
	fmt.Print("✶ After appending 4 ints to myVecIt:         ")
	myVecIt.v.append(1)
	myVecIt.v.append(2)
	myVecIt.v.append(3)
	myVecIt.v.append(4)
	fmt.Println(myVecIt.v)
	fmt.Print("✶ After appending 4 float64s to myVecInt:    ")
	myVecIt.v.append(2.5)
	myVecIt.v.append(5.5)
	myVecIt.v.append(10.5)
	myVecIt.v.append(1.5)
	fmt.Println(myVecIt.v)
	fmt.Print("✶ Using SumInt on our vector results in:     ")
	fmt.Println(SumInt(&myVecIt.v))
	fmt.Print("✶ Using SumFloat64 on our vector results in: ")
	fmt.Println(SumFloat64(&myVecIt.v))
	fmt.Println("✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶ ✶")
	fmt.Println()
	fmt.Println("❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ Testing List ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉")
	var myLisIt listIterator
	myLisIt.l = new(List)
	fmt.Print("❉ Before appending anything myLisIt:         ")
	fmt.Println(myLisIt.l, "<- Empty List")
	fmt.Print("❉ After appending 4 ints to myLisIt:         ")
	myLisIt.l.append(1)
	myLisIt.l.append(2)
	myLisIt.l.append(3)
	myLisIt.l.append(4)
	fmt.Print("{")
	myNode := myLisIt.l
	for true { //to print all nodes
		if myNode != nil && myNode.curNodeVal != nil {
			fmt.Print(myNode.curNodeVal, " ")
			myNode = myNode.nxtNodePtr
		} else {
			break
		}
	}
	fmt.Println("}")
	fmt.Print("❉ After appending 4 float64s to myListIt:    ")
	myLisIt.l.append(2.5)
	myLisIt.l.append(5.5)
	myLisIt.l.append(10.5)
	myLisIt.l.append(1.5)
	fmt.Print("{")
	myNode = myLisIt.l
	for true { //to print all nodes
		if myNode != nil && myNode.curNodeVal != nil {
			fmt.Print(myNode.curNodeVal, " ")
			myNode = myNode.nxtNodePtr
		} else {
			break
		}
	}
	fmt.Println("}")
	fmt.Print("❉ Using SumInt on our list results in:     ")
	fmt.Println(SumInt(myLisIt.l))
	fmt.Print("❉ Using SumFloat64 on our list results in: ")
	fmt.Println(SumFloat64(myLisIt.l))
	fmt.Println("❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉ ❉")
}
