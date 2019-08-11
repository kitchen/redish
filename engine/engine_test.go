package engine

import "testing"
import "github.com/stretchr/testify/assert"

func TestEngineIncr(t *testing.T) {
	engine := NewEngine()
	value, err := engine.Incr("abc")
	assert.Equal(t, int64(1), *value)
	assert.NoError(t, err)

	value, err = engine.Incr("abc")
	assert.Equal(t, int64(2), *value)
	assert.NoError(t, err)

	value, err = engine.Decr("def")
	assert.Equal(t, int64(-1), *value)
	assert.NoError(t, err)

	value, err = engine.Decr("def")
	assert.Equal(t, int64(-2), *value)
	assert.NoError(t, err)
}

func TestGetSetDel(t *testing.T) {
	engine := NewEngine()
	err := engine.Set("abc", "foo")
	assert.NoError(t, err)
	s, ok := engine.storage["abc"]
	assert.NotNil(t, s)
	assert.True(t, ok)
	assert.Equal(t, "foo", *s.stringValue)
	assert.Nil(t, s.intValue)

	value, err := engine.Get("abc")
	assert.NoError(t, err)
	assert.Equal(t, "foo", *value)

	err = engine.Set("abc", "123")
	assert.NoError(t, err)
	s, ok = engine.storage["abc"]
	assert.NotNil(t, s)
	assert.True(t, ok)
	assert.Equal(t, int64(123), *s.intValue)
	assert.Nil(t, s.stringValue)

	value, err = engine.Get("abc")
	assert.NoError(t, err)
	assert.Equal(t, "123", *value)

	err = engine.Set("key1", "abc")
	assert.NoError(t, err)
	err = engine.Set("key2", "abc")
	assert.NoError(t, err)
	err = engine.Set("key3", "abc")
	assert.NoError(t, err)

	_, ok = engine.storage["key1"]
	assert.True(t, ok)
	_, ok = engine.storage["key2"]
	assert.True(t, ok)
	_, ok = engine.storage["key3"]
	assert.True(t, ok)
	_, ok = engine.storage["doesnotexist"]
	assert.False(t, ok)

	num, err := engine.Del([]string{"key1", "key2", "key3", "doesnotexist"})
	assert.NoError(t, err)
	assert.Equal(t, int64(3), *num)

	num, err = engine.Del([]string{"anotherdoesnotexist"})
	assert.NoError(t, err)
	assert.Equal(t, int64(0), *num)

	value, err = engine.Get("doesnotexist")
	assert.Nil(t, value)
	assert.NoError(t, err)
}
