package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert_element_with_no_other_elements_on_tree_root_should_be_the_element(t *testing.T) {
	tree := BinaryTree{nil}

	tree.Insert(1)

	assert.Equal(t, 1, tree.root.value)
}

func Test_insert_element_with_value_higher_than_root_should_insert_on_right(t *testing.T) {
	elements := []int{1, 2}
	tree := BinaryTree{nil}
	tree.Insert(1)

	tree.Insert(2)

	counter := 0
	validate(t, tree.root, elements, &counter)
}

func Test_insert_element_with_value_lower_than_root_should_insert_on_left(t *testing.T) {
	elements := []int{0, 1}
	tree := BinaryTree{nil}
	tree.Insert(1)

	tree.Insert(0)

	counter := 0
	validate(t, tree.root, elements, &counter)
}

func Test_insert_two_elements_with_value_lower_than_root_should_insert_on_left(t *testing.T) {
	elements := []int{0, 1, 2}
	tree := BinaryTree{nil}
	tree.Insert(2)
	tree.Insert(1)

	tree.Insert(0)

	counter := 0
	validate(t, tree.root, elements, &counter)
}

func Test_insert_two_elements_with_value_higher_than_root_should_insert_on_right(t *testing.T) {
	elements := []int{2, 3, 4}
	tree := BinaryTree{nil}
	tree.Insert(2)
	tree.Insert(3)

	tree.Insert(4)

	counter := 0
	validate(t, tree.root, elements, &counter)
}

func Test_insert_elements_with_value_higher_and_lower_than_root_should_insert_on_right_and_left(t *testing.T) {
	elements := []int{0, 1, 2, 3, 4}
	tree := BinaryTree{nil}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(0)

	tree.Insert(4)

	counter := 0
	validate(t, tree.root, elements, &counter)
}

func Test_insert_eight_elements_with_value_higher_and_lower_than_root_should_insert_on_right_and_left(t *testing.T) {
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7}
	tree := BinaryTree{nil}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(0)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(4)

	tree.Insert(5)

	counter := 0
	validate(t, tree.root, elements, &counter)
}

func Test_search_on_tree_with_no_elements_should_return_false(t *testing.T) {
	tree := BinaryTree{nil}

	found := tree.Search(5)

	assert.False(t, found)
}

func Test_search_with_element_equal_to_root_and_with_only_one_element_should_return_true(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(5)

	found := tree.Search(5)

	assert.True(t, found)
}

func Test_search_with_element_different_from_root_and_with_only_one_element_should_return_false(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)

	found := tree.Search(5)

	assert.False(t, found)
}

func Test_search_element_in_tree_with_four_elements_to_the_left_should_return_true_for_value_on_the_edge(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)

	found := tree.Search(1)

	assert.True(t, found)
}

func Test_search_element_in_tree_with_four_elements_to_the_left_should_return_false_for_value_not_on_tree(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)

	found := tree.Search(0)

	assert.False(t, found)
}

func Test_search_element_in_tree_with_four_elements_to_the_right_should_return_true_for_value_on_the_edge(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(11)
	tree.Insert(20)

	found := tree.Search(20)

	assert.True(t, found)
}

func Test_search_element_in_tree_with_four_elements_to_the_right_should_return_false_for_value_not_on_tree(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(11)
	tree.Insert(20)

	found := tree.Search(100)

	assert.False(t, found)
}

func Test_search_element_in_tree_with_five_elements_balanced_should_return_true_for_value_on_the_edge(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(20)

	found := tree.Search(2)

	assert.True(t, found)
}

func Test_search_element_in_tree_with_five_elements_balanced_should_return_false_for_value_not_on_tree(t *testing.T) {
	tree := BinaryTree{nil}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(20)

	found := tree.Search(100)

	assert.False(t, found)
}

func validate(t *testing.T, node *Node, elements []int, counter *int) {
	if node != nil {
		validate(t, node.left, elements, counter)
		assert.Equal(t, elements[*counter], node.value)
		*counter++
		validate(t, node.right, elements, counter)
	}
}
