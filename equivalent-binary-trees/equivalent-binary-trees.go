package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // close this channel when this function returns
	
	// walk closure
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t.Left != nil {
			walk(t.Left) // recursive
		}
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right) // recursive
		}
	}
	
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	x, ok1 := <- ch1
	y, ok2 := <- ch2
	
	if x == y && ok1 == ok2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
