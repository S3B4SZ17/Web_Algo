package algorithms

import (
	"errors"
	"fmt"
)

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

type Node struct {
	Value int
	Next *Node
}

type List struct {
	LinkedList *Node
	Size int
}

func (l *List) Add(val int){
	
	if (l.LinkedList == nil){
		l.LinkedList = &Node{Value: val, Next: nil}
		l.Size ++
	}else{
		node := &Node{Value: val, Next: nil}
		temp := l.LinkedList
		
		for temp.Next != nil{
			temp = temp.Next
		}

		temp.Next = node
		l.Size ++
	}
}

func (l *List) PrintList() (res string){
	temp := l.LinkedList
	for temp.Next != nil{
		res += fmt.Sprintf("%d, ", temp.Value)
		temp = temp.Next
	}
	res += fmt.Sprintf("%d", temp.Value)
	return 
}

func SumLists(list1 *List, list2 *List) *List{
	n := 0
	resList := &List{}
	temp1, temp2 := list1.LinkedList, list2.LinkedList
	confirmLength(list1, list2)
	for (temp1 != nil) || (temp2 != nil){	
		h := temp1.Value + temp2.Value
		d := h % 10
		r := d + n
		n = h / 10

		// Another edge case
		if r >= 10{
			d = r % 10
			n = r / 10
			r = d
		}

		resList.Add(r)

		if temp1.Next == nil || temp2.Next == nil{
			break
		}
		temp1 = temp1.Next
		temp2 = temp2.Next
	}

	// Last case if 2342 + 9465 = 11807
	// in this case 9+2 spans 10 and will remain a 1, we will need to add it
	if n != 0{
		resList.Add(n)
	}

	return resList
}

func confirmLength(list1 *List, list2 *List) {
	//We will fill any list with 0 until both are of the same Size
	if list1.Size != list2.Size {
		for list1.Size < list2.Size{
			list1.Add(0)
		}
		for list2.Size < list1.Size{
			list2.Add(0)
		}
	}
	return
}

type ListVals struct {
	List []int `json:"list"`
}

func (l *List) AddFromList(list *ListVals) error{
	for _, v := range list.List {
		remain := v / 10
		if remain > 0 {
			err := errors.New("Invalid item on the list. Only numbers from 0 to 9 are allowed")
			return err
		}
		l.Add(v)
	}
	return nil
}

type Response struct {
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	Sum string `json:"sum"`
}
