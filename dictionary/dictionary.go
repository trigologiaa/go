// Package dictionary provides a generic dictionary (or map) data structure
// implemented using Go generics. It allows storing key-value pairs, where keys are
// comparable and values can be of any type.
//
// This package is useful for operations that require a key-value mapping, such as
// searching for associated values, adding or removing entries, and querying the
// size or contents of a collection.
//
// Included features:
//   - Create a new dictionary.
//   - Add or update key-value pairs.
//   - Check if a key exists in the dictionary.
//   - Retrieve values associated with keys.
//   - Remove key-value pairs from the dictionary.
//   - Get the number of key-value pairs in the dictionary.
//   - Retrieve all keys or values as slices.
//   - Clear all key-value pairs in the dictionary.
//   - Get a string representation of the dictionary contents.
//
// Most methods return an error if the dictionary receiver is nil.
package dictionary

import (
	"errors"
	"fmt"
)

// Dictionary[K comparable, V any] represents a generic dictionary structure that
// stores key-value pairs where keys are comparable and values can be any type.
type Dictionary[K comparable, V any] struct {
	dict map[K]V
}

// NewDictionary[K comparable, V any]() creates and returns a new empty dictionary.
//
// Returns:
//   - A pointer to the newly created Dictionary.
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{dict: make(map[K]V)}
}

// Put() inserts or updates the value associated with the specified key.
//
// Parameters:
//   - key: The key to add or update.
//   - value: The value to associate with the key.
//
// Returns:
//   - true if the key was already present and its value was updated.
//   - false if it was a new insertion.
func (d *Dictionary[K, V]) Put(key K, value V) bool {
	_, exists := d.dict[key]
	d.dict[key] = value
	return exists
}

// Contains() checks whether the dictionary contains the specified key.
//
// Parameters:
//   - key: The key to check for existence.
//
// Returns:
//   - true if the key exists in the dictionary.
//   - false if the key does not exist in the dictionary.
func (d *Dictionary[K, V]) Contains(key K) bool {
	_, exists := d.dict[key]
	return exists
}

// Get() retrieves the value associated with the specified key.
//
// Parameters:
//   - key: The key whose value is to be retrieved.
//
// Returns:
//   - The value associated with the key if it exists.
//   - An error if the key does not exist.
func (d *Dictionary[K, V]) Get(key K) (V, error) {
	value, exists := d.dict[key]
	if !exists {
		return value, errors.New("non-existent key")
	}
	return value, nil
}

// Remove() deletes the entry associated with the specified key.
//
// Parameters:
//   - key: The key to remove from the dictionary.
//
// Returns:
//   - true if the key was found and removed.
//   - false if the key did not exist.
func (d *Dictionary[K, V]) Remove(key K) bool {
	_, exists := d.dict[key]
	if exists {
		delete(d.dict, key)
	}
	return exists
}

// Size() returns the number of entries stored in the dictionary.
//
// Returns:
//   - The total count of key-value pairs.
func (d *Dictionary[K, V]) Size() int {
	return len(d.dict)
}

// Keys() returns a slice containing all keys currently stored in the dictionary.
//
// Returns:
//   - A slice of keys.
func (d *Dictionary[K, V]) Keys() []K {
	keys := make([]K, 0, d.Size())
	for key := range d.dict {
		keys = append(keys, key)
	}
	return keys
}

// Values() returns a slice containing all values currently stored in the
// dictionary.
//
// Returns:
//   - A slice of values.
func (d *Dictionary[K, V]) Values() []V {
	values := make([]V, 0, d.Size())
	for _, value := range d.dict {
		values = append(values, value)
	}
	return values
}

// String() returns a string representation of the dictionary's contents.
//
// Returns:
//   - A formatted string listing all key-value pairs, or an empty dictionary
//     message.
func (d *Dictionary[K, V]) String() string {
	if d.Size() == 0 {
		return "Dictionary: {}"
	}
	result := "Dictionary: {\n"
	for key, value := range d.dict {
		result += fmt.Sprintf("  %v: %v\n", key, value)
	}
	result += "}"
	return result
}

// Clear() removes all entries from the dictionary, resetting it to an empty state.
func (d *Dictionary[K, V]) Clear() {
	d.dict = make(map[K]V)
}
