// Package set provides a generic set data structure implemented using Go generics.
// It allows storing and manipulating unique elements of any comparable type (T).
//
// This package is useful for operations requiring collections of unique items,
// such as membership tests, unions, intersections, and set differences.
//
// Included features:
//   - Create a new set with initial elements.
//   - Add elements to the set (ensuring uniqueness).
//   - Remove elements from the set.
//   - Check if an element exists in the set.
//   - Get the number of elements in the set.
//   - Retrieve all elements as a slice.
//   - Clear all elements from the set.
//   - Check if the set is empty.
//   - Perform set operations: union, intersection, difference, symmetric difference.
//   - Compare sets for equality, subset, and superset relationships.
//   - Get a string representation of the set contents.
//
// Most methods return an error if the set receiver is nil.
package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSetNewSet() verifies that a newly created set is not nil and has a size of
// zero.
func TestSetNewSet(t *testing.T) {
	set := NewSet[int]()
	assert.NotNil(t, set)
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 0, size)
}

// TestSetAdd() ensures that adding a single element works correctly and increases
// the size of the set.
func TestSetAdd(t *testing.T) {
	set := NewSet[int]()
	err := set.Add(1)
	assert.NoError(t, err)
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
}

// TestSetAddMultiple() verifies that adding multiple elements at once works
// correctly and updates the size.
func TestSetAddMultiple(t *testing.T) {
	set := NewSet[int]()
	err := set.Add(1, 2, 3)
	assert.NoError(t, err)
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 3, size)
}

// TestSetAddExistingIsNotRepeated() ensures that adding an existing element does
// not duplicate it in the set.
func TestSetAddExistingIsNotRepeated(t *testing.T) {
	set := NewSet[int]()
	err := set.Add(1)
	assert.NoError(t, err)
	err = set.Add(1)
	assert.NoError(t, err)
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
}

// TestSetContains() checks that Contains() correctly reports whether an element is
// in the set.
func TestSetContains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	exists, err := set.Contains(1)
	assert.NoError(t, err)
	assert.True(t, exists)
	exists, err = set.Contains(2)
	assert.NoError(t, err)
	assert.False(t, exists)
}

// TestSetRemove() verifies that Remove() correctly deletes an existing element and
// updates the size.
func TestSetRemove(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	exists, err := set.Contains(1)
	assert.NoError(t, err)
	assert.True(t, exists)
	err = set.Remove(1)
	assert.NoError(t, err)
	exists, err = set.Contains(1)
	assert.NoError(t, err)
	assert.False(t, exists)
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
}

// TestSetRemoveNonExistent() ensures that removing a non-existent element does not
// affect the set.
func TestSetRemoveNonExistent(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
	err = set.Remove(2)
	assert.NoError(t, err)
	size, err = set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
}

// TestSetSize() checks that Size() returns the correct number of elements after
// each modification.
func TestSetSize(t *testing.T) {
	set := NewSet[int]()
	size, err := set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 0, size)
	err = set.Add(1)
	assert.NoError(t, err)
	size, err = set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
	err = set.Add(2)
	assert.NoError(t, err)
	size, err = set.Size()
	assert.NoError(t, err)
	assert.Equal(t, 2, size)
}

// TestSetValuesOnAnEmptySet() verifies that Values() returns an empty slice for an
// empty set.
func TestSetValuesOnAnEmptySet(t *testing.T) {
	set := NewSet[int]()
	values, err := set.Values()
	assert.NoError(t, err)
	assert.Empty(t, values)
}

