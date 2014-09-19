package main

import (
	"fmt"
)

type Tree struct {
	Left, Right *Tree
	Value       int
}

func (t *Tree) Set(v int) {
	if v > t.Value {
		if t.Right == nil {
			t.Right = &Tree{nil, nil, v}
		} else {
			t.Right.Set(v)
		}
	} else {
		if t.Left == nil {
			t.Left = &Tree{nil, nil, v}
		} else {
			t.Left.Set(v)
		}
	}
}

func (t *Tree) Find(v int) bool {
	if t.Right == nil && t.Left == nil {
		return false
	}
	if v == t.Value {
		return true
	} else {
		if v > t.Value {
			fmt.Print(string('R'))
			return t.Right.Find(v)
		} else {
			fmt.Print(string('L'))
			return t.Left.Find(v)
		}
	}
	return false
}

func main() {
	root := Tree{nil, nil, 5}
	for i := 0; i < 10; i++ {
		root.Set(i)
	}
	for i := 0; i < 15; i++ {
		fmt.Println(i)
		fmt.Println(root.Find(i))
	}
}
