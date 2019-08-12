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

	value, err = engine.Incrby("foo", "10")
	assert.Equal(t, "10", value)
	assert.NoError(t, err)

	value, err = engine.Incrby("foo", "10")
	assert.Equal(t, "20", value)
	assert.NoError(t, err)

	value, err = engine.Decrby("bar", "10")
	assert.Equal(t, "-10", value)
	assert.NoError(t, err)

	value, err = engine.Decrby("bar", "10")
	assert.Equal(t, "-20", value)
	assert.NoError(t, err)

	_, err = engine.Incrby("aoeuhtns", "aoeuhtns")
	assert.Error(t, err)

	_, err = engine.Decrby("aoeuhtns", "aoeuhtns")
	assert.Error(t, err)

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

	value, err = engine.GetSet("getset", "abc")
	assert.Nil(t, value)
	assert.NoError(t, err)

	value, err = engine.GetSet("getset", "def")
	assert.Equal(t, "abc", *value)
	assert.NoError(t, err)

	value, err = engine.Get("getset")
	assert.Equal(t, "def", *value)
	assert.NoError(t, err)
}

func TestStrLen(t *testing.T) {
	engine := NewEngine()

	engine.Set("stringvalue", "aoeuhtns")
	value, err := engine.Strlen("stringvalue")
	assert.NoError(t, err)
	assert.Equal(t, "8", value)

	engine.Set("intvalue", "1234567890")
	value, err = engine.Strlen("intvalue")
	assert.NoError(t, err)
	assert.Equal(t, "10", value)

	value, err = engine.Strlen("doesnotexist")
	assert.Equal(t, "0", value)
	assert.NoError(t, err)

	// TODO: test what happens when strlen on an invalid data type (e.g. fakeValueStore)
}

func TestMGet(t *testing.T) {
	engine := NewEngine()

	stringValue := "aoeuhtns"
	intValue := "123"
	engine.Set("stringvalue", stringValue)
	engine.Set("intvalue", intValue)

	values, err := engine.MGet([]string{"stringvalue", "intvalue", "doesnotexist"})
	assert.NoError(t, err)
	assert.Equal(t, []*string{&stringValue, &intValue, nil}, values)

	values, err = engine.MGet([]string{"doesnotexist", "doesnotexist2", "doesnotexist3"})
	assert.Equal(t, []*string{nil, nil, nil}, values)
	assert.NoError(t, err)

	// TODO: test what happens when mget with an invalid data type (e.g. fakeValueStore)
}

func TestMSet(t *testing.T) {
	engine := NewEngine()

	kvs := map[string]string{
		"foo": "aoeuhtns",
		"bar": "123",
	}

	err := engine.MSet(kvs)
	assert.NoError(t, err)

	value, err := engine.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, "aoeuhtns", *value)

	value, err = engine.Get("bar")
	assert.NoError(t, err)
	assert.Equal(t, "123", *value)
}

func TestType(t *testing.T) {
	engine := NewEngine()
	engine.Set("stringvalue", "abc")
	assert.Equal(t, "string", engine.Type("stringvalue"))

	engine.Set("intvalue", "123")
	assert.Equal(t, "string", engine.Type("intvalue"))

	assert.Equal(t, "none", engine.Type("doesnotexist"))
}
