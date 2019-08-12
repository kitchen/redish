package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntValue(t *testing.T) {
	intValue := int64(123)
	s := intValueStore{intValue: intValue}
	assert.Equal(t, int64(123), s.intValue, "the storage contains the correct value")

	storedValue, err := s.get()
	assert.Equal(t, "123", storedValue, "get returns the correct value as a string")
	assert.NoError(t, err, "there was no error getting the value")

	assert.Equal(t, "string", s.getType())
}

func TestIncrBy(t *testing.T) {
	intValue := int64(123)
	s := intValueStore{intValue: intValue}

	assert.Equal(t, int64(123), s.intValue)

	newValue, err := s.incrby(1)
	assert.Equal(t, "124", newValue)
	assert.NoError(t, err)

	newValue, err = s.incrby(-122)
	assert.Equal(t, "2", newValue)
	assert.NoError(t, err)
}
