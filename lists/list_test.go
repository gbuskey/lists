package lists_test

import (
	"lists/lists"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	testList(t, lists.Array)
}

func TestLinkedList(t *testing.T) {
	testList(t, lists.LinkedList)
}

func testList(t *testing.T, listType lists.ListType) {
	arr := lists.New(listType)
	assert.Zero(t, arr.Length())

	// Add
	arr.Add("one")
	assert.Equal(t, 1, arr.Length())

	arr.Add("two")
	assert.Equal(t, 2, arr.Length())

	arr.Add("three")
	assert.Equal(t, 3, arr.Length())

	// Find
	num, err := arr.Find("two")
	assert.NoError(t, err)
	assert.Equal(t, "two", num)

	num, err = arr.Find("three")
	assert.NoError(t, err)
	assert.Equal(t, "three", num)

	num, err = arr.Find("one")
	assert.NoError(t, err)
	assert.Equal(t, "one", num)

	num, err = arr.Find("six")
	assert.EqualError(t, err, "unable to find element 'six'")
	assert.Nil(t, num)

	// Remove
	num, err = arr.Remove("four")
	assert.EqualError(t, err, "unable to remove element. unable to find element 'four'")
	assert.Nil(t, num)

	num, err = arr.Remove("two")
	assert.NoError(t, err)
	assert.Equal(t, "two", num)
	assert.Equal(t, 2, arr.Length())

	num, err = arr.Remove("one")
	assert.NoError(t, err)
	assert.Equal(t, "one", num)
	assert.Equal(t, 1, arr.Length())

	num, err = arr.Remove("three")
	assert.NoError(t, err)
	assert.Equal(t, "three", num)
	assert.Equal(t, 0, arr.Length())
}
