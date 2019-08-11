package engine

import "testing"
import "github.com/stretchr/testify/assert"

func TestEngineIncr(t *testing.T) {
	engine := NewEngine()
	value, err := engine.Incr("abc")
	assert.Equal(t, "1", value)
	assert.NoError(t, err)

	value, err = engine.Incr("abc")
	assert.Equal(t, "2", value)
	assert.NoError(t, err)

	value, err = engine.Decr("def")
	assert.Equal(t, "-1", value)
	assert.NoError(t, err)

	value, err = engine.Decr("def")
	assert.Equal(t, "-2", value)
	assert.NoError(t, err)
}

func TestGetSetDelExists(t *testing.T) {
	engine := NewEngine()
	err := engine.Set("abc", "foo")
	assert.NoError(t, err)
	value, err := engine.Get("abc")
	assert.Equal(t, "foo", *value)
	assert.NoError(t, err)

	value, err = engine.Get("abc")
	assert.NoError(t, err)
	assert.Equal(t, "foo", *value)

	err = engine.Set("abc", "123")
	assert.NoError(t, err)

	value, err = engine.Get("abc")
	assert.Equal(t, "123", *value)
	assert.NoError(t, err)

	value, err = engine.Get("abc")
	assert.NoError(t, err)
	assert.Equal(t, "123", *value)

	err = engine.Set("key1", "abc")
	assert.NoError(t, err)
	err = engine.Set("key2", "abc")
	assert.NoError(t, err)
	err = engine.Set("key3", "abc")
	assert.NoError(t, err)

	exists, err := engine.Exists([]string{"key1", "key2", "key3"})
	assert.Equal(t, int64(3), exists)
	assert.NoError(t, err)

	exists, err = engine.Exists([]string{"doesnotexist"})
	assert.Equal(t, int64(0), exists)
	assert.NoError(t, err)

	deleted, err := engine.Del([]string{"key1", "key2", "key3", "doesnotexist"})
	assert.NoError(t, err)
	assert.Equal(t, int64(3), deleted)

	deleted, err = engine.Del([]string{"anotherdoesnotexist"})
	assert.NoError(t, err)
	assert.Equal(t, int64(0), deleted)

	value, err = engine.Get("doesnotexist")
	assert.Nil(t, value)
	assert.NoError(t, err)
}

// TODO: move this up to engine.Get / engine.Set since this is no longer a function of the valueStore
// func TestStringToIntAndBackAndForth(t *testing.T) {
// 	intValue := int64(123)
// 	stringValue := "abc"
// 	s := valueStore{intValue: &intValue}
// 	assert.Nil(t, s.stringValue)
// 	assert.Equal(t, int64(123), *s.intValue)
//
// 	s.set(stringValue)
// 	assert.Nil(t, s.intValue)
// 	assert.Equal(t, stringValue, *s.stringValue)
//
// 	s = valueStore{stringValue: &stringValue}
// 	assert.Nil(t, s.intValue)
// 	assert.Equal(t, stringValue, *s.stringValue)
//
// 	s.set(fmt.Sprintf("%d", intValue))
// 	assert.Nil(t, s.stringValue)
// 	assert.Equal(t, intValue, *s.intValue)
// }
