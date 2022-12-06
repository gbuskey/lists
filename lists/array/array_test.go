package array_test

import (
	"lists/lists/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	arr := array.New()
	assert.Zero(t, arr.Length())
}

func TestArray(t *testing.T) {
	arr := array.New()
	assert.Zero(t, arr.Length())

	// Add
	err := arr.Add("one")
	assert.NoError(t, err)
	assert.Equal(t, 1, arr.Length())

	err = arr.Add("two")
	assert.NoError(t, err)
	assert.Equal(t, 2, arr.Length())

	err = arr.Add("three")
	assert.NoError(t, err)
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
