package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add_element_to_empty_list_should_create_head_and_tail_with_value(t *testing.T) {
	final_list := []int{2}
	linked_list := LinkedList{nil, nil}

	error := linked_list.Add(2)

	assert.Nil(t, error)
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

	error := linked_list.Add(3)

	assert.Nil(t, error)
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

	error := linked_list.Add(4)

	assert.Nil(t, error)
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

	error := linked_list.Add(6)

	assert.Nil(t, error)
	assert.Nil(t, linked_list.head.prev)
	assert.Nil(t, linked_list.tail.next)
	assert.NotEqual(t, linked_list.head.next, linked_list.tail)
	assert.NotEqual(t, linked_list.tail.prev, linked_list.head)

	assertList(t, &linked_list, final_list)
}

func assertList(t *testing.T, list *LinkedList, validation_list []int) {
	i := 0
	for element := list.head; element != nil; element = element.next {
		assert.Equal(t, validation_list[i], element.value)
		i++
	}
}
