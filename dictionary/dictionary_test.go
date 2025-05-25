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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewDictionary() verifies that a newly created dictionary is not nil and its
// initial size is 0.
func TestNewDictionary(t *testing.T) {
	dict := NewDictionary[int, int]()
	assert.NotNil(t, dict)
	assert.Equal(t, 0, dict.Size())
}

// TestDictionaryContains() checks that the dictionary correctly identifies whether
// a key exists in the dictionary.
func TestDictionaryContains(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)
	assert.True(t, dict.Contains("Leo"))
	assert.True(t, dict.Contains("Lucas"))
	assert.False(t, dict.Contains("Fede"))
}

// TestDictionaryPut() verifies that the dictionary correctly adds elements and the
// updated flag is false when adding a new key.
func TestDictionaryPut(t *testing.T) {
	dict := NewDictionary[string, int]()
	updated := dict.Put("Leo", 55)
	assert.False(t, updated)
	updated = dict.Put("Lucas", 38)
	assert.False(t, updated)
	assert.Equal(t, 2, dict.Size())
}

// TestDictionaryPutReplace() checks that when the same key is put again, the
// updated flag is true and the value is replaced.
func TestDictionaryPutReplace(t *testing.T) {
	dict := NewDictionary[string, int]()
	updated := dict.Put("Leo", 55)
	assert.False(t, updated)
	updated = dict.Put("Leo", 38)
	assert.True(t, updated)
	assert.Equal(t, 1, dict.Size())
	value, _ := dict.Get("Leo")
	assert.Equal(t, 38, value)
}

// TestDictionaryGet() checks that the dictionary correctly retrieves a value for
// an existing key, and returns an error for a non-existent key.
func TestDictionaryGet(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Lucas", 35)
	value, err := dict.Get("Lucas")
	assert.Equal(t, 35, value)
	require.NoError(t, err)
	value, err = dict.Get("Fede")
	assert.Equal(t, 0, value)
	assert.EqualError(t, err, "non-existent key")
}

// TestDictionaryRemove() verifies that a key can be removed from the dictionary
// and that the size is updated accordingly.
func TestDictionaryRemove(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)
	assert.Equal(t, 2, dict.Size())
	removed := dict.Remove("Leo")
	assert.True(t, removed)
	assert.Equal(t, 1, dict.Size())
	assert.True(t, dict.Contains("Lucas"))
	assert.False(t, dict.Contains("Leo"))
}

// TestDictionaryRemoveNotExists() checks that trying to remove a non-existent key
// returns false and does not alter the dictionary size.
func TestDictionaryRemoveNotExists(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)
	assert.Equal(t, 2, dict.Size())
	removed := dict.Remove("Fede")
	assert.False(t, removed)
	assert.Equal(t, 2, dict.Size())
}

// TestDictionarySize() verifies that the size of the dictionary is correctly
// updated after inserting elements.
func TestDictionarySize(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)
	assert.Equal(t, 2, dict.Size())
}

// TestDictionaryValues() checks that the dictionary's values can be correctly
// retrieved as a slice.
func TestDictionaryValues(t *testing.T) {
	dic := NewDictionary[int, int]()
	dic.Put(1, 2)
	dic.Put(3, 4)
	dic.Put(5, 6)
	assert.ElementsMatch(t, []int{6, 4, 2}, dic.Values())
}

// TestDictionaryKeys() verifies that the dictionary's keys can be correctly
// retrieved as a slice.
func TestDictionaryKeys(t *testing.T) {
	dict := NewDictionary[int, int]()
	dict.Put(1, 2)
	dict.Put(3, 4)
	dict.Put(5, 6)
	assert.ElementsMatch(t, []int{1, 5, 3}, dict.Keys())
}

// TestDictionaryStringOnEmptyDictionary() checks the string representation of an
// empty dictionary.
func TestDictionaryStringOnEmptyDictionary(t *testing.T) {
	dict := NewDictionary[int, int]()
	assert.Equal(t, "Dictionary: {}", dict.String())
}

