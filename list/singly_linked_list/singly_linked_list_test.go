// Package singlylinkedlist provides a generic implementation of a singly linked
// list. A singly linked list is a linear data structure shere each element points
// to the next, with the last element pointing to nil.
//
// This package supports a variety of operations such as adding elements to the
// list, removing elements, and traversing through the list. It also provides
// functionality to check the size of the list and whether it is empty.
//
// Included features:
//   - Create a new singly linked list.
//   - Add elements to the beginning or end of the list.
//   - Search for an element.
//   - Remove elements by value or position.
//   - Check if the list is empty.
//   - Retrieve the first (head) and last (tail) elements.
//   - Reverse the list's order.
//   - Clear all elements from the list.
//   - Iterate over the list and apply a function to each element.
//   - Insert elements at a specified index.
//   - Remove all occurrences of a value from the list.
//
// Most methods handle cases where the list is empty and return nil or no-op
// accordingly. Methods like 'InsertAt()' and 'RemoveAll()' ensure safe list
// manipulation without causing invalid memory access.
package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListPrependOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListPrependOnNonEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Append("1")
	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListAppendOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListAppendOnNonEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Append("1")
	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

func TestLinkedListClear(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListIsEmpty(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.True(t, list.IsEmpty())
	list.Append(1)
	assert.False(t, list.IsEmpty())
}

func TestLinkedListSize(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.Equal(t, 0, list.Size())
	list.Append(1)
	assert.Equal(t, 1, list.Size())
	list.Append(2)
	assert.Equal(t, 2, list.Size())
}

func TestLinkedListHead(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.Nil(t, list.Head())
	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())
	list.Append(2)
	assert.Equal(t, 1, list.Head().Data())
}

func TestLinkedListTail(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.Nil(t, list.Tail())
	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())
	list.Append(2)
	assert.Equal(t, 2, list.Tail().Data())
}

func TestLinkedListFind(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	assert.NotNil(t, nodo)
	assert.Equal(t, 2, nodo.Data())
	nodo = list.Find(4)
	assert.Nil(t, nodo)
}

func TestLinkedListRemoveFirst(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveFirst()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.RemoveFirst()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListRemoveLast(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveLast()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 2, list.Tail().Data())
	list.RemoveLast()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListRemove(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.Remove(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.Remove(3)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListRemoveOnLastElement(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Remove(4)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestLinkedListRemoveNotExists(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(4)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestLinkedListRemoveFirstOnEmpty(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
}

func TestLinkedListRemoveLastOnEmpty(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
}

func TestLinkedListStringOnEmpty(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.Equal(t, "SinglyLinkedList: []", list.String())
}

func TestLinkedListString(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, "SinglyLinkedList: [1] → [2] → [3]", list.String())
}

func TestLinkedListInsertAt(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.InsertAt(0, 1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	list.InsertAt(1, 3)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.InsertAt(1, 2)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 2, list.Head().Next().Data())
}

func TestLinkedListInsertAtOutOfBounds(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.InsertAt(-1, 1)
	assert.Equal(t, 0, list.Size())
	list.InsertAt(2, 1)
	assert.Equal(t, 0, list.Size())
}

func TestLinkedListRemoveAll(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(1)
	list.Append(3)
	list.Append(1)
	list.RemoveAll(1)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestLinkedListRemoveAllNotFound(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.RemoveAll(3)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 2, list.Tail().Data())
}

func TestLinkedListReverse(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
	assert.Equal(t, "SinglyLinkedList: [3] → [2] → [1]", list.String())
}

func TestLinkedListReverseEmptyAndSingle(t *testing.T) {
	empty := NewSinglyLinkedList[int]()
	empty.Reverse()
	assert.Nil(t, empty.Head())
	assert.Nil(t, empty.Tail())
	single := NewSinglyLinkedList[int]()
	single.Append(1)
	single.Reverse()
	assert.Equal(t, 1, single.Head().Data())
	assert.Equal(t, 1, single.Tail().Data())
}

func TestLinkedListForEach(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	var result []int
	list.ForEach(func(value int) { result = append(result, value) })
	assert.Equal(t, []int{1, 2, 3}, result)
}
