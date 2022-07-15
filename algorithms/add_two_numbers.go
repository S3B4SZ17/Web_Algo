package algorithms

import "fmt"

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
	value int
	next *Node
}

type List struct {
	linkedList *Node
	size int
}

func (l *List) Add(val int){
	
	if (l.linkedList == nil){
		l.linkedList = &Node{value: val, next: nil}
		l.size ++
	}else{
		node := &Node{value: val, next: nil}
		temp := l.linkedList
		
		for temp.next != nil{
			temp = temp.next
		}

		temp.next = node
		l.size ++
	}
}

func (l *List) PrintList() (res string){
	temp := l.linkedList
	for temp.next != nil{
		res += fmt.Sprintf("%d, ", temp.value)
		temp = temp.next
	}
	res += fmt.Sprintf("%d, ", temp.value)
	return 
}

func SumLists(list1 *List, list2 *List) *List{
	n := 0
	resList := &List{}
	temp1, temp2 := list1.linkedList, list2.linkedList
	confirmLength(list1, list2)
	for (temp1 != nil) || (temp2 != nil){	
		h := temp1.value + temp2.value
		d := h % 10
		r := d + n
		n = h / 10

		resList.Add(r)

		if temp1.next == nil || temp2.next == nil{
			break
		}
		temp1 = temp1.next
		temp2 = temp2.next
	}

	return resList
}

func confirmLength(list1 *List, list2 *List) {
	//We will fill any list with 0 until both are of the same size
	if list1.size != list2.size {
		for list1.size < list2.size{
			list1.Add(0)
		}
		for list2.size < list1.size{
			list1.Add(0)
		}
	}
	return
}