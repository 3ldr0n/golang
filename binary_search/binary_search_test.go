package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search_with_element_in_the_middle_of_the_list_should_return_true(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	element := 3
	result := binary_search(element, list)

	assert.Equal(t, true, result)
}

func Test_search_for_first_element_in_the_list_should_return_true(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	element := 1
	result := binary_search(element, list)

	assert.Equal(t, true, result)
}

func Test_search_for_last_element_in_the_list_should_return_true(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	element := 5
	result := binary_search(element, list)

	assert.Equal(t, true, result)
}

func Test_search_with_number_after_list_should_return_false(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	element := 6
	result := binary_search(element, list)

	assert.Equal(t, false, result)
}

func Test_search_with_number_before_list_should_return_false(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	element := 0
	result := binary_search(element, list)

	assert.Equal(t, false, result)
}
