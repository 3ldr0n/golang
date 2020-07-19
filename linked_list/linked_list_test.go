package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add_element_to_empty_list_should_create_head_and_tail_with_value(t *testing.T) {
	final_list := []int{2}
	linked_list := LinkedList{nil, nil}

	linked_list.Add(2)

	assert.Equal(t, 2, linked_list.head.value)
	assert.Equal(t, 2, linked_list.tail.value)
	assert.Nil(t, linked_list.head.prev)
	assert.Nil(t, linked_list.tail.next)

	assertList(t, &linked_list, final_list)
}

func Test_add_element_to_list_with_one_element_should_have_different_head_and_tail(t *testing.T) {
	final_list := []int{2, 3}
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)

	linked_list.Add(3)

	assert.Nil(t, linked_list.head.prev)
	assert.Nil(t, linked_list.tail.next)
	assert.Equal(t, linked_list.head.next, linked_list.tail)
	assert.Equal(t, linked_list.tail.prev, linked_list.head)

	assertList(t, &linked_list, final_list)
}

func Test_add_element_to_list_with_two_elements_list_should_have_three_elements(t *testing.T) {
	final_list := []int{2, 3, 4}
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)

	linked_list.Add(4)

	assert.Nil(t, linked_list.head.prev)
	assert.Nil(t, linked_list.tail.next)
	assert.NotEqual(t, linked_list.head.next, linked_list.tail)
	assert.NotEqual(t, linked_list.tail.prev, linked_list.head)
	assert.Equal(t, linked_list.head.next, linked_list.tail.prev)

	assertList(t, &linked_list, final_list)
}

func Test_add_element_to_list_with_four_elements_list_should_have_five_elements(t *testing.T) {
	final_list := []int{2, 3, 4, 5, 6}
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)

	linked_list.Add(6)

	assert.Nil(t, linked_list.head.prev)
	assert.Nil(t, linked_list.tail.next)
	assert.NotEqual(t, linked_list.head.next, linked_list.tail)
	assert.NotEqual(t, linked_list.tail.prev, linked_list.head)

	assertList(t, &linked_list, final_list)
}

func Test_search_first_element_in_list_should_return_true(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)

	found := linked_list.Search(2)

	assert.True(t, found)
}

func Test_search_last_element_in_list_should_return_true(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)

	found := linked_list.Search(5)

	assert.True(t, found)
}

func Test_search_element_in_middle_of_odd_list_should_return_true(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)
	linked_list.Add(6)

	found := linked_list.Search(4)

	assert.True(t, found)
}

func Test_search_element_in_first_middle_part_of_even_list_should_return_true(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)
	linked_list.Add(6)
	linked_list.Add(7)

	found := linked_list.Search(4)

	assert.True(t, found)
}

func Test_search_element_in_second_middle_part_of_even_list_should_return_true(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)
	linked_list.Add(6)
	linked_list.Add(7)

	found := linked_list.Search(5)

	assert.True(t, found)
}

func Test_search_for_element_not_on_list_should_return_false(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)
	linked_list.Add(6)
	linked_list.Add(7)

	found := linked_list.Search(15)

	assert.False(t, found)
}

func Test_search_for_element_in_empty_list_should_return_false(t *testing.T) {
	linked_list := LinkedList{nil, nil}

	found := linked_list.Search(1)

	assert.False(t, found)
}

func Test_removing_last_element_from_list_with_multiple_elements_should_remove_last_element(t *testing.T) {
	final_list := []int{2, 3, 4, 5, 6}
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)
	linked_list.Add(6)
	linked_list.Add(7)

	error := linked_list.RemoveLast()

	assert.Nil(t, error)

	assertList(t, &linked_list, final_list)
}

func Test_removing_last_element_from_empty_list_should_keep_list_empty(t *testing.T) {
	linked_list := LinkedList{nil, nil}

	error := linked_list.RemoveLast()

	assert.Nil(t, error)
	assert.Nil(t, linked_list.head)
	assert.Nil(t, linked_list.tail)
}

func Test_removing_last_element_from_list_with_one_element_should_make_list_empty(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)

	error := linked_list.RemoveLast()

	assert.Nil(t, error)
	assert.Nil(t, linked_list.head)
	assert.Nil(t, linked_list.tail)
}

func Test_removing_first_element_from_list_with_multiple_elements_should_remove_first_element(t *testing.T) {
	final_list := []int{3, 4, 5, 6, 7}
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)
	linked_list.Add(3)
	linked_list.Add(4)
	linked_list.Add(5)
	linked_list.Add(6)
	linked_list.Add(7)

	error := linked_list.RemoveFirst()

	assert.Nil(t, error)

	assertList(t, &linked_list, final_list)
}

func Test_removing_first_element_from_empty_list_should_keep_list_empty(t *testing.T) {
	linked_list := LinkedList{nil, nil}

	error := linked_list.RemoveFirst()

	assert.Nil(t, error)
	assert.Nil(t, linked_list.head)
	assert.Nil(t, linked_list.tail)
}

func Test_removing_first_element_from_list_with_one_element_should_make_list_empty(t *testing.T) {
	linked_list := LinkedList{nil, nil}
	linked_list.Add(2)

	error := linked_list.RemoveFirst()

	assert.Nil(t, error)
	assert.Nil(t, linked_list.head)
	assert.Nil(t, linked_list.tail)
}

func assertList(t *testing.T, list *LinkedList, validation_list []int) {
	i := 0
	for element := list.head; element != nil; element = element.next {
		assert.Equal(t, validation_list[i], element.value)
		i++
	}
}
