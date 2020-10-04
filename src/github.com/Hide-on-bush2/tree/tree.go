package tree

import (
	"fmt"
	"strconv"
)

type node struct {
	val   int
	left  *node
	right *node
}

func Make_node(num int) *node {
	return &node{num, nil, nil}
}

func Insert(num int, tree *node) *node {
	if tree == nil {
		return Make_node(num)
	}
	if tree.val > num {
		tree.left = Insert(num, tree.left)
	} else if tree.val < num {
		tree.right = Insert(num, tree.right)
	}
	return tree

}

func Find(num int, tree *node) *node {
	if tree == nil {
		return nil
	}

	if tree.val == num {
		return tree
	} else if tree.val < num {
		return Find(num, tree.right)
	} else {
		return Find(num, tree.left)
	}
}

func FindMin(tree *node) *node {
	if tree == nil {
		fmt.Printf("the tree is empty")
		return nil
	}
	if tree.left == nil {
		return tree
	} else {
		return FindMin(tree.left)
	}
}

func Delete(num int, tree *node) *node {
	if tree == nil {
		fmt.Printf("can not delete node from a empty tree\n")
		return nil
	}
	if num < tree.val {
		tree.left = Delete(num, tree.left)
	} else if num > tree.val {
		tree.right = Delete(num, tree.right)
	} else {
		if tree.left != nil && tree.right != nil {
			tmp := FindMin(tree.right)
			tree.val = tmp.val
			tree.right = Delete(tmp.val, tree.right)
		} else {
			if tree.left == nil {
				tree = tree.right
			} else {
				tree = tree.left
			}
		}
	}
	return tree

}

func Preorder(tree *node) string {
	if tree == nil {
		return ""
	}

	return strconv.Itoa(tree.val) + Preorder(tree.left) + Preorder(tree.right)
}

func Inorder(tree *node) string {
	if tree == nil {
		return ""
	}

	return Inorder(tree.left) + strconv.Itoa(tree.val) + Inorder(tree.right)
}

func Postorder(tree *node) string {
	if tree == nil {
		return ""
	}

	return Postorder(tree.left) + Postorder(tree.right) + strconv.Itoa(tree.val)
}

// func main() {
// 	tree := make_node(4)
// 	tree = insert(2, tree)
// 	tree = insert(1, tree)
// 	tree = insert(3, tree)
// 	tree = insert(7, tree)
// 	tree = insert(5, tree)
// 	tree = insert(8, tree)

// 	fmt.Printf("preorder: " + preorder(tree) + "\n")
// 	fmt.Printf("inorder: " + inorder(tree) + "\n")
// 	fmt.Printf("postorder: " + postorder(tree) + "\n")
// }