// TestDictionaryString() verifies that the string representation of the dictionary
// reflects its content, and checks for multiple possible formats.
func TestDictionaryString(t *testing.T) {
	dict := NewDictionary[int, int]()
	dict.Put(1, 2)
	dict.Put(3, 4)
	possibleRepresentations := []string{"Dictionary: {\n  1: 2\n  3: 4\n}", "Dictionary: {\n  3: 4\n  1: 2\n}"}
	assert.Contains(t, possibleRepresentations, dict.String())
}

// TestDictionaryClear() checks that the Clear() function correctly empties the
// dictionary.
func TestDictionaryClear(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)
	assert.Equal(t, 2, dict.Size())
	dict.Clear()
	assert.Equal(t, 0, dict.Size())
	assert.False(t, dict.Contains("Leo"))
	assert.False(t, dict.Contains("Lucas"))
	updated := dict.Put("Fede", 100)
	assert.False(t, updated)
	assert.Equal(t, 1, dict.Size())
	value, err := dict.Get("Fede")
	assert.NoError(t, err)
	assert.Equal(t, 100, value)
}

// TestDictionaryWithFloat64KeysAndStringValues() verifies that the dictionary can
// handle float64 keys and string values correctly.
func TestDictionaryWithFloat64KeysAndStringValues(t *testing.T) {
	dict := NewDictionary[float64, string]()
	dict.Put(1.23, "One Point Two Three")
	dict.Put(4.56, "Four Point Five Six")
	assert.True(t, dict.Contains(1.23))
	assert.True(t, dict.Contains(4.56))
	value, err := dict.Get(1.23)
	assert.NoError(t, err)
	assert.Equal(t, "One Point Two Three", value)
	value, err = dict.Get(4.56)
	assert.NoError(t, err)
	assert.Equal(t, "Four Point Five Six", value)
}

type Person struct {
	Name string
	Age  int
}

// TestDictionaryWithStructValues() verifies that the dictionary can store and
// retrieve structs as values.
func TestDictionaryWithStructValues(t *testing.T) {
	dict := NewDictionary[string, Person]()
	dict.Put("Alice", Person{Name: "Alice", Age: 30})
	dict.Put("Bob", Person{Name: "Bob", Age: 25})
	person, err := dict.Get("Alice")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", person.Name)
	assert.Equal(t, 30, person.Age)
	person, err = dict.Get("Bob")
	assert.NoError(t, err)
	assert.Equal(t, "Bob", person.Name)
	assert.Equal(t, 25, person.Age)
}

// TestDictionaryGetWithEmptyDictionary() checks that attempting to retrieve a
// value from an empty dictionary returns an error with the appropriate message.
func TestDictionaryGetWithEmptyDictionary(t *testing.T) {
	dict := NewDictionary[string, int]()
	value, err := dict.Get("NonExistentKey")
	assert.Equal(t, 0, value)
	assert.EqualError(t, err, "non-existent key")
}

// TestDictionaryRemoveFromEmptyDictionary() ensures that attempting to remove a
// non-existent key from an empty dictionary returns false.
func TestDictionaryRemoveFromEmptyDictionary(t *testing.T) {
	dict := NewDictionary[string, int]()
	removed := dict.Remove("NonExistentKey")
	assert.False(t, removed)
	assert.Equal(t, 0, dict.Size())
}

// TestDictionaryPerformance() tests the performance of the dictionary when
// handling a large number of elements (1,000,000 entries).
func TestDictionaryPerformance(t *testing.T) {
	dict := NewDictionary[int, string]()
	for i := range 1000000 {
		dict.Put(i, fmt.Sprintf("Value %d", i))
	}
	assert.Equal(t, 1000000, dict.Size())
	value, err := dict.Get(999999)
	assert.NoError(t, err)
	assert.Equal(t, "Value 999999", value)
}
