package engine

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const stringValue = "foobartestvaluehtnsaoeu"
const stringValueKey = "stringvaluekey"
const intValue = 12345678901
const intValueString = "12345678901"
const intValueKey = "intvaluekey"
const fakeValueKey = "fakevaluekey"
const doesNotExistKey = "doesnotexistkey"

func (engine *engine) setFakeValue(key string) {
	engine.storage[key] = &fakeValueStore{}
}

type engineTestSuite struct {
	suite.Suite
	engine engine
}

func (suite *engineTestSuite) SetupTest() {
	suite.engine = *newEngine()
	// TODO: turn all of these strings into constants
	suite.engine.Set(intValueKey, intValueString)
	suite.engine.Set(stringValueKey, stringValue)
	suite.engine.setFakeValue(fakeValueKey)
}

func (suite *engineTestSuite) TestNewEngine() {
	var interfaceEngine Engine
	interfaceEngine = NewEngine()
	suite.NotNil(interfaceEngine)

	var rawEngine engine
	rawEngine = *newEngine()
	suite.NotNil(rawEngine)
}

func (suite *engineTestSuite) TestGetSetStringValues() {
	value, err := suite.engine.Get(stringValueKey)
	suite.NoError(err)
	suite.Equal(stringValue, *value)

	err = suite.engine.Set("key", "abc")
	suite.NoError(err)

	value, err = suite.engine.Get("key")
	suite.NoError(err)
	suite.Equal("abc", *value)
}

func (suite *engineTestSuite) TestGetSetIntValues() {
	value, err := suite.engine.Get(intValueKey)
	suite.NoError(err)
	suite.Equal(intValueString, *value)

	err = suite.engine.Set("key", "123")
	suite.NoError(err)

	value, err = suite.engine.Get("key")
	suite.NoError(err)
	suite.Equal("123", *value)
}

func (suite *engineTestSuite) TestGetOtherValues() {
	value, err := suite.engine.Get(doesNotExistKey)
	suite.NoError(err)
	suite.Nil(value)

	_, err = suite.engine.Get(fakeValueKey)
	suite.Error(err)
}

func (suite *engineTestSuite) TestGetSetMethod() {
	value, err := suite.engine.GetSet(intValueKey, "9876")
	suite.NoError(err)
	suite.Equal(intValueString, *value)

	value, err = suite.engine.Get(intValueKey)
	suite.NoError(err)
	suite.Equal("9876", *value)

	value, err = suite.engine.GetSet(stringValueKey, "abc123")
	suite.NoError(err)
	suite.Equal(stringValue, *value)

	value, err = suite.engine.Get(stringValueKey)
	suite.NoError(err)
	suite.Equal("abc123", *value)

	value, err = suite.engine.GetSet(doesNotExistKey, "it does now")
	suite.NoError(err)
	suite.Nil(value)

	value, err = suite.engine.Get(doesNotExistKey)
	suite.NoError(err)
	suite.Equal("it does now", *value)

	_, err = suite.engine.GetSet(fakeValueKey, "fake fake fake")
	suite.Error(err)

	_, err = suite.engine.Get(fakeValueKey)
	suite.Error(err)

	typeString, err := suite.engine.Type(fakeValueKey)
	suite.NoError(err)
	suite.Equal("fake", typeString)
}

func (suite *engineTestSuite) TestDel() {
	deleted, err := suite.engine.Del([]string{stringValueKey, intValueKey, doesNotExistKey})
	suite.NoError(err)
	suite.Equal(int64(2), deleted)

	value, err := suite.engine.Get(stringValueKey)
	suite.NoError(err)
	suite.Nil(value)

	value, err = suite.engine.Get(intValueKey)
	suite.NoError(err)
	suite.Nil(value)

	value, err = suite.engine.Get(doesNotExistKey)
	suite.NoError(err)
	suite.Nil(value)
}

func (suite *engineTestSuite) TestExists() {
	exists, err := suite.engine.Exists([]string{stringValueKey, intValueKey, doesNotExistKey})
	suite.NoError(err)
	suite.Equal(int64(2), exists)
}

func (suite *engineTestSuite) TestStrLen() {
	length, err := suite.engine.Strlen(stringValueKey)
	suite.NoError(err)
	suite.Equal(int64(len(stringValue)), length)

	length, err = suite.engine.Strlen(intValueKey)
	suite.NoError(err)
	suite.Equal(int64(len(intValueString)), length)

	length, err = suite.engine.Strlen(doesNotExistKey)
	suite.NoError(err)
	suite.Equal(int64(0), length)

	_, err = suite.engine.Strlen(fakeValueKey)
	suite.Error(err)
}

