// Package bitmap provides a structure and functions for handling a 32-position
// bitmap represented internally as a 32-bit unsigned integer (uint32). It allows
// you to efficiently turn on, off, toggle, and query the state of individual bits.
//
// This package is useful when you need to represent multiple boolean flags within
// a single variable, such as in permission systems, resource management, data
// compression, ensemble algorithms, and more.
//
// Included features:
//   - Turn individual bits on, off, and toggle.
//   - Check if a bit is on.
//   - Get the binary representation or the total numeric value of the map.
//   - Reset the map to zero.
//
// Attempts to access invalid positions (outside the range 0-31) return an error.
package bitmap

import (
	"errors"
	"fmt"
)

// BitmapSize defines the number of bits in the bitmap. Since we are using a uint32
// to represent the bitmap, there are 32 positions (0 to 31).
const BitmapSize uint8 = 32

// ErrInvalidPosition is returned when a bit operation is attempted on a position
// outside the valid range [0, 31].
var ErrInvalidPosition = errors.New("invalid position")

// BitMap represents a 32-bit bitmap using a uint32 value. It allows operations on
// individual bits such as setting, clearing, and querying their states.
type BitMap struct {
	bits uint32
}

// NewBitMap() creates and returns a new bitmap instance with all bits initially
// set to 0.
//
// Returns:
//   - A pointer to the newly created BitMap.
func NewBitMap() *BitMap {
	return &BitMap{bits: 0b0}
}

// On() sets the bit at the specified position to 1.
//
// Parameters:
//   - pos: The position of the bit to set (must be between 0 and 31).
//
// Returns:
//   - An error if the position is out of range.
func (bm *BitMap) On(pos uint8) error {
	if isOutOfRange(pos) {
		return ErrInvalidPosition
	}
	bm.bits |= 0b1 << pos
	return nil
}

// Off() clears the bit at the specified position (sets it to 0).
//
// Parameters:
//   - pos: The position of the bit to clear (must be between 0 and 31).
//
// Returns:
//   - An error if the position is out of range.
func (bm *BitMap) Off(pos uint8) error {
	if isOutOfRange(pos) {
		return ErrInvalidPosition
	}
	bm.bits &= ^(0b1 << pos)
	return nil
}

// Toggle() flips the bit at the specified position. If the bit is 1, it becomes 0.
// If the bit is 0, it becomes 1.
//
// Parameters:
//   - pos: The position of the bit to toggle (must be between 0 and 31).
//
// Returns:
//   - An error if the position is out of range.
func (bm *BitMap) Toggle(pos uint8) error {
	if isOutOfRange(pos) {
		return ErrInvalidPosition
	}
	bm.bits ^= 1 << pos
	return nil
}

// Reset() clears all bits in the bitmap, setting them to 0.
func (bm *BitMap) Reset() {
	bm.bits = 0
}

// IsOn() checks whether the bit at the specified position is set to 1.
//
// Parameters:
//   - pos: The position of the bit to check (must be between 0 and 31).
//
// Returns:
//   - true if the bit is set to 1.
//   - false if the bit is set to 0.
//   - An error if the position is out of range.
func (bm *BitMap) IsOn(pos uint8) (bool, error) {
	if isOutOfRange(pos) {
		return false, ErrInvalidPosition
	}
	return bm.bits&(1<<pos) != 0b0, nil
}

// GetMap() returns the underlying uint32 value to representing the current state
// of all bits in the bitmap.
//
// Returns:
//   - The 32-bit unsigned integer representing the bitmap.
func (bm *BitMap) GetMap() uint32 {
	return bm.bits
}

// String returns a 32-character binary string representation of the bitmap, padded
// with leading zeros.
//
// Returns:
//   - A string representing the 32 bits of the map as a binary number.
func (bm *BitMap) String() string {
	return fmt.Sprintf("%032b", bm.bits)
}

// isOutOfRange() checks if a given position is outside the valid range of the
// bitmap.
//
// Parameters:
//   - pos: The bit position to validate.
//
// Returns:
//   - true if the position is out of range.
//   - false if the position is valid.
func isOutOfRange(pos uint8) bool {
	return pos >= BitmapSize
}
