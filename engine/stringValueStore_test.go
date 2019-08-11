package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValue(t *testing.T) {
	stringValue := "abc"
	s := stringValueStore{stringValue: stringValue}
	assert.Equal(t, "abc", s.stringValue, "the storage contains the correct value")

	storedValue, err := s.get()
	assert.Equal(t, "abc", storedValue, "we get the correct value from the value store for a string")
	assert.NoError(t, err, "getting the value didn't error")
}

func TestIncrByStringValue(t *testing.T) {
	s := stringValueStore{stringValue: "abc"}
	assert.Equal(t, "abc", s.stringValue)

	_, err := s.incrby(1)
	assert.Error(t, err)
}