func (suite *engineTestSuite) TestIntIncrDecr() {
	err := suite.engine.Set("intvalue", "1")
	suite.NoError(err)

	value, err := suite.engine.Incr("intvalue")
	suite.NoError(err)
	suite.Equal(int64(2), value)

	// TODO: make Incrby and Decrby use int64 instead of string? :|
	value, err = suite.engine.Incrby("intvalue", int64(10))
	suite.NoError(err)
	suite.Equal(int64(12), value)

	value, err = suite.engine.Decr("intvalue")
	suite.NoError(err)
	suite.Equal(int64(11), value)

	value, err = suite.engine.Decrby("intvalue", int64(100))
	suite.NoError(err)
	suite.Equal(int64(-89), value)

	value, err = suite.engine.Incr("incr")
	suite.NoError(err)
	suite.Equal(int64(1), value)

	value, err = suite.engine.Incrby("incrby", int64(10))
	suite.NoError(err)
	suite.Equal(int64(10), value)

	value, err = suite.engine.Decr("decr")
	suite.NoError(err)
	suite.Equal(int64(-1), value)

	value, err = suite.engine.Decrby("decrby", int64(10))
	suite.NoError(err)
	suite.Equal(int64(-10), value)
}

func (suite *engineTestSuite) TestOtherIncrDecr() {
	_, err := suite.engine.Incr(stringValueKey)
	suite.Error(err)

	_, err = suite.engine.Incrby(stringValueKey, int64(10))
	suite.Error(err)

	_, err = suite.engine.Decr(stringValueKey)
	suite.Error(err)

	_, err = suite.engine.Decrby(stringValueKey, int64(10))
	suite.Error(err)

	// make sure we haven't modified the string value anywhere in there
	value, err := suite.engine.Get(stringValueKey)
	suite.NoError(err)
	suite.Equal(stringValue, *value)

	// make sure we haven't modified the int value anywhere in there
	value, err = suite.engine.Get(intValueKey)
	suite.NoError(err)
	suite.Equal(intValueString, *value)

	_, err = suite.engine.Incr(fakeValueKey)
	suite.Error(err)

	// TODO: there's some subtlety here in whether the string parsing fails or
	// the key is of the wrong type (string, fakevalue, etc) that could use some assertions
	// but right now we're not sending specific error types so we'll ignore it
	_, err = suite.engine.Incrby(fakeValueKey, int64(10))
	suite.Error(err)

	_, err = suite.engine.Decr(fakeValueKey)
	suite.Error(err)

	_, err = suite.engine.Decrby(fakeValueKey, int64(10))
	suite.Error(err)

	typeString, err := suite.engine.Type(fakeValueKey)
	suite.NoError(err)
	suite.Equal("fake", typeString)
}

func (suite *engineTestSuite) TestMGet() {
	localStringValue := stringValue
	localIntValueString := intValueString
	values, err := suite.engine.MGet([]string{stringValueKey, intValueKey, doesNotExistKey, fakeValueKey})
	suite.NoError(err)
	suite.Equal([]*string{&localStringValue, &localIntValueString, nil, nil}, values)
}

func (suite *engineTestSuite) TestMSet() {
	engine := newEngine()
	pairs := map[string]string{
		intValueKey:    intValueString,
		stringValueKey: stringValue,
	}
	err := engine.MSet(pairs)
	suite.NoError(err)

	value, err := engine.Get(intValueKey)
	suite.NoError(err)
	suite.Equal(intValueString, *value)

	value, err = engine.Get(stringValueKey)
	suite.NoError(err)
	suite.Equal(stringValue, *value)
}

func (suite *engineTestSuite) TestType() {
	typeString, err := suite.engine.Type(intValueKey)
	suite.NoError(err)
	suite.Equal("string", typeString)

	typeString, err = suite.engine.Type(stringValueKey)
	suite.NoError(err)
	suite.Equal("string", typeString)

	typeString, err = suite.engine.Type(fakeValueKey)
	suite.NoError(err)
	suite.Equal("fake", typeString)

	typeString, err = suite.engine.Type(doesNotExistKey)
	suite.NoError(err)
	suite.Equal("none", typeString)
}

func (suite *engineTestSuite) TestGetOrDefault() {
	store := suite.engine.getOrDefault(stringValueKey, "default")
	value, err := store.get()
	suite.NoError(err)
	suite.Equal(stringValue, value)

	store = suite.engine.getOrDefault(doesNotExistKey, "abc")
	value, err = store.get()
	suite.NoError(err)
	suite.Equal("abc", value)
}

func (suite *engineTestSuite) TestGetStore() {
	store := suite.engine.getStore(stringValueKey)
	suite.NotNil(store)

	value, err := store.get()
	suite.NoError(err)
	suite.Equal(stringValue, value)

	store = suite.engine.getStore(doesNotExistKey)
	suite.Nil(store)
	// set
	// del
}

func (suite *engineTestSuite) TestInternalDel() {
	ret := suite.engine.del(stringValueKey)
	suite.True(ret)

	value, err := suite.engine.Get(stringValueKey)
	suite.NoError(err)
	suite.Nil(value)

	ret = suite.engine.del(doesNotExistKey)
	suite.False(ret)

	value, err = suite.engine.Get(doesNotExistKey)
	suite.NoError(err)
	suite.Nil(value)
}

func (suite *engineTestSuite) TestInternalSet() {
	store := suite.engine.getStore(stringValueKey)
	suite.NotNil(store)

	newStore := suite.engine.set(stringValueKey, stringValue)
	suite.NotNil(newStore)
	suite.False(store == newStore)

	suite.True(newStore == suite.engine.getStore(stringValueKey))
}

func TestEngineTestSuite(t *testing.T) {
	suite.Run(t, new(engineTestSuite))
}
