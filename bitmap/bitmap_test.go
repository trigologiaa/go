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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBitMapNewAllInZero() verifies that a newly created bitmap has all bits set
// to 0.
func TestBitMapNewAllInZero(t *testing.T) {
	m := NewBitMap()
	assert.Equal(t, uint32(0), m.GetMap())
}

// TestBitMapTurnOnABit() checks that turning on a bit sets it to 1.
func TestBitMapTurnOnABit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

// TestBitMapTurnOnABitInAnInvalidPosition() ensures an error is returned when
// trying to turn on a bit out of range.
func TestBitMapTurnOnABitInAnInvalidPosition(t *testing.T) {
	m := NewBitMap()
	err := m.On(32)
	assert.EqualError(t, err, "invalid position")
}

// TestBitMapTurnOffABit() verifies that turning off a previously set bit sets it
// to 0.
func TestBitMapTurnOffABit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.Off(1)
	isOn, _ := m.IsOn(1)
	assert.False(t, isOn)
}

// TestBitMapTurnOffABitInAnInvalidPosition() ensures error is returned when
// trying to turn off a bit out of range.
func TestBitMapTurnOffABitInAnInvalidPosition(t *testing.T) {
	m := NewBitMap()
	err := m.Off(32)
	assert.EqualError(t, err, "invalid position")
}

// TestBitMapTurnOnAllBits() checks that turning on all 32 bits results in a full
// bitmap (0xffffffff).
func TestBitMapTurnOnAllBits(t *testing.T) {
	m := NewBitMap()
	for i := uint8(0); i < 32; i++ {
		m.On(i)
	}
	var max uint32 = 0xffffffff
	assert.Equal(t, max, m.GetMap())
}

// TestBitMatTurnOffAllBits() checks that all bits can be turned off after being
// turned on.
func TestBitMatTurnOffAllBits(t *testing.T) {
	m := NewBitMap()
	for i := uint8(0); i < 32; i++ {
		m.On(i)
	}
	for i := uint8(0); i < 32; i++ {
		m.Off(i)
	}
	assert.Equal(t, uint32(0), m.GetMap())
}

// TestBitMapOneBitStatus() verifies that IsOn correctly reports the status of a
// bit.
func TestBitMapOneBitStatus(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

// TestBitMapStateOfABitInAnInvalidPosition() ensures an error is returned when
// checking a bit is out of range.
func TestBitMapStateOfABitInAnInvalidPosition(t *testing.T) {
	m := NewBitMap()
	_, err := m.IsOn(32)
	assert.EqualError(t, err, "invalid position")
}

// TestBitMapTurningOnTheSameBitSeveralTimesDoesNotTurnItOff() ensures that turning
// on the same bit multiple times does not toggle it off.
func TestBitMapTurningOnTheSameBitSeveralTimesDoesNotTurnItOff(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

// TestBitMapGetMap() verifies the correct internal uint32 value after setting a
// bit.
func TestBitMapGetMap(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	assert.Equal(t, uint32(2), m.GetMap())
}

// TestBitMapToggleBit() verifies that toggling a bit changes its state and
// toggling it again restores the original state.
func TestBitMapToggleBit(t *testing.T) {
	m := NewBitMap()
	err := m.Toggle(3)
	assert.NoError(t, err)
	isOn, _ := m.IsOn(3)
	assert.True(t, isOn)
	err = m.Toggle(3)
	assert.NoError(t, err)
	isOn, _ = m.IsOn(3)
	assert.False(t, isOn)
}

// TestBitMapToggleInvalidPosition() ensures an error is returned when toggling a
// bit out of range.
func TestBitMapToggleInvalidPosition(t *testing.T) {
	m := NewBitMap()
	err := m.Toggle(32)
	assert.EqualError(t, err, "invalid position")
}

// TestBitMapReset() verifies that the Reset() method clears all bits.
func TestBitMapReset(t *testing.T) {
	m := NewBitMap()
	m.On(0)
	m.On(15)
	m.On(31)
	m.Reset()
	assert.Equal(t, uint32(0), m.GetMap())
}

// TestBitMapStringRepresentation() checks the 32-character binary string
// representation of the bitmap.
func TestBitMapStringRepresentation(t *testing.T) {
	m := NewBitMap()
	m.On(0)
	m.On(31)
	expected := "10000000000000000000000000000001"
	assert.Equal(t, expected, m.String())
}
