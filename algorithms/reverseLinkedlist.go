package algorithms

import (
	"fmt"

	pbRevList "github.com/S3B4SZ17/Web_Algo/proto/reverseLinkedlist"
)

/*
You are given a single linked lists and you just neet to reverse the order of the items

Example 1:

Input: l1 = 2 -> 4 -> 3
Output: 3 -> 4 -> 2
Explanation: Turn it over

*/

func (l *List) Add_seq(val int32) {

	if l.LinkedList == nil {
		l.LinkedList = &Node{Value: val, Next: nil}
		l.Size++
	} else {
		node := &Node{Value: val, Next: nil}
		node.Next = l.LinkedList
		l.LinkedList = node
		l.Size++
	}
}

func (l *List) Reverse() {
	temp := l.LinkedList

	if l.LinkedList == nil {
		fmt.Println("List is empty")
		return
	} else {
		next := &Node{}
		before := &Node{}
		for temp != nil {
			before = temp.Next
			temp.Next = next
			next = temp
			temp = before
		}
		l.LinkedList = next
	}
}

func (l *List) CallReverseLinkedList(linkedList *pbRevList.LinkedList) {

	// "Weird error, cant do a range over a slice of []int32
	// for val := range linkedList.GetLinkedList() {
	// 	tmp := int32(val)
	// 	fmt.Printf("%d, ", tmp)
	// 	l.Add_seq(tmp)
	// }

	for i := 0; i < len(linkedList.LinkedList); i++ {
		l.Add_seq(linkedList.LinkedList[i])
	}

}
