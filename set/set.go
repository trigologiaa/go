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
	"errors"
	"fmt"
	"sort"
)

// Set[T comparable] represents a generic set structure that stores unique
// elements, where each element is comparable.
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet[T comparable]() creates and returns a new set containing the specified
// elements.
//
// Parameters:
//   - elements: A variadic list of elements to be added to the set.
//
// Returns:
//   - A pointer to the newly created Set containing the specified elements.
func NewSet[T comparable](elements ...T) *Set[T] {
	s := &Set[T]{elements: make(map[T]struct{})}
	s.Add(elements...)
	return s
}

// Contains() Checks whether the set contains the specified element.
//
// Parameters:
//   - element: The element to check for existence.
//
// Returns:
//   - true if the element exists in the set.
//   - false if the element does not exist in the set.
//   - An error if the set is nil.
func (s *Set[T]) Contains(element T) (bool, error) {
	if s == nil {
		return false, errors.New("nil set")
	}
	_, exists := s.elements[element]
	return exists, nil
}

// Add() adds the specified elements to the set.
//
// Parameters:
//   - elements: A variadic list of elements to be added.
//
// Returns:
//   - An error if the set is nil.
func (s *Set[T]) Add(elements ...T) error {
	if s == nil {
		return errors.New("nil set")
	}
	for _, element := range elements {
		s.elements[element] = struct{}{}
	}
	return nil
}

// Remove() removes the specified element from the set.
//
// Parameters:
//   - element: The element to remove.
//
// Returns:
//   - An error if the set is nil.
func (s *Set[T]) Remove(element T) error {
	if s == nil {
		return errors.New("nil set")
	}
	delete(s.elements, element)
	return nil
}

// Size() returns the number of elements in the set.
//
// Returns:
//   - The number of elements in the set.
//   - An error if the set is nil.
func (s *Set[T]) Size() (int, error) {
	if s == nil {
		return 0, errors.New("nil set")
	}
	return len(s.elements), nil
}

// Values() returns a slice containing all the elements in the set.
//
// Returns:
//   - A slice of elements in the set.
//   - An error if the set is nil.
func (s *Set[T]) Values() ([]T, error) {
	if s == nil {
		return nil, errors.New("nil set")
	}
	size, _ := s.Size()
	values := make([]T, 0, size)
	for k := range s.elements {
		values = append(values, k)
	}
	return values, nil
}

// Clear() removes all elements from the set, resetting it to an empty state.
//
// Returns:
//   - An error if the set is nil.
func (s *Set[T]) Clear() error {
	if s == nil {
		return errors.New("nil set")
	}
	s.elements = make(map[T]struct{})
	return nil
}

// IsEmpty() checks whether the set is empty.
//
// Returns:
//   - true if the set is empty.
//   - false if the set contains at least one element.
//   - An error if the set is nil.
func (s *Set[T]) IsEmpty() (bool, error) {
	if s == nil {
		return false, errors.New("nil set")
	}
	size, _ := s.Size()
	return size == 0, nil
}

// Union() returns a new set that contains all elements from both the current set
// and the specified set (union).
//
// Parameters:
//   - other: The set to compute the union with.
//
// Returns:
//   - A new set containing the union of the two sets.
//   - An error if either set is nil.
func (s *Set[T]) Union(other *Set[T]) (*Set[T], error) {
	if s == nil || other == nil {
		return nil, errors.New("nil set")
	}
	result := NewSet[T]()
	for k := range s.elements {
		result.Add(k)
	}
	for k := range other.elements {
		result.Add(k)
	}
	return result, nil
}

// Intersection() returns a new set containing only elements that are present in
// both the current set and the specified set (intersection).
//
// Parameters:
//   - other: The set to compute the intersection with.
//
// Returns:
//   - A new set containing the intersection of the two sets.
//   - An error if either set is nil.
func (s *Set[T]) Intersection(other *Set[T]) (*Set[T], error) {
	if s == nil || other == nil {
		return nil, errors.New("nil set")
	}
	result := NewSet[T]()
	for k := range s.elements {
		exists, _ := other.Contains(k)
		if exists {
			result.Add(k)
		}
	}
	return result, nil
}

// Difference() returns a new set containing only elements that are in the current
// set but not in the specified set (difference).
//
// Parameters:
//   - other: The set to compute the difference with.
//
// Returns:
//   - A new set containing the difference of the two sets.
//   - An error if either set is nil.
func (s *Set[T]) Difference(other *Set[T]) (*Set[T], error) {
	if s == nil || other == nil {
		return nil, errors.New("nil set")
	}
	result := NewSet[T]()
	for k := range s.elements {
		exists, _ := other.Contains(k)
		if !exists {
			result.Add(k)
		}
	}
	return result, nil
}

// SymmetricDifference() returns a new set containing the elements that are in
// either of the two sets, but not in both (symmetric difference).
//
// Parameters:
//   - other: The set to compute the symmetric difference with.
//
// Returns:
//   - A new set containing the symmetric difference of the two sets.
//   - An error if either set is nil.
func (s *Set[T]) SymmetricDifference(other *Set[T]) (*Set[T], error) {
	if s == nil || other == nil {
		return nil, errors.New("nil set")
	}
	result := NewSet[T]()
	for k := range s.elements {
		exists, _ := other.Contains(k)
		if !exists {
			result.Add(k)
		}
	}
	for k := range other.elements {
		exists, _ := s.Contains(k)
		if !exists {
			result.Add(k)
		}
	}
	return result, nil
}

// Equal() checks whether the current set is equal to the specified set.
//
// Parameters:
//   - other: The set to check equality with.
//
// Returns:
//   - true if the sets are equal.
//   - false if the sets are not equal.
//   - An error if either set is nil.
func (s *Set[T]) Equal(other *Set[T]) (bool, error) {
	if s == nil || other == nil {
		return false, errors.New("nil set")
	}
	s1, _ := s.Size()
	s2, _ := other.Size()
	if s1 != s2 {
		return false, nil
	}
	for k := range s.elements {
		exists, _ := other.Contains(k)
		if !exists {
			return false, nil
		}
	}
	return true, nil
}

// Subset() checks whether the current set is a subset of the specified set.
//
// Parameters:
//   - other: The set to check if the current set is a subset of.
//
// Returns:
//   - true if the current set is a subset of the other set.
//   - false if the current set is not a subset of the other set.
//   - An error if either set is nil.
func (s *Set[T]) Subset(other *Set[T]) (bool, error) {
	if s == nil || other == nil {
		return false, errors.New("nil set")
	}
	for k := range s.elements {
		exists, _ := other.Contains(k)
		if !exists {
			return false, nil
		}
	}
	return true, nil
}

// Superset() checks whether the current set is a superset of the specified set.
//
// Parameters:
//   - other: The set to check if the current set is a superset of.
//
// Returns:
//   - true if the current set is a superset of the other set.
//   - false if the current set is not a superset of the other set.
//   - An error if either set is nil.
func (s *Set[T]) Superset(other *Set[T]) (bool, error) {
	if s == nil || other == nil {
		return false, errors.New("nil set")
	}
	subset, _ := other.Subset(s)
	return subset, nil
}

// String() Returns a string representation of the set's contents.
//
// Returns:
//   - A formatted string listing all elements in the set.
func (s *Set[T]) String() string {
	values, _ := s.Values()
	sort.Slice(values, func(i, j int) bool { return fmt.Sprintf("%v", values[i]) < fmt.Sprintf("%v", values[j]) })
	return fmt.Sprintf("Set: %v", values)
}
