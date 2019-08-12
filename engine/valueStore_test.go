package engine

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeValueStore struct {
	valueStore
}

func (s *fakeValueStore) getType() string {
	return "fake"
}

func TestDefaultFailures(t *testing.T) {
	store := fakeValueStore{}
	_, err := store.get()
	assert.Error(t, err)

	_, err = store.incrby(1)
	assert.Error(t, err)

	assert.Equal(t, "fake", store.getType())
}

func TestExpirationStorage(t *testing.T) {
	store := fakeValueStore{}
	assert.Nil(t, store.expires())
	assert.False(t, store.expired())

	duration, err := time.ParseDuration("10m")
	assert.NoError(t, err)

	expires := time.Now().Add(duration)
	store.expire(&expires)
	assert.Equal(t, expires, *store.expires())
	assert.False(t, store.expired())

	duration, err = time.ParseDuration("-10m")
	assert.NoError(t, err)

	expires = time.Now().Add(duration)
	store.expire(&expires)
	assert.Equal(t, expires, *store.expires())
	assert.True(t, store.expired())
}
