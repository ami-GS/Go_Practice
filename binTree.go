package main

import (
	"fmt"
	"math/rand"
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
	if v == t.Value {
		return true
	} else {
		if v > t.Value {
			if t.Right == nil {
				return false
			}
			fmt.Print(string('R'))
			return t.Right.Find(v)
		} else {
			if t.Left == nil {
				return false
			}
			fmt.Print(string('L'))
			return t.Left.Find(v)
		}
	}
	//return false
}

func (t *Tree) Traversal_pre() {
	fmt.Println(t.Value)
	if t.Left != nil {
		t.Left.Traversal_pre()
	}
	if t.Right != nil {
		t.Right.Traversal_pre()
	}
}

func (t *Tree) Traversal_post() {
	if t.Left == nil && t.Right == nil {
		fmt.Println(t.Value)
	}
	if t.Left != nil {
		t.Left.Traversal_post()
	}
	if t.Right != nil {
		t.Right.Traversal_post()
	}
	if t.Left != nil || t.Right != nil {
		fmt.Println(t.Value)
	}
}

func (t *Tree) Traversal_in() {
	if t.Left == nil && t.Right == nil {
		fmt.Println(t.Value)
	}
	if t.Left != nil {
		t.Left.Traversal_in()
		fmt.Println(t.Value)
	}
	if t.Right != nil {
		t.Right.Traversal_in()
		if t.Left == nil {
			fmt.Println(t.Value)
		}
	}
}

func main() {
	root := Tree{nil, nil, 5}
	randSeq := rand.Perm(10)
	fmt.Println(randSeq)
	for _, v := range randSeq {
		root.Set(v)
	}

	for i := 0; i < 15; i++ {
		fmt.Println(i)
		fmt.Println(root.Find(i))
	}
	root.Traversal_pre()
	root.Traversal_post()
	root.Traversal_in()
}
