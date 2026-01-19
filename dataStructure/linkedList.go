package datastructure

import "fmt"

type Node struct {
	Data          int
	nextNodeAddrs *Node
}

type LinkedList struct {
	rootNodeAddrs *Node
	tailNodeAddrs *Node
	Length        int
}

func (l *LinkedList) Prepend(n Node) {
	if l.Length == 0 {
		l.rootNodeAddrs = &n
		l.tailNodeAddrs = &n
		l.Length++
	} else {
		temp := l.rootNodeAddrs
		l.rootNodeAddrs = &n
		n.nextNodeAddrs = temp
		l.Length++
	}
	fmt.Println("added at the start -", n.Data)
}

func (l *LinkedList) Append(n Node) {
	if l.Length == 0 {
		l.rootNodeAddrs = &n
		l.tailNodeAddrs = &n
		l.Length++
	} else {
		l.tailNodeAddrs.nextNodeAddrs = &n
		l.tailNodeAddrs = &n
		l.Length++
	}
	fmt.Println("added at the end -", n.Data)
}

func (l LinkedList) PrintAsArray() {
	retVal := []int{}
	PrintNode := l.rootNodeAddrs
	for l.Length != 0 {
		fmt.Printf("%d,", PrintNode.Data)
		retVal = append(retVal, PrintNode.Data)
		if PrintNode.nextNodeAddrs != nil {
			PrintNode = PrintNode.nextNodeAddrs
		}
		l.Length--
	}
	fmt.Printf("\n")

	println("retVal length -", len(retVal))
	println(retVal)
}
