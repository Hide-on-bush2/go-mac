package main

import (
	"fmt"
	"github.com/Hide-on-bush2/tree"
)

func main() {
	t := tree.Make_node(4)
	t = tree.Insert(2, t)
	t = tree.Insert(1, t)
	t = tree.Insert(3, t)
	t = tree.Insert(7, t)
	t = tree.Insert(5, t)
	t = tree.Insert(8, t)

	fmt.Printf("preorder: " + tree.Preorder(t) + "\n")
	fmt.Printf("inorder: " + tree.Inorder(t) + "\n")
	fmt.Printf("postorder: " + tree.Postorder(t) + "\n")
}
