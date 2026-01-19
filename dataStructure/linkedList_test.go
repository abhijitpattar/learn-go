package datastructure

import (
	"fmt"
)

func TestPrepend() {
	fmt.Println("testing")
	anode := node{
		data: 1,
	}
	bnode := node{
		data: 2,
	}

	fmt.Println(anode.data)
	fmt.Println(bnode.data)

	var ll LinkedList

	ll.prepend(anode)
	ll.prepend(bnode)

	fmt.Println(ll)

}
