package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := Make_node(4)
	tree = Insert(2, tree)
	tree = Insert(1, tree)
	tree = Insert(3, tree)
	tree = Insert(7, tree)
	tree = Insert(5, tree)
	tree = Insert(8, tree)

	t.Run("Preoder", func(t *testing.T) {
		expected := "4213758"

		got := Preorder(tree)
		if got != expected {
			t.Errorf("expected '%q' but got '%q'", expected, got)
		}
	})

	t.Run("Inorder", func(t *testing.T) {
		expected := "1234578"

		got := Inorder(tree)
		if got != expected {
			t.Errorf("expected '%q' but got '%q'", expected, got)
		}
	})

	t.Run("porstorder", func(t *testing.T) {
		expected := "1325874"

		got := Postorder(tree)
		if got != expected {
			t.Errorf("expected '%q' but got '%q'", expected, got)
		}
	})

	t.Run("find", func(t *testing.T) {
		got := Find(9, tree)

		if got != nil {
			t.Errorf("expected nil but got real pointer")
		}
	})

	t.Run("rebuild", func(t *testing.T) {
		expected := "4213759"

		tree = Delete(8, tree)
		tree = Insert(9, tree)
		got := Preorder(tree)
		if got != expected {
			t.Errorf("expected '%q' but got '%q'", expected, got)
		}

		expected = "421395"
		tree = Delete(7, tree)
		got = Preorder(tree)
		if got != expected {
			t.Errorf("expected '%q' but got '%q'", expected, got)
		}

		expected = "4213956"
		tree = Insert(6, tree)
		got = Preorder(tree)
		if got != expected {
			t.Errorf("expected '%q' but got '%q'", expected, got)
		}
	})
}
