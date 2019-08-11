package engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValue(t *testing.T) {
	stringValue := "abc"
	s := valueStore{stringValue: &stringValue}
	assert.Equal(t, "abc", *s.stringValue, "the storage contains the correct value")

	storedValue, err := s.get()
	assert.Equal(t, "abc", *storedValue, "we get the correct value from the value store for a string")
	assert.NoError(t, err, "getting the value didn't error")

	newStringValue := "def"
	assert.NoError(t, s.set(newStringValue), "setting a new string value doesn't error")
	assert.Equal(t, "def", *s.stringValue)

	storedValue, err = s.get()
	assert.Equal(t, "def", *storedValue, "the correct value was set")
	assert.NoError(t, err, "getting the value didn't error")
}

func TestIntValue(t *testing.T) {
	intValue := int64(123)
	s := valueStore{intValue: &intValue}
	assert.Equal(t, int64(123), *s.intValue, "the storage contains the correct value")

	storedValue, err := s.get()
	assert.Equal(t, "123", *storedValue, "get returns the correct value as a string")
	assert.NoError(t, err, "there was no error getting the value")

	newIntValue := "456"
	assert.NoError(t, s.set(newIntValue), "setting a new int value doesn't error")
	assert.Equal(t, int64(456), *s.intValue)
	assert.Nil(t, s.stringValue)

	storedValue, err = s.get()
	assert.Equal(t, "456", *storedValue)
	assert.NoError(t, err, "there was no error getting the value")
}

func TestStringToIntAndBackAndForth(t *testing.T) {
	intValue := int64(123)
	stringValue := "abc"
	s := valueStore{intValue: &intValue}
	assert.Nil(t, s.stringValue)
	assert.Equal(t, int64(123), *s.intValue)

	s.set(stringValue)
	assert.Nil(t, s.intValue)
	assert.Equal(t, stringValue, *s.stringValue)

	s = valueStore{stringValue: &stringValue}
	assert.Nil(t, s.intValue)
	assert.Equal(t, stringValue, *s.stringValue)

	s.set(fmt.Sprintf("%d", intValue))
	assert.Nil(t, s.stringValue)
	assert.Equal(t, intValue, *s.intValue)
}

func TestIncrDecrIntValues(t *testing.T) {
	intValue := int64(123)
	s := valueStore{intValue: &intValue}

	assert.Equal(t, int64(123), *s.intValue)

	newValue, err := s.incr()
	assert.Equal(t, int64(124), *newValue)
	assert.NoError(t, err)

	newValue, err = s.decr()
	assert.Equal(t, int64(123), *newValue)
	assert.NoError(t, err)

	intValue = int64(0)
	s = valueStore{intValue: &intValue}
	assert.Equal(t, int64(0), *s.intValue)

	newValue, err = s.incr()
	assert.Equal(t, int64(1), *newValue)
	assert.NoError(t, err)
}

func TestIncrDecrStringValue(t *testing.T) {
	stringValue := "abc"
	s := valueStore{stringValue: &stringValue}

	assert.Equal(t, "abc", *s.stringValue)

	_, err := s.incr()
	assert.Error(t, err)

	_, err = s.decr()
	assert.Error(t, err)
}