// TestSetValuesOnANonEmptySet() checks that Values() returns all elements of the
// set.
func TestSetValuesOnANonEmptySet(t *testing.T) {
	set := NewSet(1, 2)
	values, err := set.Values()
	assert.NoError(t, err)
	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

// TestSetStringInEmptySet() verifies that the string representation of an empty
// set is formatted correctly.
func TestSetStringInEmptySet(t *testing.T) {
	set := NewSet[int]()
	str := set.String()
	assert.Equal(t, "Set: []", str)
}

// TestSetStringInNonEmptySet() checks that the string representation of a
// non-empty set contains all elements.
func TestSetStringInNonEmptySet(t *testing.T) {
	set := NewSet(1, 2)
	possibleRepresentations := []string{"Set: [1 2]", "Set: [2 1]"}
	str := set.String()
	assert.Contains(t, possibleRepresentations, str)
}

// TestSetNilSetAdd() ensures that Add() returns an error when called on a nil set.
func TestSetNilSetAdd(t *testing.T) {
	var set *Set[int]
	err := set.Add(1)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetRemove() ensures that Remove() returns an error when called on a
// nil set.
func TestSetNilSetRemove(t *testing.T) {
	var set *Set[int]
	err := set.Remove(1)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetContains() ensures that Contains() returns an error when called on
// a nil set.
func TestSetNilSetContains(t *testing.T) {
	var set *Set[int]
	_, err := set.Contains(1)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetSize() ensures that Size() returns an error when called on a nil
// set.
func TestSetNilSetSize(t *testing.T) {
	var set *Set[int]
	_, err := set.Size()
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetValues() ensures that Values() returns an error when called on a
// nil set.
func TestSetNilSetValues(t *testing.T) {
	var set *Set[int]
	_, err := set.Values()
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetClear() ensures that Clear() returns an error when called on a nil
// set.
func TestSetNilSetClear(t *testing.T) {
	var set *Set[int]
	err := set.Clear()
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetUnion() ensures that Union() returns an error when called with nil
// sets.
func TestSetNilSetUnion(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.Union(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetIntersection() ensures that Intersection() returns an error when
// called with nil sets.
func TestSetNilSetIntersection(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.Intersection(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetDifference() ensures that Difference() returns an error when called
// with nil sets.
func TestSetNilSetDifference(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.Difference(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetSymmetricDifference() verifies that calling SymmetricDifference on
// nil sets returns an error.
func TestSetNilSetSymmetricDifference(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.SymmetricDifference(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetEqual() ensures that calling Equal on nil sets returns an error.
func TestSetNilSetEqual(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.Equal(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetSubset() verifies that calling Subset on nil sets returns an error.
func TestSetNilSetSubset(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.Subset(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetNilSetSuperset() checks that calling Superset on nil sets returns an
// error.
func TestSetNilSetSuperset(t *testing.T) {
	var set1, set2 *Set[int]
	_, err := set1.Superset(set2)
	assert.Error(t, err)
	assert.Equal(t, "nil set", err.Error())
}

// TestSetValues() confirms that Values returns all elements in the set correctly.
func TestSetValues(t *testing.T) {
	set := NewSet(1, 2, 3)
	values, err := set.Values()
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3}, values)
}

// TestSetClear() verifies that Clear removes all elements from the set and results
// in an empty set.
func TestSetClear(t *testing.T) {
	set := NewSet(1, 2, 3)
	err := set.Clear()
	assert.NoError(t, err)
	isEmpty, err := set.IsEmpty()
	assert.NoError(t, err)
	assert.True(t, isEmpty)
}

// TestSetIsEmpty() checks that IsEmpty returns true for an empty set and false
// otherwise.
func TestSetIsEmpty(t *testing.T) {
	set := NewSet[int]()
	isEmpty, err := set.IsEmpty()
	assert.NoError(t, err)
	assert.True(t, isEmpty)
	set.Add(1)
	isEmpty, err = set.IsEmpty()
	assert.NoError(t, err)
	assert.False(t, isEmpty)
}

// TestSetUnion() ensures that Union returns a new set containing all unique
// elements from both sets.
func TestSetUnion(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(2, 3)
	union, err := a.Union(b)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3}, getValues(t, union))
}

// TestSetIntersection() checks that Intersection returns a set with only the
// elements common to both sets.
func TestSetIntersection(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(2, 3)
	inter, err := a.Intersection(b)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int{2}, getValues(t, inter))
}

// TestSetDifference() verifies that Difference returns a set with elements present
// in the first set but not the second.
func TestSetDifference(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(2, 4)
	diff, err := a.Difference(b)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int{1, 3}, getValues(t, diff))
}

// TestSetSymmetricDifference() checks that SymmetricDifference returns elements
// present in either set, but not both.
func TestSetSymmetricDifference(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(2, 4)
	symDiff, err := a.SymmetricDifference(b)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int{1, 3, 4}, getValues(t, symDiff))
}

// TestSetEqual() verifies that Equal returns true if two sets contain the same
// elements, regardless of order.
func TestSetEqual(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(2, 1)
	equal, err := a.Equal(b)
	assert.NoError(t, err)
	assert.True(t, equal)
	c := NewSet(1, 3)
	equal, err = a.Equal(c)
	assert.NoError(t, err)
	assert.False(t, equal)
}

// TestSetSubset() confirms that Subset returns true if the first set is a subset
// of the second.
func TestSetSubset(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(1, 2, 3)
	isSubset, err := a.Subset(b)
	assert.NoError(t, err)
	assert.True(t, isSubset)
}

// TestSetSuperset() verifies that Superset returns true if the first set is a
// superset of the second.// TestSetString() ensures that the string representation
// of the set includes all its elements.
func TestSetSuperset(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(1, 2)
	isSuperset, err := a.Superset(b)
	assert.NoError(t, err)
	assert.True(t, isSuperset)
}

// TestSetString() ensures that the string representation of the set includes all
// its elements.
func TestSetString(t *testing.T) {
	set := NewSet(3, 1, 2)
	str := set.String()
	assert.Contains(t, str, "Set: [")
	assert.Contains(t, str, "1")
	assert.Contains(t, str, "2")
	assert.Contains(t, str, "3")
}

// TestSetNilSetOperations() tests that all operations on a nil set return
// appropriate errors.
func TestSetNilSetOperations(t *testing.T) {
	var nilSet *Set[int]
	err := nilSet.Add(1)
	assert.Error(t, err)
	_, err = nilSet.Size()
	assert.Error(t, err)
	_, err = nilSet.Contains(1)
	assert.Error(t, err)
	_, err = nilSet.Values()
	assert.Error(t, err)
	_, err = nilSet.IsEmpty()
	assert.Error(t, err)
	err = nilSet.Clear()
	assert.Error(t, err)
	_, err = nilSet.Union(NewSet[int]())
	assert.Error(t, err)
	_, err = nilSet.Intersection(NewSet[int]())
	assert.Error(t, err)
	_, err = nilSet.Difference(NewSet[int]())
	assert.Error(t, err)
	_, err = nilSet.SymmetricDifference(NewSet[int]())
	assert.Error(t, err)
	_, err = nilSet.Equal(NewSet[int]())
	assert.Error(t, err)
	_, err = nilSet.Subset(NewSet[int]())
	assert.Error(t, err)
	_, err = nilSet.Superset(NewSet[int]())
	assert.Error(t, err)
}

// getValues() is a helper function that extracts and returns the values from a
// set, failing the test if an error occurs.
func getValues[T comparable](t *testing.T, set *Set[T]) []T {
	values, err := set.Values()
	assert.NoError(t, err)
	return values
}
