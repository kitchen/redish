package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeValueStore struct {
	valueStore
}

func TestDefaultFailures(t *testing.T) {
	store := fakeValueStore{}
	_, err := store.get()
	assert.Error(t, err)

	_, err = store.incrby(1)
	assert.Error(t, err)
}
