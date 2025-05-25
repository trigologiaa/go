// Package heap provides a generic implementation of a binary heap data structure,
// supporting both min-heap and max-heap configurations.
//
// A heap is a complete binary tree where the value of each node is ordered with
// respect to its children according to a comparator function. This package allows
// storing elements of any type and defines custom behavior via a comparator
// function.
//
// Included features:
//   - Create a generic heap using a custom comparator.
//   - Create a min-heap or max-heap.
//   - Insert elements into the heap.
//   - Remove and return the root element (minimum or maximum depending on the
//     heap).
//   - Retrieve the current size of the heap.
//   - Access the internal slice of elements for inspection or testing purposes.
//
// The implementation ensures the heap property is maintained on insertions and
// removals using up-heap and down-heap operations.
package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	name string
	age  int
}

func peopleFromOldestToYoungest(a Person, b Person) int {
	if a.age < b.age {
		return 1
	} else if a.age > b.age {
		return -1
	}
	return 0
}

func intComparator(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// TestHeapCreteEmpty() verifies that a newly created generic heap has a size of 0
// and is properly initialized.
func TestHeapCreteEmpty(t *testing.T) {
	m := NewGenericHeap(peopleFromOldestToYoungest)
	assert.Equal(t, 0, m.Size())
}

// TestHeapRemoveEmpty() ensures that attempting to remove an element from an empty
// heap results in an error.
func TestHeapRemoveEmpty(t *testing.T) {
	m := NewGenericHeap(peopleFromOldestToYoungest)
	_, err := m.Remove()
	assert.NotNil(t, err)
}

// TestHeapCreateInsertAndExtract() verifies that inserting elements into the heap
// maintains the heap property, and that removing elements returns them in the
// correct order as defined by the comparator.
func TestHeapCreateInsertAndExtract(t *testing.T) {
	insertionSequence := []Person{
		{"Ana", 44},
		{"Juan", 29},
		{"Pedro", 58},
		{"Maria", 2},
		{"Jose", 98},
		{"Lucia", 11},
		{"Carlos", 65},
		{"Sofia", 3},
		{"Miguel", 68},
		{"Laura", 99},
	}
	orderExpectedAfterInserting := [][]Person{
		{{"Ana", 44}},
		{{"Ana", 44}, {"Juan", 29}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}, {"Maria", 2}},
		{{"Jose", 98}, {"Pedro", 58}, {"Ana", 44}, {"Maria", 2}, {"Juan", 29}},
		{{"Jose", 98}, {"Pedro", 58}, {"Ana", 44}, {"Maria", 2}, {"Juan", 29}, {"Lucia", 11}},
		{{"Jose", 98}, {"Pedro", 58}, {"Carlos", 65}, {"Maria", 2}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}},
		{{"Jose", 98}, {"Pedro", 58}, {"Carlos", 65}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}},
		{{"Jose", 98}, {"Miguel", 68}, {"Carlos", 65}, {"Pedro", 58}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}},
		{{"Laura", 99}, {"Jose", 98}, {"Carlos", 65}, {"Pedro", 58}, {"Miguel", 68}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}, {"Juan", 29}},
	}
	m := NewGenericHeap(peopleFromOldestToYoungest)
	assert.Equal(t, 0, m.Size())
	for i := range insertionSequence {
		m.Insert(insertionSequence[i])
		assert.Equal(t, orderExpectedAfterInserting[i], m.elements)
	}
	expectedOrderAfterDeleting := [][]Person{
		{{"Jose", 98}, {"Miguel", 68}, {"Carlos", 65}, {"Pedro", 58}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}},
		{{"Miguel", 68}, {"Pedro", 58}, {"Carlos", 65}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}},
		{{"Carlos", 65}, {"Pedro", 58}, {"Ana", 44}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Maria", 2}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}, {"Sofia", 3}, {"Maria", 2}, {"Lucia", 11}},
		{{"Ana", 44}, {"Juan", 29}, {"Lucia", 11}, {"Sofia", 3}, {"Maria", 2}},
		{{"Juan", 29}, {"Sofia", 3}, {"Lucia", 11}, {"Maria", 2}},
		{{"Lucia", 11}, {"Sofia", 3}, {"Maria", 2}},
		{{"Sofia", 3}, {"Maria", 2}},
		{{"Maria", 2}},
		{},
	}
	for i := range insertionSequence {
		_, err := m.Remove()
		assert.Equal(t, expectedOrderAfterDeleting[i], m.elements)
		assert.NoError(t, err)
	}
}

