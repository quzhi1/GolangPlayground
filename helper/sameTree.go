package helper

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// Note that ch is closed inside Walk
func Walk(t *tree.Tree, ch chan int) {
	recursive(t, ch)
	close(ch)
}

func recursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		recursive(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		recursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	v1, ok1 := <-ch1
	v2, ok2 := <-ch2
	for ok1 && ok2 {
		if v1 != v2 {
			return false
		}
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
	}
	return ok1 == ok2
}

// PrintExampleTree prints example tree node in order
func PrintExampleTree() {
	fmt.Println("Tree model: ", tree.New(1).String())
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Print("Tree traversal: ")
	for i := range ch {
		fmt.Print(i, " ")
	}
}

// VerifyTreeSame verify Same function
func VerifyTreeSame() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