// TestMaxHeapCreateEmpty() verifies that a newly created max-heap has a size of 0.
func TestMaxHeapCreateEmpty(t *testing.T) {
	m := NewMaxHeap(intComparator)
	assert.Equal(t, 0, m.Size())
}

// TestMaxHeapRemoveMaxEmpty() ensures that removing an element from an empty
// max-heap returns an error.
func TestMaxHeapRemoveMaxEmpty(t *testing.T) {
	m := NewMaxHeap(intComparator)
	_, err := m.Remove()
	assert.NotNil(t, err)
}

// TestMaxHeapCreateInsertAndExtract() validates the correct behavior of a max-heap
// during insertions and removals. It verifies that the internal structure
// maintains the heap property and that elements are removed from largest to
// smallest.
func TestMaxHeapCreateInsertAndExtract(t *testing.T) {
	insertionSequence := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}
	orderExpectedAfterInsert := [][]int{
		{44},
		{44, 29},
		{58, 29, 44},
		{58, 29, 44, 2},
		{98, 58, 44, 2, 29},
		{98, 58, 44, 2, 29, 11},
		{98, 58, 65, 2, 29, 11, 44},
		{98, 58, 65, 3, 29, 11, 44, 2},
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{99, 98, 65, 58, 68, 11, 44, 2, 3, 29},
	}
	m := NewMaxHeap(intComparator)
	assert.Equal(t, 0, m.Size())
	for i := range insertionSequence {
		m.Insert(insertionSequence[i])
		assert.Equal(t, orderExpectedAfterInsert[i], m.Elements())
	}
	orderExpectedAfterDelete := [][]int{
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{68, 58, 65, 3, 29, 11, 44, 2},
		{65, 58, 44, 3, 29, 11, 2},
		{58, 29, 44, 3, 2, 11},
		{44, 29, 11, 3, 2},
		{29, 3, 11, 2},
		{11, 3, 2},
		{3, 2},
		{2},
		{},
	}
	for i := range insertionSequence {
		_, err := m.Remove()
		assert.Equal(t, orderExpectedAfterDelete[i], m.Elements())
		assert.NoError(t, err)
	}
}

// TestMinHeapCreateEmpty() verifies that a newly created min-heap has a size of 0.
func TestMinHeapCreateEmpty(t *testing.T) {
	m := NewMinHeap(intComparator)
	assert.Equal(t, 0, m.Size())
}

// TestMinHeapRemoveMaxEmpty() ensures that removing an element from an empty
// min-heap returns an error.
func TestMinHeapRemoveMaxEmpty(t *testing.T) {
	m := NewMinHeap(intComparator)
	_, err := m.Remove()
	assert.NotNil(t, err)
}

// TestMinHeapCreateInsertAndExtract() tests the insertion and extraction
// operations on a min-heap, ensuring the heap property is maintained and elements
// are returned in increasing order.
func TestMinHeapCreateInsertAndExtract(t *testing.T) {
	insertionSequence := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}
	orderExpectedAfterInsert := [][]int{
		{44},
		{29, 44},
		{29, 44, 58},
		{2, 29, 58, 44},
		{2, 29, 58, 44, 98},
		{2, 29, 11, 44, 98, 58},
		{2, 29, 11, 44, 98, 58, 65},
		{2, 3, 11, 29, 98, 58, 65, 44},
		{2, 3, 11, 29, 98, 58, 65, 44, 68},
		{2, 3, 11, 29, 98, 58, 65, 44, 68, 99},
	}
	m := NewMinHeap(intComparator)
	assert.Equal(t, 0, m.Size())
	for i := range insertionSequence {
		m.Insert(insertionSequence[i])
		assert.Equal(t, orderExpectedAfterInsert[i], m.Elements())
	}
}
